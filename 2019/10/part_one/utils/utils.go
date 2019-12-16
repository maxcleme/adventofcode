package utils

import (
	"bufio"
	"io"
	"math"
)

const (
	Asteroid = '#'
)

type Coordinate struct {
	X int
	Y int
}

type Map struct {
	Asteroids []*Coordinate
}

func Best(m Map) (*Coordinate, int) {
	max := 0
	var best *Coordinate
	for _, origin := range m.Asteroids {
		viewport := make(map[float64]bool)
		for _, a := range m.Asteroids {
			if origin != a {
				angle := Angle(*origin, *a)
				viewport[angle] = true
			}
		}
		if len(viewport) > max {
			max = len(viewport)
			best = origin
		}
	}
	return best, max
}

func Parse(reader io.Reader) (*Map, error) {
	m := Map{
		Asteroids: make([]*Coordinate, 0),
	}
	y := 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range []rune(line) {
			if c == Asteroid {
				m.Asteroids = append(m.Asteroids, &Coordinate{
					X: x,
					Y: y,
				})
			}
		}
		y++
	}
	return &m, nil
}

func Angle(c1, c2 Coordinate) float64 {
	return math.Atan2(float64(c2.Y-c1.Y), float64(c2.X-c1.X)) * 180 / math.Pi
}
