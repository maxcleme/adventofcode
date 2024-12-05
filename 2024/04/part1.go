package _04

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func count(input, word string) int {
	rows := strings.Split(input, "\n")
	h := strings.Join(rows, "#")
	v := strings.Join(verticals(rows), "#")
	d := strings.Join(append(diagonalsLeft(rows), diagonalsRight(rows)...), "#")
	count := 0
	count += strings.Count(h, word)
	count += strings.Count(h, reverse(word))
	count += strings.Count(v, word)
	count += strings.Count(v, reverse(word))
	count += strings.Count(d, word)
	count += strings.Count(d, reverse(word))
	return count
}

func verticals(rows []string) []string {
	var result []string
	var vertical string
	for i := 0; i < len(rows[0]); i++ {
		for _, r := range rows {
			vertical += string(r[i])
		}
		result = append(result, vertical)
		vertical = ""
	}
	return result
}

func diagonalsLeft(rows []string) []string {
	if len(rows) == 0 {
		return nil
	}
	cols := len(rows[0])
	var diagonals []string
	for start := 0; start < len(rows)+cols-1; start++ {
		var diagonal []rune
		for y := 0; y <= start; y++ {
			x := start - y
			if y >= 0 && y < len(rows) && x >= 0 && x < cols {
				diagonal = append(diagonal, rune(rows[y][x]))
			}
		}
		diagonals = append(diagonals, string(diagonal))
	}
	return diagonals
}
func diagonalsRight(rows []string) []string {
	if len(rows) == 0 {
		return nil
	}
	cols := len(rows[0])
	var diagonals []string
	for start := 0; start < len(rows)+cols-1; start++ {
		var diagonal []rune
		for y := 0; y <= start; y++ {
			x := cols - 1 - (start - y)
			if y >= 0 && y < len(rows) && x >= 0 && x < cols {
				diagonal = append(diagonal, rune(rows[y][x]))
			}
		}
		diagonals = append(diagonals, string(diagonal))
	}
	return diagonals
}

func reverse(word string) string {
	result := ""
	for _, v := range word {
		result = string(v) + result
	}
	return result
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/04/input")
		if err != nil {
			return err
		}
		fmt.Println(count(string(f), "XMAS"))
		return nil
	},
}
