package _06

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		payload, err := os.ReadFile("./2022/06/input")
		if err != nil {
			return err
		}
		device := NewDevice()
		device.onStartPacket = func(i int) {
			fmt.Println(i)
			os.Exit(0)
		}
		go device.Start()
		for _, r := range payload {
			device.input <- rune(r)
		}
		return nil
	},
}
