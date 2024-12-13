package _13

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func part2(input string) int {
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
		prize.X += 10000000000000
		prize.Y += 10000000000000
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

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/13/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
