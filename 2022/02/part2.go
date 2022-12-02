package _02

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
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
			output, err := ParseOutput(draws[1])
			if err != nil {
				return err
			}
			switch output {
			case Win:
				point += int(Win) + int(cheatEngine.Win(elveInput))
			case Loose:
				point += int(Loose) + int(cheatEngine.Loose(elveInput))
			default:
				point += int(Draw) + int(cheatEngine.Draw(elveInput))
			}
		}
		fmt.Println(point)
		return nil
	},
}
