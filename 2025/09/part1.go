package _09

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
		f, err := os.ReadFile("./2025/09/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

type Pair struct {
	A, B [2]int
	Area int
}

func part1(payload string) int {
	var tiles [][2]int
	for _, line := range strings.Split(payload, "\n") {
		positions := strings.Split(line, ",")
		x, _ := strconv.Atoi(positions[0])
		y, _ := strconv.Atoi(positions[1])
		tiles = append(tiles, [2]int{x, y})
	}
	pairs := distances(tiles)
	slices.SortFunc(pairs, func(a, b Pair) int {
		return b.Area - a.Area
	})
	return pairs[0].Area
}

func distances(tiles [][2]int) []Pair {
	var pairs []Pair
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			pairs = append(pairs, Pair{
				A:    tiles[i],
				B:    tiles[j],
				Area: area(tiles[i], tiles[j]),
			})
		}
	}
	return pairs
}

func area(a, b [2]int) int {
	return (int(math.Abs(float64(a[0]-b[0]))) + 1) * (int(math.Abs(float64(a[1]-b[1]))) + 1)
}
