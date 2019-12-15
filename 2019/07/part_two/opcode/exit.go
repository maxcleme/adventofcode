package opcode

import (
	"io"

	"github.com/maxcleme/adventofcode/2019/07/part_two/modes"
)

type Exit struct {
}

func (o Exit) Evaluate(input []int, offset int, modes []modes.Mode, in io.Reader, out io.Writer) (int, error) {
	return len(input), nil
}
