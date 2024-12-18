// Code generated from AOC.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type AOCLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var AOCLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func aoclexerLexerInit() {
  staticData := &AOCLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "'mul'", "'('", "','", "')'", "'do()'", "'don't()'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "NUM", "INVALID",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "NUM", "INVALID",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 8, 49, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1, 
	1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 
	5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 43, 8, 6, 10, 
	6, 12, 6, 46, 9, 6, 1, 7, 1, 7, 0, 0, 8, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 
	11, 6, 13, 7, 15, 8, 1, 0, 1, 1, 0, 48, 57, 49, 0, 1, 1, 0, 0, 0, 0, 3, 
	1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 
	1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 
	21, 1, 0, 0, 0, 5, 23, 1, 0, 0, 0, 7, 25, 1, 0, 0, 0, 9, 27, 1, 0, 0, 0, 
	11, 32, 1, 0, 0, 0, 13, 40, 1, 0, 0, 0, 15, 47, 1, 0, 0, 0, 17, 18, 5, 
	109, 0, 0, 18, 19, 5, 117, 0, 0, 19, 20, 5, 108, 0, 0, 20, 2, 1, 0, 0, 
	0, 21, 22, 5, 40, 0, 0, 22, 4, 1, 0, 0, 0, 23, 24, 5, 44, 0, 0, 24, 6, 
	1, 0, 0, 0, 25, 26, 5, 41, 0, 0, 26, 8, 1, 0, 0, 0, 27, 28, 5, 100, 0, 
	0, 28, 29, 5, 111, 0, 0, 29, 30, 5, 40, 0, 0, 30, 31, 5, 41, 0, 0, 31, 
	10, 1, 0, 0, 0, 32, 33, 5, 100, 0, 0, 33, 34, 5, 111, 0, 0, 34, 35, 5, 
	110, 0, 0, 35, 36, 5, 39, 0, 0, 36, 37, 5, 116, 0, 0, 37, 38, 5, 40, 0, 
	0, 38, 39, 5, 41, 0, 0, 39, 12, 1, 0, 0, 0, 40, 44, 7, 0, 0, 0, 41, 43, 
	7, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 
	44, 45, 1, 0, 0, 0, 45, 14, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 9, 
	0, 0, 0, 48, 16, 1, 0, 0, 0, 2, 0, 44, 0,
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

// AOCLexerInit initializes any static state used to implement AOCLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAOCLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AOCLexerInit() {
  staticData := &AOCLexerLexerStaticData
  staticData.once.Do(aoclexerLexerInit)
}

// NewAOCLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAOCLexer(input antlr.CharStream) *AOCLexer {
  AOCLexerInit()
	l := new(AOCLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &AOCLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AOC.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AOCLexer tokens.
const (
	AOCLexerT__0 = 1
	AOCLexerT__1 = 2
	AOCLexerT__2 = 3
	AOCLexerT__3 = 4
	AOCLexerT__4 = 5
	AOCLexerT__5 = 6
	AOCLexerNUM = 7
	AOCLexerINVALID = 8
)

