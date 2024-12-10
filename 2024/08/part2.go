package _08

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

func part2(input string) int {
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
				g.Map(func(t *grid.Tile[AntennaTile]) {
					// Compute the area of triangle formed by the 3 points, if the area is 0, the point is on the line
					if (t.Y-a1.Y)*(a2.X-a1.X)-(t.X-a1.X)*(a2.Y-a1.Y) == 0 {
						t.Value.Antinode = true
					}
				})
			}
		}
	}
	fmt.Println(g)
	return len(g.Find(func(t *grid.Tile[AntennaTile]) bool {
		return t.Value.Antinode
	}))
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/08/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
