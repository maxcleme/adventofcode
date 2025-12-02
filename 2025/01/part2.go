package _01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2025/01/input")
		if err != nil {
			return err
		}

		lock := &lock{position: 50}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			l := scanner.Text()
			direction := l[0:1]
			step, err := strconv.Atoi(l[1:])
			if err != nil {
				return err
			}
			lock.turn(direction, step)
		}
		fmt.Println(lock.part2)
		return nil
	},
}
