package _02

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Report []int

func parse(input string) ([]Report, error) {
	scanner := bufio.NewScanner(bytes.NewBufferString(input))
	var res []Report
	for scanner.Scan() {
		l := scanner.Text()
		arr := strings.Split(l, " ")
		var row []int
		for _, v := range arr {
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			row = append(row, i)
		}
		res = append(res, row)
	}
	return res, nil
}

func (r Report) safe() bool {
	sumDiff := 0
	for i := 1; i < len(r); i++ {
		a := r[i-1]
		b := r[i]
		delta := a - b
		if delta == 0 {
			return false
		}
		if math.Abs(float64(delta)) > 3 {
			return false
		}
		if sumDiff > 0 && delta < 0 {
			return false
		}
		if sumDiff < 0 && delta > 0 {
			return false
		}
		sumDiff += delta
	}
	return true
}

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2024/02/input")
		if err != nil {
			return err
		}
		reports, err := parse(string(f))
		if err != nil {
			return err
		}
		count := 0
		for _, r := range reports {
			if r.safe() {
				count++
			}
		}
		fmt.Println(count)
		return nil
	},
}
