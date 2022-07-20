// Code generated from Abc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package grammar // Abc
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AbcParser struct {
	*antlr.BaseParser
}

var abcParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func abcParserInit() {
	staticData := &abcParserStaticData
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
		"line", "unit", "note", "de", "triplet",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 132, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 4, 0, 12, 8, 0, 11, 0, 12, 0, 13, 1, 0, 3, 0, 17, 8, 0, 1, 0,
		3, 0, 20, 8, 0, 1, 0, 3, 0, 23, 8, 0, 1, 1, 3, 1, 26, 8, 1, 1, 1, 1, 1,
		3, 1, 30, 8, 1, 1, 1, 3, 1, 33, 8, 1, 4, 1, 35, 8, 1, 11, 1, 12, 1, 36,
		1, 2, 3, 2, 40, 8, 2, 1, 2, 3, 2, 43, 8, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2,
		1, 2, 3, 2, 50, 8, 2, 1, 2, 5, 2, 53, 8, 2, 10, 2, 12, 2, 56, 9, 2, 1,
		2, 3, 2, 59, 8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 64, 8, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 70, 8, 3, 1, 3, 4, 3, 73, 8, 3, 11, 3, 12, 3, 74, 1, 3, 1,
		3, 1, 3, 1, 3, 4, 3, 81, 8, 3, 11, 3, 12, 3, 82, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 89, 8, 3, 1, 3, 1, 3, 3, 3, 93, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 99, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		110, 8, 3, 1, 3, 5, 3, 113, 8, 3, 10, 3, 12, 3, 116, 9, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 122, 8, 4, 3, 4, 124, 8, 4, 1, 4, 1, 4, 3, 4, 128, 8,
		4, 3, 4, 130, 8, 4, 1, 4, 0, 1, 6, 5, 0, 2, 4, 6, 8, 0, 1, 1, 0, 15, 17,
		159, 0, 11, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 98, 1,
		0, 0, 0, 8, 117, 1, 0, 0, 0, 10, 12, 3, 2, 1, 0, 11, 10, 1, 0, 0, 0, 12,
		13, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 16, 1, 0, 0,
		0, 15, 17, 5, 4, 0, 0, 16, 15, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 19,
		1, 0, 0, 0, 18, 20, 5, 6, 0, 0, 19, 18, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0,
		20, 22, 1, 0, 0, 0, 21, 23, 5, 5, 0, 0, 22, 21, 1, 0, 0, 0, 22, 23, 1,
		0, 0, 0, 23, 1, 1, 0, 0, 0, 24, 26, 5, 19, 0, 0, 25, 24, 1, 0, 0, 0, 25,
		26, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 30, 3, 6, 3, 0, 28, 30, 5, 7, 0,
		0, 29, 27, 1, 0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 33,
		5, 6, 0, 0, 32, 31, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 35, 1, 0, 0, 0,
		34, 25, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1,
		0, 0, 0, 37, 3, 1, 0, 0, 0, 38, 40, 5, 20, 0, 0, 39, 38, 1, 0, 0, 0, 39,
		40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 43, 7, 0, 0, 0, 42, 41, 1, 0, 0,
		0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 46, 5, 1, 0, 0, 45, 47,
		5, 18, 0, 0, 46, 45, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 49, 1, 0, 0, 0,
		48, 50, 5, 2, 0, 0, 49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 54, 1,
		0, 0, 0, 51, 53, 5, 3, 0, 0, 52, 51, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54,
		52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0,
		0, 57, 59, 5, 2, 0, 0, 58, 57, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 5, 1,
		0, 0, 0, 60, 61, 6, 3, -1, 0, 61, 63, 3, 8, 4, 0, 62, 64, 5, 6, 0, 0, 63,
		62, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 3, 6, 3,
		9, 66, 99, 1, 0, 0, 0, 67, 69, 5, 8, 0, 0, 68, 70, 5, 3, 0, 0, 69, 68,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 73, 3, 4, 2, 0,
		72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1,
		0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 5, 9, 0, 0, 77, 99, 1, 0, 0, 0, 78,
		80, 5, 10, 0, 0, 79, 81, 3, 6, 3, 0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0,
		0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85,
		5, 11, 0, 0, 85, 99, 1, 0, 0, 0, 86, 88, 3, 4, 2, 0, 87, 89, 5, 6, 0, 0,
		88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 99, 1, 0, 0, 0, 90, 92, 5,
		19, 0, 0, 91, 93, 5, 6, 0, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		94, 1, 0, 0, 0, 94, 99, 3, 6, 3, 2, 95, 96, 3, 4, 2, 0, 96, 97, 5, 21,
		0, 0, 97, 99, 1, 0, 0, 0, 98, 60, 1, 0, 0, 0, 98, 67, 1, 0, 0, 0, 98, 78,
		1, 0, 0, 0, 98, 86, 1, 0, 0, 0, 98, 90, 1, 0, 0, 0, 98, 95, 1, 0, 0, 0,
		99, 114, 1, 0, 0, 0, 100, 101, 10, 7, 0, 0, 101, 102, 5, 13, 0, 0, 102,
		113, 3, 6, 3, 8, 103, 104, 10, 6, 0, 0, 104, 105, 5, 14, 0, 0, 105, 113,
		3, 6, 3, 7, 106, 109, 10, 5, 0, 0, 107, 108, 5, 19, 0, 0, 108, 110, 5,
		6, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0,
		0, 111, 113, 5, 7, 0, 0, 112, 100, 1, 0, 0, 0, 112, 103, 1, 0, 0, 0, 112,
		106, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115,
		1, 0, 0, 0, 115, 7, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 10,
		0, 0, 118, 129, 5, 2, 0, 0, 119, 121, 5, 12, 0, 0, 120, 122, 5, 2, 0, 0,
		121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123,
		119, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127,
		5, 12, 0, 0, 126, 128, 5, 2, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0,
		0, 0, 128, 130, 1, 0, 0, 0, 129, 123, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0,
		130, 9, 1, 0, 0, 0, 28, 13, 16, 19, 22, 25, 29, 32, 36, 39, 42, 46, 49,
		54, 58, 63, 69, 74, 82, 88, 92, 98, 109, 112, 114, 121, 123, 127, 129,
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

// AbcParserInit initializes any static state used to implement AbcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAbcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AbcParserInit() {
	staticData := &abcParserStaticData
	staticData.once.Do(abcParserInit)
}

// NewAbcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAbcParser(input antlr.TokenStream) *AbcParser {
	AbcParserInit()
	this := new(AbcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &abcParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Abc.g4"

	return this
}

// AbcParser tokens.
const (
	AbcParserEOF     = antlr.TokenEOF
	AbcParserPITCH   = 1
	AbcParserNUMBER  = 2
	AbcParserDIVIDER = 3
	AbcParserCONT    = 4
	AbcParserNL      = 5
	AbcParserWS      = 6
	AbcParserSEP     = 7
	AbcParserGS      = 8
	AbcParserGE      = 9
	AbcParserPS      = 10
	AbcParserPE      = 11
	AbcParserCOLON   = 12
	AbcParserLT      = 13
	AbcParserGT      = 14
	AbcParserSHARP   = 15
	AbcParserNATURAL = 16
	AbcParserFLAT    = 17
	AbcParserDOWN    = 18
	AbcParserH       = 19
	AbcParserROLL    = 20
	AbcParserTIE     = 21
)

// AbcParser rules.
const (
	AbcParserRULE_line    = 0
	AbcParserRULE_unit    = 1
	AbcParserRULE_note    = 2
	AbcParserRULE_de      = 3
	AbcParserRULE_triplet = 4
)

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbcParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbcParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) AllUnit() []IUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnitContext); ok {
			len++
		}
	}

	tst := make([]IUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnitContext); ok {
			tst[i] = t.(IUnitContext)
			i++
		}
	}

	return tst
}

func (s *LineContext) Unit(i int) IUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *LineContext) CONT() antlr.TerminalNode {
	return s.GetToken(AbcParserCONT, 0)
}

