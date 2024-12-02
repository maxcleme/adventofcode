package _01

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
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
		sort.Ints(left)
		sort.Ints(right)
		var sum int
		for i, _ := range left {
			diff := left[i] - right[i]
			sum += int(math.Abs(float64(diff)))
		}
		fmt.Println(sum)
		return nil
	},
}
