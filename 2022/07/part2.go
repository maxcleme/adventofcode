package _07

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/07/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		fs := NewFS("/")
		for scanner.Scan() {
			l := scanner.Text()
			fs = fs.Parse(l)
		}
		total := 70_000_000
		required := 30_000_000
		left := total - fs.Root().size
		miss := required - left
		var candidate int
		fs.Root().Walk(func(path string, dir bool, size int) {
			if path != "/" && dir && size > miss && (candidate == 0 || size < candidate){
				candidate = size
			}
		})
		fmt.Println(candidate)
		return nil
	},
}
