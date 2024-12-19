package _18

import (
	"fmt"
	"os"
	"strings"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

type Direction int

const (
	Up = iota
	Down
	Left
	Right
)

type State struct {
	position *grid.Tile[string]
	cost     int
}

func part1(input string, width int, height, byteCount int) int {
	g := grid.Empty[string](width, height)
	start, _ := g.Get(0, 0)
	end, _ := g.Get(width-1, height-1)
	for _, row := range strings.Split(input, "\n")[:byteCount] {
		x, y := 0, 0
		_, err := fmt.Sscanf(row, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		if t, ok := g.Get(x, y); ok {
			t.Value = "#"
		}
	}
	count, _ := g.Shortest(start, end, func(t *grid.Tile[string]) bool {
		return t.Value != "#"
	})
	return count
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/18/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f), 71, 71, 1024))
		return nil
	},
}
