package _24

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func part2(input string) string {
	parts := strings.Split(input, "\n\n")
	wires := map[string]bool{}
	for _, row := range strings.Split(parts[0], "\n") {
		parts := strings.Split(row, ":")
		wire := parts[0]
		wires[wire] = strings.TrimSpace(parts[1]) == "1"
	}
	gates := map[string]struct {
		a, b, op, c string
	}{}
	for _, row := range strings.Split(parts[1], "\n") {
		a, op, b, c := "", "", "", ""
		if _, err := fmt.Sscanf(row, "%s %s %s -> %s", &a, &op, &b, &c); err != nil {
			fmt.Println(err)
		}
		gates[row] = struct {
			a, b, op, c string
		}{a, b, op, c}
	}

	// https://en.wikipedia.org/wiki/Adder_(electronics)#Ripple-carry_adder
	var invalid []string
	// all z gates should be XOR except last one
	for _, v := range gates {
		if v.c[:1] != "z" {
			continue
		}
		if v.op != "XOR" && v.c != "z45" {
			invalid = append(invalid, v.c)
			continue
		}
	}
	// all temporary gates should be AND/OR but not XOR
	for _, v := range gates {
		if v.c[:1] == "z" {
			continue
		}
		if v.a[:1] == "x" || v.b[:1] == "y" {
			continue
		}
		if v.a[:1] == "y" || v.b[:1] == "x" {
			continue
		}
		if v.op == "XOR" {
			invalid = append(invalid, v.c)
		}
	}
	// all x(i) XOR y(i) results should be used in another XOR (except 00)
	for _, v := range gates {
		if v.a[:1] != "x" && v.a[:1] != "y" {
			continue
		}
		if v.b[:1] != "x" && v.b[:1] != "y" {
			continue
		}
		if v.op != "XOR" {
			continue
		}
		if v.a[1:] == "00" || v.b[1:] == "00" || v.c[1:] == "00" {
			continue
		}
		found := false
		for _, v2 := range gates {
			if v2.op == "XOR" && (v2.a == v.c || v2.b == v.c) {
				found = true
				break
			}
		}
		if !found {
			invalid = append(invalid, v.c)
		}
	}
	// all x(i) AND y(i) results should be used in another OR (except 00)
	for _, v := range gates {
		if v.a[:1] != "x" && v.a[:1] != "y" {
			continue
		}
		if v.b[:1] != "x" && v.b[:1] != "y" {
			continue
		}
		if v.op != "AND" {
			continue
		}
		if v.a[1:] == "00" || v.b[1:] == "00" || v.c[1:] == "00" {
			continue
		}
		found := false
		for _, v2 := range gates {
			if v2.op == "OR" && (v2.a == v.c || v2.b == v.c) {
				found = true
				break
			}
		}
		if !found {
			invalid = append(invalid, v.c)
		}
	}
	sort.Strings(invalid)
	invalid = slices.Compact(invalid)
	return strings.Join(invalid, ",")
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/24/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
