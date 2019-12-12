package opcode

import "github.com/maxcleme/adventofcode/2019/05/part_two/modes"

type Lt struct {
}

func (o Lt) Evaluate(input []int, offset int, modes []modes.Mode) (int, error) {
	val1 := modes[0].Read(input, offset+1)
	val2 := modes[1].Read(input, offset+2)
	dest := input[offset+3]
	if val1 < val2 {
		input[dest] = 1
	} else {
		input[dest] = 0
	}
	return offset + 4, nil
}
