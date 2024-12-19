package _18

import (
	"fmt"
	"os"
	"strings"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

func part2(input string, width int, height int) string {
	g := grid.Empty[string](width, height)
	start, _ := g.Get(0, 0)
	end, _ := g.Get(width-1, height-1)
	for _, row := range strings.Split(input, "\n") {
		x, y := 0, 0
		_, err := fmt.Sscanf(row, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		if t, ok := g.Get(x, y); ok {
			t.Value = "#"
		}
		_, ok := g.Shortest(start, end, func(t *grid.Tile[string]) bool {
			return t.Value != "#"
		})
		if !ok {
			return fmt.Sprintf("%d,%d", x, y)
		}
	}
	return "impossible input"
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/18/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f), 71, 71))
		return nil
	},
}
