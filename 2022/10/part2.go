package _10

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/10/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)

		var wg sync.WaitGroup
		wg.Add(1)
		registry := NewRegistry()
		clock := NewClock(func(n int) {
			switch (n - 1) % 40 {
			case registry["X"], registry["X"] - 1, registry["X"] + 1:
				fmt.Print("#")
			default:
				fmt.Print(" ")
			}
			if n%40 == 0 {
				fmt.Println()
			}
			if n == 240 {
				wg.Done()
				return
			}
		})
		tick := clock.Start()
		for scanner.Scan() {
			l := scanner.Text()
			i := ParseInstruction(l)
			i.Run(tick, clock.next, registry)
		}
		wg.Wait()
		return nil
	},
}
