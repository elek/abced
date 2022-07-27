package music

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMultiply(t *testing.T) {
	require.Equal(t, NewBeat(1, 2), NewBeat(2, 3).Multiply(NewBeat(3, 4)))
}

func TestDivide(t *testing.T) {
	require.Equal(t, NewBeat(2, 1), NewBeat(1, 2).Divide(NewBeat(1, 4)))
}

func TestGCD(t *testing.T) {
	require.Equal(t, 1, gcd(5, 6))
	require.Equal(t, 2, gcd(2, 6))
	require.Equal(t, 8, gcd(16, 8))
}
