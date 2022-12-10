package _10

type Registry map[string]int

func NewRegistry() Registry {
	return map[string]int{
		"X": 1,
	}
}
