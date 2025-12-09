package _03

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/03/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(payload string) int {
	sum := 0
	rows := strings.Split(payload, "\n")
	for _, row := range rows {
		left := slices.Max([]rune(row)[:len(row)-1])
		i := strings.IndexRune(row, left)
		right := slices.Max([]rune(row)[i+1:])
		joltage, _ := strconv.Atoi(string(left) + string(right))
		sum += joltage
	}
	return sum
}
