package _03

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/03/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var rs []Rucksack
		var shared []rune
		var p int
		for scanner.Scan() {
			l := scanner.Text()
			r := ParseRucksack(l)
			rs = append(rs, r)
			if len(rs) != 3 {
				continue
			}
			shared = Shared(rs...)
			p += Priorities(shared)
			shared = nil
			rs = nil
		}
		fmt.Println(p)
		return nil
	},
}
