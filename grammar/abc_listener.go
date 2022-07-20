// Code generated from Abc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package grammar // Abc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// AbcListener is a complete listener for a parse tree produced by AbcParser.
type AbcListener interface {
	antlr.ParseTreeListener

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterNote is called when entering the note production.
	EnterNote(c *NoteContext)

	// EnterDe is called when entering the de production.
	EnterDe(c *DeContext)

	// EnterTriplet is called when entering the triplet production.
	EnterTriplet(c *TripletContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitNote is called when exiting the note production.
	ExitNote(c *NoteContext)

	// ExitDe is called when exiting the de production.
	ExitDe(c *DeContext)

	// ExitTriplet is called when exiting the triplet production.
	ExitTriplet(c *TripletContext)
}
