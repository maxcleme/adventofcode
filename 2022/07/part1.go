package _07

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
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
		var total int
		fs.Root().Walk(func(path string, dir bool, size int) {
			if path != "/" && dir && size <= 100000 {
				fmt.Println(path, size)
				total += size
			}
		})
		fmt.Println(total)
		return nil
	},
}
