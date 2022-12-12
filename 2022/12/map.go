package _12

import (
	"fmt"
)

type Tile struct {
	Position  Position
	Key       string
	Elevation int
}

type Position struct {
	X int
	Y int
}

func (p Position) Equal(p2 Position) bool {
	return p.X == p2.X && p.Y == p2.Y
}

type Map struct {
	Tiles  [][]*Tile
	Source *Tile
	Target *Tile
}

func MakeMap() Map {
	return Map{}
}

func (m Map) AppendRow(row string) Map {
	var elevations []*Tile
	for i, s := range row {
		x, y := i, len(m.Tiles)
		p := Position{X: i, Y: len(m.Tiles)}
		tile := &Tile{
			Key:      fmt.Sprintf("%d#%d", y, x),
			Position: p,
		}
		switch s {
		case 'S':
			tile.Elevation = 'a'
			m.Source = tile
		case 'E':
			tile.Elevation = 'z'
			m.Target = tile
		default:
			tile.Elevation = int(s)
		}
		elevations = append(elevations, tile)
	}
	m.Tiles = append(m.Tiles, elevations)
	return m
}

func (m Map) Get(p Position) *Tile {
	return m.Tiles[p.Y][p.X]
}

func (m Map) All() []*Tile {
	var all []*Tile
	for i, _ := range m.Tiles {
		for j, _ := range m.Tiles[i] {
			all = append(all, m.Tiles[i][j])
		}
	}
	return all
}

func (m Map) Up(t *Tile) *Tile {
	if t == nil {
		return nil
	}
	if t.Position.Y == 0 {
		return nil
	}
	next := m.Get(Position{t.Position.X, t.Position.Y - 1})
	if next.Elevation > t.Elevation+1 {
		return nil
	}
	return next
}
func (m Map) Down(t *Tile) *Tile {
	if t == nil {
		return nil
	}
	if t.Position.Y >= len(m.Tiles)-1 {
		return nil
	}
	next := m.Get(Position{t.Position.X, t.Position.Y + 1})
	if next.Elevation > t.Elevation+1 {
		return nil
	}
	return next
}
func (m Map) Left(t *Tile) *Tile {
	if t == nil {
		return nil
	}
	if t.Position.X == 0 {
		return nil
	}
	next := m.Get(Position{t.Position.X - 1, t.Position.Y})
	if next.Elevation > t.Elevation+1 {
		return nil
	}
	return next
}
func (m Map) Right(t *Tile) *Tile {
	if t == nil {
		return nil
	}
	if t.Position.X >= len(m.Tiles[0])-1 {
		return nil
	}
	next := m.Get(Position{t.Position.X + 1, t.Position.Y})
	if next.Elevation > t.Elevation+1 {
		return nil
	}
	return next
}

func (m Map) Neighbours(t *Tile) []*Tile {
	if t == nil {
		return nil
	}
	var neighbours []*Tile
	if next := m.Up(t); next != nil {
		neighbours = append(neighbours, next)
	}
	if next := m.Right(t); next != nil {
		neighbours = append(neighbours, next)
	}
	if next := m.Down(t); next != nil {
		neighbours = append(neighbours, next)
	}
	if next := m.Left(t); next != nil {
		neighbours = append(neighbours, next)
	}
	return neighbours
}

func (m Map) Find(startCandidate func(t *Tile) bool) []Position {
	// Basically a c/c from https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm

	all := m.All()
	index := make(map[string]*Tile)
	dist := make(map[string]int)
	prev := make(map[string]*Tile)
	for _, t := range all {
		dist[t.Key] = 1_000_000
		prev[t.Key] = nil
		index[t.Key] = t
		if startCandidate(t) {
			dist[t.Key] = 0
		}
	}

	for len(index) > 0 {
		var u *Tile
		for _, t := range index {
			if u == nil || dist[t.Key] < dist[u.Key] {
				u = t
			}
		}

		if u.Position.Equal(m.Target.Position) {
			u = m.Target
			var path []Position
			if _, ok := prev[u.Key]; ok || u.Position.Equal(m.Source.Position) {
				for u != nil {
					path = append([]Position{u.Position}, path...)
					u = prev[u.Key]
				}
				return path
			}
		}

		delete(index, u.Key)
		for _, next := range m.Neighbours(u) {
			if _, ok := index[next.Key]; !ok {
				continue
			}

			alt := dist[u.Key] + 1
			if alt < dist[next.Key] {
				dist[next.Key] = alt
				prev[next.Key] = u
			}
		}
	}
	return nil
}
