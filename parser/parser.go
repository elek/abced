package parser

import (
	"fmt"
	"github.com/elek/abced/abcfile"
	"github.com/elek/abced/ast"
	"github.com/elek/abced/music"
	"github.com/zeebo/errs/v2"
	"strconv"
	"strings"
)

func ParseAST(line *ast.Line, unit music.Beat, key string) (*music.Score, error) {
	p := Parser{
		score:     music.NewScore(),
		unit:      unit,
		key:       key,
		modifiers: map[music.Pitch]int{},
	}
	ast.Walk(&p, line)
	return p.score, p.err
}

func ParseLine(score abcfile.Score, unit music.Beat, key string) (*music.Score, error) {
	line := string(score)
	if len(strings.TrimSpace(line)) == 0 {
		return music.NewScore(), nil
	}
	parsed, err := ast.BuildAST(line)
	if err != nil {
		return nil, err
	}
	p := Parser{
		score:     music.NewScore(),
		unit:      unit,
		key:       key,
		modifiers: map[music.Pitch]int{},
	}
	ast.Walk(&p, parsed)
	return p.score, p.err
}

type Parser struct {
	score            *music.Score
	unit             music.Beat
	err              error
	modifiers        map[music.Pitch]int
	key              string
	tripletRemaining int
	triplet          ast.Triplet
}

func (p *Parser) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return
	}
	switch n := node.(type) {
	case *ast.Line:
	case ast.Note:
		note := p.createNote(n)
		p.score.Notes = append(p.score.Notes, note)
		if p.tripletRemaining > 0 {
			p.tripletRemaining--
			if p.tripletRemaining == 0 {
				updated := 0
				for i := len(p.score.Notes) - 1; i >= 0 && updated < p.triplet.Length(); i-- {
					mod := p.triplet.Modifier()
					p.score.Notes[i].Length = p.score.Notes[i].Length.Multiply(music.NewBeat(mod.Denominator, mod.Nominator))
					updated++
				}
			}
		}
		return nil
	case ast.Broken:
		n1 := p.createNote(n.First)
		n2 := p.createNote(n.Second)

		if n.Symbol.Symbol == ">" {
			n1.Length = n1.Length.Multiply(music.NewBeat(3, 2))
			n2.Length = n2.Length.Multiply(music.NewBeat(1, 2))
		} else if n.Symbol.Symbol == "<" {
			n2.Length = n2.Length.Multiply(music.NewBeat(3, 2))
			n1.Length = n1.Length.Multiply(music.NewBeat(1, 2))

		}
		p.score.Notes = append(p.score.Notes, n1)
		p.score.Notes = append(p.score.Notes, n2)
		return nil
	case ast.Bar:
		p.modifiers = map[music.Pitch]int{}
	case ast.Triplet:
		p.tripletRemaining = n.Length()
		p.triplet = n
	default:
		panic(fmt.Sprintf("unsupported node %T", node))
	}
	return p
}

func (p *Parser) createNote(note ast.Note) *music.Note {

	pitch := music.NewPitchFromRune(rune(note.Letter.Pitch[0]))
	switch note.Accidental.Type {
	case ast.Flat:
		p.modifiers[pitch] = -1
	case ast.Natural:
		p.modifiers[pitch] = 0
	case ast.Sharp:
		p.modifiers[pitch] = 1

	}
	if mod, found := p.modifiers[pitch]; found {
		pitch = music.Pitch(pitch + music.Pitch(mod))
	} else if keyTrafo, f := keyTransformation[p.key]; f {
		if mod, found := p.modifiers[pitch]; !found || mod != 0 {
			if nv, f2 := keyTrafo[pitch]; f2 {
				pitch = nv
			}
		}
	}

	return &music.Note{
		Pitch:  pitch,
		Length: toBeat(p.unit, note.Length),
	}
}

func ParseTune(tune abcfile.Tune) (*music.Tune, error) {
	t := music.NewTune()

	headers, lines := tune.HeadersAndScores()
	key := headers.Get("K")
	unit := music.NewBeat(1, 4)
	u := strings.Split(strings.TrimSpace(headers.Get("M")), "/")
	if len(u) >= 2 {
		nom, _ := strconv.Atoi(u[0])
		denom, _ := strconv.Atoi(u[1])
		unit = music.NewBeat(nom, denom)
	}

	for _, line := range lines {
		score, err := ParseLine(line, unit, key)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		t.Lines = append(t.Lines, score)
	}

	return t, nil
}

func toBeat(unit music.Beat, length *ast.Length) music.Beat {
	if length == nil {
		return unit
	}
	if length.Divider != nil {
		if length.First != nil && length.Last != nil {
			return unit.Multiply(music.NewBeat(length.First.Value, length.Last.Value))
		}
		if length.Last == nil {
			t := unit
			for i := 0; i < length.Divider.Number; i++ {
				t = t.Multiply(music.NewBeat(1, 2))
			}
			return t
		}
		if length.Last != nil {
			return unit.Multiply(music.NewBeat(1, length.Last.Value))
		}
	} else if length.First != nil {
		return unit.Multiply(music.NewBeat(length.First.Value, 1))
	}

	return unit
}
