package _2022

import (
	"github.com/maxcleme/adventofcode/2022/01"
	"github.com/maxcleme/adventofcode/2022/02"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2022",
}

func init() {
	RootCmd.AddCommand(_01.RootCmd)
	RootCmd.AddCommand(_02.RootCmd)
}
