package _01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/01/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var elves []int
		sum := 0
		for scanner.Scan() {
			l := scanner.Text()
			if l == "" {
				elves = append(elves, sum)
				sum = 0
				continue
			}
			value, err := strconv.Atoi(l)
			if err != nil {
				return err
			}
			sum += value
		}
		elves = append(elves, sum)
		sort.Ints(elves)
		fmt.Println(elves[len(elves)-3] + elves[len(elves)-2] + elves[len(elves)-1])
		return nil
	},
}