func (s *LineContext) WS() antlr.TerminalNode {
	return s.GetToken(AbcParserWS, 0)
}

func (s *LineContext) NL() antlr.TerminalNode {
	return s.GetToken(AbcParserNL, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *AbcParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AbcParserRULE_line)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbcParserPITCH)|(1<<AbcParserSEP)|(1<<AbcParserGS)|(1<<AbcParserPS)|(1<<AbcParserSHARP)|(1<<AbcParserNATURAL)|(1<<AbcParserFLAT)|(1<<AbcParserH)|(1<<AbcParserROLL))) != 0) {
		{
			p.SetState(10)
			p.Unit()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbcParserCONT {
		{
			p.SetState(15)
			p.Match(AbcParserCONT)
		}

	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbcParserWS {
		{
			p.SetState(18)
			p.Match(AbcParserWS)
		}

	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbcParserNL {
		{
			p.SetState(21)
			p.Match(AbcParserNL)
		}

	}

	return localctx
}

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbcParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbcParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) AllDe() []IDeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeContext); ok {
			len++
		}
	}

	tst := make([]IDeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeContext); ok {
			tst[i] = t.(IDeContext)
			i++
		}
	}

	return tst
}

func (s *UnitContext) De(i int) IDeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeContext)
}

func (s *UnitContext) AllSEP() []antlr.TerminalNode {
	return s.GetTokens(AbcParserSEP)
}

func (s *UnitContext) SEP(i int) antlr.TerminalNode {
	return s.GetToken(AbcParserSEP, i)
}

func (s *UnitContext) AllH() []antlr.TerminalNode {
	return s.GetTokens(AbcParserH)
}

func (s *UnitContext) H(i int) antlr.TerminalNode {
	return s.GetToken(AbcParserH, i)
}

func (s *UnitContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(AbcParserWS)
}

func (s *UnitContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(AbcParserWS, i)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *AbcParser) Unit() (localctx IUnitContext) {
	this := p
	_ = this

	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AbcParserRULE_unit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(25)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(24)
					p.Match(AbcParserH)
				}

			}
			p.SetState(29)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case AbcParserPITCH, AbcParserGS, AbcParserPS, AbcParserSHARP, AbcParserNATURAL, AbcParserFLAT, AbcParserH, AbcParserROLL:
				{
					p.SetState(27)
					p.de(0)
				}

			case AbcParserSEP:
				{
					p.SetState(28)
					p.Match(AbcParserSEP)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			p.SetState(32)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(31)
					p.Match(AbcParserWS)
				}

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// INoteContext is an interface to support dynamic dispatch.
type INoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoteContext differentiates from other interfaces.
	IsNoteContext()
}

type NoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoteContext() *NoteContext {
	var p = new(NoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbcParserRULE_note
	return p
}

func (*NoteContext) IsNoteContext() {}

func NewNoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteContext {
	var p = new(NoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbcParserRULE_note

	return p
}

func (s *NoteContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteContext) PITCH() antlr.TerminalNode {
	return s.GetToken(AbcParserPITCH, 0)
}

func (s *NoteContext) ROLL() antlr.TerminalNode {
	return s.GetToken(AbcParserROLL, 0)
}

func (s *NoteContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AbcParserDOWN, 0)
}

func (s *NoteContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(AbcParserNUMBER)
}

func (s *NoteContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(AbcParserNUMBER, i)
}

func (s *NoteContext) AllDIVIDER() []antlr.TerminalNode {
	return s.GetTokens(AbcParserDIVIDER)
}

func (s *NoteContext) DIVIDER(i int) antlr.TerminalNode {
	return s.GetToken(AbcParserDIVIDER, i)
}

func (s *NoteContext) SHARP() antlr.TerminalNode {
	return s.GetToken(AbcParserSHARP, 0)
}

func (s *NoteContext) NATURAL() antlr.TerminalNode {
	return s.GetToken(AbcParserNATURAL, 0)
}

func (s *NoteContext) FLAT() antlr.TerminalNode {
	return s.GetToken(AbcParserFLAT, 0)
}

func (s *NoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.EnterNote(s)
	}
}

