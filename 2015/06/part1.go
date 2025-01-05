package _06

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/06/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

var rgx = regexp.MustCompile(`^(turn on|turn off|toggle) (\w+),(\w+) through (\w+),(\w+)$`)

func part1(input string) int {
	g := grid.Empty[bool](1000, 1000)
	for _, row := range strings.Split(input, "\n") {
		groups := rgx.FindStringSubmatch(row)
		if len(groups) != 6 {
			panic("invalid input:" + row)
		}
		x1, _ := strconv.Atoi(groups[2])
		y1, _ := strconv.Atoi(groups[3])
		x2, _ := strconv.Atoi(groups[4])
		y2, _ := strconv.Atoi(groups[5])
		switch groups[1] {
		case "turn on":
			m(g, x1, y1, x2, y2, func(_ bool) bool {
				return true
			})
		case "turn off":
			m(g, x1, y1, x2, y2, func(_ bool) bool {
				return false
			})
		case "toggle":
			m(g, x1, y1, x2, y2, func(curr bool) bool {
				return !curr
			})
		}
	}
	return len(g.Find(func(t *grid.Tile[bool]) bool {
		return t.Value
	}))
}

func m[T any](g *grid.Grid[T], x1, y1, x2, y2 int, fn func(T) T) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			t, _ := g.Get(x, y)
			t.Value = fn(t.Value)
		}
	}
}
