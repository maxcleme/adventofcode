package _03

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/03/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(input string) int {
	xSanta, ySanta, xRobot, yRobot := 0, 0, 0, 0
	visited := map[[2]int]struct{}{
		[2]int{0, 0}: {},
	}
	for i, r := range input {
		x, y := &xSanta, &ySanta
		if i%2 == 1 {
			x, y = &xRobot, &yRobot
		}
		switch r {
		case '^':
			*y--
		case 'v':
			*y++
		case '>':
			*x++
		case '<':
			*x--
		}
		visited[[2]int{*x, *y}] = struct{}{}
	}
	return len(visited)
}
