package _02

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/02/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(input string) int {
	total := 0
	for _, row := range strings.Split(input, "\n") {
		var l, w, h int
		if _, err := fmt.Sscanf(row, "%dx%dx%d", &l, &w, &h); err != nil {
			panic(err)
		}
		total += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}
	return total
}
