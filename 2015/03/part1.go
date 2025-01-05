package _03

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/03/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(input string) int {
	x, y := 0, 0
	visited := map[[2]int]struct{}{
		[2]int{x, y}: {},
	}
	for _, r := range input {
		switch r {
		case '^':
			y--
		case 'v':
			y++
		case '>':
			x++
		case '<':
			x--
		}
		visited[[2]int{x, y}] = struct{}{}
	}
	return len(visited)
}
