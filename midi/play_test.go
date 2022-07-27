package midi

import (
	"github.com/elek/abced/abcfile"
	"github.com/elek/abced/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlay(t *testing.T) {
	t.Skip()
	s := abcfile.Tune{Raw: `M: 1/4
K:C
C/E/ C/E/ G G (3:2:3 C/C/E C/E/ G G"`}
	parsed, err := parser.ParseTune(s)
	require.NoError(t, err)
	require.NoError(t, PlayTune(parsed))
}

func TestPlay2(t *testing.T) {
	t.Skip()
	s := abcfile.Tune{
		Raw: `X:5
T: Jobbról-balra sirítem
O: Gyimesközéplok-Hidegség (Csík), Tímár Jánosné - Csorba Anna (63), 1968.07
M:4/4
L:1/4
Q: 1/4=108
K:C
(3:2:2 G/D G G B/<B/d/d/ | d<c B z | 
G/G/B/B/ G/<G/d/d/ | A<A A z | 
G/<G/ (3:2:2 B/B B/G/ B | d/<d/ (3:2:2 A/A (3:2:2 B/B | G/z/ |
e/e/e/<d/ e/<d/c/c/ | A<G G z | `,
	}
	parsed, err := parser.ParseTune(s)
	require.NoError(t, err)
	require.NoError(t, PlayTune(parsed))
}
