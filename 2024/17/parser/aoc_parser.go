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
    "", "'\\n'", "'Register '", "': '", "'Program: '", "','", "'0'", "'1'", 
    "'2'", "'3'", "'4'", "'5'", "'6'", "'7'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "INTEGER", 
    "DIGIT",
  }
  staticData.RuleNames = []string{
    "program", "registers", "register", "instructions", "instruction", "adv", 
    "bxl", "bst", "jnz", "bxc", "out", "bdv", "cdv", "operand", "literal_value", 
    "reg_a", "reg_b", "reg_c",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 16, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 
	1, 44, 8, 1, 10, 1, 12, 1, 47, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 
	1, 3, 1, 3, 1, 3, 5, 3, 58, 8, 3, 10, 3, 12, 3, 61, 9, 3, 1, 4, 1, 4, 1, 
	4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 71, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 
	1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 
	1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 109, 
	8, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 0, 
	0, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 
	0, 1, 1, 0, 6, 9, 112, 0, 36, 1, 0, 0, 0, 2, 45, 1, 0, 0, 0, 4, 48, 1, 
	0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 72, 1, 0, 0, 0, 12, 
	76, 1, 0, 0, 0, 14, 80, 1, 0, 0, 0, 16, 84, 1, 0, 0, 0, 18, 88, 1, 0, 0, 
	0, 20, 92, 1, 0, 0, 0, 22, 96, 1, 0, 0, 0, 24, 100, 1, 0, 0, 0, 26, 108, 
	1, 0, 0, 0, 28, 110, 1, 0, 0, 0, 30, 112, 1, 0, 0, 0, 32, 114, 1, 0, 0, 
	0, 34, 116, 1, 0, 0, 0, 36, 37, 3, 2, 1, 0, 37, 38, 5, 1, 0, 0, 38, 39, 
	3, 6, 3, 0, 39, 1, 1, 0, 0, 0, 40, 41, 3, 4, 2, 0, 41, 42, 5, 1, 0, 0, 
	42, 44, 1, 0, 0, 0, 43, 40, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 
	0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 
	49, 5, 2, 0, 0, 49, 50, 5, 14, 0, 0, 50, 51, 5, 3, 0, 0, 51, 52, 5, 15, 
	0, 0, 52, 5, 1, 0, 0, 0, 53, 54, 5, 4, 0, 0, 54, 59, 3, 8, 4, 0, 55, 56, 
	5, 5, 0, 0, 56, 58, 3, 8, 4, 0, 57, 55, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 
	59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 7, 1, 0, 0, 0, 61, 59, 1, 0, 
	0, 0, 62, 71, 3, 10, 5, 0, 63, 71, 3, 12, 6, 0, 64, 71, 3, 14, 7, 0, 65, 
	71, 3, 16, 8, 0, 66, 71, 3, 18, 9, 0, 67, 71, 3, 20, 10, 0, 68, 71, 3, 
	22, 11, 0, 69, 71, 3, 24, 12, 0, 70, 62, 1, 0, 0, 0, 70, 63, 1, 0, 0, 0, 
	70, 64, 1, 0, 0, 0, 70, 65, 1, 0, 0, 0, 70, 66, 1, 0, 0, 0, 70, 67, 1, 
	0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 9, 1, 0, 0, 0, 72, 
	73, 5, 6, 0, 0, 73, 74, 5, 5, 0, 0, 74, 75, 3, 26, 13, 0, 75, 11, 1, 0, 
	0, 0, 76, 77, 5, 7, 0, 0, 77, 78, 5, 5, 0, 0, 78, 79, 3, 26, 13, 0, 79, 
	13, 1, 0, 0, 0, 80, 81, 5, 8, 0, 0, 81, 82, 5, 5, 0, 0, 82, 83, 3, 26, 
	13, 0, 83, 15, 1, 0, 0, 0, 84, 85, 5, 9, 0, 0, 85, 86, 5, 5, 0, 0, 86, 
	87, 3, 26, 13, 0, 87, 17, 1, 0, 0, 0, 88, 89, 5, 10, 0, 0, 89, 90, 5, 5, 
	0, 0, 90, 91, 3, 26, 13, 0, 91, 19, 1, 0, 0, 0, 92, 93, 5, 11, 0, 0, 93, 
	94, 5, 5, 0, 0, 94, 95, 3, 26, 13, 0, 95, 21, 1, 0, 0, 0, 96, 97, 5, 12, 
	0, 0, 97, 98, 5, 5, 0, 0, 98, 99, 3, 26, 13, 0, 99, 23, 1, 0, 0, 0, 100, 
	101, 5, 13, 0, 0, 101, 102, 5, 5, 0, 0, 102, 103, 3, 26, 13, 0, 103, 25, 
	1, 0, 0, 0, 104, 109, 3, 28, 14, 0, 105, 109, 3, 30, 15, 0, 106, 109, 3, 
	32, 16, 0, 107, 109, 3, 34, 17, 0, 108, 104, 1, 0, 0, 0, 108, 105, 1, 0, 
	0, 0, 108, 106, 1, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 27, 1, 0, 0, 0, 
	110, 111, 7, 0, 0, 0, 111, 29, 1, 0, 0, 0, 112, 113, 5, 10, 0, 0, 113, 
	31, 1, 0, 0, 0, 114, 115, 5, 11, 0, 0, 115, 33, 1, 0, 0, 0, 116, 117, 5, 
	12, 0, 0, 117, 35, 1, 0, 0, 0, 4, 45, 59, 70, 108,
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
	AOCParserT__6 = 7
	AOCParserT__7 = 8
	AOCParserT__8 = 9
	AOCParserT__9 = 10
	AOCParserT__10 = 11
	AOCParserT__11 = 12
	AOCParserT__12 = 13
	AOCParserID = 14
	AOCParserINTEGER = 15
	AOCParserDIGIT = 16
)

