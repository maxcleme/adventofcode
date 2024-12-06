package _06

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func part1(input string) int {
	m, xStart, yStart := NewMap(input)
	x, y := xStart, yStart
	walk(m, x, y)
	return len(m.history)
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/06/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
