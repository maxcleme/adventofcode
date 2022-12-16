package _15

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
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
		_, c := RowCoverage(sensors, math.MinInt, math.MaxInt, 2000000)
		fmt.Println(c)
		return nil
	},
}
