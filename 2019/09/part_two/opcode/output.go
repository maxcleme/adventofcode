package opcode

import (
	"strconv"

	"github.com/maxcleme/adventofcode/2019/09/part_two/model"
	"github.com/maxcleme/adventofcode/2019/09/part_two/modes"
)

type Output struct {
}

func (o Output) Evaluate(p *model.Program, offset int, modes []modes.Mode) (int, error) {
	_, err := p.Out.Write([]byte(strconv.Itoa(modes[0].AtRead(p, p.Statements[offset+1])) + "\n"))
	if err != nil {
		return 0, err
	}
	return offset + 2, nil
}
