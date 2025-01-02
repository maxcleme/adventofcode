package _01

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/01/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(input string) int {
	floor := 0
	for _, r := range input {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}
