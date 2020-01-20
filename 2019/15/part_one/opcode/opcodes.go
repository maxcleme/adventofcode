package opcode

import (
	"github.com/maxcleme/adventofcode/2019/15/part_one/modes"
	"github.com/maxcleme/adventofcode/2019/15/part_one/model"
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
	9: RelativeBase{},

	99: Exit{},
}

type Opcode interface {
	Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error)
}
