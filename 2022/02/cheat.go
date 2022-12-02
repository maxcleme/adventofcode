package _02

type CheatEngine struct {
	Inputs []Input
	index  map[Input]int
}

func NewCheatEngine(inputs []Input) CheatEngine {
	index := make(map[Input]int, len(inputs))
	for i, input := range inputs {
		index[input] = i
	}
	return CheatEngine{
		Inputs: inputs,
		index:  index,
	}
}

func (c CheatEngine) Win(input Input) Input {
	idx, ok := c.index[input]
	if !ok {
		return -1
	}
	return c.Inputs[(idx+1)%len(c.Inputs)]
}
func (c CheatEngine) Loose(input Input) Input {
	idx, ok := c.index[input]
	if !ok {
		return -1
	}
	if idx == 0 {
		return c.Inputs[len(c.Inputs)-1]
	}
	return c.Inputs[(idx-1)%len(c.Inputs)]
}
func (c CheatEngine) Draw(i Input) Input {
	return i
}

func (c CheatEngine) Point(elveInput, myInput Input) int {
	if c.Win(elveInput) == myInput {
		return int(Win) + int(myInput)
	}
	if c.Loose(elveInput) == myInput {
		return int(Loose) + int(myInput)
	}
	return int(Draw) + int(myInput)
}
