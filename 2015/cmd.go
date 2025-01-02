package _2015

import (
	_01 "github.com/maxcleme/adventofcode/2015/01"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2015",
}

func init() {
	RootCmd.AddCommand(_01.RootCmd)
}
