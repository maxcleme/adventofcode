package _05

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Tree map[string]Nodes

type Nodes struct {
	id       string
	parent   *Tree
	children *Tree
}

func (t Tree) validates(pages []string) bool {
	for i, page := range pages[:len(pages)-1] {
		if _, ok := t[page]; !ok {
			continue
		}
		if _, ok := (*t[page].children)[pages[i+1]]; !ok {
			return false
		}
	}
	return true
}

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	rules, updates := parts[0], parts[1]

	tree := makeTree(rules)
	var valids [][]string
	for _, update := range strings.Split(updates, "\n") {
		pages := strings.Split(update, ",")
		if tree.validates(pages) {
			valids = append(valids, pages)
		}
	}

	sum := 0
	for _, pages := range valids {
		middle := pages[len(pages)/2]
		v, err := strconv.Atoi(middle)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	return sum
}

func makeTree(rules string) Tree {
	tree := Tree{}
	for _, line := range strings.Split(rules, "\n") {
		parts := strings.Split(line, "|")
		a, b := parts[0], parts[1]
		if _, ok := tree[a]; !ok {
			tree[a] = Nodes{id: a, parent: &Tree{}, children: &Tree{}}
		}
		if _, ok := tree[b]; !ok {
			tree[b] = Nodes{id: b, parent: &Tree{}, children: &Tree{}}
		}
		t := *tree[a].children
		t[b] = tree[b]
		t = *tree[b].parent
		t[a] = tree[a]
	}
	return tree
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/05/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
