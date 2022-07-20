package abcfile

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadAbc(t *testing.T) {
	a := `T:asd

X:1
T:qwe
K:G
abc d

X:2
T:qwex
K:G
abc d
`
	abc, err := Read(a)
	require.NoError(t, err)
	require.Equal(t, 2, len(abc.Scores))
	require.Equal(t, 1, len(abc.Headers))
}
