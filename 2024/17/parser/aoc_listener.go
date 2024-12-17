// Code generated from AOC.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AOC

import "github.com/antlr4-go/antlr/v4"


// AOCListener is a complete listener for a parse tree produced by AOCParser.
type AOCListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterRegisters is called when entering the registers production.
	EnterRegisters(c *RegistersContext)

	// EnterRegister is called when entering the register production.
	EnterRegister(c *RegisterContext)

	// EnterInstructions is called when entering the instructions production.
	EnterInstructions(c *InstructionsContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterAdv is called when entering the adv production.
	EnterAdv(c *AdvContext)

	// EnterBxl is called when entering the bxl production.
	EnterBxl(c *BxlContext)

	// EnterBst is called when entering the bst production.
	EnterBst(c *BstContext)

	// EnterJnz is called when entering the jnz production.
	EnterJnz(c *JnzContext)

	// EnterBxc is called when entering the bxc production.
	EnterBxc(c *BxcContext)

	// EnterOut is called when entering the out production.
	EnterOut(c *OutContext)

	// EnterBdv is called when entering the bdv production.
	EnterBdv(c *BdvContext)

	// EnterCdv is called when entering the cdv production.
	EnterCdv(c *CdvContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterLiteral_value is called when entering the literal_value production.
	EnterLiteral_value(c *Literal_valueContext)

	// EnterReg_a is called when entering the reg_a production.
	EnterReg_a(c *Reg_aContext)

	// EnterReg_b is called when entering the reg_b production.
	EnterReg_b(c *Reg_bContext)

	// EnterReg_c is called when entering the reg_c production.
	EnterReg_c(c *Reg_cContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitRegisters is called when exiting the registers production.
	ExitRegisters(c *RegistersContext)

	// ExitRegister is called when exiting the register production.
	ExitRegister(c *RegisterContext)

	// ExitInstructions is called when exiting the instructions production.
	ExitInstructions(c *InstructionsContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitAdv is called when exiting the adv production.
	ExitAdv(c *AdvContext)

	// ExitBxl is called when exiting the bxl production.
	ExitBxl(c *BxlContext)

	// ExitBst is called when exiting the bst production.
	ExitBst(c *BstContext)

	// ExitJnz is called when exiting the jnz production.
	ExitJnz(c *JnzContext)

	// ExitBxc is called when exiting the bxc production.
	ExitBxc(c *BxcContext)

	// ExitOut is called when exiting the out production.
	ExitOut(c *OutContext)

	// ExitBdv is called when exiting the bdv production.
	ExitBdv(c *BdvContext)

	// ExitCdv is called when exiting the cdv production.
	ExitCdv(c *CdvContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitLiteral_value is called when exiting the literal_value production.
	ExitLiteral_value(c *Literal_valueContext)

	// ExitReg_a is called when exiting the reg_a production.
	ExitReg_a(c *Reg_aContext)

	// ExitReg_b is called when exiting the reg_b production.
	ExitReg_b(c *Reg_bContext)

	// ExitReg_c is called when exiting the reg_c production.
	ExitReg_c(c *Reg_cContext)
}
