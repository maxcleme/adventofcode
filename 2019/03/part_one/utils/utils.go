package utils

import (
	"math"
	"sort"
	"strconv"
)

const (
	UP    = 'U'
	RIGHT = 'R'
	DOWN  = 'D'
	LEFT  = 'L'

	ErrParseSegment   = Error("cannot parse segment")
	ErrNoIntersection = Error("no intersection found")

	LayoutSize = 30000
)

type Segment struct {
	Orientation int
	Length      int
}

type Wire struct {
	Segments []Segment
}

type Error string

func (e Error) Error() string {
	return string(e)
}

func NewSegment(stringSegment string) (*Segment, error) {
	if len(stringSegment) < 2 {
		return nil, ErrParseSegment
	}
	o := stringSegment[0]
	length, err := strconv.Atoi(stringSegment[1:])
	if err != nil {
		return nil, ErrParseSegment
	}
	switch o {
	case UP:
		return &Segment{UP, length}, nil
	case RIGHT:
		return &Segment{RIGHT, length}, nil
	case DOWN:
		return &Segment{DOWN, length}, nil
	case LEFT:
		return &Segment{LEFT, length}, nil
	default:
		return nil, ErrParseSegment
	}
}

func NewWire(stringSegments []string) (*Wire, error) {
	segments := make([]Segment, len(stringSegments))

	for i, segment := range stringSegments {
		s, err := NewSegment(segment)
		if err != nil {
			return nil, err
		}
		segments[i] = *s
	}

	return &Wire{segments}, nil
}

func ClosestManhattanDistance(stringWire1, stringWire2 []string) (int, error) {
	wire1, err := NewWire(stringWire1)
	if err != nil {
		return 0, err
	}
	wire2, err := NewWire(stringWire2)
	if err != nil {
		return 0, err
	}

	// create distances slice (will store distance of each intersections)
	distances := make([]int, 0)

	// use an arbitrary size of 10000*10000
	layout := make([][]bool, LayoutSize)
	for i, _ := range layout {
		layout[i] = make([]bool, LayoutSize)
	}

	// start from center of layout for first wire
	x, y := len(layout)/2, len(layout)/2
	for _, segment := range wire1.Segments {
		newX, newY, _ := Trace(layout, x, y, segment)
		x, y = newX, newY
	}

	// start from center of layout for second wire
	x, y = len(layout)/2, len(layout)/2
	for _, segment := range wire2.Segments {
		newX, newY, intersectionDistances := Trace(layout, x, y, segment)
		x, y = newX, newY
		distances = append(distances, intersectionDistances...)
	}

	sort.Ints(distances)
	if len(distances) == 0 {
		return 0, ErrNoIntersection
	}
	return distances[0], nil
}

func Trace(layout [][]bool, x, y int, segment Segment) (int, int, []int) {
	// create distances slice (will store distance of each intersections)
	distances := make([]int, 0)

	xOffset := 0
	yOffset := 0
	switch segment.Orientation {
	case UP:
		yOffset = 1
	case RIGHT:
		xOffset = 1
	case DOWN:
		yOffset = -1
	case LEFT:
		xOffset = -1
	}

	for i := 0; i < segment.Length; i++ {
		checkX, checkY := x+i*xOffset, y+i*yOffset
		if layout[checkX][checkY] {
			// intersection found
			distance := math.Abs(float64(len(layout)/2-checkX)) + math.Abs(float64(len(layout)/2-checkY))
			// intersection at origin is ignored
			if distance != 0 {
				distances = append(distances, int(distance))
			}
		} else {
			layout[checkX][checkY] = true
		}
	}

	return x + segment.Length*xOffset, y + segment.Length*yOffset, distances
}
