package _11

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/11/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

type Node struct {
	ID       string
	Parents  []*Node
	Children []*Node
}

type Graph struct {
	Nodes map[string]*Node
}

func (g *Graph) AddEdge(from, to string) {
	if g.Nodes == nil {
		g.Nodes = map[string]*Node{}
	}
	if _, exists := g.Nodes[from]; !exists {
		g.Nodes[from] = &Node{ID: from}
	}
	if _, exists := g.Nodes[to]; !exists {
		g.Nodes[to] = &Node{ID: to}
	}
	g.Nodes[from].Children = append(g.Nodes[from].Children, g.Nodes[to])
	g.Nodes[to].Parents = append(g.Nodes[to].Parents, g.Nodes[from])
}

func (g *Graph) Clone() *Graph {
	newGraph := Graph{Nodes: map[string]*Node{}}
	for id, node := range g.Nodes {
		for _, child := range node.Children {
			newGraph.AddEdge(id, child.ID)
		}
	}
	return &newGraph
}

func (g *Graph) TopologicalSort() []string {
	var result []string
	var starts []string
	parentsCount := map[string]int{}
	for id, node := range g.Nodes {
		if len(node.Parents) == 0 {
			starts = append(starts, id)
		}
		parentsCount[id] = len(node.Parents)
	}
	for len(starts) > 0 {
		currentID := starts[0]
		starts = starts[1:]
		result = append(result, currentID)
		for _, child := range g.Nodes[currentID].Children {
			parentsCount[child.ID]--
			if parentsCount[child.ID] == 0 {
				starts = append(starts, child.ID)
			}
		}
	}
	return result
}

func (g *Graph) Paths(from, to string) int {
	sorted := g.TopologicalSort()
	counts := map[string]int{
		from: 1,
	}
	for _, id := range sorted {
		if counts[id] == 0 {
			continue
		}
		for _, child := range g.Nodes[id].Children {
			counts[child.ID] += counts[id]
		}
	}
	return counts[to]
}

func part1(payload string) int {
	var g Graph
	for _, line := range strings.Split(payload, "\n") {
		parts := strings.Split(line, ":")
		from := parts[0]
		for _, to := range strings.Fields(parts[1]) {
			g.AddEdge(from, to)
		}
	}
	return g.Paths("you", "out")
}
