package grid

import (
	"fmt"
	"iter"
	"math"
	"strings"

	"github.com/maxcleme/adventofcode/queue"
)

type Tile[T any] struct {
	X, Y  int
	Value T
}

type Grid[T any] struct {
	layout [][]*Tile[T]
}

func Empty[T any](width, height int) *Grid[T] {
	g := &Grid[T]{}
	for y := 0; y < height; y++ {
		var r []*Tile[T]
		for x := 0; x < width; x++ {
			r = append(r, &Tile[T]{X: x, Y: y, Value: *new(T)})
		}
		g.layout = append(g.layout, r)
	}
	return g
}

func New[T any](input string, fn func(rune) T) *Grid[T] {
	g := &Grid[T]{}
	rows := strings.Split(input, "\n")
	for y, row := range rows {
		var r []*Tile[T]
		for x, c := range row {
			r = append(r, &Tile[T]{X: x, Y: y, Value: fn(c)})
		}
		g.layout = append(g.layout, r)
	}
	return g
}

func (g Grid[T]) Get(x, y int) (*Tile[T], bool) {
	if y < 0 || y >= len(g.layout) {
		return nil, false
	}
	if x < 0 || x >= len(g.layout[y]) {
		return nil, false
	}
	return g.layout[y][x], true
}

func (g Grid[T]) Find(fn func(*Tile[T]) bool) []*Tile[T] {
	var res []*Tile[T]
	for y, row := range g.layout {
		for x := range row {
			if t, _ := g.Get(x, y); fn(t) {
				res = append(res, t)
			}
		}
	}
	return res
}

func (g Grid[T]) Map(fn func(t *Tile[T])) {
	for _, row := range g.layout {
		for _, t := range row {
			fn(t)
		}
	}
}

func (g Grid[T]) All() iter.Seq[*Tile[T]] {
	return func(yield func(*Tile[T]) bool) {
		for y, row := range g.layout {
			for x := range row {
				if t, _ := g.Get(x, y); !yield(t) {
					return
				}
			}
		}
	}
}

func (g Grid[T]) Height() int {
	return len(g.layout)
}

func (g Grid[T]) Width() int {
	if len(g.layout) == 0 {
		return 0
	}
	return len(g.layout[0])
}

func (g Grid[T]) Shortest(from, to *Tile[T], walk func(*Tile[T]) bool) (int, bool) {
	var costs [][]int
	for range g.layout {
		var c []int
		for range g.layout[0] {
			c = append(c, math.MaxInt)
		}
		costs = append(costs, c)
	}
	costs[from.Y][from.X] = 0
	q := queue.New[*Tile[T]]()
	q.Push(from)
	for q.Len() > 0 {
		curr := q.Pop()
		if curr == to {
			return costs[to.Y][to.X], true
		}
		directions := []struct{ X, Y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
		for _, dir := range directions {
			nextX, nextY := curr.X+dir.X, curr.Y+dir.Y
			next, ok := g.Get(nextX, nextY)
			if !ok || !walk(next) {
				continue
			}
			newCost := costs[curr.Y][curr.X] + 1
			if newCost < costs[nextY][nextX] {
				costs[nextY][nextX] = newCost
				q.Push(next)
			}
		}
	}
	return 0, false
}

func (g Grid[T]) String() string {
	var s strings.Builder
	for _, row := range g.layout {
		for _, t := range row {
			s.WriteString(fmt.Sprintf("%s", t.Value))
		}
		s.WriteByte('\n')
	}
	return s.String()
}
