package _02

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
		f, err := os.ReadFile("./2025/02/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(payload string) int {
	sum := 0
	ranges := strings.Split(payload, ",")
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		for i := from; i <= to; i++ {
			v := fmt.Sprint(i)
			if len(v)%2 != 0 {
				continue
			}
			if v[:len(v)/2] != v[len(v)/2:] {
				continue
			}
			sum += i
		}
	}
	return sum
}
