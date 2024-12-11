package _11

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func part1(input string) int {
	return countAfter(input, 25)
}

func countAfter(input string, time int) int {
	stoneCounts := make(map[int]int)
	for _, p := range strings.Split(input, " ") {
		n, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		stoneCounts[n]++
	}
	for range time {
		stoneCounts = blink(stoneCounts)
	}
	totalStones := 0
	for _, count := range stoneCounts {
		totalStones += count
	}
	return totalStones
}

func blink(stoneCounts map[int]int) map[int]int {
	nextCounts := make(map[int]int)
	for stone, count := range stoneCounts {
		switch {
		case stone == 0:
			nextCounts[1] += count
		case len(strconv.Itoa(stone))%2 == 0:
			left, right := splitNumber(stone)
			nextCounts[left] += count
			nextCounts[right] += count
		default:
			nextCounts[stone*2024] += count
		}
	}
	return nextCounts
}

func splitNumber(n int) (int, int) {
	s := strconv.Itoa(n)
	left, _ := strconv.Atoi(s[:len(s)/2])
	right, _ := strconv.Atoi(s[len(s)/2:])
	return left, right
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/11/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
