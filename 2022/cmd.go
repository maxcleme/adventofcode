package _2022

import (
	"github.com/maxcleme/adventofcode/2022/01"
	"github.com/maxcleme/adventofcode/2022/02"
	"github.com/maxcleme/adventofcode/2022/03"
	"github.com/maxcleme/adventofcode/2022/04"
	"github.com/maxcleme/adventofcode/2022/05"
	"github.com/maxcleme/adventofcode/2022/06"
	"github.com/maxcleme/adventofcode/2022/07"
	"github.com/maxcleme/adventofcode/2022/08"
	"github.com/maxcleme/adventofcode/2022/09"
	"github.com/maxcleme/adventofcode/2022/10"
	"github.com/maxcleme/adventofcode/2022/11"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "2022",
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
	RootCmd.AddCommand(_10.RootCmd)
	RootCmd.AddCommand(_11.RootCmd)
}
