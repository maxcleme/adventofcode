package _03

type Rucksack struct {
	first  Compartment
	second Compartment
}

type Compartment struct {
	content string
	index   map[rune]struct{}
}

func (c Compartment) Contains(items []rune) []rune {
	var shared []rune
	for _, i := range items {
		if _, ok := c.index[i]; ok {
			shared = append(shared, i)
		}
	}
	return shared
}

func (r Rucksack) Shared() []rune {
	var shared []rune
	for i := range r.first.index {
		if _, ok := r.second.index[i]; ok {
			shared = append(shared, i)
		}
	}
	return shared
}

func (r Rucksack) Contains(items []rune) []rune {
	var contains []rune
	contains = append(contains, r.first.Contains(items)...)
	contains = append(contains, r.second.Contains(items)...)
	return contains
}

func (r Rucksack) Unique() []rune {
	s := []rune(r.first.content + r.second.content)
	index := map[rune]struct{}{}
	var unique []rune
	for _, i := range s {
		if _, ok := index[i]; ok {
			continue
		}
		index[i] = struct{}{}
		unique = append(unique, i)
	}
	return unique
}

func Shared(rs ...Rucksack) []rune {
	shared := rs[0].Unique()
	for _, r := range rs[1:] {
		shared = r.Contains(shared)
	}
	return shared
}

func Priority(i rune) int {
	if i >= 'a' {
		return int(i - 96)
	}
	return int(i - 65 + 27)
}

func Priorities(items []rune) int {
	var s int
	for _, i := range items {
		s += Priority(i)
	}
	return s
}

func ParseRucksack(s string) Rucksack {
	return Rucksack{
		first:  ParseCompartment(s[:len(s)/2]),
		second: ParseCompartment(s[len(s)/2:]),
	}
}

func ParseCompartment(content string) Compartment {
	index := map[rune]struct{}{}
	for _, c := range content {
		index[c] = struct{}{}
	}
	return Compartment{content: content, index: index}
}
