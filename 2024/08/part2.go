package _08

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (m *Map) MarkAntinodes(fn func(int, int) bool) {
	for _, row := range m.Tiles {
		for _, t := range row {
			if fn(t.X, t.Y) {
				t.Antinode = true
			}
		}
	}
}

func findAntinodeFunc(a1 *Tile, a2 *Tile) func(int, int) bool {
	return func(x, y int) bool {
		// Compute the cross product of the vectors (a1, a2) and (a1, (x, y))
		return (y-a1.Y)*(a2.X-a1.X)-(x-a1.X)*(a2.Y-a1.Y) == 0
	}
}

func part2(input string) int {
	m := NewMake(input)
	for _, antennas := range m.Antennas {
		if len(antennas) == 1 {
			continue
		}
		for i, a1 := range antennas[:len(antennas)-1] {
			for _, a2 := range antennas[i+1:] {
				fn := findAntinodeFunc(a1, a2)
				m.MarkAntinodes(fn)
			}
		}
	}
	fmt.Println(m)
	return m.CountAntinode()
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
