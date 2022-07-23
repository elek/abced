package abcfile

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestItems(t *testing.T) {
	tu := Tune{
		Raw: `a:x
b: f 
abcde
w:xx
efef

`}

	res := tu.Items()
	require.Equal(t, 5, len(res))
	require.Equal(t, Header{Key: "a", Value: "x"}, res[0])
	require.Equal(t, Header{Key: "b", Value: "f"}, res[1])
	require.Equal(t, Score("abcde"), res[2])
	require.Equal(t, Header{Key: "w", Value: "xx"}, res[3])
}
