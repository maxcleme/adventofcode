package opcode

import (
	"github.com/maxcleme/adventofcode/2019/15/part_two/model"
	"github.com/maxcleme/adventofcode/2019/15/part_two/modes"
)

type RelativeBase struct {
}

func (o RelativeBase) Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error) {
	val := modes[0].AtRead(p, p.Statements[offset+1])
	p.Base += val
	return offset + 2, nil
}
