package _14

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func part2(input string, width, height int) int {
	var guards []*Guard
	for _, row := range strings.Split(input, "\n") {
		g := &Guard{}
		if _, err := fmt.Sscanf(row, "p=%d,%d v=%d,%d", &g.X, &g.Y, &g.VX, &g.VY); err != nil {
			panic(err)
		}
		guards = append(guards, g)
	}

	iteration := 0
	for {
		iteration++
		for _, g := range guards {
			g.MoveOnce(width, height)
		}
		// Stop when many guard are within the center of the grid
		// 300 is arbitrary, but it works ¯\_(ツ)_/¯
		if countWithin(guards, width/4, height/4, width/2, height/2) > 300 {
			break
		}
	}
	draw(guards, width, height)
	return iteration
}

func draw(guards []*Guard, width int, height int) {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for _, g := range guards {
		grid[g.Y][g.X] = '#'
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/14/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f), 101, 103))
		return nil
	},
}