// AOCParser rules.
const (
	AOCParserRULE_program = 0
	AOCParserRULE_registers = 1
	AOCParserRULE_register = 2
	AOCParserRULE_instructions = 3
	AOCParserRULE_instruction = 4
	AOCParserRULE_adv = 5
	AOCParserRULE_bxl = 6
	AOCParserRULE_bst = 7
	AOCParserRULE_jnz = 8
	AOCParserRULE_bxc = 9
	AOCParserRULE_out = 10
	AOCParserRULE_bdv = 11
	AOCParserRULE_cdv = 12
	AOCParserRULE_operand = 13
	AOCParserRULE_literal_value = 14
	AOCParserRULE_reg_a = 15
	AOCParserRULE_reg_b = 16
	AOCParserRULE_reg_c = 17
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Registers() IRegistersContext
	Instructions() IInstructionsContext

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

func (s *ProgramContext) Registers() IRegistersContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegistersContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegistersContext)
}

func (s *ProgramContext) Instructions() IInstructionsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Registers()
	}
	{
		p.SetState(37)
		p.Match(AOCParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Instructions()
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


// IRegistersContext is an interface to support dynamic dispatch.
type IRegistersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRegister() []IRegisterContext
	Register(i int) IRegisterContext

	// IsRegistersContext differentiates from other interfaces.
	IsRegistersContext()
}

type RegistersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegistersContext() *RegistersContext {
	var p = new(RegistersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_registers
	return p
}

func InitEmptyRegistersContext(p *RegistersContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_registers
}

func (*RegistersContext) IsRegistersContext() {}

func NewRegistersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegistersContext {
	var p = new(RegistersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_registers

	return p
}

func (s *RegistersContext) GetParser() antlr.Parser { return s.parser }

func (s *RegistersContext) AllRegister() []IRegisterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterContext); ok {
			len++
		}
	}

	tst := make([]IRegisterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterContext); ok {
			tst[i] = t.(IRegisterContext)
			i++
		}
	}

	return tst
}

func (s *RegistersContext) Register(i int) IRegisterContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
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

	return t.(IRegisterContext)
}

