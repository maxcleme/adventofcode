package _02

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
		f, err := os.ReadFile("./2025/02/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(payload string) int {
	sum := 0
	ranges := strings.Split(payload, ",")
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		for i := from; i <= to; i++ {
			v := fmt.Sprint(i)
			if !invalid(v) {
				continue
			}
			sum += i
		}
	}
	return sum
}

func invalid(v string) bool {
	for i := 1; i <= len(v)/2; i++ {
		if len(v)%i != 0 {
			continue
		}
		s := v[0:i]
		repeat := true
		for j := i; j < len(v); j += i {
			if v[j:j+i] != s {
				repeat = false
				break
			}
		}
		if repeat {
			return true
		}
	}
	return false
}
