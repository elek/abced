package abc

import (
	"strconv"
)

var scaleModes = []string{"major", "dorian", "phrygian", "lydian", "mixolydian", "minor", "locrian"}
var scaleModesPreference = []string{"major", "minor", "mixolydian", "dorian", "phrygian", "lydian", "locrian"}

var scaleTypes = map[int]string{
	4:  "tertratonic",
	5:  "pentatonic",
	6:  "hexatonic",
	7:  "heptatonic",
	8:  "octatonic",
	9:  "nonatonic",
	12: "chromatic",
}
var keyTransformation = map[string]map[int]int{
	"D": {
		5:  6,
		17: 18,
		12: 11,
	},
	"G": {
		5:  6,
		17: 18,
	},
	"C": {},
	"F": {
		11: 10,
		23: 22,
	},
	"Bb": {
		4:  3,
		11: 10,
		16: 15,
		23: 22,
	},
	"Eb": {
		4:  3,
		11: 10,
		16: 15,
		23: 22,
		9:  8,
		21: 20,
	},
}

func diffToSymbol(diff int) string {
	if diff+12 >= len(cadences) {
		panic("Difference couldn't be encoded: " + strconv.Itoa(diff))
	}
	return cadences[diff+12]
}

func sliceToPitch(strings []string) []Pitch {
	res := make([]Pitch, len(strings))
	for ix, note := range strings {
		v := Pitch(0)
		for _, r := range note {
			if r == '♭' {
				v--
			} else if r == '♯' {
				v++
			}
			for pr, pv := range pitches {
				if pr == r {
					v += pv
				}
			}
		}
		res[ix] = v
	}
	return res
}
