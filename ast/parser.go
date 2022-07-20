package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elek/abced/grammar"
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
	if e.err {
		return l.line, fmt.Errorf("eror on parsing the abc line: " + line)
	}
	return l.line, nil

}
