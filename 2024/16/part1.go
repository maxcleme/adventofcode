package _16

import (
	"fmt"
	"math"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/maxcleme/adventofcode/queue"
	"github.com/spf13/cobra"
)

type Direction int

const (
	Up = iota
	Down
	Left
	Right
)

type State struct {
	position  *grid.Tile[string]
	direction Direction
	cost      int
}

func part1(input string) int {
	g := grid.New[string](input, func(r rune) string {
		return string(r)
	})
	start := g.Find(func(t *grid.Tile[string]) bool {
		return t.Value == "S"
	})[0]
	end := g.Find(func(t *grid.Tile[string]) bool {
		return t.Value == "E"
	})[0]
	var costs [][][]int
	for range g.Height() {
		var c [][]int
		for range g.Width() {
			c = append(c, []int{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt})
		}
		costs = append(costs, c)
	}
	costs[start.Y][start.X][Right] = 0

	cost := math.MaxInt
	q := queue.New[State]()
	q.Push(State{position: start, cost: 0, direction: Right})
	for q.Len() > 0 {
		curr := q.Pop()
		if curr.position == end && curr.cost < cost {
			cost = curr.cost
		}
		directions := []Direction{Up, Down, Left, Right}
		for _, dir := range directions {
			var nextX, nextY int
			switch dir {
			case Up:
				nextX, nextY = curr.position.X, curr.position.Y-1
			case Down:
				nextX, nextY = curr.position.X, curr.position.Y+1
			case Left:
				nextX, nextY = curr.position.X-1, curr.position.Y
			case Right:
				nextX, nextY = curr.position.X+1, curr.position.Y
			}
			next, ok := g.Get(nextX, nextY)
			if !ok || next.Value == "#" {
				continue
			}
			moveCost := 1
			if curr.direction != dir {
				moveCost += 1000
			}
			newCost := curr.cost + moveCost
			if newCost < costs[nextY][nextX][dir] {
				costs[nextY][nextX][dir] = newCost
				q.Push(State{position: next, cost: newCost, direction: dir})
			}
		}
	}
	return cost
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/16/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
