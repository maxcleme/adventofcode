package _13

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	costA = 3
	costB = 1
)

type ClawMachineConfig struct {
	A             Button
	B             Button
	PrizeEquation Equation
}

// See https://en.wikipedia.org/wiki/Cramer%27s_rule
func (conf ClawMachineConfig) Resolve() (*Solution, bool) {
	// Coefficients matrix A
	a, b := conf.A.X, conf.B.X
	c, d := conf.A.Y, conf.B.Y

	// Constants vector B
	e, f := conf.PrizeEquation.X, conf.PrizeEquation.Y

	// Calculate determinant of A
	detA := determinant(a, b, c, d)

	// Check if determinant is zero (no solution or infinite solutions)
	if detA == 0 {
		return nil, false
	}

	// Calculate determinants of A_x and A_y
	detAx := determinant(e, b, f, d)
	detAy := determinant(a, e, c, f)

	// Solve for x and y
	x := float64(detAx) / float64(detA)
	y := float64(detAy) / float64(detA)

	// We cannot move by a fraction of a unit
	if x != float64(int(x)) || y != float64(int(y)) {
		return nil, false
	}
	return &Solution{int(x), int(y)}, true
}

// Function to calculate determinant of a 2x2 matrix
func determinant(a, b, c, d int) int {
	return a*d - b*c
}

type Button struct {
	X int
	Y int
}

type Equation struct {
	X int
	Y int
}

type Solution struct {
	A int
	B int
}

func part1(input string) int {
	var configs []ClawMachineConfig
	for _, claw := range strings.Split(input, "\n\n") {
		lines := strings.Split(claw, "\n")
		a := Button{}
		if _, err := fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &a.X, &a.Y); err != nil {
			panic(err)
		}
		b := Button{}
		if _, err := fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &b.X, &b.Y); err != nil {
			panic(err)
		}
		prize := Equation{}
		if _, err := fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &prize.X, &prize.Y); err != nil {
			panic(err)
		}
		configs = append(configs, ClawMachineConfig{A: a, B: b, PrizeEquation: prize})
	}

	total := 0
	for _, config := range configs {
		if solution, ok := config.Resolve(); ok {
			total += costA*solution.A + costB*solution.B
		}
	}
	return total
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/13/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
