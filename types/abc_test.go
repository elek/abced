package abc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCadence(t *testing.T) {
	base := Note{Pitch: NewPitchFromRune('G')}

	n := Note{Pitch: NewPitchFromRune('G')}
	require.Equal(t, "1", n.Cadence(base))

	n = Note{Pitch: NewPitchFromRune('g')}
	require.Equal(t, "8", n.Cadence(base))

	n = Note{Pitch: NewPitchFromRune('A')}
	require.Equal(t, "2", n.Cadence(base))
}

func repeat(list []int, times int, value int) []int {
	res := make([]int, 0)
	res = append(res, list...)
	for i := 0; i < times; i++ {
		res = append(res, value)
	}
	return res
}

func TestSliceToPitch(t *testing.T) {
	require.Equal(t, []int{7, 19, 8}, sliceToPitch([]string{"G", "g", "♭A"}))
}

func addNotes(tune *Tune, runes [][]rune, beat Beat) {
	for _, line := range runes {
		s := NewScore()
		for _, n := range line {
			s.AddNote(n, beat)
		}
		tune.Lines = append(tune.Lines, s)
	}
}

func createTune(runes [][]rune) *Tune {
	tune := NewTune()
	addNotes(tune, runes, NewBeat(1, 16))
	return tune
}
