package _21

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func part2(input string) int {
	result := 0
	for _, code := range strings.Split(input, "\n") {
		result += complexity(code, 25)
	}
	return result
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/21/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
