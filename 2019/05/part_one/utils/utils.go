package utils

import (
	"fmt"
	"strconv"
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
	Opcode Opcode
	Modes  []Mode
}

var Modes = map[int]Mode{
	0: PositionMode{},
	1: ImmediateMode{},
}

var Opcodes = map[int]Opcode{
	1: Add{},
	2: Mult{},
	3: Input{},
	4: Output{},

	99: Exit{},
}

type Mode interface {
	Read(input []int, src int) int
}

type PositionMode struct {
}

func (m PositionMode) Read(input []int, src int) int {
	return input[input[src]]
}

type ImmediateMode struct {
}

func (m ImmediateMode) Read(input []int, src int) int {
	return input[src]
}

type Opcode interface {
	Evaluate(input []int, offset int, modes []Mode) (int, error)
}

type Add struct {
}

func (o Add) Evaluate(input []int, offset int, modes []Mode) (int, error) {
	val1 := modes[0].Read(input, offset+1)
	val2 := modes[1].Read(input, offset+2)
	dest := input[offset+3]
	input[dest] = val1 + val2
	return offset + 4, nil
}

type Mult struct {
}

func (o Mult) Evaluate(input []int, offset int, modes []Mode) (int, error) {
	val1 := modes[0].Read(input, offset+1)
	val2 := modes[1].Read(input, offset+2)
	dest := input[offset+3]
	input[dest] = val1 * val2
	return offset + 4, nil
}

type Input struct {
}

func (o Input) Evaluate(input []int, offset int, modes []Mode) (int, error) {
	dest := input[offset+1]
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return 0, err
	}
	input[dest] = i
	return offset + 2, nil
}

type Output struct {
}

func (o Output) Evaluate(input []int, offset int, modes []Mode) (int, error) {
	src := modes[0].Read(input, offset+1)
	fmt.Println(">", src)
	return offset + 2, nil
}

type Exit struct {
}

func (o Exit) Evaluate(input []int, offset int, modes []Mode) (int, error) {
	return len(input), nil
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

	opcode, ok := Opcodes[o]
	if !ok {
		return nil, ErrUnknownOpCode
	}

	modes := []Mode{
		PositionMode{},
		PositionMode{},
		PositionMode{},
	}

	for i, m := range reverse(runes[:len(runes)-2]) {
		mi, err := strconv.Atoi(string(m))
		if err != nil {
			return nil, ErrInvalidMode
		}
		mode, ok := Modes[mi]
		if !ok {
			return nil, ErrUnknownMode
		}
		modes[i] = mode
	}

	return &Instruction{
		Opcode: opcode,
		Modes:  modes,
	}, nil

}

func Run(program []int) ([]int, error) {
	// clone program
	input := make([]int, len(program))
	for i, statement := range program {
		input[i] = statement
	}

	offset := 0
	for offset < len(input) {
		instruction, err := ParseInstruction(input[offset])
		if err != nil {
			return nil, err
		}

		next, err := instruction.Opcode.Evaluate(input, offset, instruction.Modes)
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
