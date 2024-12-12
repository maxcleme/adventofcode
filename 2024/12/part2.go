package _12

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

// BEWARE THIS CODE IS A FREAKING MESS
func (r Region) Sides(g *grid.Grid[Flower]) int {
	g.Map(func(t *grid.Tile[Flower]) {
		t.Value.Up = false
		t.Value.Down = false
		t.Value.Left = false
		t.Value.Right = false
	})
	var outsides []*grid.Tile[Flower]
	for _, t := range r.Flowers {
		// find all tiles having at least one neighbor outside the region
		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if n, ok := g.Get(t.X+dir[0], t.Y+dir[1]); !ok || n.Value.Kind != r.Kind {
				outsides = append(outsides, t)
				break
			}
		}
	}

	total := 0
	for _, t := range outsides {
		// up
		if n, ok := g.Get(t.X, t.Y-1); !t.Value.Up && (!ok || n.Value.Kind != r.Kind) {
			total++
			// mark all tiles in same horizontal line as visited
			for x := t.X - 1; x >= 0; x-- {
				n, ok := g.Get(x, t.Y)
				if !ok || n.Value.Kind != r.Kind {
					break
				}
				if n, ok := g.Get(x, t.Y-1); ok && n.Value.Kind == r.Kind {
					break
				}
				n.Value.Up = true
			}
			for x := t.X + 1; x < g.Width(); x++ {
				n, ok := g.Get(x, t.Y)
				if !ok || n.Value.Kind != r.Kind {
					break
				}
				if n, ok := g.Get(x, t.Y-1); ok && n.Value.Kind == r.Kind {
					break
				}
				n.Value.Up = true
			}
		}
		// down
		if n, ok := g.Get(t.X, t.Y+1); !t.Value.Down && (!ok || n.Value.Kind != r.Kind) {
			total++
			// mark all tiles in same horizontal line as visited
			for x := t.X - 1; x >= 0; x-- {
				n, ok := g.Get(x, t.Y)
				if !ok || n.Value.Kind != r.Kind {
					break
				}
				if n, ok := g.Get(x, t.Y+1); ok && n.Value.Kind == r.Kind {
					break
				}
				n.Value.Down = true
			}
			for x := t.X + 1; x < g.Width(); x++ {
				n, ok := g.Get(x, t.Y)
				if !ok || n.Value.Kind != r.Kind {
					break
				}
				if n, ok := g.Get(x, t.Y+1); ok && n.Value.Kind == r.Kind {
					break
				}
				n.Value.Down = true
			}
		}
		// left
		if n, ok := g.Get(t.X-1, t.Y); !t.Value.Left && (!ok || n.Value.Kind != r.Kind) {
			total++
			// mark all tiles in same vertical line as visited
			for y := t.Y - 1; y >= 0; y-- {
				n, ok := g.Get(t.X, y)
				if !ok || n.Value.Kind != r.Kind {
					break
				}
				if n, ok := g.Get(t.X-1, y); ok && n.Value.Kind == r.Kind {
					break
				}
				n.Value.Left = true
			}
			for y := t.Y + 1; y < g.Height(); y++ {
				n, ok := g.Get(t.X, y)
				if !ok || n.Value.Kind != r.Kind {
					break
				}
				if n, ok := g.Get(t.X-1, y); ok && n.Value.Kind == r.Kind {
					break
				}
				n.Value.Left = true
			}
		}
		// right
		if n, ok := g.Get(t.X+1, t.Y); !t.Value.Right && (!ok || n.Value.Kind != r.Kind) {
			total++
			// mark all tiles in same vertical line as visited
			for y := t.Y - 1; y >= 0; y-- {
				n, ok := g.Get(t.X, y)
				if !ok || n.Value.Kind != r.Kind {
					break
				}
				if n, ok := g.Get(t.X+1, y); ok && n.Value.Kind == r.Kind {
					break
				}
				n.Value.Right = true
			}
			for y := t.Y + 1; y < g.Height(); y++ {
				n, ok := g.Get(t.X, y)
				if !ok || n.Value.Kind != r.Kind {
					break
				}
				if n, ok := g.Get(t.X+1, y); ok && n.Value.Kind == r.Kind {
					break
				}
				n.Value.Right = true
			}
		}
	}
	return total
}

func part2(input string) int {
	g := grid.New[Flower](input, func(r rune) Flower {
		return Flower{Kind: string(r)}
	})

	var regions []Region
	for t := range g.All() {
		if t.Value.Visited {
			continue
		}
		regions = append(regions, group(g, t))
	}
	total := 0
	for _, r := range regions {
		total += r.Area() * r.Sides(g)
	}
	return total
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/12/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
