package parser

import (
	"fmt"
	"github.com/elek/abced/abcfile"
	abc "github.com/elek/abced/music"
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

func TestPaseLineHSep(t *testing.T) {
	_, err := ParseLine("a H |\n", abc.NewBeat(1, 4), "C")
	require.Nil(t, err)
}

func TestParseLineAntlr(t *testing.T) {
	s, err := ParseLine("g<f d | g a/g/ |", abc.NewBeat(1, 4), "C")
	fmt.Println(s.AbcString(abc.NewBeat(1, 4), abc.NewBeat(1, 4)))
	require.Nil(t, err)
	require.Equal(t, abc.NewBeat(1, 8), s.Notes[0].Length)
	require.Equal(t, abc.NewBeat(1, 4), s.Notes[3].Length)
	require.Equal(t, abc.NewBeat(3, 8), s.Notes[1].Length)
}

func TestParseLineLength(t *testing.T) {
	s, err := ParseLine("a>b a/2 a3/2 a/ a// a2", abc.NewBeat(1, 4), "C")
	require.Nil(t, err)
	require.Equal(t, abc.NewBeat(3, 8), s.Notes[0].Length)
	require.Equal(t, abc.NewBeat(1, 8), s.Notes[1].Length)
	require.Equal(t, abc.NewBeat(1, 8), s.Notes[2].Length)
	require.Equal(t, abc.NewBeat(3, 8), s.Notes[3].Length)
	require.Equal(t, abc.NewBeat(1, 8), s.Notes[4].Length)
	require.Equal(t, abc.NewBeat(1, 16), s.Notes[5].Length)
	require.Equal(t, abc.NewBeat(1, 2), s.Notes[6].Length)

}

func TestParseLineWithKeyAntlr(t *testing.T) {
	s, err := ParseLine("A B c", abc.NewBeat(1, 4), "F")
	require.Nil(t, err)
	require.Equal(t, abc.Pitch(9), s.Notes[0].Pitch)
	require.Equal(t, abc.Pitch(10), s.Notes[1].Pitch)
	require.Equal(t, abc.Pitch(12), s.Notes[2].Pitch)
}

func TestParseLineWithModifiers(t *testing.T) {
	s, err := ParseLine("A B =B c ^c _d", abc.NewBeat(1, 4), "F")
	require.Nil(t, err)
	require.Equal(t, abc.Pitch(9), s.Notes[0].Pitch)
	require.Equal(t, abc.Pitch(10), s.Notes[1].Pitch)
	require.Equal(t, abc.Pitch(11), s.Notes[2].Pitch)
	require.Equal(t, abc.Pitch(12), s.Notes[3].Pitch)
	require.Equal(t, abc.Pitch(13), s.Notes[4].Pitch)
	require.Equal(t, abc.Pitch(13), s.Notes[5].Pitch)
}

func TestParseLineWithModifiersAndBar(t *testing.T) {
	s, err := ParseLine("A _A | A", abc.NewBeat(1, 4), "F")
	require.Nil(t, err)
	require.Equal(t, abc.Pitch(9), s.Notes[0].Pitch)
	require.Equal(t, abc.Pitch(8), s.Notes[1].Pitch)
	require.Equal(t, abc.Pitch(9), s.Notes[2].Pitch)
}

func TestParseLineWithNatural(t *testing.T) {
	s, err := ParseLine("B =B | B", abc.NewBeat(1, 4), "F")
	require.Nil(t, err)
	require.Equal(t, abc.Pitch(10), s.Notes[0].Pitch)
	require.Equal(t, abc.Pitch(11), s.Notes[1].Pitch)
	require.Equal(t, abc.Pitch(10), s.Notes[2].Pitch)
}

func TestParseTriplet(t *testing.T) {
	s, err := ParseLine("a (3:2:3 a a a", abc.NewBeat(1, 4), "F")
	require.Nil(t, err)
	require.Equal(t, abc.NewBeat(1, 4), s.Notes[0].Length)
	require.Equal(t, abc.NewBeat(1, 6), s.Notes[1].Length)
	require.Equal(t, abc.NewBeat(1, 6), s.Notes[2].Length)
	require.Equal(t, abc.NewBeat(1, 6), s.Notes[3].Length)

}

func TestParseSample1(t *testing.T) {
	s := abcfile.Tune{Raw: `M: 1/4
K:C
C/E/ C/E/ G G (3:2:3 C/C/E C/E/ G G"`}
	parsed, err := ParseTune(s)
	require.Nil(t, err)
	line := parsed.Lines[0].Notes
	require.Equal(t, abc.NewBeat(1, 8), line[0].Length)
	require.Equal(t, abc.NewBeat(1, 8), line[1].Length)
	require.Equal(t, abc.NewBeat(1, 8), line[2].Length)
}

//func TestParseLineGraceAntlr(t *testing.T) {
//	s, err := parseLineAntlr("{A}BA", abc.NewBeat(1, 4), "C")
//	require.Nil(t, err)
//	require.Equal(t, 2, len(s.Notes))
//	require.Equal(t, 16, s.Notes[0].Size)
//	require.Equal(t, 16, s.Notes[1].Size)
//}
//
//func TestParseLineTied(t *testing.T) {
//	s, err := parseLineAntlr("a (ab) c", abc.NewBeat(1, 4), "C")
//	require.Nil(t, err)
//	require.Equal(t, 4, len(s.Notes))
//	require.Equal(t, false, s.Notes[0].Tied)
//	require.Equal(t, false, s.Notes[1].Tied)
//	require.Equal(t, true, s.Notes[2].Tied)
//	require.Equal(t, false, s.Notes[3].Tied)
//
//	s, err = parseLineAntlr("a ((ab) c)", abc.NewBeat(1, 4), "C")
//	require.Nil(t, err)
//	require.Equal(t, 4, len(s.Notes))
//	require.Equal(t, false, s.Notes[0].Tied)
//	require.Equal(t, false, s.Notes[1].Tied)
//	require.Equal(t, true, s.Notes[2].Tied)
//	require.Equal(t, true, s.Notes[3].Tied)
//}
//
//func TestParseLower(t *testing.T) {
//	s, err := parseLineAntlr("C ^G,", abc.NewBeat(1, 4), "C")
//	require.Nil(t, err)
//	require.Equal(t, 2, len(s.Notes))
//	require.Equal(t, 0, s.Notes[0].Pitch)
//	require.Equal(t, -4, s.Notes[1].Pitch)
//}
//
//func TestParseBreak(t *testing.T) {
//	s, err := parseLineAntlr("A z,", abc.NewBeat(1, 4), "C")
//	require.Nil(t, err)
//	require.Equal(t, 2, len(s.Notes))
//	require.Equal(t, 9, s.Notes[0].Pitch)
//	require.Equal(t, 0, s.Notes[1].Pitch)
//	require.False(t, s.Notes[0].Break)
//	require.True(t, s.Notes[1].Break)
//}
