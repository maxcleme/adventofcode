package _06

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
		f, err := os.ReadFile("./2025/06/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(payload string) int {
	rows := strings.Split(payload, "\n")
	var matrix [][]int
	for _, row := range rows[:len(rows)-1] {
		var values []int
		for _, v := range strings.Fields(row) {
			n, _ := strconv.Atoi(v)
			values = append(values, n)
		}
		matrix = append(matrix, values)
	}
	total := 0
	for i := 0; i < len(matrix[0]); i++ {
		op := strings.Fields(rows[len(rows)-1])[i]
		sum := 0
		if op == "*" {
			sum = 1
		}
		for j := 0; j < len(matrix); j++ {
			switch op {
			case "+":
				sum += matrix[j][i]
			case "*":
				sum *= matrix[j][i]
			}
		}
		total += sum
	}
	return total
}
