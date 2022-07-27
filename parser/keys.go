package parser

import (
	abc "github.com/elek/abced/music"
)

var keyTransformation = map[string]map[abc.Pitch]abc.Pitch{
	"D": {
		abc.Pitch(5):  abc.Pitch(6),
		abc.Pitch(17): abc.Pitch(18),
		abc.Pitch(12): abc.Pitch(11),
	},
	"G": {
		abc.Pitch(5):  abc.Pitch(6),
		abc.Pitch(17): abc.Pitch(18),
	},
	"C": {},
	"F": {
		abc.Pitch(11): abc.Pitch(10),
		abc.Pitch(23): abc.Pitch(22),
	},
	"Bb": {
		abc.Pitch(4):  abc.Pitch(3),
		abc.Pitch(11): abc.Pitch(10),
		abc.Pitch(16): abc.Pitch(15),
		abc.Pitch(23): abc.Pitch(22),
	},
	"Eb": {
		abc.Pitch(4):  abc.Pitch(3),
		abc.Pitch(11): abc.Pitch(10),
		abc.Pitch(16): abc.Pitch(15),
		abc.Pitch(23): abc.Pitch(22),
		abc.Pitch(9):  abc.Pitch(8),
		abc.Pitch(21): abc.Pitch(20),
	},
}
