package _05

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
		f, err := os.ReadFile("./2025/05/input")
		if err != nil {
			return err
		}
		fmt.Println(part1(string(f)))
		return nil
	},
}

type Range struct {
	Min int
	Max int
}

type Ranges []Range

func (rs Ranges) overlaps(r Range) (bool, int) {
	for i, existing := range rs {
		if r.Min <= existing.Max && r.Max >= existing.Min {
			return true, i
		}
	}
	return false, -1
}

func (rs Ranges) add(r Range) Ranges {
	ok, existing := rs.overlaps(r)
	if !ok {
		return append(rs, r)
	}
	if r.Min < rs[existing].Min {
		rs[existing].Min = r.Min
	}
	if r.Max > rs[existing].Max {
		rs[existing].Max = r.Max
	}
	return rs
}

func (rs Ranges) compact() Ranges {
	curr := len(rs)
	for {
		var compacted Ranges
		for _, r := range rs {
			compacted = compacted.add(r)
		}
		if len(compacted) == curr {
			break
		}
		curr = len(compacted)
		rs = compacted
	}
	return rs
}

func (rs Ranges) contains(v int) bool {
	for _, r := range rs {
		if v >= r.Min && v <= r.Max {
			return true
		}
	}
	return false
}

func part1(payload string) int {
	var ranges Ranges
	for _, r := range strings.Split(strings.Split(payload, "\n\n")[0], "\n") {
		parts := strings.Split(r, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		ranges = ranges.add(Range{a, b})
	}
	ranges = ranges.compact()
	count := 0
	for _, i := range strings.Split(strings.Split(payload, "\n\n")[1], "\n") {
		v, _ := strconv.Atoi(i)
		if ranges.contains(v) {
			count++
		}
	}
	return count
}