func (s *RegistersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegistersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RegistersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterRegisters(s)
	}
}

func (s *RegistersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitRegisters(s)
	}
}




func (p *AOCParser) Registers() (localctx IRegistersContext) {
	localctx = NewRegistersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AOCParserRULE_registers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == AOCParserT__1 {
		{
			p.SetState(40)
			p.Register()
		}
		{
			p.SetState(41)
			p.Match(AOCParserT__0)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
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


// IRegisterContext is an interface to support dynamic dispatch.
type IRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	INTEGER() antlr.TerminalNode

	// IsRegisterContext differentiates from other interfaces.
	IsRegisterContext()
}

type RegisterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegisterContext() *RegisterContext {
	var p = new(RegisterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_register
	return p
}

func InitEmptyRegisterContext(p *RegisterContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_register
}

func (*RegisterContext) IsRegisterContext() {}

func NewRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterContext {
	var p = new(RegisterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_register

	return p
}

func (s *RegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterContext) ID() antlr.TerminalNode {
	return s.GetToken(AOCParserID, 0)
}

func (s *RegisterContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(AOCParserINTEGER, 0)
}

func (s *RegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterRegister(s)
	}
}

func (s *RegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitRegister(s)
	}
}




func (p *AOCParser) Register() (localctx IRegisterContext) {
	localctx = NewRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AOCParserRULE_register)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(AOCParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(AOCParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(AOCParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(AOCParserINTEGER)
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


// IInstructionsContext is an interface to support dynamic dispatch.
type IInstructionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsInstructionsContext differentiates from other interfaces.
	IsInstructionsContext()
}

type InstructionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionsContext() *InstructionsContext {
	var p = new(InstructionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_instructions
	return p
}

func InitEmptyInstructionsContext(p *InstructionsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_instructions
}

func (*InstructionsContext) IsInstructionsContext() {}

func NewInstructionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsContext {
	var p = new(InstructionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_instructions

	return p
}

func (s *InstructionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionsContext) AllInstruction() []IInstructionContext {
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

func (s *InstructionsContext) Instruction(i int) IInstructionContext {
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

func (s *InstructionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InstructionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterInstructions(s)
	}
}

func (s *InstructionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitInstructions(s)
	}
}




func (p *AOCParser) Instructions() (localctx IInstructionsContext) {
	localctx = NewInstructionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AOCParserRULE_instructions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(AOCParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Instruction()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == AOCParserT__4 {
		{
			p.SetState(55)
			p.Match(AOCParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Instruction()
		}


		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
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
	Adv() IAdvContext
	Bxl() IBxlContext
	Bst() IBstContext
	Jnz() IJnzContext
	Bxc() IBxcContext
	Out() IOutContext
	Bdv() IBdvContext
	Cdv() ICdvContext

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

func (s *InstructionContext) Adv() IAdvContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdvContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdvContext)
}

func (s *InstructionContext) Bxl() IBxlContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBxlContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBxlContext)
}

func (s *InstructionContext) Bst() IBstContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBstContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBstContext)
}

func (s *InstructionContext) Jnz() IJnzContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJnzContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJnzContext)
}

func (s *InstructionContext) Bxc() IBxcContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBxcContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBxcContext)
}

func (s *InstructionContext) Out() IOutContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutContext)
}

func (s *InstructionContext) Bdv() IBdvContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBdvContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBdvContext)
}

func (s *InstructionContext) Cdv() ICdvContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICdvContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICdvContext)
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
	p.EnterRule(localctx, 8, AOCParserRULE_instruction)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AOCParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Adv()
		}


	case AOCParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Bxl()
		}


	case AOCParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Bst()
		}


	case AOCParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Jnz()
		}


	case AOCParserT__9:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(66)
			p.Bxc()
		}


	case AOCParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(67)
			p.Out()
		}


	case AOCParserT__11:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(68)
			p.Bdv()
		}


	case AOCParserT__12:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(69)
			p.Cdv()
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


