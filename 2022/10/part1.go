package _10

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/10/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)

		var sum int
		var wg sync.WaitGroup
		wg.Add(6)
		registry := NewRegistry()
		clock := NewClock(func(n int) {
			switch n {
			case 20, 60, 100, 140, 180, 220:
				sum += n * registry["X"]
				wg.Done()
			}
		})
		tick := clock.Start()
		for scanner.Scan() {
			l := scanner.Text()
			i := ParseInstruction(l)
			i.Run(tick, clock.next, registry)
		}
		wg.Wait()
		fmt.Println(sum)
		return nil
	},
}