func (s *NoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.ExitNote(s)
	}
}

func (p *AbcParser) Note() (localctx INoteContext) {
	this := p
	_ = this

	localctx = NewNoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AbcParserRULE_note)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbcParserROLL {
		{
			p.SetState(38)
			p.Match(AbcParserROLL)
		}

	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbcParserSHARP)|(1<<AbcParserNATURAL)|(1<<AbcParserFLAT))) != 0 {
		{
			p.SetState(41)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbcParserSHARP)|(1<<AbcParserNATURAL)|(1<<AbcParserFLAT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(44)
		p.Match(AbcParserPITCH)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(45)
			p.Match(AbcParserDOWN)
		}

	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(48)
			p.Match(AbcParserNUMBER)
		}

	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(51)
				p.Match(AbcParserDIVIDER)
			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(57)
			p.Match(AbcParserNUMBER)
		}

	}

	return localctx
}

// IDeContext is an interface to support dynamic dispatch.
type IDeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeContext differentiates from other interfaces.
	IsDeContext()
}

type DeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeContext() *DeContext {
	var p = new(DeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbcParserRULE_de
	return p
}

func (*DeContext) IsDeContext() {}

func NewDeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeContext {
	var p = new(DeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbcParserRULE_de

	return p
}

func (s *DeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeContext) Triplet() ITripletContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITripletContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITripletContext)
}

func (s *DeContext) AllDe() []IDeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeContext); ok {
			len++
		}
	}

	tst := make([]IDeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeContext); ok {
			tst[i] = t.(IDeContext)
			i++
		}
	}

	return tst
}

func (s *DeContext) De(i int) IDeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeContext)
}

func (s *DeContext) WS() antlr.TerminalNode {
	return s.GetToken(AbcParserWS, 0)
}

func (s *DeContext) GS() antlr.TerminalNode {
	return s.GetToken(AbcParserGS, 0)
}

func (s *DeContext) GE() antlr.TerminalNode {
	return s.GetToken(AbcParserGE, 0)
}

func (s *DeContext) DIVIDER() antlr.TerminalNode {
	return s.GetToken(AbcParserDIVIDER, 0)
}

func (s *DeContext) AllNote() []INoteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INoteContext); ok {
			len++
		}
	}

	tst := make([]INoteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INoteContext); ok {
			tst[i] = t.(INoteContext)
			i++
		}
	}

	return tst
}

func (s *DeContext) Note(i int) INoteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoteContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoteContext)
}

func (s *DeContext) PS() antlr.TerminalNode {
	return s.GetToken(AbcParserPS, 0)
}

func (s *DeContext) PE() antlr.TerminalNode {
	return s.GetToken(AbcParserPE, 0)
}

func (s *DeContext) H() antlr.TerminalNode {
	return s.GetToken(AbcParserH, 0)
}

func (s *DeContext) TIE() antlr.TerminalNode {
	return s.GetToken(AbcParserTIE, 0)
}

func (s *DeContext) LT() antlr.TerminalNode {
	return s.GetToken(AbcParserLT, 0)
}

func (s *DeContext) GT() antlr.TerminalNode {
	return s.GetToken(AbcParserGT, 0)
}

func (s *DeContext) SEP() antlr.TerminalNode {
	return s.GetToken(AbcParserSEP, 0)
}

func (s *DeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.EnterDe(s)
	}
}

func (s *DeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.ExitDe(s)
	}
}

func (p *AbcParser) De() (localctx IDeContext) {
	return p.de(0)
}

