package _2025

import (
	_01 "github.com/maxcleme/adventofcode/2025/01"
	_02 "github.com/maxcleme/adventofcode/2025/02"
	_03 "github.com/maxcleme/adventofcode/2025/03"
	_04 "github.com/maxcleme/adventofcode/2025/04"
	_05 "github.com/maxcleme/adventofcode/2025/05"
	_06 "github.com/maxcleme/adventofcode/2025/06"
	_07 "github.com/maxcleme/adventofcode/2025/07"
	_08 "github.com/maxcleme/adventofcode/2025/08"
	_09 "github.com/maxcleme/adventofcode/2025/09"

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
	RootCmd.AddCommand(_05.RootCmd)
	RootCmd.AddCommand(_06.RootCmd)
	RootCmd.AddCommand(_07.RootCmd)
	RootCmd.AddCommand(_08.RootCmd)
	RootCmd.AddCommand(_09.RootCmd)
}
