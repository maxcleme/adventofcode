package _24

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "24",
}

func init() {
	RootCmd.AddCommand(part1Cmd)
	RootCmd.AddCommand(part2Cmd)
}
