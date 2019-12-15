package opcode

import (
	"fmt"

	"github.com/maxcleme/adventofcode/2019/09/part_two/modes"
	"github.com/maxcleme/adventofcode/2019/09/part_two/model"
)

type Input struct {
}

func (o Input) Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error) {
	var i int
	_, err := fmt.Fscanf(p.In, "%d\n", &i)
	if err != nil {
		return 0, err
	}
	p.Statements[modes[0].AtWrite(p, p.Statements[offset+1])] = i
	return offset + 2, nil
}
