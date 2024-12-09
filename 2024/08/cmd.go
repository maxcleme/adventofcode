package _08

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "08",
}

func init() {
	RootCmd.AddCommand(part1Cmd)
	RootCmd.AddCommand(part2Cmd)
}
