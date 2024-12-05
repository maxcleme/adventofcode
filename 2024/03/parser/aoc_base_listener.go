// Code generated from AOC.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AOC

import "github.com/antlr4-go/antlr/v4"

// BaseAOCListener is a complete listener for a parse tree produced by AOCParser.
type BaseAOCListener struct{}

var _ AOCListener = &BaseAOCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAOCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAOCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAOCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAOCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseAOCListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAOCListener) ExitProgram(ctx *ProgramContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseAOCListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseAOCListener) ExitInstruction(ctx *InstructionContext) {}

// EnterMul is called when production mul is entered.
func (s *BaseAOCListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production mul is exited.
func (s *BaseAOCListener) ExitMul(ctx *MulContext) {}

// EnterDo is called when production do is entered.
func (s *BaseAOCListener) EnterDo(ctx *DoContext) {}

// ExitDo is called when production do is exited.
func (s *BaseAOCListener) ExitDo(ctx *DoContext) {}

// EnterDon_t is called when production don_t is entered.
func (s *BaseAOCListener) EnterDon_t(ctx *Don_tContext) {}

// ExitDon_t is called when production don_t is exited.
func (s *BaseAOCListener) ExitDon_t(ctx *Don_tContext) {}
