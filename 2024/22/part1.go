package _22

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func part1(input string) int {
	total := 0
	for _, row := range strings.Split(input, "\n") {
		secret, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
		total += nth(secret, 2000)
	}
	return total
}

func nth(secret int, i int) int {
	for range i {
		secret = next(secret)
	}
	return secret
}

func next(secret int) int {
	secret = prune(mix(secret, secret*64))
	secret = prune(mix(secret, secret/32))
	secret = prune(mix(secret, secret*2048))
	return secret
}

func mix(secret int, n int) int {
	return secret ^ n
}

func prune(secret int) int {
	return secret % 16777216
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/22/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}
