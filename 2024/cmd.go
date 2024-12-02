package _2024

import (
	_01 "github.com/maxcleme/adventofcode/2024/01"
	_02 "github.com/maxcleme/adventofcode/2024/02"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2024",
}

func init() {
	RootCmd.AddCommand(_01.RootCmd)
	RootCmd.AddCommand(_02.RootCmd)
}
