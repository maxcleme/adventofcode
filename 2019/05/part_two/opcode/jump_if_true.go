package opcode

import "github.com/maxcleme/adventofcode/2019/05/part_two/modes"

type JumpIfTrue struct {
}

func (o JumpIfTrue) Evaluate(input []int, offset int, modes []modes.Mode) (int, error) {
	val1 := modes[0].Read(input, offset+1)
	val2 := modes[1].Read(input, offset+2)

	if val1 != 0 {
		return val2, nil
	}
	return offset + 3, nil
}
