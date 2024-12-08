package _07

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Equation struct {
	Result int
	Parts  []int
}

type Operator interface {
	Apply(int, int) int
}

type Plus struct{}

func (p Plus) Apply(i1 int, i2 int) int {
	return i1 + i2
}

type Mult struct{}

func (m Mult) Apply(i1 int, i2 int) int {
	return i1 * i2
}

func (e Equation) Solve(operators ...Operator) bool {
	if len(e.Parts) == 1 {
		return e.Result == e.Parts[0]
	}
	for _, op := range operators {
		e := Equation{Result: e.Result, Parts: append([]int{op.Apply(e.Parts[0], e.Parts[1])}, e.Parts[2:]...)}
		if e.Solve(operators...) {
			return true
		}
	}
	return false
}

func part1(input string) int {
	rows := strings.Split(input, "\n")
	sum := 0
	for _, row := range rows {
		if eq := makeEquation(row); eq.Solve(Plus{}, Mult{}) {
			sum += eq.Result
		}
	}
	return sum
}

func makeEquation(row string) Equation {
	parts := strings.Split(row, ":")
	return Equation{
		Result: parseNumber(parts[0]),
		Parts:  parseNumbers(strings.Split(strings.TrimSpace(parts[1]), " ")),
	}
}

func parseNumber(number string) int {
	n, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return n
}

func parseNumbers(numbers []string) []int {
	var n []int
	for _, number := range numbers {
		n = append(n, parseNumber(number))
	}
	return n
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/07/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
