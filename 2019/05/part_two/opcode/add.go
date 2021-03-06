package opcode

import (
	"github.com/maxcleme/adventofcode/2019/05/part_two/modes"
)

type Add struct {
}

func (o Add) Evaluate(input []int, offset int, modes []modes.Mode) (int, error) {
	val1 := modes[0].Read(input, offset+1)
	val2 := modes[1].Read(input, offset+2)
	dest := input[offset+3]
	input[dest] = val1 + val2
	return offset + 4, nil
}
