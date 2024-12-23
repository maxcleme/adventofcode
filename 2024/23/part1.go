package _23

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

type Node struct {
	ID    string
	Nodes []*Node
}

func part1(input string) int {
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

	// https://www.cs.cornell.edu/courses/cs6241/2019sp/readings/Chiba-1985-arboricity.pdf
	for _, node := range nodes {
		sort.Slice(node.Nodes, func(i, j int) bool {
			return len(node.Nodes[i].Nodes) < len(node.Nodes[j].Nodes)
		})
	}
	triangles := map[string][]string{}
	seen := map[string]struct{}{}
	for _, u := range nodes {
		seen[u.ID] = struct{}{}
		for _, v := range u.Nodes {
			if _, ok := seen[v.ID]; ok {
				continue
			}
			for _, w := range v.Nodes {
				if _, ok := seen[w.ID]; ok {
					continue
				}
				if slices.Contains(u.Nodes, w) {
					t := []string{u.ID, v.ID, w.ID}
					sort.Strings(t)
					triangles[fmt.Sprint(t)] = t
				}
			}
		}
	}
	total := 0
	for _, triangle := range triangles {
		if !slices.ContainsFunc(triangle, func(s string) bool {
			return strings.HasPrefix(s, "t")
		}) {
			continue
		}
		total++
	}
	return total
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/23/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
