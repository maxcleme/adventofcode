package _12

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/12/input")
		if err != nil {
			return err
		}
		m := MakeMap()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			l := scanner.Text()
			m = m.AppendRow(l)
		}
		p := m.Find(func(t *Tile) bool {
			return t == m.Source
		})
		fmt.Println(len(p) - 1)
		return nil
	},
}
