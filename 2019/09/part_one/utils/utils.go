package utils

import (
	"strconv"

	"github.com/maxcleme/adventofcode/2019/09/part_one/model"
	"github.com/maxcleme/adventofcode/2019/09/part_one/modes"
	"github.com/maxcleme/adventofcode/2019/09/part_one/opcode"
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

func ParseInstruction(instruction int) (*opcode.Instruction, error) {
	s := strconv.Itoa(instruction)
	if len(s) < 2 {
		s = "0" + s
	}

	runes := []rune(s)
	o, err := strconv.Atoi(string(runes[len(runes)-2:]))
	if err != nil {
		return nil, ErrInvalidOpCode
	}

	code, ok := opcode.Opcodes[o]
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

	return &opcode.Instruction{
		Opcode: code,
		Modes:  paramModes,
	}, nil

}

func Run(program model.Program) ([]int, error) {
	// clone model, also add more memory
	input := make([]int, 1024*1024+len(program.Statements)*2)
	for i, statement := range program.Statements {
		input[i] = statement
	}

	p := model.Program{
		Statements: input,
		In:         program.In,
		Out:        program.Out,
		Base:       0,
	}

	offset := 0
	for offset < len(p.Statements) {
		instruction, err := ParseInstruction(p.Statements[offset])
		if err != nil {
			return nil, err
		}

		next, err := instruction.Opcode.Evaluate(&p, offset, instruction.Modes)
		if err != nil {
			return nil, err
		}
		offset = next
	}
	// remove extra memory from output
	return p.Statements[:len(program.Statements)], nil
}

func reverse(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}
