package _21

import "github.com/maxcleme/adventofcode/grid"

var (
	numericKeypad = newKeypad(`789
456
123
 0A`)
	directionalKeypad = newKeypad(` ^A
<v>`)
)

type keypad struct {
	g     *grid.Grid[string]
	index map[string]*grid.Tile[string]
}

func newKeypad(input string) *keypad {
	g := grid.New[string](input, func(r rune) string {
		return string(r)
	})
	index := map[string]*grid.Tile[string]{}
	for t := range g.All() {
		index[t.Value] = t
	}
	return &keypad{g: g, index: index}
}

func (k *keypad) paths(from, to string) []string {
	paths := k.g.ShortestN(k.index[from], k.index[to], func(t *grid.Tile[string]) bool {
		return t.Value != " "
	})
	var result []string
	for _, path := range paths {
		s := ""
		for i := range path[:len(path)-1] {
			s += step(path[i], path[i+1])
		}
		result = append(result, s+"A")
	}
	return result
}

func step(from *grid.Tile[string], to *grid.Tile[string]) string {
	if from.X < to.X {
		return ">"
	}
	if from.X > to.X {
		return "<"
	}
	if from.Y < to.Y {
		return "v"
	}
	if from.Y > to.Y {
		return "^"
	}
	return ""
}
