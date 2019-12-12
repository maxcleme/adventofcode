package opcode

import (
	"fmt"

	"github.com/maxcleme/adventofcode/2019/05/part_two/modes"
)

type Output struct {
}

func (o Output) Evaluate(input []int, offset int, modes []modes.Mode) (int, error) {
	src := modes[0].Read(input, offset+1)
	fmt.Println(">", src)
	return offset + 2, nil
}
