// Code generated from Abc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package grammar // Abc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAbcListener is a complete listener for a parse tree produced by AbcParser.
type BaseAbcListener struct{}

var _ AbcListener = &BaseAbcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAbcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAbcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAbcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAbcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLine is called when production line is entered.
func (s *BaseAbcListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseAbcListener) ExitLine(ctx *LineContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseAbcListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseAbcListener) ExitUnit(ctx *UnitContext) {}

// EnterNote is called when production note is entered.
func (s *BaseAbcListener) EnterNote(ctx *NoteContext) {}

// ExitNote is called when production note is exited.
func (s *BaseAbcListener) ExitNote(ctx *NoteContext) {}

// EnterDe is called when production de is entered.
func (s *BaseAbcListener) EnterDe(ctx *DeContext) {}

// ExitDe is called when production de is exited.
func (s *BaseAbcListener) ExitDe(ctx *DeContext) {}

// EnterTriplet is called when production triplet is entered.
func (s *BaseAbcListener) EnterTriplet(ctx *TripletContext) {}

// ExitTriplet is called when production triplet is exited.
func (s *BaseAbcListener) ExitTriplet(ctx *TripletContext) {}
