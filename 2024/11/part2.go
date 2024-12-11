package _11

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func part2(input string) int {
	return countAfter(input, 75)
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/11/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
