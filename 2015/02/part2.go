package _01

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/02/input")
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
		var l, w, h int
		if _, err := fmt.Sscanf(row, "%dx%dx%d", &l, &w, &h); err != nil {
			panic(err)
		}
		total += l*w*h + min(2*l+2*w, 2*w+2*h, 2*h+2*l)
	}
	return total
}
