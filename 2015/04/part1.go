package _01

import (
	"crypto/md5"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/04/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(input string) int {
	i := 0
	for {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		if hex := fmt.Sprintf("%x", hash); hex[:5] == "00000" {
			return i
		}
		i++
	}
}
