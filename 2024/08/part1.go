package _08

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

type AntennaTile struct {
	Symbol   string
	Antinode bool
}

func (t AntennaTile) String() string {
	if t.Antinode {
		return "#"
	}
	return t.Symbol
}

func isAntenna(t *grid.Tile[AntennaTile]) bool {
	return t.Value.Symbol != ""
}

func part1(input string) int {
	g := grid.New[AntennaTile](input, func(r rune) AntennaTile {
		return AntennaTile{Symbol: string(r)}
	})
	antennas := map[string][]*grid.Tile[AntennaTile]{}
	for _, t := range g.Find(isAntenna) {
		antennas[t.Value.Symbol] = append(antennas[t.Value.Symbol], t)
	}
	for symbol, antennas := range antennas {
		if symbol == "." || len(antennas) == 1 {
			continue
		}
		for i, a1 := range antennas[:len(antennas)-1] {
			for _, a2 := range antennas[i+1:] {
				x1, y1, x2, y2 := computeAntinode(a1, a2)
				if antinode1, ok := g.Get(x1, y1); ok {
					antinode1.Value.Antinode = true
				}
				if antinode2, ok := g.Get(x2, y2); ok {
					antinode2.Value.Antinode = true
				}
			}
		}
	}
	fmt.Println(g)
	return len(g.Find(func(t *grid.Tile[AntennaTile]) bool {
		return t.Value.Antinode
	}))
}

func computeAntinode(a1, a2 *grid.Tile[AntennaTile]) (int, int, int, int) {
	return a2.X + (a2.X - a1.X), a2.Y + (a2.Y - a1.Y), a1.X + (a1.X - a2.X), a1.Y + (a1.Y - a2.Y)
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/08/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
