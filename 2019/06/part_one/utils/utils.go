package utils

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type Node struct {
	Parent   *Node
	Children map[string]*Node
}

const (
	Separator                 = ")"
	ErrInvalidOrbitDefinition = Error("invalid orbit definition")
)

func TotalOrbits(input []string) (int, error) {
	index := map[string]*Node{}

	// initialize graph
	for _, i := range input {
		def := strings.Split(i, Separator)
		if len(def) != 2 {
			return 0, ErrInvalidOrbitDefinition
		}
		parent, ok := index[def[0]]
		if !ok {
			parent = &Node{
				Children: make(map[string]*Node),
			}
			index[def[0]] = parent
		}

		child, ok := index[def[1]]
		if !ok {
			child = &Node{
				Parent:   parent,
				Children: make(map[string]*Node),
			}
			index[def[1]] = child
		}

		child.Parent = parent
		parent.Children[def[1]] = child
	}

	// compute orbits from each node by going up
	count := 0
	for key, node := range index {
		c := 0
		parent := node.Parent
		for parent != nil {
			c++
			parent = parent.Parent
		}

		logrus.
			WithField("node", key).
			WithField("orbits", c).
			Info("")

		count += c
	}

	return count, nil
}
