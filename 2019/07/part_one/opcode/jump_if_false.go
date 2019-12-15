package opcode

import (
	"io"

	"github.com/maxcleme/adventofcode/2019/07/part_one/modes"
)

type JumpIfFalse struct {
}

func (o JumpIfFalse) Evaluate(input []int, offset int, modes []modes.Mode, in io.Reader, out io.Writer) (int, error) {
	val1 := modes[0].Read(input, offset+1)
	val2 := modes[1].Read(input, offset+2)

	if val1 == 0 {
		return val2, nil
	}
	return offset + 3, nil
}
