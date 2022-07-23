package parser

import (
	"fmt"
	"github.com/elek/abced/abcfile"
	"github.com/elek/abced/ast"
	abc "github.com/elek/abced/types"
	"strings"
)

func ParseAST(line *ast.Line, unit abc.Beat, key string) (*abc.Score, error) {
	p := Parser{
		score: abc.NewScore(),
		unit:  unit,
	}
	ast.Walk(&p, line)
	return p.score, p.err
}

func ParseLine(score abcfile.Score, unit abc.Beat, key string) (*abc.Score, error) {
	line := string(score)
	if len(strings.TrimSpace(line)) == 0 {
		return abc.NewScore(), nil
	}
	parsed, err := ast.BuildAST(line)
	if err != nil {
		return nil, err
	}
	p := Parser{
		score: abc.NewScore(),
		unit:  unit,
	}
	ast.Walk(&p, parsed)
	return p.score, p.err
}

type Parser struct {
	score *abc.Score
	unit  abc.Beat
	err   error
}

func (p *Parser) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return
	}
	switch n := node.(type) {
	case *ast.Line:
	case ast.Note:
		note := &abc.Note{
			Pitch:  abc.NewPitchFromString(n.Letter.Pitch),
			Length: toBeat(p.unit, n.Length),
		}
		p.score.Notes = append(p.score.Notes, note)
		return nil
	case ast.Broken:
		n1 := &abc.Note{
			Pitch:  abc.NewPitchFromString(n.First.Letter.Pitch),
			Length: toBeat(p.unit, n.First.Length),
		}

		n2 := &abc.Note{
			Pitch:  abc.NewPitchFromString(n.Second.Letter.Pitch),
			Length: toBeat(p.unit, n.Second.Length),
		}

		if n.Symbol.Symbol == ">" {
			n1.Length = n1.Length.Multiply(abc.NewBeat(3, 2))
			n2.Length = n2.Length.Multiply(abc.NewBeat(1, 2))
		} else if n.Symbol.Symbol == "<" {
			n2.Length = n2.Length.Multiply(abc.NewBeat(3, 2))
			n1.Length = n1.Length.Multiply(abc.NewBeat(1, 2))

		}

		p.score.Notes = append(p.score.Notes, n1)
		p.score.Notes = append(p.score.Notes, n2)
		return nil
	case ast.Bar:
	default:
		panic(fmt.Sprintf("unsupported node %T", node))
	}
	return p
}

func toBeat(unit abc.Beat, length *ast.Length) abc.Beat {
	if length == nil {
		return unit
	}
	if length.Divider != nil {
		if length.First != nil && length.Last != nil {
			return unit.Multiply(abc.NewBeat(length.First.Value, length.Last.Value))
		}
		if length.Last == nil {
			t := unit
			for i := 0; i < length.Divider.Number; i++ {
				t = t.Multiply(abc.NewBeat(1, 2))
			}
			return t
		}
		if length.Last != nil {
			return unit.Multiply(abc.NewBeat(1, length.Last.Value))
		}
	} else if length.First != nil {
		return unit.Multiply(abc.NewBeat(length.First.Value, 1))
	}

	return unit
}
