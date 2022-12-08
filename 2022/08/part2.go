package _08

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/08/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var m Map
		for scanner.Scan() {
			l := scanner.Text()
			m = append(m, ParseRow(l))
		}
		fmt.Println(m.BestScore())
		return nil
	},
}
