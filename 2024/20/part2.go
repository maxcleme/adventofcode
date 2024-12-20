package _20

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

func part2(input string, minSave int) int {
	g := grid.New[string](input, func(r rune) string {
		return string(r)
	})
	start := g.Find(func(t *grid.Tile[string]) bool {
		return t.Value == "S"
	})[0]
	end := g.Find(func(t *grid.Tile[string]) bool {
		return t.Value == "E"
	})[0]
	path := g.Shortest(start, end, func(t *grid.Tile[string]) bool {
		return t.Value != "#"
	})
	return bestCheat(path, minSave, 20)
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/20/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f), 100))
		return nil
	},
}
