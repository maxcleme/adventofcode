package _10

import (
	"fmt"
	"os"
	"strconv"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func climb(g *grid.Grid[int], x, y int, visited [][]bool, onSummit func(tile *grid.Tile[int])) {
	currentTile, _ := g.Get(x, y)
	if currentTile.Value == 9 {
		onSummit(currentTile)
		return
	}
	visited[x][y] = true
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		nextSlope, ok := g.Get(nx, ny)
		if ok && nextSlope.Value == currentTile.Value+1 && !visited[nx][ny] {
			climb(g, nx, ny, visited, onSummit)
		}
	}
	visited[x][y] = false
}

func part1(input string) int {
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
		reachedSummits := make(map[*grid.Tile[int]]bool)
		visited := make([][]bool, g.Height())
		for i := range visited {
			visited[i] = make([]bool, g.Width())
		}
		climb(g, trailhead.X, trailhead.Y, visited, func(t *grid.Tile[int]) {
			reachedSummits[t] = true
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
