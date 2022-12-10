package _09

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/09/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		r := MakeRope(10)
		history := map[int]map[int]bool{
			r.Head().x: {r.Head().y: true},
		}
		count := 1
		for scanner.Scan() {
			l := scanner.Text()
			m := ParseMove(l)
			r = m.Apply(r, func(r Rope) {
				if _, ok := history[r.Tail().x]; !ok {
					history[r.Tail().x] = map[int]bool{}
				}
				if !history[r.Tail().x][r.Tail().y] {
					history[r.Tail().x][r.Tail().y] = true
					count++
				}
			})
		}
		fmt.Println(count)
		return nil
	},
}