// IAdvContext is an interface to support dynamic dispatch.
type IAdvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext

	// IsAdvContext differentiates from other interfaces.
	IsAdvContext()
}

type AdvContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdvContext() *AdvContext {
	var p = new(AdvContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_adv
	return p
}

func InitEmptyAdvContext(p *AdvContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_adv
}

func (*AdvContext) IsAdvContext() {}

func NewAdvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdvContext {
	var p = new(AdvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_adv

	return p
}

func (s *AdvContext) GetParser() antlr.Parser { return s.parser }

func (s *AdvContext) Operand() IOperandContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *AdvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AdvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterAdv(s)
	}
}

func (s *AdvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitAdv(s)
	}
}




func (p *AOCParser) Adv() (localctx IAdvContext) {
	localctx = NewAdvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AOCParserRULE_adv)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(AOCParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Operand()
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


// IBxlContext is an interface to support dynamic dispatch.
type IBxlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext

	// IsBxlContext differentiates from other interfaces.
	IsBxlContext()
}

type BxlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBxlContext() *BxlContext {
	var p = new(BxlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_bxl
	return p
}

func InitEmptyBxlContext(p *BxlContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_bxl
}

func (*BxlContext) IsBxlContext() {}

func NewBxlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BxlContext {
	var p = new(BxlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_bxl

	return p
}

func (s *BxlContext) GetParser() antlr.Parser { return s.parser }

func (s *BxlContext) Operand() IOperandContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *BxlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BxlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BxlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterBxl(s)
	}
}

func (s *BxlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitBxl(s)
	}
}




func (p *AOCParser) Bxl() (localctx IBxlContext) {
	localctx = NewBxlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AOCParserRULE_bxl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(AOCParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Operand()
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


// IBstContext is an interface to support dynamic dispatch.
type IBstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext

	// IsBstContext differentiates from other interfaces.
	IsBstContext()
}

type BstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBstContext() *BstContext {
	var p = new(BstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_bst
	return p
}

func InitEmptyBstContext(p *BstContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_bst
}

func (*BstContext) IsBstContext() {}

func NewBstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BstContext {
	var p = new(BstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_bst

	return p
}

func (s *BstContext) GetParser() antlr.Parser { return s.parser }

func (s *BstContext) Operand() IOperandContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *BstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterBst(s)
	}
}

func (s *BstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitBst(s)
	}
}




func (p *AOCParser) Bst() (localctx IBstContext) {
	localctx = NewBstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AOCParserRULE_bst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(AOCParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Operand()
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


// IJnzContext is an interface to support dynamic dispatch.
type IJnzContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext

	// IsJnzContext differentiates from other interfaces.
	IsJnzContext()
}

type JnzContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJnzContext() *JnzContext {
	var p = new(JnzContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_jnz
	return p
}

func InitEmptyJnzContext(p *JnzContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_jnz
}

func (*JnzContext) IsJnzContext() {}

func NewJnzContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JnzContext {
	var p = new(JnzContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_jnz

	return p
}

func (s *JnzContext) GetParser() antlr.Parser { return s.parser }

func (s *JnzContext) Operand() IOperandContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *JnzContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JnzContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *JnzContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterJnz(s)
	}
}

func (s *JnzContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitJnz(s)
	}
}




func (p *AOCParser) Jnz() (localctx IJnzContext) {
	localctx = NewJnzContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AOCParserRULE_jnz)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(AOCParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Operand()
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


// IBxcContext is an interface to support dynamic dispatch.
type IBxcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext

	// IsBxcContext differentiates from other interfaces.
	IsBxcContext()
}

type BxcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBxcContext() *BxcContext {
	var p = new(BxcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_bxc
	return p
}

func InitEmptyBxcContext(p *BxcContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_bxc
}

func (*BxcContext) IsBxcContext() {}

func NewBxcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BxcContext {
	var p = new(BxcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_bxc

	return p
}

func (s *BxcContext) GetParser() antlr.Parser { return s.parser }

func (s *BxcContext) Operand() IOperandContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *BxcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BxcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BxcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterBxc(s)
	}
}

func (s *BxcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitBxc(s)
	}
}




func (p *AOCParser) Bxc() (localctx IBxcContext) {
	localctx = NewBxcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AOCParserRULE_bxc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(AOCParserT__9)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Operand()
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


// IOutContext is an interface to support dynamic dispatch.
type IOutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext

	// IsOutContext differentiates from other interfaces.
	IsOutContext()
}

type OutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutContext() *OutContext {
	var p = new(OutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_out
	return p
}

func InitEmptyOutContext(p *OutContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_out
}

func (*OutContext) IsOutContext() {}

func NewOutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutContext {
	var p = new(OutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_out

	return p
}

func (s *OutContext) GetParser() antlr.Parser { return s.parser }

func (s *OutContext) Operand() IOperandContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *OutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterOut(s)
	}
}

func (s *OutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitOut(s)
	}
}




func (p *AOCParser) Out() (localctx IOutContext) {
	localctx = NewOutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AOCParserRULE_out)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(AOCParserT__10)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Operand()
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


// IBdvContext is an interface to support dynamic dispatch.
type IBdvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext

	// IsBdvContext differentiates from other interfaces.
	IsBdvContext()
}

type BdvContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBdvContext() *BdvContext {
	var p = new(BdvContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_bdv
	return p
}

func InitEmptyBdvContext(p *BdvContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_bdv
}

func (*BdvContext) IsBdvContext() {}

func NewBdvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BdvContext {
	var p = new(BdvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_bdv

	return p
}

func (s *BdvContext) GetParser() antlr.Parser { return s.parser }

func (s *BdvContext) Operand() IOperandContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *BdvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BdvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BdvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterBdv(s)
	}
}

func (s *BdvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitBdv(s)
	}
}




func (p *AOCParser) Bdv() (localctx IBdvContext) {
	localctx = NewBdvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AOCParserRULE_bdv)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(AOCParserT__11)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Operand()
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


// ICdvContext is an interface to support dynamic dispatch.
type ICdvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext

	// IsCdvContext differentiates from other interfaces.
	IsCdvContext()
}

type CdvContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCdvContext() *CdvContext {
	var p = new(CdvContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_cdv
	return p
}

func InitEmptyCdvContext(p *CdvContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_cdv
}

func (*CdvContext) IsCdvContext() {}

func NewCdvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CdvContext {
	var p = new(CdvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_cdv

	return p
}

func (s *CdvContext) GetParser() antlr.Parser { return s.parser }

func (s *CdvContext) Operand() IOperandContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *CdvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CdvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CdvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterCdv(s)
	}
}

func (s *CdvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitCdv(s)
	}
}




func (p *AOCParser) Cdv() (localctx ICdvContext) {
	localctx = NewCdvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AOCParserRULE_cdv)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(AOCParserT__12)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(AOCParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Operand()
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


// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal_value() ILiteral_valueContext
	Reg_a() IReg_aContext
	Reg_b() IReg_bContext
	Reg_c() IReg_cContext

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *OperandContext) Reg_a() IReg_aContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg_aContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg_aContext)
}

func (s *OperandContext) Reg_b() IReg_bContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg_bContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg_bContext)
}

func (s *OperandContext) Reg_c() IReg_cContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReg_cContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReg_cContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitOperand(s)
	}
}




func (p *AOCParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AOCParserRULE_operand)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AOCParserT__5, AOCParserT__6, AOCParserT__7, AOCParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Literal_value()
		}


	case AOCParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Reg_a()
		}


	case AOCParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.Reg_b()
		}


	case AOCParserT__11:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(107)
			p.Reg_c()
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


