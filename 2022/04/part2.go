package _04

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/04/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var p int
		for scanner.Scan() {
			l := scanner.Text()
			p1, p2, err := ParsePairs(l)
			if err != nil {
				return err
			}
			if p1.Overlaps(p2) {
				p += 1
			}
		}
		fmt.Println(p)
		return nil
	},
}
