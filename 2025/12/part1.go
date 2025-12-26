package _12

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use: "part1",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile("./2025/12/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

type Region struct {
	Width, Height int
	Constraints   []int
}

func (r Region) Area() int {
	return r.Width * r.Height
}

type Shape struct {
	Pattern []string
}

func (s Shape) Area() int {
	area := 0
	for _, row := range s.Pattern {
		area += strings.Count(row, "#")
	}
	return area
}

func part1(payload string) int {
	rows := strings.Split(payload, "\n")
	i := 0
	var shapes []Shape
	var regions []Region
	for ; i < len(rows); i++ {
		switch {
		case strings.Contains(rows[i], "x"):
			parts := strings.Split(rows[i], ":")
			w, _ := strconv.Atoi(strings.Split(parts[0], "x")[0])
			h, _ := strconv.Atoi(strings.Split(parts[0], "x")[1])
			r := Region{Width: w, Height: h}
			for _, c := range strings.Fields(parts[1]) {
				cv, _ := strconv.Atoi(c)
				r.Constraints = append(r.Constraints, cv)
			}
			regions = append(regions, r)
		case strings.Contains(rows[i], ":"):
			i++
			var shape []string
			for ; i < len(rows); i++ {
				next := strings.TrimSpace(rows[i])
				if next == "" {
					break
				}
				shape = append(shape, next)
			}
			shapes = append(shapes, Shape{Pattern: shape})
		}
	}
	total := 0
	for _, region := range regions {
		needed := 0
		for i, constraint := range region.Constraints {
			needed += constraint * shapes[i].Area()
		}
		if needed <= region.Area() {
			total++
		}
	}
	return total
}
