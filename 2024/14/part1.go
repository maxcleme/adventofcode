package _14

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Guard struct {
	X, Y   int
	VX, VY int
}

func (g *Guard) Move(iteration, width, height int) {
	startX, startY := g.X, g.Y
	cycle := 0
	for i := 1; i <= iteration; i++ {
		g.MoveOnce(width, height)
		if g.X == startX && g.Y == startY {
			cycle = i
			break
		}
	}
	if cycle == 0 {
		return
	}
	for i := iteration - cycle*(iteration/cycle); i > 0; i-- {
		g.MoveOnce(width, height)
	}
}

func (g *Guard) MoveOnce(width, height int) {
	g.X = (g.X + g.VX + width) % width
	g.Y = (g.Y + g.VY + height) % height
}

func part1(input string, iteration, width, height int) int {
	var guards []*Guard
	for _, row := range strings.Split(input, "\n") {
		g := &Guard{}
		if _, err := fmt.Sscanf(row, "p=%d,%d v=%d,%d", &g.X, &g.Y, &g.VX, &g.VY); err != nil {
			panic(err)
		}
		guards = append(guards, g)
	}

	for _, g := range guards {
		g.Move(iteration, width, height)
	}
	total := countWithin(guards, 0, 0, width/2, height/2)
	total *= countWithin(guards, 1+width/2, 0, width/2, height/2)
	total *= countWithin(guards, 0, 1+height/2, width/2, height/2)
	total *= countWithin(guards, 1+width/2, 1+height/2, width/2, height/2)
	return total
}

func countWithin(guards []*Guard, x, y, width, height int) int {
	total := 0
	for _, g := range guards {
		if g.X >= x && g.X < x+width && g.Y >= y && g.Y < y+height {
			total++
		}
	}
	return total
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/14/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f), 100, 101, 103))
		return nil
	},
}
