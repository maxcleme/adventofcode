package _05

import "strings"

type crate struct {
	content []rune
}

func MakeCrate() crate {
	return crate{}
}

func (c *crate) Pop() rune {
	p := c.content[0]
	c.content = c.content[1:]
	return p
}

func (c *crate) Add(i string) {
	i = strings.TrimSpace(i)
	if i == "" {
		return
	}
	c.content = append(c.content, rune(i[1]))
}

func (c *crate) Push(r rune) {
	i := []rune{r}
	c.content = append(i, c.content...)
}

func (c crate) String() string {
	return "[" + string(c.content) + "]"
}

func Move9000(crates []crate, count int, from int, to int) []crate {
	for i := 0; i < count; i++ {
		r := crates[from-1].Pop()
		crates[to-1].Push(r)
	}
	return crates
}

func Move9001(crates []crate, count int, from int, to int) []crate {
	var r []rune
	for i := 0; i < count; i++ {
		r = append(r, crates[from-1].Pop())
	}
	for i := len(r)-1; i >= 0; i-- {
		crates[to-1].Push(r[i])
	}
	return crates
}
