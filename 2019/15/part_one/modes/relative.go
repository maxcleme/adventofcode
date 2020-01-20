package modes

import "github.com/maxcleme/adventofcode/2019/15/part_one/model"

type Relative struct {
}

func (m Relative) AtRead(p *model.Program, src int) int {
	return p.Statements[src+p.Base]
}

func (m Relative) AtWrite(p *model.Program, src int) int {
	return src + p.Base
}
