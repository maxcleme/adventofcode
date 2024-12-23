package _21

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func part1(input string) int {
	result := 0
	for _, code := range strings.Split(input, "\n") {
		result += complexity(code, 2)
	}
	return result
}

func complexity(code string, depth int) int {
	code = "A" + code
	total := 0
	cache := map[string]int{}
	for i := range code[:len(code)-1] {
		score := math.MaxInt
		for _, path := range numericKeypad.paths(string(code[i]), string(code[i+1])) {
			score = min(score, pathComplexity(path, depth, cache))
		}
		total += score
	}
	n, err := strconv.Atoi(code[1 : len(code)-1])
	if err != nil {
		panic(err)
	}
	return n * total
}

func pathComplexity(path string, depth int, cache map[string]int) int {
	path = "A" + path
	total := 0
	for j := range path[:len(path)-1] {
		total += translate(string(path[j]), string(path[j+1]), depth, cache)
	}
	return total
}

func translate(from, to string, depth int, cache map[string]int) int {
	key := fmt.Sprintf("%s-%s-%d", from, to, depth)
	if total, ok := cache[key]; ok {
		return total
	}
	if depth == 1 {
		return len(slices.MinFunc(directionalKeypad.paths(from, to), func(a, b string) int {
			return len(a) - len(b)
		}))
	}
	total := math.MaxInt
	for _, s := range directionalKeypad.paths(from, to) {
		s = "A" + s
		c := 0
		for i := range s[:len(s)-1] {
			t := translate(string(s[i]), string(s[i+1]), depth-1, cache)
			c += t
		}
		total = min(total, c)
	}
	cache[key] = total
	return total
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
