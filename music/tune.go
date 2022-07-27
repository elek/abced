package music

import (
	"math"
)

type Tune struct {
	Fields map[string]string
	Lines  []*Score
}

func (t *Tune) MinMax() (Pitch, Pitch) {
	max := Pitch(math.MinInt8)
	min := Pitch(math.MaxInt8)
	for _, l := range t.Lines {
		n, x := l.MinMax()
		if x > max {
			max = x
		}
		if n < min {
			min = n
		}
	}
	return min, max
}

func (t *Tune) LastNote() *Note {
	return t.LastLine().LastNote()
}

func (t *Tune) LastLine() *Score {
	if len(t.Lines) == 0 {
		return nil
	}

	return t.Lines[len(t.Lines)-1]
}

func (t *Tune) Syllable() []int {
	res := make([]int, 0)
	for _, l := range t.Lines {
		res = append(res, l.Size())
	}
	return res
}

func NewTune() *Tune {
	return &Tune{
		Fields: make(map[string]string),
		Lines:  make([]*Score, 0),
	}
}
