package _11

import (
	"strconv"
	"strings"
)

type Monkey struct {
	Items           []int
	Operation       Operation
	Divider         int
	Test            Test
	InspectionCount int
}

type Operation func(priority int, d int) int

type Test func(priority int, d int, monkeys []Monkey)

func ParseMonkey(lines []string) Monkey {
	return Monkey{
		Items:     ParseItems(strings.TrimSpace(lines[0])),
		Operation: ParseOperation(strings.TrimSpace(lines[1])),
		Divider:   ParseDivider(strings.TrimSpace(lines[2])),
		Test:      ParseTest(lines[3:]),
	}
}

func ParseTest(lines []string) Test {
	t := ParseThrow(strings.TrimPrefix(strings.TrimSpace(lines[0]), "If true: "))
	f := ParseThrow(strings.TrimPrefix(strings.TrimSpace(lines[1]), "If false: "))
	return func(priority int, d int, monkeys []Monkey) {
		if priority%d == 0 {
			t(priority, monkeys)
			return
		}
		f(priority, monkeys)
	}
}

func ParseThrow(l string) func(priority int, monkeys []Monkey) {
	l = strings.TrimPrefix(l, "throw to monkey ")
	dest, _ := strconv.Atoi(l)
	return func(priority int, monkeys []Monkey) {
		monkeys[dest].Items = append(monkeys[dest].Items, priority)
	}
}

func ParseDivider(l string) int {
	d, _ := strconv.Atoi(strings.TrimPrefix(l, "Test: divisible by "))
	return d
}

func ParseOperation(l string) Operation {
	l = strings.TrimPrefix(l, "Operation: new = old ")
	op := l[0]
	v := func(p int) int {
		return p
	}
	if l[2:] != "old" {
		n, _ := strconv.Atoi(l[2:])
		v = func(p int) int {
			return n
		}
	}
	switch op {
	case '+':
		return func(priority int, d int) int {
			return (priority % d) + (v(priority) % d)
		}
	case '*':
		return func(priority int, d int) int {
			return (priority % d) * (v(priority) % d)
		}
	}
	return nil
}

func ParseItems(l string) []int {
	l = strings.TrimPrefix(l, "Starting items: ")
	args := strings.Split(l, ",")
	var items []int
	for _, arg := range args {
		n, _ := strconv.Atoi(strings.TrimSpace(arg))
		items = append(items, n)
	}
	return items
}
