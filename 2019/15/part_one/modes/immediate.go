package modes

import "github.com/maxcleme/adventofcode/2019/15/part_one/model"

type Immediate struct {
}

func (m Immediate) AtRead(p *model.Program, src int) int {
	return src
}


func (m Immediate) AtWrite(p *model.Program, src int) int {
	return src
}
