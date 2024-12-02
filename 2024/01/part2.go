package _01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2024/01/input")
		if err != nil {
			return err
		}

		var left []int
		var right []int
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			l := scanner.Text()
			arr := strings.Split(l, "   ")
			i, err := strconv.Atoi(arr[0])
			if err != nil {
				return err
			}
			left = append(left, i)
			j, err := strconv.Atoi(arr[1])
			if err != nil {
				return err
			}
			right = append(right, j)
		}
		index := map[int]int{}
		for _, n := range right {
			_, ok := index[n]
			if ok {
				index[n] = index[n] + 1
			} else {
				index[n] = 1
			}
		}
		var sum int
		for _, n := range left {
			sum += n * index[n]
		}
		fmt.Println(sum)
		return nil
	},
}
