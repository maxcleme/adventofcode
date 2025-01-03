package _2024

import (
	_01 "github.com/maxcleme/adventofcode/2024/01"
	_02 "github.com/maxcleme/adventofcode/2024/02"
	_03 "github.com/maxcleme/adventofcode/2024/03"
	_04 "github.com/maxcleme/adventofcode/2024/04"
	_05 "github.com/maxcleme/adventofcode/2024/05"
	_06 "github.com/maxcleme/adventofcode/2024/06"
	_07 "github.com/maxcleme/adventofcode/2024/07"
	_08 "github.com/maxcleme/adventofcode/2024/08"
	_09 "github.com/maxcleme/adventofcode/2024/09"
	_10 "github.com/maxcleme/adventofcode/2024/10"
	_11 "github.com/maxcleme/adventofcode/2024/11"
	_12 "github.com/maxcleme/adventofcode/2024/12"
	_13 "github.com/maxcleme/adventofcode/2024/13"
	_14 "github.com/maxcleme/adventofcode/2024/14"
	_15 "github.com/maxcleme/adventofcode/2024/15"
	_16 "github.com/maxcleme/adventofcode/2024/16"
	_17 "github.com/maxcleme/adventofcode/2024/17"
	_18 "github.com/maxcleme/adventofcode/2024/18"
	_19 "github.com/maxcleme/adventofcode/2024/19"
	_20 "github.com/maxcleme/adventofcode/2024/20"
	_21 "github.com/maxcleme/adventofcode/2024/21"
	_22 "github.com/maxcleme/adventofcode/2024/22"
	_23 "github.com/maxcleme/adventofcode/2024/23"
	_24 "github.com/maxcleme/adventofcode/2024/24"
	_25 "github.com/maxcleme/adventofcode/2024/25"
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
	RootCmd.AddCommand(_06.RootCmd)
	RootCmd.AddCommand(_07.RootCmd)
	RootCmd.AddCommand(_08.RootCmd)
	RootCmd.AddCommand(_09.RootCmd)
	RootCmd.AddCommand(_10.RootCmd)
	RootCmd.AddCommand(_11.RootCmd)
	RootCmd.AddCommand(_12.RootCmd)
	RootCmd.AddCommand(_13.RootCmd)
	RootCmd.AddCommand(_14.RootCmd)
	RootCmd.AddCommand(_15.RootCmd)
	RootCmd.AddCommand(_16.RootCmd)
	RootCmd.AddCommand(_17.RootCmd)
	RootCmd.AddCommand(_18.RootCmd)
	RootCmd.AddCommand(_19.RootCmd)
	RootCmd.AddCommand(_20.RootCmd)
	RootCmd.AddCommand(_21.RootCmd)
	RootCmd.AddCommand(_22.RootCmd)
	RootCmd.AddCommand(_23.RootCmd)
	RootCmd.AddCommand(_24.RootCmd)
	RootCmd.AddCommand(_25.RootCmd)
}
