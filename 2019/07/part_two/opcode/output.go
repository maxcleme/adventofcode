package opcode

import (
	"io"
	"strconv"

	"github.com/maxcleme/adventofcode/2019/07/part_two/modes"
)

type Output struct {
}

func (o Output) Evaluate(input []int, offset int, modes []modes.Mode, in io.Reader, out io.Writer) (int, error) {
	src := modes[0].Read(input, offset+1)
	_, err := out.Write([]byte(strconv.Itoa(src) + "\n"))
	if err != nil {
		return 0, err
	}
	return offset + 2, nil
}
