package ui

import (
	"github.com/elek/abced/music"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestAddBar(t *testing.T) {
	tests := []struct {
		orig   string
		result string
	}{
		{
			orig:   "aaaa ",
			result: "aaaa | ",
		},
		{
			orig:   "a",
			result: "a",
		},
		{
			orig:   "aaaa | ",
			result: "aaaa | ",
		},
		{
			orig:   "aaaa | a b c d | e e ee ",
			result: "aaaa | a b c d | e e ee | ",
		},
		{
			orig:   "f<f e/d/c/d/ | B<G B d ",
			result: "f<f e/d/c/d/ | B<G B d | ",
		},
	}

	for _, ts := range tests {
		t.Run(strings.ReplaceAll(ts.orig, "|", "X"), func(t *testing.T) {
			l := NewLine(ts.orig, music.NewBeat(1, 4))
			l.ReIndex()
			l.addBar()
			require.Equal(t, ts.result, l.content.Value())
		})

	}

}
