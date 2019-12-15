package modes

var Modes = map[int]Mode{
	0: Position{},
	1: Immediate{},
}

type Mode interface {
	Read(input []int, src int) int
}
