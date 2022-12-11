package _11

func Rounds(monkeys []Monkey, w WorryManagement, length int) []Monkey {
	for i := 0; i < length; i++ {
		monkeys = Round(monkeys, w)
	}
	return monkeys
}

func Round(monkeys []Monkey, w WorryManagement) []Monkey {
	s := 1
	for _, m := range monkeys {
		s *= m.Divider
	}
	for i, m := range monkeys {
		for _, item := range m.Items {
			item = m.Operation(item, s)
			item = w(item)
			m.Test(item, m.Divider, monkeys)
		}
		monkeys[i].InspectionCount += len(monkeys[i].Items)
		monkeys[i].Items = nil
	}
	return monkeys
}
