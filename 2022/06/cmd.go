package _06

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "06",
}

func init() {
	RootCmd.AddCommand(part1Cmd)
	RootCmd.AddCommand(part2Cmd)
}
