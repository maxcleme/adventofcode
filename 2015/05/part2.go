package _05

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/05/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(input string) int {
	total := 0
	for _, row := range strings.Split(input, "\n") {
		if !isNice2(row) {
			continue
		}
		total++
	}
	return total
}

func isNice2(row string) bool {
	if len(row) < 4 {
		return false
	}
	if !hasPair(row) {
		return false
	}
	if !hasRepeat(row) {
		return false
	}
	return true
}

func hasPair(row string) bool {
	for i := 0; i < len(row)-1; i++ {
		if strings.Contains(row[i+2:], row[i:i+2]) {
			return true
		}
	}
	return false
}

func hasRepeat(row string) bool {
	for i := 0; i < len(row)-2; i++ {
		if row[i] == row[i+2] {
			return true
		}
	}
	return false
}
