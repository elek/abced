package ast

import (
	"fmt"
	"github.com/elek/abced/music"
)

type Pos int

type Node interface {
	Pos() Pos
}

type Positioned struct {
	Position Pos
}

func (p Positioned) Pos() Pos {
	return p.Position
}

type Line struct {
	Positioned
	Nodes []Node
}

type Triplet struct {
	Positioned
	P uint8
	Q uint8
	R uint8
}

func (t Triplet) Modifier() music.Beat {
	return music.NewBeat(int(t.P), int(t.Q))
}

func (t Triplet) Length() int {
	if t.R == 0 {
		return int(t.P)
	}
	return int(t.R)
}

type Divider struct {
	Positioned
	Number int
}

func (d Divider) String() string {
	return fmt.Sprintf("//-%d", d.Number)
}

type Number struct {
	Positioned
	Value int
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

type Length struct {
	Positioned
	First   *Number
	Divider *Divider
	Last    *Number
}

func (l Length) String() string {
	return fmt.Sprintf("%s %s %s", l.First, l.Divider, l.Last)
}

type Letter struct {
	Positioned
	Pitch string
}

type Accidental struct {
	Positioned
	Type AccidentalType
}

type AccidentalType uint8

const (
	Nothing AccidentalType = iota
	Flat
	Natural
	Sharp
)

type Note struct {
	Positioned
	Letter     Letter
	Accidental Accidental
	Length     *Length
}

type Rest struct {
	Positioned
}

type Broken struct {
	Positioned
	First  Note
	Second Note
	Symbol Symbol
}

type Symbol struct {
	Positioned
	Symbol string
}

type Bar struct {
	Positioned
	Symbol string
}
