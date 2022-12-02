package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read input file
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	point := 0
	for scanner.Scan() {
		l := scanner.Text()
		draws := strings.Split(l, " ")
		p, err := ParseRound(draws[0], draws[1])
		if err != nil {
			panic(err)
		}
		point += p
	}
	fmt.Println(point)
}

type Input int
type Result int

var cheat = NewCheat([]Input{Rock, Paper, Scissors})

type Cheat struct {
	Inputs []Input
	index  map[Input]int
}

func NewCheat(inputs []Input) Cheat {
	index := make(map[Input]int, len(inputs))
	for i, input := range inputs {
		index[input] = i
	}
	return Cheat{
		Inputs: inputs,
		index:  index,
	}
}

func (c Cheat) Win(input Input) Input {
	idx, ok := c.index[input]
	if !ok {
		return -1
	}
	return c.Inputs[(idx+1)%len(c.Inputs)]
}
func (c Cheat) Loose(input Input) Input {
	idx, ok := c.index[input]
	if !ok {
		return -1
	}
	if idx == 0 {
		return c.Inputs[len(c.Inputs)-1]
	}
	return c.Inputs[(idx-1)%len(c.Inputs)]
}
func (c Cheat) Draw(i Input) Input {
	return i
}

var Rock = Input(1)
var Paper = Input(2)
var Scissors = Input(3)

var Win = Result(6)
var Draw = Result(3)
var Loose = Result(0)

var (
	drawMapping = map[string]Input{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	resultMapping = map[string]Result{
		"X": Loose,
		"Y": Draw,
		"Z": Win,
	}
)

func ParseInput(s string) (Input, error) {
	draw, ok := drawMapping[s]
	if !ok {
		return 0, fmt.Errorf("parsing draw: %s", s)
	}
	return draw, nil
}

func ParseResult(s string) (Result, error) {
	result, ok := resultMapping[s]
	if !ok {
		return 0, fmt.Errorf("parsing result: %s", s)
	}
	return result, nil
}

func ParseRound(elveStr, myStr string) (int, error) {
	elveInput, err := ParseInput(elveStr)
	if err != nil {
		return 0, err
	}
	expectedResult, err := ParseResult(myStr)
	if err != nil {
		return 0, err
	}
	switch expectedResult {
	case Win:
		return int(Win) + int(cheat.Win(elveInput)), nil
	case Loose:
		return int(Loose) + int(cheat.Loose(elveInput)), nil
	}
	return int(Draw) + int(cheat.Draw(elveInput)), nil
}
