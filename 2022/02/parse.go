package _02

import "fmt"

var (
	inputMapping = map[string]Input{
		"A": Rock,
		"X": Rock,
		"B": Paper,
		"Y": Paper,
		"C": Scissors,
		"Z": Scissors,
	}
)

var (
	outputMapping = map[string]Output{
		"X": Loose,
		"Y": Draw,
		"Z": Win,
	}
)

func ParseInput(s string) (Input, error) {
	draw, ok := inputMapping[s]
	if !ok {
		return 0, fmt.Errorf("parsing input: %s", s)
	}
	return draw, nil
}

func ParseOutput(s string) (Output, error) {
	draw, ok := outputMapping[s]
	if !ok {
		return 0, fmt.Errorf("parsing input: %s", s)
	}
	return draw, nil
}