func (p *AbcParser) de(_p int) (localctx IDeContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, AbcParserRULE_de, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(61)
			p.Triplet()
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbcParserWS {
			{
				p.SetState(62)
				p.Match(AbcParserWS)
			}

		}
		{
			p.SetState(65)
			p.de(9)
		}

	case 2:
		{
			p.SetState(67)
			p.Match(AbcParserGS)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbcParserDIVIDER {
			{
				p.SetState(68)
				p.Match(AbcParserDIVIDER)
			}

		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbcParserPITCH)|(1<<AbcParserSHARP)|(1<<AbcParserNATURAL)|(1<<AbcParserFLAT)|(1<<AbcParserROLL))) != 0) {
			{
				p.SetState(71)
				p.Note()
			}

			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(76)
			p.Match(AbcParserGE)
		}

	case 3:
		{
			p.SetState(78)
			p.Match(AbcParserPS)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbcParserPITCH)|(1<<AbcParserGS)|(1<<AbcParserPS)|(1<<AbcParserSHARP)|(1<<AbcParserNATURAL)|(1<<AbcParserFLAT)|(1<<AbcParserH)|(1<<AbcParserROLL))) != 0) {
			{
				p.SetState(79)
				p.de(0)
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(84)
			p.Match(AbcParserPE)
		}

	case 4:
		{
			p.SetState(86)
			p.Note()
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(87)
				p.Match(AbcParserWS)
			}

		}

	case 5:
		{
			p.SetState(90)
			p.Match(AbcParserH)
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbcParserWS {
			{
				p.SetState(91)
				p.Match(AbcParserWS)
			}

		}
		{
			p.SetState(94)
			p.de(2)
		}

	case 6:
		{
			p.SetState(95)
			p.Note()
		}
		{
			p.SetState(96)
			p.Match(AbcParserTIE)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDeContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbcParserRULE_de)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(101)
					p.Match(AbcParserLT)
				}
				{
					p.SetState(102)
					p.de(8)
				}

			case 2:
				localctx = NewDeContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbcParserRULE_de)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(104)
					p.Match(AbcParserGT)
				}
				{
					p.SetState(105)
					p.de(7)
				}

			case 3:
				localctx = NewDeContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbcParserRULE_de)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(109)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == AbcParserH {
					{
						p.SetState(107)
						p.Match(AbcParserH)
					}
					{
						p.SetState(108)
						p.Match(AbcParserWS)
					}

				}
				{
					p.SetState(111)
					p.Match(AbcParserSEP)
				}

			}

		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// ITripletContext is an interface to support dynamic dispatch.
type ITripletContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTripletContext differentiates from other interfaces.
	IsTripletContext()
}

type TripletContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTripletContext() *TripletContext {
	var p = new(TripletContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbcParserRULE_triplet
	return p
}

func (*TripletContext) IsTripletContext() {}

func NewTripletContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TripletContext {
	var p = new(TripletContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbcParserRULE_triplet

	return p
}

func (s *TripletContext) GetParser() antlr.Parser { return s.parser }

func (s *TripletContext) PS() antlr.TerminalNode {
	return s.GetToken(AbcParserPS, 0)
}

func (s *TripletContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(AbcParserNUMBER)
}

func (s *TripletContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(AbcParserNUMBER, i)
}

func (s *TripletContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(AbcParserCOLON)
}

func (s *TripletContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(AbcParserCOLON, i)
}

func (s *TripletContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TripletContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TripletContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.EnterTriplet(s)
	}
}

func (s *TripletContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbcListener); ok {
		listenerT.ExitTriplet(s)
	}
}

func (p *AbcParser) Triplet() (localctx ITripletContext) {
	this := p
	_ = this

	localctx = NewTripletContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AbcParserRULE_triplet)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(AbcParserPS)
	}
	{
		p.SetState(118)
		p.Match(AbcParserNUMBER)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbcParserCOLON {
		p.SetState(123)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(119)
				p.Match(AbcParserCOLON)
			}
			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AbcParserNUMBER {
				{
					p.SetState(120)
					p.Match(AbcParserNUMBER)
				}

			}

		}
		{
			p.SetState(125)
			p.Match(AbcParserCOLON)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbcParserNUMBER {
			{
				p.SetState(126)
				p.Match(AbcParserNUMBER)
			}

		}

	}

	return localctx
}

func (p *AbcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *DeContext = nil
		if localctx != nil {
			t = localctx.(*DeContext)
		}
		return p.De_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AbcParser) De_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
