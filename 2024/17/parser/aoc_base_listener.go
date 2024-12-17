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

// EnterRegisters is called when production registers is entered.
func (s *BaseAOCListener) EnterRegisters(ctx *RegistersContext) {}

// ExitRegisters is called when production registers is exited.
func (s *BaseAOCListener) ExitRegisters(ctx *RegistersContext) {}

// EnterRegister is called when production register is entered.
func (s *BaseAOCListener) EnterRegister(ctx *RegisterContext) {}

// ExitRegister is called when production register is exited.
func (s *BaseAOCListener) ExitRegister(ctx *RegisterContext) {}

// EnterInstructions is called when production instructions is entered.
func (s *BaseAOCListener) EnterInstructions(ctx *InstructionsContext) {}

// ExitInstructions is called when production instructions is exited.
func (s *BaseAOCListener) ExitInstructions(ctx *InstructionsContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseAOCListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseAOCListener) ExitInstruction(ctx *InstructionContext) {}

// EnterAdv is called when production adv is entered.
func (s *BaseAOCListener) EnterAdv(ctx *AdvContext) {}

// ExitAdv is called when production adv is exited.
func (s *BaseAOCListener) ExitAdv(ctx *AdvContext) {}

// EnterBxl is called when production bxl is entered.
func (s *BaseAOCListener) EnterBxl(ctx *BxlContext) {}

// ExitBxl is called when production bxl is exited.
func (s *BaseAOCListener) ExitBxl(ctx *BxlContext) {}

// EnterBst is called when production bst is entered.
func (s *BaseAOCListener) EnterBst(ctx *BstContext) {}

// ExitBst is called when production bst is exited.
func (s *BaseAOCListener) ExitBst(ctx *BstContext) {}

// EnterJnz is called when production jnz is entered.
func (s *BaseAOCListener) EnterJnz(ctx *JnzContext) {}

// ExitJnz is called when production jnz is exited.
func (s *BaseAOCListener) ExitJnz(ctx *JnzContext) {}

// EnterBxc is called when production bxc is entered.
func (s *BaseAOCListener) EnterBxc(ctx *BxcContext) {}

// ExitBxc is called when production bxc is exited.
func (s *BaseAOCListener) ExitBxc(ctx *BxcContext) {}

// EnterOut is called when production out is entered.
func (s *BaseAOCListener) EnterOut(ctx *OutContext) {}

// ExitOut is called when production out is exited.
func (s *BaseAOCListener) ExitOut(ctx *OutContext) {}

// EnterBdv is called when production bdv is entered.
func (s *BaseAOCListener) EnterBdv(ctx *BdvContext) {}

// ExitBdv is called when production bdv is exited.
func (s *BaseAOCListener) ExitBdv(ctx *BdvContext) {}

// EnterCdv is called when production cdv is entered.
func (s *BaseAOCListener) EnterCdv(ctx *CdvContext) {}

// ExitCdv is called when production cdv is exited.
func (s *BaseAOCListener) ExitCdv(ctx *CdvContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseAOCListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseAOCListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral_value is called when production literal_value is entered.
func (s *BaseAOCListener) EnterLiteral_value(ctx *Literal_valueContext) {}

// ExitLiteral_value is called when production literal_value is exited.
func (s *BaseAOCListener) ExitLiteral_value(ctx *Literal_valueContext) {}

// EnterReg_a is called when production reg_a is entered.
func (s *BaseAOCListener) EnterReg_a(ctx *Reg_aContext) {}

// ExitReg_a is called when production reg_a is exited.
func (s *BaseAOCListener) ExitReg_a(ctx *Reg_aContext) {}

// EnterReg_b is called when production reg_b is entered.
func (s *BaseAOCListener) EnterReg_b(ctx *Reg_bContext) {}

// ExitReg_b is called when production reg_b is exited.
func (s *BaseAOCListener) ExitReg_b(ctx *Reg_bContext) {}

// EnterReg_c is called when production reg_c is entered.
func (s *BaseAOCListener) EnterReg_c(ctx *Reg_cContext) {}

// ExitReg_c is called when production reg_c is exited.
func (s *BaseAOCListener) ExitReg_c(ctx *Reg_cContext) {}
