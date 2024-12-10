package _10

import (
	"fmt"
	"os"
	"strconv"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

func part2(input string) int {
	g := grid.New[int](input, func(r rune) int {
		if r == '.' {
			return -1
		}
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		return n
	})
	trailheads := g.Find(func(t *grid.Tile[int]) bool {
		return t.Value == 0
	})
	total := 0
	for _, trailhead := range trailheads {
		visited := make([][]bool, g.Height())
		for i := range visited {
			visited[i] = make([]bool, g.Width())
		}
		climb(g, trailhead.X, trailhead.Y, visited, func(t *grid.Tile[int]) {
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
