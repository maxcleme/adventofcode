package _01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/01/input")
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(f)
		max := 0
		sum := 0
		for scanner.Scan() {
			l := scanner.Text()
			if l == "" {
				sum = 0
				continue
			}
			value, err := strconv.Atoi(l)
			if err != nil {
				return err
			}
			sum += value
			if sum > max {
				max = sum
			}
		}
		fmt.Println(max)
		return nil
	},
}
