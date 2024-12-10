package _10

import (
	"strconv"
	"strings"
)

type Grid [][]int

func NewGrid(input string) Grid {
	var m [][]int
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		var r []int
		for _, c := range row {
			switch c {
			case '.':
				r = append(r, -1)
			default:
				n, err := strconv.Atoi(string(c))
				if err != nil {
					panic(err)
				}
				r = append(r, n)
			}
		}
		m = append(m, r)
	}
	return m
}

type Pos struct {
	X, Y int
}

func (g Grid) Get(x, y int) (int, bool) {
	if y < 0 || y >= len(g) {
		return 0, false
	}
	if x < 0 || x >= len(g[y]) {
		return 0, false
	}
	return g[y][x], true
}

func (g Grid) Find(fn func(p Pos) bool) []Pos {
	var res []Pos
	for y, row := range g {
		for x := range row {
			if fn(Pos{x, y}) {
				res = append(res, Pos{x, y})
			}
		}
	}
	return res
}
