// Code generated from Abc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package grammar

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AbcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var abclexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func abclexerLexerInit() {
	staticData := &abclexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "'/'", "'\\'", "'\\n'", "' '", "'|'", "'{'", "'}'", "'('",
		"')'", "':'", "'<'", "'>'", "'^'", "'='", "'_'", "','", "'H'", "'!roll!'",
		"'-'",
	}
	staticData.symbolicNames = []string{
		"", "PITCH", "NUMBER", "DIVIDER", "CONT", "NL", "WS", "SEP", "GS", "GE",
		"PS", "PE", "COLON", "LT", "GT", "SHARP", "NATURAL", "FLAT", "DOWN",
		"H", "ROLL", "TIE",
	}
	staticData.ruleNames = []string{
		"PITCH", "NUMBER", "DIVIDER", "CONT", "NL", "WS", "SEP", "GS", "GE",
		"PS", "PE", "COLON", "LT", "GT", "SHARP", "NATURAL", "FLAT", "DOWN",
		"H", "ROLL", "TIE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 92, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 0, 0, 21, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 1, 0, 2, 3, 0, 65,
		71, 97, 103, 122, 122, 1, 0, 48, 57, 91, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		1, 43, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5, 47, 1, 0, 0, 0, 7, 49, 1, 0, 0,
		0, 9, 51, 1, 0, 0, 0, 11, 55, 1, 0, 0, 0, 13, 57, 1, 0, 0, 0, 15, 59, 1,
		0, 0, 0, 17, 61, 1, 0, 0, 0, 19, 63, 1, 0, 0, 0, 21, 65, 1, 0, 0, 0, 23,
		67, 1, 0, 0, 0, 25, 69, 1, 0, 0, 0, 27, 71, 1, 0, 0, 0, 29, 73, 1, 0, 0,
		0, 31, 75, 1, 0, 0, 0, 33, 77, 1, 0, 0, 0, 35, 79, 1, 0, 0, 0, 37, 81,
		1, 0, 0, 0, 39, 83, 1, 0, 0, 0, 41, 90, 1, 0, 0, 0, 43, 44, 7, 0, 0, 0,
		44, 2, 1, 0, 0, 0, 45, 46, 7, 1, 0, 0, 46, 4, 1, 0, 0, 0, 47, 48, 5, 47,
		0, 0, 48, 6, 1, 0, 0, 0, 49, 50, 5, 92, 0, 0, 50, 8, 1, 0, 0, 0, 51, 52,
		5, 10, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 6, 4, 0, 0, 54, 10, 1, 0, 0, 0,
		55, 56, 5, 32, 0, 0, 56, 12, 1, 0, 0, 0, 57, 58, 5, 124, 0, 0, 58, 14,
		1, 0, 0, 0, 59, 60, 5, 123, 0, 0, 60, 16, 1, 0, 0, 0, 61, 62, 5, 125, 0,
		0, 62, 18, 1, 0, 0, 0, 63, 64, 5, 40, 0, 0, 64, 20, 1, 0, 0, 0, 65, 66,
		5, 41, 0, 0, 66, 22, 1, 0, 0, 0, 67, 68, 5, 58, 0, 0, 68, 24, 1, 0, 0,
		0, 69, 70, 5, 60, 0, 0, 70, 26, 1, 0, 0, 0, 71, 72, 5, 62, 0, 0, 72, 28,
		1, 0, 0, 0, 73, 74, 5, 94, 0, 0, 74, 30, 1, 0, 0, 0, 75, 76, 5, 61, 0,
		0, 76, 32, 1, 0, 0, 0, 77, 78, 5, 95, 0, 0, 78, 34, 1, 0, 0, 0, 79, 80,
		5, 44, 0, 0, 80, 36, 1, 0, 0, 0, 81, 82, 5, 72, 0, 0, 82, 38, 1, 0, 0,
		0, 83, 84, 5, 33, 0, 0, 84, 85, 5, 114, 0, 0, 85, 86, 5, 111, 0, 0, 86,
		87, 5, 108, 0, 0, 87, 88, 5, 108, 0, 0, 88, 89, 5, 33, 0, 0, 89, 40, 1,
		0, 0, 0, 90, 91, 5, 45, 0, 0, 91, 42, 1, 0, 0, 0, 1, 0, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AbcLexerInit initializes any static state used to implement AbcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAbcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AbcLexerInit() {
	staticData := &abclexerLexerStaticData
	staticData.once.Do(abclexerLexerInit)
}

// NewAbcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAbcLexer(input antlr.CharStream) *AbcLexer {
	AbcLexerInit()
	l := new(AbcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &abclexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Abc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AbcLexer tokens.
const (
	AbcLexerPITCH   = 1
	AbcLexerNUMBER  = 2
	AbcLexerDIVIDER = 3
	AbcLexerCONT    = 4
	AbcLexerNL      = 5
	AbcLexerWS      = 6
	AbcLexerSEP     = 7
	AbcLexerGS      = 8
	AbcLexerGE      = 9
	AbcLexerPS      = 10
	AbcLexerPE      = 11
	AbcLexerCOLON   = 12
	AbcLexerLT      = 13
	AbcLexerGT      = 14
	AbcLexerSHARP   = 15
	AbcLexerNATURAL = 16
	AbcLexerFLAT    = 17
	AbcLexerDOWN    = 18
	AbcLexerH       = 19
	AbcLexerROLL    = 20
	AbcLexerTIE     = 21
)
