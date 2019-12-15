package opcode

import (
	"fmt"
	"io"

	"github.com/maxcleme/adventofcode/2019/07/part_two/modes"
)

type Input struct {
}

func (o Input) Evaluate(input []int, offset int, modes []modes.Mode, in io.Reader, out io.Writer) (int, error) {
	dest := input[offset+1]
	var i int
	_, err := fmt.Fscanf(in, "%d\n", &i)
	if err != nil {
		return 0, err
	}
	input[dest] = i
	return offset + 2, nil
}
