package music

import (
	"math"
	"strings"
)

type Score struct {
	Notes []*Note
}

func NewScore() *Score {
	return &Score{
		Notes: make([]*Note, 0),
	}
}
func (s *Score) AddNote(pitch rune, length Beat) {
	s.Notes = append(s.Notes, &Note{
		Pitch:  NewPitchFromRune(pitch),
		Length: length,
	})
}

func (s *Score) AddBreak(length Beat) {
	s.Notes = append(s.Notes, &Note{
		Length: length,
		Break:  true,
	})
}

func (s *Score) AddRawNote(i Pitch, length Beat) {
	s.Notes = append(s.Notes, &Note{
		Pitch:  i,
		Length: length,
	})
}

func (s *Score) MinMax() (Pitch, Pitch) {
	max := Pitch(math.MinInt8)
	min := Pitch(math.MaxInt8)
	for _, n := range s.Notes {
		if n.Pitch == -1 {
			continue
		}
		if n.Pitch > max {
			max = n.Pitch
		}
		if n.Pitch < min {
			min = n.Pitch
		}
	}
	return min, max
}

func (s *Score) String() string {
	k := make([]string, 0)
	for _, n := range s.Notes {
		k = append(k, n.String())
	}
	return strings.Join(k, " ")
}

func (s *Score) AbcString(unit Beat, meter Beat) string {
	meter = meter.Simplify()
	k := make([]string, 0)
	cm := NewBeat(0, 0)
	for _, n := range s.Notes {
		relLength := n.Length.Divide(unit)
		k = append(k, n.Letter()+relLength.AbcString())
		cm = cm.Add(n.Length)
		if cm == meter {
			k = append(k, "|")
			cm = NewBeat(0, 0)
		}
	}
	return strings.Join(k, " ")
}

func (s *Score) Append(n *Score) {
	s.Notes = append(s.Notes, n.Notes...)
}

func (s *Score) Transpose(diff int) *Score {
	res := NewScore()
	for _, n := range s.Notes {
		if n.Pitch == -1 {
			res.AddRawNote(n.Pitch, n.Length)
		} else {
			res.AddRawNote(n.Pitch+Pitch(diff), n.Length)
		}
	}
	return res
}

func (s *Score) LastNote() *Note {
	for o := len(s.Notes) - 1; o >= 0; o-- {
		if !s.Notes[o].Break {
			return s.Notes[o]
		}
	}
	panic("empty line")
}

func (s *Score) Size() int {
	i := 0
	for _, k := range s.Notes {
		if k.Pitch != -1 && !k.Tied {
			i++
		}
	}
	return i
}

func (s *Score) Length() Beat {
	i := NewBeat(0, 0)
	for _, k := range s.Notes {
		if k.Pitch != -1 && !k.Tied {
			i = i.Add(k.Length)
		}
	}
	return i.Simplify()
}
