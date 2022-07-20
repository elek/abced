package parser

import (
	"fmt"
	abc "github.com/elek/abced/types"
	"github.com/stretchr/testify/require"
	"testing"
)

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

//
//func TestParseLineWithKeyAntlr(t *testing.T) {
//	s, err := parseLineAntlr("A B c", abc.NewBeat(1, 4), "F")
//	require.Nil(t, err)
//	require.Equal(t, 9, s.Notes[0].Pitch)
//	require.Equal(t, 10, s.Notes[1].Pitch)
//	require.Equal(t, 12, s.Notes[2].Pitch)
//}
//
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
