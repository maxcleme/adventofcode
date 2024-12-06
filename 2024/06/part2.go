package _06

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func part2(input string) int {
	m, xStart, yStart := NewMap(input)
	walk(m, xStart, yStart)
	loop := 0
	for tile := range m.history {
		m.history = map[*Tile]map[Direction]struct{}{}
		oldKind := tile.kind
		tile.kind = Wall
		if walk(m, xStart, yStart) {
			loop++
		}
		tile.kind = oldKind
	}
	return loop
}

func walk(m *Map, startX, startY int) bool {
	d := Up
	x, y := startX, startY
	m.Visit(x, y, d)
	for {
		xNext, yNext := d.Next(x, y)
		ok, pos := m.Get(xNext, yNext)
		if !ok {
			break
		}
		if pos.kind == Wall {
			d = d.TurnClockwise()
			continue
		}
		x, y = xNext, yNext
		if m.history != nil && m.history[pos] != nil {
			if _, ok := m.history[pos][d]; ok {
				return true
			}
		}
		m.Visit(x, y, d)
	}
	return false
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/06/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
