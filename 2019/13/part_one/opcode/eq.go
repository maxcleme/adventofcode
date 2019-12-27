package opcode

import (
	"github.com/maxcleme/adventofcode/2019/13/part_one/modes"
	"github.com/maxcleme/adventofcode/2019/13/part_one/model"
)

type Eq struct {
}

func (o Eq) Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error) {
	val1 := modes[0].AtRead(p, p.Statements[offset+1])
	val2 := modes[1].AtRead(p, p.Statements[offset+2])
	if val1 == val2 {
		p.Statements[modes[2].AtWrite(p, p.Statements[offset+3])] = 1
	} else {
		p.Statements[modes[2].AtWrite(p, p.Statements[offset+3])] = 0
	}
	return offset + 4, nil
}
