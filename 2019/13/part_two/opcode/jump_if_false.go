package opcode

import (
	"github.com/maxcleme/adventofcode/2019/13/part_two/modes"
	"github.com/maxcleme/adventofcode/2019/13/part_two/model"
)

type JumpIfFalse struct {
}

func (o JumpIfFalse) Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error) {
	val1 := modes[0].AtRead(p, p.Statements[offset+1])
	val2 := modes[1].AtRead(p, p.Statements[offset+2])
	if val1 == 0 {
		return val2, nil
	}
	return offset + 3, nil
}
