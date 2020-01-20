package opcode

import (
	"github.com/maxcleme/adventofcode/2019/15/part_one/modes"
	"github.com/maxcleme/adventofcode/2019/15/part_one/model"
)

type JumpIfTrue struct {
}

func (o JumpIfTrue) Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error) {
	val1 := modes[0].AtRead(p, p.Statements[offset+1])
	val2 := modes[1].AtRead(p, p.Statements[offset+2])
	if val1 != 0 {
		return val2, nil
	}
	return offset + 3, nil
}
