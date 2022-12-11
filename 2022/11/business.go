package _11

import (
	"sort"
)

func Business(monkeys []Monkey) int {
	m := append(monkeys) // copy
	sort.Slice(m, func(i, j int) bool {
		return m[i].InspectionCount > m[j].InspectionCount
	})
	return m[0].InspectionCount * m[1].InspectionCount
}
