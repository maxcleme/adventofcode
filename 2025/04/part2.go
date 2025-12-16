package _04

import (
	"fmt"
	"os"
	"slices"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/04/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(payload string) int {
	g := grid.New[string](payload, func(r rune) string {
		return string(r)
	})
	total := 0
	for {
		toBeRemoved := g.Find(func(t *grid.Tile[string]) bool {
			if t.Value != "@" {
				return false
			}
			return len(slices.DeleteFunc(g.Neighbors(t), func(t *grid.Tile[string]) bool {
				return t.Value != "@"
			})) < 4
		})
		for _, t := range toBeRemoved {
			t.Value = "x"
		}
		count := len(toBeRemoved)
		if count == 0 {
			return total
		}
		total += count
		g.Delete(func(t *grid.Tile[string]) bool {
			return t.Value == "x"
		}, ".")
	}
}
