package _08

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
		f, err := os.ReadFile("./2025/08/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(payload string) int {
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
	for _, p := range pairs {
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
		if len(circuits) == 1 && len(circuits[0]) == len(boxes) {
			return a.X * b.X
		}
	}
	return -1
}
