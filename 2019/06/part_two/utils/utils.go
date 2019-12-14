package utils

import (
	"math"
	"strings"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type Node struct {
	Identifier string
	Parent     *Node
	Children   map[string]*Node
}

const (
	Separator                 = ")"
	ErrInvalidOrbitDefinition = Error("invalid orbit definition")
	ErrUnknownNode            = Error("unknown node")
	ErrReferenceNeedParent    = Error("reference node needs a parent")
)

func TotalOrbits(input []string) (uint64, error) {
	graph, err := create(input)
	if err != nil {
		return 0, err
	}

	// compute orbits from each node by going up
	count := uint64(0)
	for _, node := range graph {
		c := uint64(0)
		parent := node.Parent
		for parent != nil {
			c++
			parent = parent.Parent
		}

		count += c
	}

	return count, nil
}

func Between(input []string, from, to string) (uint64, error) {
	graph, err := create(input)
	if err != nil {
		return 0, err
	}

	// check presence of from
	srcNode, ok := graph[from]
	if !ok {
		return 0, ErrUnknownNode
	}
	if srcNode.Parent == nil {
		return 0, ErrReferenceNeedParent
	}

	// check presence of to
	destNode, ok := graph[to]
	if !ok {
		return 0, ErrUnknownNode
	}
	if destNode.Parent == nil {
		return 0, ErrReferenceNeedParent
	}

	// pick orbit reference for source / destination node
	src := srcNode.Parent.Identifier
	dest := destNode.Parent.Identifier

	// init data required by dijkstra
	distances := make(map[string]uint64)
	visited := make(map[string]bool)
	nodes := make(map[string]*Node)
	for k, n := range graph {
		nodes[k] = n
		distances[k] = math.MaxUint64
	}
	distances[src] = 0

	for len(nodes) != 0 {
		// find node with minimal distance from origin
		min := uint64(math.MaxUint64)
		var u string
		for k, d := range distances {
			if !visited[k] && d < min {
				min = d
				u = k
			}
		}

		// short circuit since only distance between from and to is what we want
		if u == dest {
			return distances[u], nil
		}

		// remove u from nodes index
		delete(nodes, u)

		current := graph[u]
		visited[u] = true
		// compute/update distance for children and parent
		// since graph is not weighted, all distance will be the same
		d := distances[u] + 1
		if current.Parent != nil && d < distances[current.Parent.Identifier] {
			distances[current.Parent.Identifier] = d
		}
		for _, n := range current.Children {
			if d < distances[n.Identifier] {
				distances[n.Identifier] = d
			}
		}

	}

	// worst case scenario, to node was the fairest node
	return distances[to], nil
}

func create(input []string) (map[string]*Node, error) {
	index := map[string]*Node{}

	// initialize graph
	for _, i := range input {
		def := strings.Split(i, Separator)
		if len(def) != 2 {
			return nil, ErrInvalidOrbitDefinition
		}
		parent, ok := index[def[0]]
		if !ok {
			parent = &Node{
				Identifier: def[0],
				Children:   make(map[string]*Node),
			}
			index[def[0]] = parent
		}

		child, ok := index[def[1]]
		if !ok {
			child = &Node{
				Identifier: def[1],
				Parent:     parent,
				Children:   make(map[string]*Node),
			}
			index[def[1]] = child
		}

		child.Parent = parent
		parent.Children[def[1]] = child
	}

	return index, nil
}
