package _14

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Type rune

const (
	SandSource = '+'

	Air  = '.'
	Rock = '#'
	Sand = 'o'
)

func (t Type) Solid() bool {
	return t == Sand || t == Rock
}

type Map struct {
	Tiles  [][]*Tile
	Source *Tile
}

func (m Map) String() string {
	var s string
	for _, row := range m.Tiles {
		for _, t := range row {
			s += string(t.Type)
		}
		s += "\n"
	}
	return s
}

func (m Map) Down(t *Tile) *Tile {
	if t == nil {
		return nil
	}
	return m.Get(Position{t.RelPosition.X, t.RelPosition.Y + 1})
}

func (m Map) Left(t *Tile) *Tile {
	if t == nil {
		return nil
	}
	return m.Get(Position{t.RelPosition.X - 1, t.RelPosition.Y})
}

func (m Map) Right(t *Tile) *Tile {
	if t == nil {
		return nil
	}
	return m.Get(Position{t.RelPosition.X + 1, t.RelPosition.Y})
}

func (m Map) Get(p Position) *Tile {
	if p.Y < 0 || p.Y >= len(m.Tiles) {
		return nil
	}
	if p.X < 0 || p.X >= len(m.Tiles[0]) {
		return nil
	}
	return m.Tiles[p.Y][p.X]
}

func (m Map) Fall(src *Tile) *Tile {
	if src == nil {
		return nil
	}
	last := src
	for {
		next := m.Down(last)
		if next == nil {
			return nil
		}
		if next.Type.Solid() {
			dl, dr := m.Left(next), m.Right(next)
			if dl == nil || dr == nil {
				return nil
			}
			if dl.Type.Solid() && dr.Type.Solid() {
				return last
			}
			if !dl.Type.Solid() {
				return m.Fall(m.Left(last))
			}
			if !dr.Type.Solid() {
				return m.Fall(m.Right(last))
			}
		}
		last = next
	}
}

func (m Map) Tick() bool {
	if m.Source.Type == Sand {
		return false
	}

	next := m.Fall(m.Source)
	if next == nil {
		return false
	}
	next.Type = Sand
	return true
}

type Tile struct {
	RelPosition Position
	AbsPosition Position
	Type        Type
}

type Position struct {
	X, Y int
}

type Path []Position

func ParsePosition(l string) Position {
	parts := strings.Split(l, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Position{X: x, Y: y}
}

func ParsePath(l string) Path {
	parts := strings.Split(l, " -> ")
	var path Path
	for _, p := range parts {
		path = append(path, ParsePosition(p))
	}
	return path
}

func MakeMap(rocks []Path, source Position, floor bool) Map {
	var positions []Position
	for _, path := range rocks {
		for _, p := range path {
			positions = append(positions, p)
		}
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].X < positions[j].X
	})
	minX, maxX := positions[0].X, positions[len(positions)-1].X
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].Y < positions[j].Y
	})
	minY, maxY := positions[0].Y, positions[len(positions)-1].Y
	if floor {
		maxY += 2
		minX = minX - maxY
		maxX = maxX + maxY
	}

	var tiles [][]*Tile
	for y := 0; y <= maxY; y++ {
		tiles = append(tiles, []*Tile{})
		for x := 0; x <= maxX-minX; x++ {
			typ := Air
			if floor && y == maxY {
				typ = Rock
			}
			tiles[y] = append(tiles[y], &Tile{
				RelPosition: Position{X: x, Y: y},
				AbsPosition: Position{X: x + minX, Y: y + minY},
				Type:        Type(typ),
			})
		}
	}
	for _, path := range rocks {
		if len(path) == 1 {

		}
		for i := 0; i < len(path)-1; i++ {
			from, to := path[i], path[i+1]
			startX, startY := int(math.Min(float64(from.X), float64(to.X))), int(math.Min(float64(from.Y), float64(to.Y)))
			endX, endY := int(math.Max(float64(from.X), float64(to.X))), int(math.Max(float64(from.Y), float64(to.Y)))
			for y := startY; y <= endY; y++ {
				for x := startX - minX; x <= endX-minX; x++ {
					tiles[y][x].Type = Rock
				}
			}
		}
	}
	tiles[source.Y][source.X-minX].Type = SandSource
	return Map{Tiles: tiles, Source: tiles[source.Y][source.X-minX]}
}
