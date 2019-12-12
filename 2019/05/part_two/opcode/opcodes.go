package opcode

import (
	"github.com/maxcleme/adventofcode/2019/05/part_two/modes"
)

var Opcodes = map[int]Opcode{
	1: Add{},
	2: Mult{},
	3: Input{},
	4: Output{},
	5: JumpIfTrue{},
	6: JumpIfFalse{},
	7: Lt{},
	8: Eq{},

	99: Exit{},
}

type Opcode interface {
	Evaluate(input []int, offset int, modes []modes.Mode) (int, error)
}
