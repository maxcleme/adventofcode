package _10

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/10/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

type Machine struct {
	Lights  string
	Joltage []int
	Buttons [][]int
}

func part1(payload string) int {
	var machines []Machine
	for _, line := range strings.Split(payload, "\n") {
		fields := strings.Fields(line)
		m := Machine{Lights: fields[0][1 : len(fields[0])-1]}
		for _, inputs := range fields[1 : len(fields)-1] {
			var buttons []int
			for _, input := range strings.Split(inputs[1:len(inputs)-1], ",") {
				v, _ := strconv.Atoi(input)
				buttons = append(buttons, v)
			}
			m.Buttons = append(m.Buttons, buttons)
		}
		machines = append(machines, m)
	}

	total := 0
	for _, m := range machines {
		from := ""
		for range len(m.Lights) {
			from += "."
		}
		distance := BFS(from, m.Lights, m.Buttons)
		total += distance
	}
	return total
}

func BFS(start, target string, buttons [][]int) int {
	if start == target {
		return 0
	}
	visited := map[string]bool{start: true}
	queue := []string{start}
	distance := 0
	for len(queue) > 0 {
		distance++
		var nextQueue []string
		for _, current := range queue {
			for _, b := range buttons {
				next := Press(current, b)
				if next == target {
					return distance
				}
				if !visited[next] {
					visited[next] = true
					nextQueue = append(nextQueue, next)
				}
			}
		}
		queue = nextQueue
	}
	return -1 // Target not reachable
}

func Press(lights string, buttons []int) string {
	s := lights
	for _, b := range buttons {
		if s[b] == '#' {
			s = s[:b] + "." + s[b+1:]
		} else {
			s = s[:b] + "#" + s[b+1:]
		}
	}
	return s
}
