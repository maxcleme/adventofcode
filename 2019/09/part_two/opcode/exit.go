package opcode

import (
	"github.com/maxcleme/adventofcode/2019/09/part_two/modes"
	"github.com/maxcleme/adventofcode/2019/09/part_two/model"
)

type Exit struct {
}

func (o Exit) Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error) {
	return len(p.Statements), nil
}
