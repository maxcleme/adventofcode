package _06

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/06/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(payload string) int {
	// parse column to only get max length
	rows := strings.Split(payload, "\n")
	var matrix [][]string
	for _, row := range rows[:len(rows)-1] {
		var values []string
		for _, v := range strings.Fields(row) {
			values = append(values, v)
		}
		matrix = append(matrix, values)
	}
	var lengths []int
	for i := 0; i < len(matrix[0]); i++ {
		l := 0
		for j := 0; j < len(matrix); j++ {
			l = max(l, len(matrix[j][i]))
		}
		lengths = append(lengths, l)
	}
	// parse again but with padding
	matrix = [][]string{}
	for _, row := range rows[:len(rows)-1] {
		var values []string
		s := 0
		for _, l := range lengths {
			v := row[s : s+l]
			values = append(values, v)
			s += l + 1
		}
		matrix = append(matrix, values)
	}

	total := 0
	for i := 0; i < len(matrix[0]); i++ {
		// compute actual values
		var values []int
		for k := 0; k < len(matrix[0][i]); k++ {
			s := ""
			for j := 0; j < len(matrix); j++ {
				s += string(matrix[j][i][k])
			}
			n, _ := strconv.Atoi(strings.TrimSpace(s))
			values = append(values, n)
		}
		// compute operation
		op := strings.Fields(rows[len(rows)-1])[i]
		result := 0
		if op == "*" {
			result = 1
		}
		for _, v := range values {
			switch op {
			case "+":
				result += v
			case "*":
				result *= v
			}
		}
		total += result
	}
	return total
}
