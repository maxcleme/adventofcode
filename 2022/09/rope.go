package _09

import (
	"math"
	"strconv"
	"strings"
)

type Direction string

func (m Move) Apply(r Rope, callback func(Rope)) Rope {
	for i := 0; i < m.Count; i++ {
		switch m.Direction {
		case "L":
			r = r.Follow(Position{x: r[0].x - 1, y: r[0].y})
		case "R":
			r = r.Follow(Position{x: r[0].x + 1, y: r[0].y})
		case "U":
			r = r.Follow(Position{x: r[0].x, y: r[0].y + 1})
		case "D":
			r = r.Follow(Position{x: r[0].x, y: r[0].y - 1})
		}
		callback(r)
	}
	return r
}

type Move struct {
	Direction Direction
	Count     int
}

type Rope []Position

func (r Rope) Head() Position {
	return r[0]
}

func (r Rope) Tail() Position {
	return r[len(r)-1]
}

type Position struct {
	x, y int
}

func MakeRope(length int) Rope {
	return make([]Position, length)
}

func ParseMove(s string) *Move {
	parts := strings.Split(s, " ")
	count, _ := strconv.Atoi(parts[1])
	return &Move{
		Direction: Direction(parts[0]),
		Count:     count,
	}
}

func (r Rope) Follow(p Position) Rope {
	r[0] = p
	for i := 1; i < len(r); i++ {
		r[i] = r[i].Follow(r[i-1])
	}
	return r
}

func (p Position) Follow(head Position) Position {
	accX, accY := float64(head.x-p.x), float64(head.y-p.y)
	linear := math.Abs(accX) == math.Abs(accY)
	if math.Abs(accX) > 1 {
		p.x += int(accX - math.Copysign(1, accX))
		if !linear && math.Abs(accY) > 0 {
			p.y += int(accY)
		}
	}
	if math.Abs(accY) > 1 {
		p.y += int(accY - math.Copysign(1, accY))
		if !linear && math.Abs(accX) > 0 {
			p.x += int(accX)
		}
	}
	return p
}
