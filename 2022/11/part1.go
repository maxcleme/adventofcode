package _11

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
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
		fmt.Println(Business(Rounds(monkeys, func(p int) int {
			return int(math.Floor(float64(p) / 3))
		}, 20)))
		return nil
	},
}
