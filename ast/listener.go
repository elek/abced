package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elek/abced/grammar"
	"strconv"
)

type abcListener struct {
	*grammar.BaseAbcListener
	line *Line
	skip bool
}

func (s *abcListener) EnterDe(ctx *grammar.DeContext) {

}
func (s *abcListener) ExitDe(ctx *grammar.DeContext) {
	if ctx.GT() != nil || ctx.LT() != nil {
		if len(s.line.Nodes) < 2 {
			// unfinished line
			return
		}
		a := s.line.Nodes[len(s.line.Nodes)-2]
		b := s.line.Nodes[len(s.line.Nodes)-1]
		s.line.Nodes = s.line.Nodes[:len(s.line.Nodes)-2]
		if _, ok := a.(Note); !ok {
			return
		}

		if _, ok := b.(Note); !ok {
			return
		}

		br := Broken{
			First:  a.(Note),
			Second: b.(Note),
		}
		if ctx.GT() != nil {
			br.Symbol = Symbol{
				Symbol:     ">",
				Positioned: position(ctx.GT()),
			}
		}
		if ctx.LT() != nil {
			br.Symbol = Symbol{
				Symbol:     "<",
				Positioned: position(ctx.LT()),
			}
		}
		s.line.Nodes = append(s.line.Nodes, br)
	}

	if ctx.SEP() != nil {
		s.line.Nodes = append(s.line.Nodes, Bar{
			Symbol:     ctx.SEP().GetText(),
			Positioned: position(ctx.SEP()),
		})
	}

}

func (s *abcListener) ExitNote(ctx *grammar.NoteContext) {
	n := Note{
		Letter: Letter{
			Pitch:      ctx.PITCH().GetText(),
			Positioned: position(ctx.PITCH()),
		},
		Length: &Length{},
	}
	divider := ctx.AllDIVIDER()
	if len(divider) > 0 {
		n.Length.Divider = &Divider{
			Number:     len(divider),
			Positioned: position(ctx.DIVIDER(0)),
		}
	}

	if len(ctx.AllNUMBER()) > 0 {
		val, _ := strconv.Atoi(ctx.NUMBER(0).GetText())
		n.Length.First = &Number{
			Value: val,
			Positioned: Positioned{
				Position: Pos(ctx.GetSourceInterval().Start),
			},
		}
	}

	if len(ctx.AllNUMBER()) > 1 {
		val, _ := strconv.Atoi(ctx.NUMBER(1).GetText())
		n.Length.Last = &Number{
			Value: val,
			Positioned: Positioned{
				Position: Pos(ctx.GetSourceInterval().Start),
			},
		}
	}

	s.line.Nodes = append(s.line.Nodes, n)
}

func position(gt antlr.TerminalNode) Positioned {
	return Positioned{
		Position: Pos(gt.GetSourceInterval().Start),
	}

}
