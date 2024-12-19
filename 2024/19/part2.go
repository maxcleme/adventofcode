package _19

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	designs := map[string]struct{}{}
	for _, d := range strings.Split(parts[0], ",") {
		designs[strings.TrimSpace(d)] = struct{}{}
	}
	count := 0
	for _, p := range strings.Split(parts[1], "\n") {
		_, ok := minimumDesignRequired(p, designs)
		if !ok {
			continue
		}
		count += allPossibleCombinations(p, designs)
	}
	return count
}

func allPossibleCombinations(p string, designs map[string]struct{}) int {
	if len(p) == 0 {
		return 1
	}
	var knownMin []int
	for range len(p) + 1 {
		knownMin = append(knownMin, 0)
	}
	knownMin[0] = 1
	for i := 1; i <= len(p); i++ {
		for j := 0; j < i; j++ {
			substring := p[j:i]
			if _, exists := designs[substring]; exists {
				knownMin[i] += knownMin[j]
			}
		}
	}
	return knownMin[len(p)]
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/19/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
