package opcode

import (
	"fmt"

	"github.com/maxcleme/adventofcode/2019/05/part_two/modes"
)

type Input struct {
}

func (o Input) Evaluate(input []int, offset int, modes []modes.Mode) (int, error) {
	dest := input[offset+1]
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return 0, err
	}
	input[dest] = i
	return offset + 2, nil
}