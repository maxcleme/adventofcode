package _25

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "25",
}

func init() {
	RootCmd.AddCommand(part1Cmd)
}
