package _15

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var part2Cmd = &cobra.Command{
	Use: "part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open("./2022/15/input")
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		var sensors []Sensor
		for scanner.Scan() {
			l := scanner.Text()
			sensors = append(sensors, MustParseSensor(l))
		}
		fmt.Println(Locate(sensors, 0, 4000000).TuningFrequency())
		return nil
	},
}
