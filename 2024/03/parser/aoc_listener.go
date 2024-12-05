// Code generated from AOC.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AOC

import "github.com/antlr4-go/antlr/v4"


// AOCListener is a complete listener for a parse tree produced by AOCParser.
type AOCListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterMul is called when entering the mul production.
	EnterMul(c *MulContext)

	// EnterDo is called when entering the do production.
	EnterDo(c *DoContext)

	// EnterDon_t is called when entering the don_t production.
	EnterDon_t(c *Don_tContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitMul is called when exiting the mul production.
	ExitMul(c *MulContext)

	// ExitDo is called when exiting the do production.
	ExitDo(c *DoContext)

	// ExitDon_t is called when exiting the don_t production.
	ExitDon_t(c *Don_tContext)
}
