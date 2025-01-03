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
    "", "'\\n'", "'Register '", "': '", "'Program: '", "','", "'0'", "'1'", 
    "'2'", "'3'", "'4'", "'5'", "'6'", "'7'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "INTEGER", 
    "DIGIT",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "ID", "INTEGER", "INT", "DIGIT",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 16, 94, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 
	3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 
	7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 
	13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 86, 8, 15, 10, 15, 
	12, 15, 89, 9, 15, 3, 15, 91, 8, 15, 1, 16, 1, 16, 0, 0, 17, 1, 1, 3, 2, 
	5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 
	13, 27, 14, 29, 15, 31, 0, 33, 16, 1, 0, 2, 1, 0, 49, 57, 1, 0, 48, 57, 
	94, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 
	0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 
	0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 
	0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 33, 
	1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 37, 1, 0, 0, 0, 5, 47, 1, 0, 0, 0, 7, 
	50, 1, 0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 62, 1, 0, 0, 0, 13, 64, 1, 0, 0, 
	0, 15, 66, 1, 0, 0, 0, 17, 68, 1, 0, 0, 0, 19, 70, 1, 0, 0, 0, 21, 72, 
	1, 0, 0, 0, 23, 74, 1, 0, 0, 0, 25, 76, 1, 0, 0, 0, 27, 78, 1, 0, 0, 0, 
	29, 80, 1, 0, 0, 0, 31, 90, 1, 0, 0, 0, 33, 92, 1, 0, 0, 0, 35, 36, 5, 
	10, 0, 0, 36, 2, 1, 0, 0, 0, 37, 38, 5, 82, 0, 0, 38, 39, 5, 101, 0, 0, 
	39, 40, 5, 103, 0, 0, 40, 41, 5, 105, 0, 0, 41, 42, 5, 115, 0, 0, 42, 43, 
	5, 116, 0, 0, 43, 44, 5, 101, 0, 0, 44, 45, 5, 114, 0, 0, 45, 46, 5, 32, 
	0, 0, 46, 4, 1, 0, 0, 0, 47, 48, 5, 58, 0, 0, 48, 49, 5, 32, 0, 0, 49, 
	6, 1, 0, 0, 0, 50, 51, 5, 80, 0, 0, 51, 52, 5, 114, 0, 0, 52, 53, 5, 111, 
	0, 0, 53, 54, 5, 103, 0, 0, 54, 55, 5, 114, 0, 0, 55, 56, 5, 97, 0, 0, 
	56, 57, 5, 109, 0, 0, 57, 58, 5, 58, 0, 0, 58, 59, 5, 32, 0, 0, 59, 8, 
	1, 0, 0, 0, 60, 61, 5, 44, 0, 0, 61, 10, 1, 0, 0, 0, 62, 63, 5, 48, 0, 
	0, 63, 12, 1, 0, 0, 0, 64, 65, 5, 49, 0, 0, 65, 14, 1, 0, 0, 0, 66, 67, 
	5, 50, 0, 0, 67, 16, 1, 0, 0, 0, 68, 69, 5, 51, 0, 0, 69, 18, 1, 0, 0, 
	0, 70, 71, 5, 52, 0, 0, 71, 20, 1, 0, 0, 0, 72, 73, 5, 53, 0, 0, 73, 22, 
	1, 0, 0, 0, 74, 75, 5, 54, 0, 0, 75, 24, 1, 0, 0, 0, 76, 77, 5, 55, 0, 
	0, 77, 26, 1, 0, 0, 0, 78, 79, 2, 65, 67, 0, 79, 28, 1, 0, 0, 0, 80, 81, 
	3, 31, 15, 0, 81, 30, 1, 0, 0, 0, 82, 91, 5, 48, 0, 0, 83, 87, 7, 0, 0, 
	0, 84, 86, 3, 33, 16, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 
	1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 
	90, 82, 1, 0, 0, 0, 90, 83, 1, 0, 0, 0, 91, 32, 1, 0, 0, 0, 92, 93, 7, 
	1, 0, 0, 93, 34, 1, 0, 0, 0, 3, 0, 87, 90, 0,
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
	AOCLexerT__6 = 7
	AOCLexerT__7 = 8
	AOCLexerT__8 = 9
	AOCLexerT__9 = 10
	AOCLexerT__10 = 11
	AOCLexerT__11 = 12
	AOCLexerT__12 = 13
	AOCLexerID = 14
	AOCLexerINTEGER = 15
	AOCLexerDIGIT = 16
)

