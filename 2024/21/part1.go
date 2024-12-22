package _21

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/maxcleme/adventofcode/grid"
	"github.com/spf13/cobra"
)

func part1(input string) int {
	result := 0
	for _, code := range strings.Split(input, "\n") {
		last := ""
		sumComplexity := 0
		for i := range code {
			minComplexity := math.MaxInt
			sequences := n2d(last, string(code[i]))
			for _, sequence := range sequences {
				lastSequence := ""
				t := 0
				for j := range sequence {
					t += translate2(lastSequence, string(sequence[j]), 2, map[string]int{})
					lastSequence = string(sequence[j])
				}
				if t < minComplexity {
					minComplexity = t
				}
			}
			last = string(code[i])
			sumComplexity += minComplexity
		}
		n, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			panic(err)
		}
		result += n * sumComplexity
	}
	return result
}

func complexity(code, s string) int {
	n, err := strconv.Atoi(code[:len(code)-1])
	if err != nil {
		panic(err)
	}
	return n * len(s)
}

func translate(code string, fns ...func(from, to string) string) string {
	for _, fn := range fns {
		result, last := "", ""
		for _, c := range code {
			result += fn(last, string(c))
			last = string(c)
		}
		code = result
	}
	return code
}

func n2d(from, to string) []string {
	if from == "" {
		from = "A"
	}
	numericKeypad := grid.New[string](`789
456
123
 0A`, func(r rune) string {
		return string(r)
	})
	seen := map[string]bool{}
	for t := range numericKeypad.All() {
		seen[t.Value] = false
	}
	seen[from] = true
	return dfs(from, to, numericKeypad, seen, "", []string{})
}

func d2d(from, to string) []string {
	if from == "" {
		from = "A"
	}
	directionalKeypad := grid.New[string](` ^A
<v>`, func(r rune) string {
		return string(r)
	})
	seen := map[string]bool{}
	for t := range directionalKeypad.All() {
		seen[t.Value] = false
	}
	seen[from] = true
	return dfs(from, to, directionalKeypad, seen, "", []string{})
}

func dfs(from string, to string, keypad *grid.Grid[string], seen map[string]bool, currentPath string, paths []string) []string {
	if from == to {
		return append(paths, currentPath+"A")
	}
	curr := keypad.Find(func(t *grid.Tile[string]) bool {
		return t.Value == from
	})[0]
	seen[from] = true
	for _, dir := range [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
		nextX, nextY := curr.X+dir[0], curr.Y+dir[1]
		next, ok := keypad.Get(nextX, nextY)
		if !ok || next.Value == " " {
			continue
		}
		if seen[next.Value] {
			continue
		}
		paths = dfs(next.Value, to, keypad, seen, currentPath+relative(curr, next), paths)
	}
	seen[from] = false
	return paths
}

func relative(from *grid.Tile[string], to *grid.Tile[string]) string {
	if from.X < to.X {
		return ">"
	}
	if from.X > to.X {
		return "<"
	}
	if from.Y < to.Y {
		return "v"
	}
	if from.Y > to.Y {
		return "^"
	}
	return ""
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/21/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
