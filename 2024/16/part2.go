package _16

import (
	"fmt"
	"math"
	"os"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/maxcleme/adventofcode/queue"
	"github.com/spf13/cobra"
)

type StateWithPath struct {
	State
	paths []*grid.Tile[string]
}

func part2(input string) int {
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
	var bestPaths [][]*grid.Tile[string]
	q := queue.New[StateWithPath]()
	q.Push(StateWithPath{State: State{position: start, cost: 0, direction: Right}, paths: []*grid.Tile[string]{start}})
	for q.Len() > 0 {
		curr := q.Pop()
		if curr.position == end {
			if curr.cost < cost {
				bestPaths = nil
			}
			if curr.cost <= cost {
				cost = curr.cost
				bestPaths = append(bestPaths, curr.paths)
			}
			continue
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
			if newCost <= costs[nextY][nextX][dir] {
				costs[nextY][nextX][dir] = newCost
				var nextPath []*grid.Tile[string]
				for _, p := range curr.paths {
					nextPath = append(nextPath, p)
				}
				nextPath = append(nextPath, next)
				q.Push(StateWithPath{State: State{position: next, cost: newCost, direction: dir}, paths: nextPath})
			}
		}
	}

	uniq := map[*grid.Tile[string]]struct{}{}
	for _, path := range bestPaths {
		for _, p := range path {
			uniq[p] = struct{}{}
		}
	}
	return len(uniq)
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/16/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
