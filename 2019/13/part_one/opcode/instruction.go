package opcode

import (
	"github.com/maxcleme/adventofcode/2019/13/part_one/modes"
)

type Instruction struct {
	Opcode Opcode
	Modes  []modes.Mode
}
