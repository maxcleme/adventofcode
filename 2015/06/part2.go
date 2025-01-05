package _06

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/06/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(input string) int {
	g := grid.Empty[int](1000, 1000)
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
			m(g, x1, y1, x2, y2, func(curr int) int {
				return curr + 1
			})
		case "turn off":
			m(g, x1, y1, x2, y2, func(curr int) int {
				return max(0, curr-1)
			})
		case "toggle":
			m(g, x1, y1, x2, y2, func(curr int) int {
				return curr + 2
			})
		}
	}
	total := 0
	for t := range g.All() {
		total += t.Value
	}
	return total
}
