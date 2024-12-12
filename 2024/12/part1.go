package _12

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

type Flower struct {
	Kind    string
	Visited bool

	// part 2
	Up, Down, Left, Right bool
}

func (f Flower) String() string {
	return f.Kind
}

type Region struct {
	Kind    string
	Flowers []*grid.Tile[Flower]
}

func (r Region) Area() int {
	return len(r.Flowers)
}

func (r Region) Perimeter(g *grid.Grid[Flower]) int {
	total := 0
	for _, t := range r.Flowers {
		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if n, ok := g.Get(t.X+dir[0], t.Y+dir[1]); !ok || n.Value.Kind != r.Kind {
				total++
			}
		}
	}
	return total
}

func part1(input string) int {
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
		total += r.Area() * r.Perimeter(g)
	}
	return total
}

func group(g *grid.Grid[Flower], t *grid.Tile[Flower]) Region {
	region := Region{Kind: t.Value.Kind}
	region.Flowers = append(region.Flowers, t)
	t.Value.Visited = true
	for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		if n, ok := g.Get(t.X+dir[0], t.Y+dir[1]); ok && !n.Value.Visited && n.Value.Kind == t.Value.Kind {
			region.Flowers = append(region.Flowers, group(g, n).Flowers...)
		}
	}
	return region
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/12/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
