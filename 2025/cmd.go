package _2025

import (
	_01 "github.com/maxcleme/adventofcode/2025/01"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2025",
}

func init() {
	RootCmd.AddCommand(_01.RootCmd)
}
