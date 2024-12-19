package _19

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	designs := map[string]struct{}{}
	for _, d := range strings.Split(parts[0], ",") {
		designs[strings.TrimSpace(d)] = struct{}{}
	}
	count := 0
	for _, p := range strings.Split(parts[1], "\n") {
		if _, ok := minimumDesignRequired(p, designs); !ok {
			continue
		}
		count++
	}
	return count
}

func minimumDesignRequired(p string, designs map[string]struct{}) (int, bool) {
	if len(p) == 0 {
		return 0, true
	}
	var knownMin []int
	for range len(p) + 1 {
		knownMin = append(knownMin, math.MaxInt32)
	}
	knownMin[0] = 0
	for i := 1; i <= len(p); i++ {
		for j := 0; j < i; j++ {
			substring := p[j:i]
			if _, exists := designs[substring]; exists {
				knownMin[i] = min(knownMin[i], knownMin[j]+1)
			}
		}
	}
	if knownMin[len(p)] == math.MaxInt32 {
		return -1, false
	}
	return knownMin[len(p)], true
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/19/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
