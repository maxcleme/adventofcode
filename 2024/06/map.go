package _06

import "strings"

type Map struct {
	tiles   [][]*Tile
	history map[*Tile]map[Direction]struct{}
}

func NewMap(input string) (*Map, int, int) {
	rows := strings.Split(input, "\n")
	m := Map{}
	xStart, yStart := 0, 0
	for y, row := range rows {
		var r []*Tile
		for x, c := range row {
			switch TileKind(c) {
			case Empty, TileKind('^'):
				r = append(r, &Tile{kind: Empty})
				if c == '^' {
					xStart, yStart = x, y
				}
			case Wall:
				r = append(r, &Tile{kind: Wall})
			}
		}
		m.tiles = append(m.tiles, r)
	}
	return &m, xStart, yStart
}

func (m *Map) Get(x, y int) (bool, *Tile) {
	if y < 0 || y >= len(m.tiles) {
		return false, nil
	}
	if x < 0 || x >= len(m.tiles[y]) {
		return false, nil
	}
	return true, m.tiles[y][x]
}

func (m *Map) Visit(x, y int, d Direction) {
	_, tile := m.Get(x, y)
	if m.history == nil {
		m.history = map[*Tile]map[Direction]struct{}{}
	}
	if _, ok := m.history[tile]; !ok {
		m.history[tile] = map[Direction]struct{}{}
	}
	m.history[tile][d] = struct{}{}
}

var tileID = 0

func makeTileID() int {
	tileID++
	return tileID
}

type Tile struct {
	kind TileKind
}

type TileKind string

const (
	Empty = TileKind('.')
	Wall  = TileKind('#')
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) TurnClockwise() Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	default:
		return Up
	}
}

func (d Direction) Next(x, y int) (int, int) {
	switch d {
	case Up:
		return x, y - 1
	case Right:
		return x + 1, y
	case Down:
		return x, y + 1
	default:
		return x - 1, y
	}
}
