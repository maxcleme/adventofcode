package modes

type Position struct {
}

func (m Position) Read(input []int, src int) int {
	return input[input[src]]
}