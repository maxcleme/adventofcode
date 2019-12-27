package modes

import "github.com/maxcleme/adventofcode/2019/13/part_two/model"

type Position struct {
}

func (m Position) AtRead(p *model.Program, src int) int {
	return p.Statements[src]
}

func (m Position) AtWrite(p *model.Program, src int) int {
	return src
}
