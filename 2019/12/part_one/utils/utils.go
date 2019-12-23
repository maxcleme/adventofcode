package utils

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

type Object struct {
	Position [3]int
	Velocity [3]int
}

func (o Object) Potential() int {
	p := 0
	for v := 0; v < 3; v++ {
		p += int(math.Abs(float64(o.Position[v])))
	}
	return p
}

func (o Object) Kinetic() int {
	p := 0
	for v := 0; v < 3; v++ {
		p += int(math.Abs(float64(o.Velocity[v])))
	}
	return p
}

func (o Object) Total() int {
	return o.Potential() * o.Kinetic()
}

func Parse(reader io.Reader) ([]*Object, error) {
	objects := make([]*Object, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		x, y, z := 0, 0, 0
		if _, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z); err != nil {
			return nil, err
		}
		objects = append(objects, &Object{Position: [3]int{x, y, z}})
	}
	return objects, nil
}

func Simulate(objects []*Object, nth int) []*Object {
	for n := 1; n <= nth; n++ {
		for i := 0; i < len(objects)-1; i++ {
			for j := i + 1; j < len(objects); j++ {
				a := objects[i]
				b := objects[j]
				for v := 0; v < 3; v++ {
					// apply gravity
					if a.Position[v] > b.Position[v] {
						a.Velocity[v]--
						b.Velocity[v]++
					} else if a.Position[v] < b.Position[v] {
						a.Velocity[v]++
						b.Velocity[v]--
					}
				}
			}
		}

		for _, o := range objects {
			for v := 0; v < 3; v++ {
				// apply velocity
				o.Position[v] += o.Velocity[v]
			}
		}
	}
	return objects
}
