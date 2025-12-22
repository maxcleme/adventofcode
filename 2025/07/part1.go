package _07

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/07/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(payload string) int {
	d := grid.New[string](payload, func(r rune) string {
		return string(r)
	})
	start := d.Find(func(t *grid.Tile[string]) bool {
		return t.Value == "S"
	})

	total := 0
	var sources = start
	var newBeams []*grid.Tile[string]
	for {
		newBeams = []*grid.Tile[string]{}
		for _, src := range sources {
			b, ok := d.Get(src.X, src.Y+1)
			if !ok {
				continue
			}
			switch b.Value {
			case ".":
				b.Value = "|"
				newBeams = append(newBeams, b)
			case "^":
				total += 1
				left, ok := d.Get(b.X-1, b.Y)
				if ok && left.Value == "." {
					left.Value = "|"
					newBeams = append(newBeams, left)
				}
				right, ok := d.Get(b.X+1, b.Y)
				if ok && right.Value == "." {
					right.Value = "|"
					newBeams = append(newBeams, right)
				}
			case "|":
				continue
			}
		}
		if len(newBeams) == 0 {
			break
		}
		sources = newBeams
	}
	return total
}
