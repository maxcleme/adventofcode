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
		elveInput, myInput, err := ParseRound(draws[0], draws[1])
		if err != nil {
			panic(err)
		}
		point += round(elveInput, myInput)
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
		"X": Rock,
		"B": Paper,
		"Y": Paper,
		"C": Scissors,
		"Z": Scissors,
	}
)

func ParseInput(s string) (Input, error) {
	draw, ok := drawMapping[s]
	if !ok {
		return 0, fmt.Errorf("parsing draw: %s", s)
	}
	return draw, nil
}

func ParseRound(elveStr, myStr string) (Input, Input, error) {
	elveInput, err := ParseInput(elveStr)
	if err != nil {
		return 0, 0, err
	}
	myInput, err := ParseInput(myStr)
	if err != nil {
		return 0, 0, err
	}
	return elveInput, myInput, nil
}

func round(elveInput, myInput Input) int {
	if cheat.Win(elveInput) == myInput {
		return int(Win) + int(myInput)
	}
	if cheat.Loose(elveInput) == myInput {
		return int(Loose) + int(myInput)
	}
	return int(Draw) + int(myInput)
}
