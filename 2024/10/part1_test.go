package _10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	input := `0123
1234
8765
9876`
	assert.Equal(t, 1, part1(input))
}

func TestPart1_2(t *testing.T) {
	input := `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`
	assert.Equal(t, 2, part1(input))
}

func TestPart1_3(t *testing.T) {
	input := `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
	assert.Equal(t, 4, part1(input))
}

func TestPart1_4(t *testing.T) {
	input := `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`
	assert.Equal(t, 3, part1(input))
}

func TestPart1_5(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	assert.Equal(t, 36, part1(input))
}
