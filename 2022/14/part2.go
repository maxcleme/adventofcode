package _14

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/14/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var paths []Path
		for scanner.Scan() {
			l := scanner.Text()
			paths = append(paths, ParsePath(l))
		}
		m := MakeMap(paths, Position{X: 500}, true)

		var count int
		for m.Tick() {
			count++
		}
		fmt.Println(count)
		return nil
	},
}
