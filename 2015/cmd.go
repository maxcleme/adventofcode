package _2015

import (
	_01 "github.com/maxcleme/adventofcode/2015/01"
	_02 "github.com/maxcleme/adventofcode/2015/02"
	_03 "github.com/maxcleme/adventofcode/2015/03"
	_04 "github.com/maxcleme/adventofcode/2015/04"
	_05 "github.com/maxcleme/adventofcode/2015/05"
	_06 "github.com/maxcleme/adventofcode/2015/06"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2015",
}

func init() {
	RootCmd.AddCommand(_01.RootCmd)
	RootCmd.AddCommand(_02.RootCmd)
	RootCmd.AddCommand(_03.RootCmd)
	RootCmd.AddCommand(_04.RootCmd)
	RootCmd.AddCommand(_05.RootCmd)
	RootCmd.AddCommand(_06.RootCmd)
}
