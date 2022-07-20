package ast

import "fmt"

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

type LetterModifier struct {
}

type Note struct {
	Positioned
	Letter   Letter
	Modifier *LetterModifier
	Length   *Length
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
