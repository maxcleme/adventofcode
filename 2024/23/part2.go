package _23

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func part2(input string) string {
	nodes := map[string]*Node{}
	for _, row := range strings.Split(input, "\n") {
		ids := strings.Split(row, "-")
		a, b := ids[0], ids[1]
		if _, ok := nodes[a]; !ok {
			nodes[a] = &Node{ID: a}
		}
		if _, ok := nodes[b]; !ok {
			nodes[b] = &Node{ID: b}
		}
		nodes[a].Nodes = append(nodes[a].Nodes, nodes[b])
		nodes[b].Nodes = append(nodes[b].Nodes, nodes[a])
	}
	cliques := allCliques([]*Node{}, slices.Collect(maps.Values(nodes)), []*Node{})
	slices.SortFunc(cliques, func(a, b []*Node) int {
		return len(b) - len(a)
	})
	var ids []string
	for _, node := range cliques[0] {
		ids = append(ids, node.ID)
	}
	sort.Strings(ids)
	return strings.Join(ids, ",")
}

// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func allCliques(clique, nodes, excluded []*Node) [][]*Node {
	var all [][]*Node
	if len(nodes) == 0 && len(excluded) == 0 {
		return append(all, append([]*Node{}, clique...))
	}
	for _, v := range append([]*Node{}, nodes...) {
		neighbors := v.Nodes
		all = append(all, allCliques(
			append(clique, v),
			intersect(nodes, neighbors),
			intersect(excluded, neighbors),
		)...)
		nodes = slices.DeleteFunc(nodes, func(a *Node) bool {
			return a == v
		})
		excluded = append(excluded, v)
	}
	return all
}

func intersect(a, b []*Node) []*Node {
	set := make(map[*Node]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}
	var result []*Node
	for _, v := range b {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/23/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
