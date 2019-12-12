package opcode

import (
	"github.com/maxcleme/adventofcode/2019/05/part_two/modes"
)

type Exit struct {
}

func (o Exit) Evaluate(input []int, offset int, modes []modes.Mode) (int, error) {
	return len(input), nil
}
