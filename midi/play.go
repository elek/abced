package midi

import (
	"bytes"
	"fmt"
	"github.com/elek/abced/music"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/portmididrv"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
)

func PlayTune(tune *music.Tune) error {
	defer midi.CloseDriver()

	fmt.Println(midi.GetOutPorts())
	// create a SMF
	rd := bytes.NewReader(createMIDI(tune.Lines))

	// read and play it
	err := smf.ReadTracksFrom(rd).Play(midi.GetOutPorts()[1])
	return err
}

func createMIDI(scores []*music.Score) []byte {
	var (
		bf    bytes.Buffer
		clock = smf.MetricTicks(960)
		tr    smf.Track
	)

	// first track must have tempo and meter informations
	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(200))
	tr.Add(0, smf.MetaInstrument("Piano"))
	tr.Add(0, midi.ProgramChange(0, gm.Instr_AcousticGrandPiano.Value()))

	init := uint32(clock.Resolution())
	for _, score := range scores {
		for _, n := range score.Notes {
			if n.Break {
				init += uint32(clock.Resolution()) * 10 * uint32(n.Length.Nominator) / uint32(n.Length.Denominator)
			} else {
				p := n.Pitch + 48
				tr.Add(init, midi.NoteOn(0, uint8(p), 127))
				deltaticks := uint32(clock.Resolution()) * 10 * uint32(n.Length.Nominator) / uint32(n.Length.Denominator)
				tr.Add(deltaticks, midi.NoteOff(0, uint8(p)))
				init = 0

			}
		}
	}

	tr.Close(0)

	// create the SMF and add the tracks
	s := smf.New()
	s.TimeFormat = clock
	s.Add(tr)
	s.WriteTo(&bf)
	return bf.Bytes()
}
