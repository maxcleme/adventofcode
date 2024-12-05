package _2024

import (
	_01 "github.com/maxcleme/adventofcode/2024/01"
	_02 "github.com/maxcleme/adventofcode/2024/02"
	_03 "github.com/maxcleme/adventofcode/2024/03"
	_04 "github.com/maxcleme/adventofcode/2024/04"
	_05 "github.com/maxcleme/adventofcode/2024/05"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2024",
}

func init() {
	RootCmd.AddCommand(_01.RootCmd)
	RootCmd.AddCommand(_02.RootCmd)
	RootCmd.AddCommand(_03.RootCmd)
	RootCmd.AddCommand(_04.RootCmd)
	RootCmd.AddCommand(_05.RootCmd)
}
