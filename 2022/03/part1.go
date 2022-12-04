package _03

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/03/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var p int
		for scanner.Scan() {
			l := scanner.Text()
			r := ParseRucksack(l)
			s := r.Shared()
			p += Priorities(s)
		}
		fmt.Println(p)
		return nil
	},
}
