package _04

import (
	"fmt"
	"os"
	"slices"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/04/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(payload string) int {
	g := grid.New[rune](payload, func(r rune) rune {
		return r
	})
	return len(g.Find(func(t *grid.Tile[rune]) bool {
		if t.Value != '@' {
			return false
		}
		return len(slices.DeleteFunc(g.Neighbors(t), func(t *grid.Tile[rune]) bool {
			return t.Value != '@'
		})) < 4
	}))
}
