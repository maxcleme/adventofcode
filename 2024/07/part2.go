package _07

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Concat struct{}

func (c Concat) Apply(i1 int, i2 int) int {
	return parseNumber(strconv.Itoa(i1) + strconv.Itoa(i2))
}

func part2(input string) int {
	rows := strings.Split(input, "\n")
	sum := 0
	for _, row := range rows {
		if eq := makeEquation(row); eq.Solve(Plus{}, Mult{}, Concat{}) {
			sum += eq.Result
		}
	}
	return sum
}

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/07/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}
