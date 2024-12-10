package _10

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func climb(g Grid, x, y int, visited [][]bool, onSummit func(int, int)) {
	slope, _ := g.Get(x, y)
	if slope == 9 {
		onSummit(x, y)
		return
	}
	visited[x][y] = true
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		nextSlope, ok := g.Get(nx, ny)
		if ok && nextSlope == slope+1 && !visited[nx][ny] {
			climb(g, nx, ny, visited, onSummit)
		}
	}
	visited[x][y] = false
}

func part1(input string) int {
	grid := NewGrid(input)
	trailheads := grid.Find(func(p Pos) bool {
		slope, ok := grid.Get(p.X, p.Y)
		return ok && slope == 0
	})
	total := 0
	for _, trailhead := range trailheads {
		reachedSummits := make(map[Pos]bool)
		visited := make([][]bool, len(grid))
		for i := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}
		climb(grid, trailhead.X, trailhead.Y, visited, func(x, y int) {
			reachedSummits[Pos{x, y}] = true
		})
		total += len(reachedSummits)
	}
	return total
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/10/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
