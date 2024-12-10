package _04

import (
	"fmt"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/04/input")
		if err != nil {
			return err
		}
		fmt.Println(countMASPart2(string(f)))
		return nil
	},
}

func countMASPart2(input string) int {
	g := grid.New[rune](input, func(r rune) rune {
		return r
	})
	centers := g.Find(func(t *grid.Tile[rune]) bool {
		return t.Value == 'A'
	})
	count := 0
	for _, center := range centers {
		topLeft, ok := g.Get(center.X-1, center.Y-1)
		if !ok {
			continue
		}
		topRight, ok := g.Get(center.X+1, center.Y-1)
		if !ok {
			continue
		}
		bottomLeft, ok := g.Get(center.X-1, center.Y+1)
		if !ok {
			continue
		}
		bottomRight, ok := g.Get(center.X+1, center.Y+1)
		if !ok {
			continue
		}
		a := string(topLeft.Value) + string(center.Value) + string(bottomRight.Value)
		b := string(topRight.Value) + string(center.Value) + string(bottomLeft.Value)
		if (a == "MAS" || a == "SAM") && (b == "MAS" || b == "SAM") {
			count++
		}
	}
	return count
}
