package utils

import (
	"bufio"
	"fmt"
	"io"
	"math"

	"github.com/sirupsen/logrus"
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
		objects = append(objects, &Object{
			Position: [3]int{x, y, z},
		})
	}
	return objects, nil
}

// This is the tricky part
// My first implementation tried to look for periods for each objects/axis (12 periods to found)
// However, despite passing the first two tests, the results for the input doesn't work.
// Based on some hints over Reddit I tried to look for periods for each axis (3 periods to found), and it works.
//
// Once periods a found, using lcm of all periods is enough to know when all periods will occurred at the same time.
//
// I still don't know what the issue for my first implementation, I might look it up later.
func NextIdentity(objects []*Object) uint64 {
	ref := make([][]int, 3)
	for i := 0; i < 3; i++ {
		ref[i] = make([]int, 2*len(objects))
		for j, o := range objects {
			ref[i][2*j] = o.Position[i]
			ref[i][2*j+1] = o.Velocity[i]
		}
	}
	periods := make([]uint64, 3)

	iteration := uint64(0)
	for {
		Simulate(objects, 1)
		iteration++

		for i := 0; i < 3; i++ {
			if periods[i] != 0 {
				continue
			}
			found := true
			for j, o := range objects {
				if o.Position[i] != ref[i][2*j] || o.Velocity[i] != ref[i][2*j+1] {
					found = false
				}
			}
			if found {
				periods[i] = iteration
			}
		}

		found := true
		for _, p := range periods {
			if p == 0 {
				found = false
			}
		}

		if found {
			break
		}
	}

	logrus.Info(periods)
	return LCM(periods[0], periods[1], periods[2:]...)
}

// greatest common divisor (GCD) via Euclidean algorithm
// found on google : https://play.golang.org/p/SmzvkDjYlb
func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// find on google : https://play.golang.org/p/SmzvkDjYlb
func LCM(a, b uint64, integers ...uint64) uint64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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
