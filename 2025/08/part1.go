package _08

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/08/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f), 1000))
		return nil
	},
}

type Box struct {
	ID      int
	X, Y, Z int
}

type Pair struct {
	A, B     Box
	Distance int
}

type Circuit map[int]Box

func part1(payload string, step int) int {
	var boxes []Box
	lines := strings.Split(payload, "\n")
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		boxes = append(boxes, Box{ID: i, X: x, Y: y, Z: z})
	}

	pairs := distances(boxes)
	slices.SortFunc(pairs, func(a, b Pair) int {
		return a.Distance - b.Distance
	})
	var circuits []Circuit
	for _, p := range pairs[:step] {
		a, b := p.A, p.B
		clusterIdxA := slices.IndexFunc(circuits, func(c Circuit) bool {
			_, ok := c[a.ID]
			return ok
		})
		clusterIdxB := slices.IndexFunc(circuits, func(c Circuit) bool {
			_, ok := c[b.ID]
			return ok
		})
		if clusterIdxA != -1 && clusterIdxA == clusterIdxB {
			continue
		}
		switch {
		case clusterIdxA == -1 && clusterIdxB == -1:
			circuits = append(circuits, map[int]Box{a.ID: a, b.ID: b})
		case clusterIdxA != -1 && clusterIdxB == -1:
			circuits[clusterIdxA][b.ID] = b
		case clusterIdxA == -1 && clusterIdxB != -1:
			circuits[clusterIdxB][a.ID] = a
		case clusterIdxA != -1 && clusterIdxB != -1 && clusterIdxA != clusterIdxB:
			for id, box := range circuits[clusterIdxB] {
				circuits[clusterIdxA][id] = box
			}
			circuits = slices.Delete(circuits, clusterIdxB, clusterIdxB+1)
		}
	}
	slices.SortFunc(circuits, func(a, b Circuit) int {
		return len(b) - len(a)
	})
	total := 1
	for _, c := range circuits[:3] {
		total *= len(c)
	}
	return total
}

func distances(boxes []Box) []Pair {
	var pairs []Pair
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			a, b := boxes[i], boxes[j]
			pairs = append(pairs, Pair{A: a, B: b, Distance: int(distance(a, b))})
		}
	}
	return pairs
}

func distance(a, b Box) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2))
}
