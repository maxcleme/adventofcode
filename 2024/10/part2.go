package _10

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func part2(input string) int {
	grid := NewGrid(input)
	trailheads := grid.Find(func(p Pos) bool {
		slope, ok := grid.Get(p.X, p.Y)
		return ok && slope == 0
	})
	total := 0
	for _, trailhead := range trailheads {
		visited := make([][]bool, len(grid))
		for i := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}
		climb(grid, trailhead.X, trailhead.Y, visited, func(_, _ int) {
			total++
		})
	}
	return total
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/10/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
