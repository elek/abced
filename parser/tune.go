package parser

import (
	"github.com/elek/abced/abcfile"
	abc "github.com/elek/abced/types"
	"github.com/zeebo/errs/v2"
	"strconv"
	"strings"
)

func ParseTune(tune abcfile.Tune) (*abc.Tune, error) {
	t := abc.NewTune()

	headers, lines := tune.HeadersAndLines()
	key := headers.Get("K")
	u := strings.Split(strings.TrimSpace(headers.Get("M")), "/")
	if len(u) < 2 {
		return nil, errs.Errorf("Invalid M: %s", u)
	}
	nom, _ := strconv.Atoi(u[0])
	denom, _ := strconv.Atoi(u[1])
	unit := abc.NewBeat(nom, denom)

	for _, line := range lines {
		score, err := ParseLine(line, unit, key)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		t.Lines = append(t.Lines, score)
	}

	return t, nil
}
