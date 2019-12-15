package modes

type Immediate struct {
}

func (m Immediate) Read(input []int, src int) int {
	return input[src]
}
