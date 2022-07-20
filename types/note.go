package abc

import "fmt"

type Note struct {
	Break  bool
	Pitch  Pitch
	Length Beat
	Tied   bool
}

func NewNote(pitch rune, beat Beat) Note {
	return Note{
		Pitch:  NewPitchFromRune(pitch),
		Length: beat,
	}
}
func (n *Note) Cadence(base Note) string {
	diff := n.Pitch - base.Pitch
	return cadences[diff+12]
}

func (n *Note) String() string {
	return fmt.Sprintf("%s%s", n.Letter(), n.Length)
}

func (n *Note) Letter() string {
	if n.Break {
		return "z"
	}
	return n.Pitch.String()
}

func (n *Note) Transpose(i int) {
	n.Pitch = n.Pitch + Pitch(i)
}

func (n *Note) MultiplyLength(beat Beat) {
	n.Length = n.Length.Multiply(beat)
}