// ILiteral_valueContext is an interface to support dynamic dispatch.
type ILiteral_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteral_valueContext differentiates from other interfaces.
	IsLiteral_valueContext()
}

type Literal_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_valueContext() *Literal_valueContext {
	var p = new(Literal_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_literal_value
	return p
}

func InitEmptyLiteral_valueContext(p *Literal_valueContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_literal_value
}

func (*Literal_valueContext) IsLiteral_valueContext() {}

func NewLiteral_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_valueContext {
	var p = new(Literal_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_literal_value

	return p
}

func (s *Literal_valueContext) GetParser() antlr.Parser { return s.parser }
func (s *Literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Literal_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterLiteral_value(s)
	}
}

func (s *Literal_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitLiteral_value(s)
	}
}




func (p *AOCParser) Literal_value() (localctx ILiteral_valueContext) {
	localctx = NewLiteral_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AOCParserRULE_literal_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 960) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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


// IReg_aContext is an interface to support dynamic dispatch.
type IReg_aContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReg_aContext differentiates from other interfaces.
	IsReg_aContext()
}

type Reg_aContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReg_aContext() *Reg_aContext {
	var p = new(Reg_aContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_reg_a
	return p
}

func InitEmptyReg_aContext(p *Reg_aContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_reg_a
}

func (*Reg_aContext) IsReg_aContext() {}

func NewReg_aContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reg_aContext {
	var p = new(Reg_aContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_reg_a

	return p
}

func (s *Reg_aContext) GetParser() antlr.Parser { return s.parser }
func (s *Reg_aContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reg_aContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Reg_aContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterReg_a(s)
	}
}

func (s *Reg_aContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitReg_a(s)
	}
}




func (p *AOCParser) Reg_a() (localctx IReg_aContext) {
	localctx = NewReg_aContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AOCParserRULE_reg_a)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(AOCParserT__9)
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


// IReg_bContext is an interface to support dynamic dispatch.
type IReg_bContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReg_bContext differentiates from other interfaces.
	IsReg_bContext()
}

type Reg_bContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReg_bContext() *Reg_bContext {
	var p = new(Reg_bContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_reg_b
	return p
}

func InitEmptyReg_bContext(p *Reg_bContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_reg_b
}

func (*Reg_bContext) IsReg_bContext() {}

func NewReg_bContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reg_bContext {
	var p = new(Reg_bContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_reg_b

	return p
}

func (s *Reg_bContext) GetParser() antlr.Parser { return s.parser }
func (s *Reg_bContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reg_bContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Reg_bContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterReg_b(s)
	}
}

func (s *Reg_bContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitReg_b(s)
	}
}




func (p *AOCParser) Reg_b() (localctx IReg_bContext) {
	localctx = NewReg_bContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AOCParserRULE_reg_b)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(AOCParserT__10)
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


// IReg_cContext is an interface to support dynamic dispatch.
type IReg_cContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReg_cContext differentiates from other interfaces.
	IsReg_cContext()
}

type Reg_cContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReg_cContext() *Reg_cContext {
	var p = new(Reg_cContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_reg_c
	return p
}

func InitEmptyReg_cContext(p *Reg_cContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AOCParserRULE_reg_c
}

func (*Reg_cContext) IsReg_cContext() {}

func NewReg_cContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reg_cContext {
	var p = new(Reg_cContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AOCParserRULE_reg_c

	return p
}

func (s *Reg_cContext) GetParser() antlr.Parser { return s.parser }
func (s *Reg_cContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reg_cContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Reg_cContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.EnterReg_c(s)
	}
}

func (s *Reg_cContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AOCListener); ok {
		listenerT.ExitReg_c(s)
	}
}




func (p *AOCParser) Reg_c() (localctx IReg_cContext) {
	localctx = NewReg_cContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AOCParserRULE_reg_c)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(AOCParserT__11)
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


