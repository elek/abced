package music

import "fmt"

type Pitch int

var pitches = map[rune]Pitch{
	'C': Pitch(0),
	'D': Pitch(2),
	'E': Pitch(4),
	'F': Pitch(5),
	'G': Pitch(7),
	'A': Pitch(9),
	'B': Pitch(11),
	'c': Pitch(12),
	'd': Pitch(14),
	'e': Pitch(16),
	'f': Pitch(17),
	'g': Pitch(19),
	'a': Pitch(21),
	'b': Pitch(23),
}

func NewPitchFromRune(note rune) Pitch {
	for k, v := range pitches {
		if k == note {
			return v
		}
	}
	return 999
}

func NewPitchFromString(s string) Pitch {
	note := rune(s[0])
	for k, v := range pitches {
		if k == note {
			return v
		}
	}
	return 999
}

func (p Pitch) Distance(o Pitch) int {
	if p > o {
		return int(p - o)
	}
	return int(o - p)
}

func (p Pitch) String() string {
	if p < 0 {
		return Pitch(p+12).String() + ","
	}
	for k, v := range pitches {
		if v == p {
			return string(k)
		}
	}
	for k, v := range pitches {
		if v == p-1 {
			return "♯" + string(k)
		}
	}
	return fmt.Sprintf("?{%d}", p)
}
