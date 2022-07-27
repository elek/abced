package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elek/abced/grammar"
	"github.com/zeebo/errs/v2"
)

func BuildAST(line string) (*Line, error) {
	input := antlr.NewInputStream(line)
	lexer := grammar.NewAbcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := grammar.NewAbcParser(stream)

	e := errorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(&e)

	p.BuildParseTrees = true

	l := abcListener{
		line: &Line{
			Nodes: make([]Node, 0),
		},
	}
	antlr.ParseTreeWalkerDefault.Walk(&l, p.Line())
	if e.err != nil {
		return l.line, errs.Errorf("error on parsing the abc line: %s %v", line, e.err)
	}
	return l.line, nil

}
