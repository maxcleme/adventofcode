package _08

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func part1(input string) int {
	m := NewMake(input)
	for _, antennas := range m.Antennas {
		if len(antennas) == 1 {
			continue
		}
		for i, a1 := range antennas[:len(antennas)-1] {
			for _, a2 := range antennas[i+1:] {
				x1, y1, x2, y2 := findAntinode(a1, a2)
				m.MarkAntinode(x1, y1)
				m.MarkAntinode(x2, y2)
			}
		}
	}
	fmt.Println(m)
	return m.CountAntinode()
}

func findAntinode(a1 *Tile, a2 *Tile) (int, int, int, int) {
	return a2.X + (a2.X - a1.X), a2.Y + (a2.Y - a1.Y), a1.X + (a1.X - a2.X), a1.Y + (a1.Y - a2.Y)
}

type Map struct {
	Tiles    [][]*Tile
	Antennas map[string][]*Tile
}

func (m *Map) String() string {
	var s strings.Builder
	for _, row := range m.Tiles {
		for _, t := range row {
			if t.Antinode {
				s.WriteByte('#')
			} else if t.Antenna != "" {
				s.WriteString(t.Antenna)
			} else {
				s.WriteByte('.')
			}
		}
		s.WriteByte('\n')
	}
	return s.String()
}

func (m *Map) MarkAntinode(x int, y int) {
	if y < 0 || y >= len(m.Tiles) {
		return
	}
	if x < 0 || x >= len(m.Tiles[y]) {
		return
	}
	m.Tiles[y][x].Antinode = true
}

func (m *Map) CountAntinode() int {
	count := 0
	for _, row := range m.Tiles {
		for _, t := range row {
			if t.Antinode {
				count++
			}
		}
	}
	return count
}

type Tile struct {
	X, Y     int
	Antenna  string
	Antinode bool
}

func NewMake(input string) *Map {
	m := &Map{Antennas: map[string][]*Tile{}}
	for y, row := range strings.Split(input, "\n") {
		var r []*Tile
		for x, c := range row {
			t := &Tile{X: x, Y: y}
			if c != '.' {
				t.Antenna = string(c)
				m.Antennas[string(c)] = append(m.Antennas[string(c)], t)
			}
			r = append(r, t)
		}
		m.Tiles = append(m.Tiles, r)
	}
	return m
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
