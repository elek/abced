package parser

import (
	"github.com/elek/abced/abcfile"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	tune := abcfile.Tune{
		Raw: `X:6
T: Hortobágyi kocsmárosné
K:Eb
L:1/4
M:4/4
g/g/g/g/ g/g/g/g/ | g e c2 |
f/f/f/f/ e/e/g/e/ | d d d2 |
=B c d e | f/e/d/c/  | =B A F2 |
f/e/d/c/ | =B/B/c/A/ | G G G2 |`,
	}

	parsed, err := ParseTune(tune)
	require.NoError(t, err)
	require.Equal(t, 4, len(parsed.Lines))
}
