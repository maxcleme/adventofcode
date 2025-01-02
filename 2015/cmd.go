package _2015

import (
	_01 "github.com/maxcleme/adventofcode/2015/01"
	_02 "github.com/maxcleme/adventofcode/2015/02"
	_03 "github.com/maxcleme/adventofcode/2015/03"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2015",
}

func init() {
	RootCmd.AddCommand(_01.RootCmd)
	RootCmd.AddCommand(_02.RootCmd)
	RootCmd.AddCommand(_03.RootCmd)
}
