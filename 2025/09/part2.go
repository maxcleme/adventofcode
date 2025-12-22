package _09

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/09/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func (p Pair) Contains(t [2]int) bool {
	return (p.A[0] == t[0] && p.A[1] == t[1]) || (p.B[0] == t[0] && p.B[1] == t[1])
}

type Segment struct {
	From, To [2]int
}

func (l Segment) Contains(t [2]int) bool {
	if l.From[0] == l.To[0] {
		return t[0] == l.From[0] && min(l.From[1], l.To[1]) <= t[1] &&
			t[1] <= max(l.From[1], l.To[1])
	}
	return t[1] == l.From[1] && min(l.From[0], l.To[0]) <= t[0] &&
		t[0] <= max(l.From[0], l.To[0])
}

type Rectangle struct {
	X1, Y1, X2, Y2 int
}

type Polygon []Segment

func (p Polygon) ContainsRectangle(r Rectangle) bool {
	// all rectangles corners inside polygon
	minX := min(r.X1, r.X2)
	maxX := max(r.X1, r.X2)
	minY := min(r.Y1, r.Y2)
	maxY := max(r.Y1, r.Y2)
	topLeft := [2]int{minX, minY}
	topRight := [2]int{maxX, minY}
	bottomLeft := [2]int{minX, maxY}
	bottomRight := [2]int{maxX, maxY}
	if !(p.Contains(topLeft) &&
		p.Contains(topRight) &&
		p.Contains(bottomLeft) &&
		p.Contains(bottomRight)) {
		return false
	}
	// no polygon edges intersect rectangle edges
	// as even if all corners are inside, edges could cross outside
	rect := []Segment{
		{From: topLeft, To: topRight},
		{From: topRight, To: bottomRight},
		{From: bottomRight, To: bottomLeft},
		{From: bottomLeft, To: topLeft},
	}
	for _, re := range rect {
		for _, pe := range p {
			if intersect(re, pe) {
				return false
			}
		}
	}
	return true
}

func (p Polygon) Contains(t [2]int) bool {
	// Incredibly efficient
	// https://en.wikipedia.org/wiki/Point_in_polygon
	var inside bool
	for _, segment := range p {
		if segment.Contains(t) {
			return true
		}
		if segment.From[0] != segment.To[0] {
			continue
		}
		x := segment.From[0]
		if x <= t[0] {
			continue
		}
		yMin := min(segment.From[1], segment.To[1])
		yMax := max(segment.From[1], segment.To[1])
		if yMin <= t[1] && t[1] < yMax {
			inside = !inside
		}
	}
	return inside
}

func intersect(a, b Segment) bool {
	var v, h Segment
	if isVertical(a) && isHorizontal(b) {
		v, h = a, b
	} else if isHorizontal(a) && isVertical(b) {
		v, h = b, a
	} else {
		return false
	}
	xv := v.From[0]
	yh := h.From[1]
	vMinY := min(v.From[1], v.To[1])
	vMaxY := max(v.From[1], v.To[1])
	hMinX := min(h.From[0], h.To[0])
	hMaxX := max(h.From[0], h.To[0])
	return hMinX < xv && xv < hMaxX && vMinY < yh && yh < vMaxY
}

func isVertical(s Segment) bool   { return s.From[0] == s.To[0] }
func isHorizontal(s Segment) bool { return s.From[1] == s.To[1] }

func part2(payload string) int {
	var tiles [][2]int
	var maxX, maxY int
	for _, line := range strings.Split(payload, "\n") {
		positions := strings.Split(line, ",")
		x, _ := strconv.Atoi(positions[0])
		y, _ := strconv.Atoi(positions[1])
		tiles = append(tiles, [2]int{x, y})
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	var p Polygon
	last := tiles[len(tiles)-1]
	for _, t := range tiles {
		p = append(p, Segment{From: last, To: t})
		last = t
	}
	pairs := distances(tiles)
	slices.SortFunc(pairs, func(a, b Pair) int {
		return b.Area - a.Area
	})
	for _, pair := range pairs {
		if !p.ContainsRectangle(Rectangle{X1: pair.A[0], Y1: pair.A[1], X2: pair.B[0], Y2: pair.B[1]}) {
			continue
		}
		return pair.Area
	}
	return -1
}
