package _24

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	wires := map[string]bool{}
	for _, row := range strings.Split(parts[0], "\n") {
		parts := strings.Split(row, ":")
		wire := parts[0]
		wires[wire] = strings.TrimSpace(parts[1]) == "1"
	}
	gates := map[string]func() bool{}
	for _, row := range strings.Split(parts[1], "\n") {
		a, op, b, c := "", "", "", ""
		if _, err := fmt.Sscanf(row, "%s %s %s -> %s", &a, &op, &b, &c); err != nil {
			fmt.Println(err)
		}
		gates[row] = func() bool {
			if _, ok := wires[a]; !ok {
				return false
			}
			if _, ok := wires[b]; !ok {
				return false
			}
			switch op {
			case "AND":
				wires[c] = wires[a] && wires[b]
			case "OR":
				wires[c] = wires[a] || wires[b]
			case "XOR":
				wires[c] = wires[a] != wires[b]
			default:
				panic("unknown op")
			}
			return true
		}
	}

	for len(gates) > 0 {
		for k, v := range gates {
			if v() {
				delete(gates, k)
			}
		}
	}
	result := make([]rune, 1000)
	size := 0
	for k, v := range wires {
		if !strings.HasPrefix(k, "z") {
			continue
		}
		n, err := strconv.Atoi(strings.TrimPrefix(k, "z"))
		if err != nil {
			panic(err)
		}
		if v {
			result[n] = '1'
		} else {
			result[n] = '0'
		}
		size = max(size, n)
	}
	result = result[:size+1]
	slices.Reverse(result)
	n, err := strconv.ParseInt(string(result), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/24/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
