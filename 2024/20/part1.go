package _20

import (
	"fmt"
	"math"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

func part1(input string, minSave int) int {
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
	return bestCheat(path, minSave, 2)
}

func bestCheat(path []*grid.Tile[string], minSave, cheatDuration int) int {
	distances := map[*grid.Tile[string]]int{}
	for i, t := range path {
		distances[t] = i
	}
	total := 0
	for i, s := range path[:len(path)-1] {
		for _, e := range path[i+1:] {
			d := taxicabDistance(s, e)
			if d > cheatDuration {
				continue
			}
			saved := distances[e] - distances[s]
			// check if cheating is worth it
			if d == saved {
				continue
			}
			if saved-d < minSave {
				continue
			}
			total++
		}
	}
	return total
}

func taxicabDistance(a, b *grid.Tile[string]) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/20/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f), 100))
		return nil
	},
}
