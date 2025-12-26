package _11

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/11/input")
		if err != nil {
			return err
		}
		fmt.Println(part2(string(f)))
		return nil
	},
}

func part2(payload string) int {
	var g Graph
	for _, line := range strings.Split(payload, "\n") {
		parts := strings.Split(line, ":")
		from := parts[0]
		for _, to := range strings.Fields(parts[1]) {
			g.AddEdge(from, to)
		}
	}
	return g.Paths("svr", "dac")*g.Paths("dac", "fft")*g.Paths("fft", "out") +
		g.Paths("svr", "fft")*g.Paths("fft", "dac")*g.Paths("dac", "out")
}
