package opcode

import (
	"github.com/maxcleme/adventofcode/2019/15/part_two/modes"
)

type Instruction struct {
	Opcode Opcode
	Modes  []modes.Mode
}
