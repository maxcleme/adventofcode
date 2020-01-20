package modes

import "github.com/maxcleme/adventofcode/2019/15/part_two/model"

var Modes = map[int]Mode{
	0: Position{},
	1: Immediate{},
	2: Relative{},
}

type Mode interface {
	AtRead(p *model.Program, src int) int
	AtWrite(p *model.Program, src int) int
}
