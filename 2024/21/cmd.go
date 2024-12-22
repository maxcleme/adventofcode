package _21

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "21",
}

func init() {
	RootCmd.AddCommand(part1Cmd)
	RootCmd.AddCommand(part2Cmd)
}
