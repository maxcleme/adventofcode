package _25

import (
	"fmt"
	"os"
	"strings"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

type pins [5]int

func part1(input string) int {
	parts := strings.Split(input, "\n\n")

	var locks []pins
	var keys []pins
	for _, part := range parts {
		g := grid.New(part, func(r rune) string {
			return string(r)
		})
		if isLock(g) {
			locks = append(locks, readLockPins(g))
		} else {
			keys = append(keys, readKeyPins(g))
		}
	}
	total := 0
	for _, key := range keys {
		for _, lock := range locks {
			if fit(key, lock) {
				total++
			}
		}
	}
	return total
}

func fit(key pins, lock pins) bool {
	for i := 0; i < 5; i++ {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}

func readLockPins(g *grid.Grid[string]) pins {
	// count the number of '#' in each column starting from the top
	var p pins
	for i := 0; i < g.Width(); i++ {
		for j := 1; j < g.Height(); j++ {
			if t, _ := g.Get(i, j); t.Value == "#" {
				p[i]++
			}
		}
	}
	return p
}

func readKeyPins(g *grid.Grid[string]) pins {
	// count the number of '#' in each column starting from the bottom
	var p pins
	for i := 0; i < g.Width(); i++ {
		for j := g.Height() - 2; j >= 0; j-- {
			if t, _ := g.Get(i, j); t.Value == "#" {
				p[i]++
			}
		}
	}
	return p
}

func isLock(g *grid.Grid[string]) bool {
	// is lock if first row is full of #
	for i := 0; i < g.Width(); i++ {
		if t, _ := g.Get(i, 0); t.Value != "#" {
			return false
		}
	}
	return true
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/25/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
