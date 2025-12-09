package _03

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/03/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(payload string) int {
	sum := 0
	rows := strings.Split(payload, "\n")
	for _, row := range rows {
		sum += maxJoltage(row)
	}
	return sum
}

func maxJoltage(row string) int {
	s, _ := strconv.Atoi(maxJoltageN(row, 12))
	return s
}

func maxJoltageN(row string, n int) string {
	left := slices.Max([]rune(row)[:len(row)-(n-1)])
	if n == 1 {
		return string(left)
	}
	i := strings.IndexRune(row, left)
	return string(append([]rune{left}, []rune(maxJoltageN(row[i+1:], n-1))...)[:n])
}
