package ast

import "github.com/antlr/antlr4/runtime/Go/antlr"

type errorListener struct {
	*antlr.DefaultErrorListener
	err bool
}

func (d *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.err = true
}
