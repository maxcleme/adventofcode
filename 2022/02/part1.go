package _02

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/02/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		point := 0
		var cheatEngine = NewCheatEngine([]Input{Rock, Paper, Scissors})
		for scanner.Scan() {
			l := scanner.Text()
			draws := strings.Split(l, " ")
			elveInput, err := ParseInput(draws[0])
			if err != nil {
				return err
			}
			myInput, err := ParseInput(draws[1])
			if err != nil {
				return err
			}
			point += cheatEngine.Point(elveInput, myInput)
		}
		fmt.Println(point)
		return nil
	},
}
