package _15

import "fmt"

type Position struct {
	X, Y int
}

func (p Position) TuningFrequency() int {
	return p.X*4000000 + p.Y
}

type Sensor struct {
	Position
	Beacon Position
}

func (s Sensor) Bounds() []Position {
	d := Distance(s.Position, s.Beacon)
	return []Position{
		{s.X - d, s.Y},
		{s.X + d, s.Y},
		{s.X, s.Y - d},
		{s.X, s.Y + d},
	}
}

func MustParseSensor(l string) Sensor {
	s := Sensor{}
	_, err := fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.Position.X, &s.Position.Y, &s.Beacon.X, &s.Beacon.Y)
	if err != nil {
		panic(err)
	}
	return s
}
