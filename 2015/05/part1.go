package _05

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2015/05/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

func part1(input string) int {
	total := 0
	for _, row := range strings.Split(input, "\n") {
		if !isNice(row) {
			continue
		}
		total++
	}
	return total
}

func isNice(word string) bool {
	if strings.Contains(word, "ab") || strings.Contains(word, "cd") || strings.Contains(word, "pq") || strings.Contains(word, "xy") {
		return false
	}
	vowels := 0
	for _, r := range word {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
	}
	if vowels < 3 {
		return false
	}
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			return true
		}
	}
	return false
}
