// Code generated from AOC.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AOC

import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type AOCParser struct {
	*antlr.BaseParser
}

var AOCParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func aocParserInit() {
  staticData := &AOCParserStaticData
  staticData.LiteralNames = []string{
    "", "'mul'", "'('", "','", "')'", "'do()'", "'don't()'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "NUM", "INVALID",
  }
  staticData.RuleNames = []string{
    "program", "instruction", "mul", "do", "don_t",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 8, 36, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 
	1, 0, 5, 0, 12, 8, 0, 10, 0, 12, 0, 15, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 
	1, 1, 1, 3, 1, 23, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 
	1, 3, 1, 4, 1, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 0, 34, 0, 13, 1, 0, 
	0, 0, 2, 22, 1, 0, 0, 0, 4, 24, 1, 0, 0, 0, 6, 31, 1, 0, 0, 0, 8, 33, 1, 
	0, 0, 0, 10, 12, 3, 2, 1, 0, 11, 10, 1, 0, 0, 0, 12, 15, 1, 0, 0, 0, 13, 
	11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 16, 1, 0, 0, 0, 15, 13, 1, 0, 0, 
	0, 16, 17, 5, 0, 0, 1, 17, 1, 1, 0, 0, 0, 18, 23, 3, 4, 2, 0, 19, 23, 3, 
	6, 3, 0, 20, 23, 3, 8, 4, 0, 21, 23, 5, 8, 0, 0, 22, 18, 1, 0, 0, 0, 22, 
	19, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 21, 1, 0, 0, 0, 23, 3, 1, 0, 0, 
	0, 24, 25, 5, 1, 0, 0, 25, 26, 5, 2, 0, 0, 26, 27, 5, 7, 0, 0, 27, 28, 
	5, 3, 0, 0, 28, 29, 5, 7, 0, 0, 29, 30, 5, 4, 0, 0, 30, 5, 1, 0, 0, 0, 
	31, 32, 5, 5, 0, 0, 32, 7, 1, 0, 0, 0, 33, 34, 5, 6, 0, 0, 34, 9, 1, 0, 
	0, 0, 2, 13, 22,
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

// AOCParserInit initializes any static state used to implement AOCParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAOCParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AOCParserInit() {
  staticData := &AOCParserStaticData
  staticData.once.Do(aocParserInit)
}

// NewAOCParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAOCParser(input antlr.TokenStream) *AOCParser {
	AOCParserInit()
	this := new(AOCParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &AOCParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AOC.g4"

	return this
}


// AOCParser tokens.
const (
	AOCParserEOF = antlr.TokenEOF
	AOCParserT__0 = 1
	AOCParserT__1 = 2
	AOCParserT__2 = 3
	AOCParserT__3 = 4
	AOCParserT__4 = 5
	AOCParserT__5 = 6
	AOCParserNUM = 7
	AOCParserINVALID = 8
)

// AOCParser rules.
const (
	AOCParserRULE_program = 0
	AOCParserRULE_instruction = 1
	AOCParserRULE_mul = 2
	AOCParserRULE_do = 3
	AOCParserRULE_don_t = 4
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(AOCParserEOF, 0)
}

func (s *ProgramContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitProgram(s)
	}
}




func (p *AOCParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AOCParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 354) != 0) {
		{
			p.SetState(10)
			p.Instruction()
		}


		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(16)
		p.Match(AOCParserEOF)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mul() IMulContext
	Do() IDoContext
	Don_t() IDon_tContext
	INVALID() antlr.TerminalNode

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Mul() IMulContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulContext)
}

func (s *InstructionContext) Do() IDoContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoContext)
}

func (s *InstructionContext) Don_t() IDon_tContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDon_tContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDon_tContext)
}

func (s *InstructionContext) INVALID() antlr.TerminalNode {
	return s.GetToken(AOCParserINVALID, 0)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitInstruction(s)
	}
}




func (p *AOCParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AOCParserRULE_instruction)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AOCParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Mul()
		}


	case AOCParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Do()
		}


	case AOCParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(20)
			p.Don_t()
		}


	case AOCParserINVALID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(21)
			p.Match(AOCParserINVALID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMulContext is an interface to support dynamic dispatch.
type IMulContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNUM() []antlr.TerminalNode
	NUM(i int) antlr.TerminalNode

	// IsMulContext differentiates from other interfaces.
	IsMulContext()
}

type MulContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulContext() *MulContext {
	var p = new(MulContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_mul
	return p
}

func InitEmptyMulContext(p *MulContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_mul
}

func (*MulContext) IsMulContext() {}

func NewMulContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulContext {
	var p = new(MulContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_mul

	return p
}

func (s *MulContext) GetParser() antlr.Parser { return s.parser }

func (s *MulContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(AOCParserNUM)
}

func (s *MulContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(AOCParserNUM, i)
}

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterMul(s)
	}
}

func (s *MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitMul(s)
	}
}




func (p *AOCParser) Mul() (localctx IMulContext) {
	localctx = NewMulContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AOCParserRULE_mul)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(AOCParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(25)
		p.Match(AOCParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(AOCParserNUM)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(AOCParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(AOCParserNUM)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Match(AOCParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDoContext is an interface to support dynamic dispatch.
type IDoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDoContext differentiates from other interfaces.
	IsDoContext()
}

type DoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoContext() *DoContext {
	var p = new(DoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_do
	return p
}

func InitEmptyDoContext(p *DoContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_do
}

func (*DoContext) IsDoContext() {}

func NewDoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoContext {
	var p = new(DoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_do

	return p
}

func (s *DoContext) GetParser() antlr.Parser { return s.parser }
func (s *DoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterDo(s)
	}
}

func (s *DoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitDo(s)
	}
}




func (p *AOCParser) Do() (localctx IDoContext) {
	localctx = NewDoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AOCParserRULE_do)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDon_tContext is an interface to support dynamic dispatch.
type IDon_tContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDon_tContext differentiates from other interfaces.
	IsDon_tContext()
}

type Don_tContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDon_tContext() *Don_tContext {
	var p = new(Don_tContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_don_t
	return p
}

func InitEmptyDon_tContext(p *Don_tContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_don_t
}

func (*Don_tContext) IsDon_tContext() {}

func NewDon_tContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Don_tContext {
	var p = new(Don_tContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_don_t

	return p
}

func (s *Don_tContext) GetParser() antlr.Parser { return s.parser }
func (s *Don_tContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Don_tContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Don_tContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterDon_t(s)
	}
}

func (s *Don_tContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitDon_t(s)
	}
}




func (p *AOCParser) Don_t() (localctx IDon_tContext) {
	localctx = NewDon_tContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AOCParserRULE_don_t)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(AOCParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


