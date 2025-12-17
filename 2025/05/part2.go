package _05

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
		f, err := os.ReadFile("./2025/05/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(payload string) int {
	var ranges Ranges
	for _, r := range strings.Split(strings.Split(payload, "\n\n")[0], "\n") {
		parts := strings.Split(r, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		ranges = ranges.add(Range{a, b})
	}
	ranges = ranges.compact()
	count := 0
	for r := range ranges {
		count += ranges[r].Max - ranges[r].Min + 1
	}
	return count
}
