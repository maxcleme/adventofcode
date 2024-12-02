package _02

import (
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/02/input")
		if err != nil {
			return err
		}
		reports, err := parse(string(f))
		if err != nil {
			return err
		}
		count := 0
		for _, r := range reports {
			if r.safeWithDampener() {
				count++
			}
		}
		fmt.Println(count)
		return nil
	},
}

func (r Report) safeWithDampener() bool {
	if r.safe() {
		return true
	}
	for i := 0; i < len(r); i++ {
		copy := slices.Clone(r)
		a := slices.Delete(copy, i, i+1)
		if a.safe() {
			return true
		}
	}
	return false
}
