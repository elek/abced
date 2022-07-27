package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/zeebo/errs/v2"
)

type errorListener struct {
	*antlr.DefaultErrorListener
	err error
}

func (d *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.err = errs.Errorf("%s", msg)
}
