package opcode

import (
	"github.com/maxcleme/adventofcode/2019/15/part_one/modes"
	"github.com/maxcleme/adventofcode/2019/15/part_one/model"
)

type Add struct {
}

func (o Add) Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error) {
	val1 := modes[0].AtRead(p, p.Statements[offset+1])
	val2 := modes[1].AtRead(p, p.Statements[offset+2])
	p.Statements[modes[2].AtWrite(p, p.Statements[offset+3])] = val1 + val2
	return offset + 4, nil
}
