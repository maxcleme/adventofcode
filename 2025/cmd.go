package _2025

import (
	_01 "github.com/maxcleme/adventofcode/2025/01"
	_02 "github.com/maxcleme/adventofcode/2025/02"
	_03 "github.com/maxcleme/adventofcode/2025/03"
	_04 "github.com/maxcleme/adventofcode/2025/04"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2025",
}

func init() {
	RootCmd.AddCommand(_01.RootCmd)
	RootCmd.AddCommand(_02.RootCmd)
	RootCmd.AddCommand(_03.RootCmd)
	RootCmd.AddCommand(_04.RootCmd)
}
