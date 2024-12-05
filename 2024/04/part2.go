package _04

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/04/input")
		if err != nil {
			return err
		}
		fmt.Println(countMASPart2(string(f)))
		return nil
	},
}

func countMASPart2(input string) int {
	count := 0
	rows := strings.Split(input, "\n")
	for i, row := range rows[1 : len(rows)-1] {
		for j, r := range row[1 : len(row)-1] {
			if r == 'A' {
				a := string(rows[i][j]) + string(rows[i+1][j+1]) + string(rows[i+2][j+2])
				b := string(rows[i][j+2]) + string(rows[i+1][j+1]) + string(rows[i+2][j])
				if (a == "MAS" || a == "SAM") && (b == "MAS" || b == "SAM") {
					count++
				}
			}
		}
	}
	return count
}
