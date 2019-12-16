package utils

import (
	"bufio"
	"io"
	"math"
	"sort"
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

type Station struct {
	Position *Coordinate
	Size     int
	Viewport map[float64][]*Coordinate
}

type Distances struct {
	Asteroids []*Coordinate
	Origin    Coordinate
}

func (d Distances) Len() int {
	return len(d.Asteroids)
}
func (d Distances) Less(i, j int) bool {
	d1 := Distance(d.Origin, *d.Asteroids[i])
	d2 := Distance(d.Origin, *d.Asteroids[j])
	return d1 < d2
}
func (d Distances) Swap(i, j int) {
	d.Asteroids[i], d.Asteroids[j] = d.Asteroids[j], d.Asteroids[i]
}

func NthVaporized(station *Station, nth int) *Coordinate {

	angles := make([]float64, 0, len(station.Viewport))
	for k, _ := range station.Viewport {
		angles = append(angles, k)
	}

	// sort viewport by degrees (clockwise)
	sort.Float64s(angles)

	// start rotating and vaporized
	i := 1
	// rotate until all asteroids are vaporized
	for len(station.Viewport) != 0 {

		// rotate
		for _, angle := range angles {
			asteroids := station.Viewport[angle]
			if len(asteroids) == 0 {
				continue
			}

			asteroid := asteroids[0]
			if i == nth {
				return asteroid
			}

			// vaporized
			i++
			station.Viewport[angle] = asteroids[1:]

			if len(station.Viewport[angle]) == 0 {
				delete(station.Viewport, angle)
			}
		}

	}
	return nil
}

func Best(m Map) *Station {
	max := 0
	var best *Station
	for _, origin := range m.Asteroids {
		viewport := make(map[float64][]*Coordinate)
		for _, a := range m.Asteroids {
			if origin != a {
				angle := Angle(*origin, *a)
				viewport[angle] = append(viewport[angle], a)
			}
		}
		if len(viewport) > max {
			max = len(viewport)
			best = &Station{
				Position: origin,
				Size:     len(viewport),
				Viewport: viewport,
			}
		}
	}

	// sort each viewport angle by distance between origin and asteroids
	if best.Viewport != nil && best.Position != nil {
		for k, _ := range best.Viewport {
			sort.Sort(Distances{
				Asteroids: best.Viewport[k],
				Origin:    *best.Position,
			})
		}
	}

	return best
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
	r := math.Atan2(float64(c2.Y-c1.Y), float64(c2.X-c1.X)) * 180 / math.Pi
	if r < 0 {
		r += 360.0
	}
	return math.Mod(r+90, 360)
}

func Distance(c1, c2 Coordinate) float64 {
	return math.Sqrt(math.Pow(float64(c1.X-c2.X), 2) + math.Pow(float64(c1.Y-c2.Y), 2))
}
