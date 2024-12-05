package _05

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/05/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	rules, updates := parts[0], parts[1]
	var fixed [][]string
	for _, update := range strings.Split(updates, "\n") {
		tree := makeTree(rules)
		pages := strings.Split(update, ",")
		if tree.validates(pages) {
			continue
		}
		fixed = append(fixed, fix(tree, pages))
	}
	sum := 0
	for _, pages := range fixed {
		middle := pages[len(pages)/2]
		v, err := strconv.Atoi(middle)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	return sum
}

// https://en.wikipedia.org/wiki/Topological_sorting
// Kahn's algorithm
func fix(t Tree, paths []string) []string {
	slices.Sort(paths)
	var possibleStart []string
	uniqueNodes := slices.Compact(slices.Clone(paths))
	for _, node := range uniqueNodes {
		count := 0
		for _, parent := range *t[node].parent {
			if slices.Contains(paths, parent.id) {
				count++
			}
		}
		if count == 0 {
			possibleStart = append(possibleStart, node)
		}
	}
	var fixed []string
	for len(possibleStart) > 0 {
		start := possibleStart[0]
		possibleStart = possibleStart[1:]
		fixed = append(fixed, start)
		for _, child := range *t[start].children {
			if !slices.Contains(paths, child.id) {
				continue
			}
			delete(*t[child.id].parent, start)
			count := 0
			for _, parent := range *t[child.id].parent {
				if slices.Contains(paths, parent.id) {
					count++
				}
			}
			if count == 0 {
				possibleStart = append(possibleStart, child.id)
			}
		}
	}
	return fixed
}
