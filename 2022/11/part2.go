package _11

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/11/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var monkeys []Monkey
		for scanner.Scan() {
			var lines []string
			for i := 0; i < 6; i++ {
				lines = append(lines, scanner.Text())
				scanner.Scan()
			}
			monkeys = append(monkeys, ParseMonkey(lines[1:]))
		}
		monkeys = Rounds(monkeys, func(p int) int {
			return p
		}, 10000)
		fmt.Println(Business(monkeys))
		return nil
	},
}
