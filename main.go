package main

import (
	"fmt"
	"os"

	_2015 "github.com/maxcleme/adventofcode/2015"
	"github.com/maxcleme/adventofcode/2022"
	"github.com/maxcleme/adventofcode/2024"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(_2015.RootCmd)
	rootCmd.AddCommand(_2022.RootCmd)
	rootCmd.AddCommand(_2024.RootCmd)
}
