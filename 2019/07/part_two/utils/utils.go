package utils

import (
	"io"
	"strconv"

	"github.com/maxcleme/adventofcode/2019/07/part_two/modes"
	"github.com/maxcleme/adventofcode/2019/07/part_two/opcode"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrInvalidOpCode = Error("invalid opcode")
	ErrUnknownOpCode = Error("unknown opcode")
	ErrUnknownMode   = Error("unknown mode")
	ErrInvalidMode   = Error("invalid mode")
)

type Instruction struct {
	Opcode opcode.Opcode
	Modes  []modes.Mode
}

type Program struct {
	Statements []int
	In         io.Reader
	Out        io.Writer
}

func ParseInstruction(instruction int) (*Instruction, error) {
	s := strconv.Itoa(instruction)
	if len(s) < 2 {
		s = "0" + s
	}

	runes := []rune(s)
	o, err := strconv.Atoi(string(runes[len(runes)-2:]))
	if err != nil {
		return nil, ErrInvalidOpCode
	}

	opcode, ok := opcode.Opcodes[o]
	if !ok {
		return nil, ErrUnknownOpCode
	}

	paramModes := []modes.Mode{
		modes.Position{},
		modes.Position{},
		modes.Position{},
	}

	for i, m := range reverse(runes[:len(runes)-2]) {
		mi, err := strconv.Atoi(string(m))
		if err != nil {
			return nil, ErrInvalidMode
		}
		mode, ok := modes.Modes[mi]
		if !ok {
			return nil, ErrUnknownMode
		}
		paramModes[i] = mode
	}

	return &Instruction{
		Opcode: opcode,
		Modes:  paramModes,
	}, nil

}

func Run(program Program) ([]int, error) {
	// clone program
	input := make([]int, len(program.Statements))
	for i, statement := range program.Statements {
		input[i] = statement
	}

	offset := 0
	for offset < len(input) {
		instruction, err := ParseInstruction(input[offset])
		if err != nil {
			return nil, err
		}

		next, err := instruction.Opcode.Evaluate(input, offset, instruction.Modes, program.In, program.Out)
		if err != nil {
			return nil, err
		}
		offset = next
	}
	return input, nil
}

func reverse(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}
