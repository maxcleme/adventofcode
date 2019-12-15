package opcode

import (
	"io"

	"github.com/maxcleme/adventofcode/2019/07/part_one/modes"
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
	Evaluate(input []int, offset int, modes []modes.Mode, in io.Reader, out io.Writer) (int, error)
}
