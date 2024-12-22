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

func part2(input string) int {
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
					t += translate2(lastSequence, string(sequence[j]), 25, map[string]int{})
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

func translate2(from, to string, depth int, cache map[string]int) int {
	key := fmt.Sprintf("%s-%s-%d", from, to, depth)
	if total, ok := cache[key]; ok {
		return total
	}
	if depth == 1 {
		return len(slices.MinFunc(d2d(from, to), func(a, b string) int {
			return len(a) - len(b)
		}))
	}
	total := math.MaxInt
	for _, s := range d2d(from, to) {
		last := ""
		c := 0
		for i := range s {
			t := translate2(last, string(s[i]), depth-1, cache)
			last = string(s[i])
			c += t
		}
		if c < total {
			total = c
		}
	}
	cache[key] = total
	return total
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/21/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
