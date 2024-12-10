package _10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2_1(t *testing.T) {
	input := `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`
	assert.Equal(t, 3, part2(input))
}

func TestPart2_2(t *testing.T) {
	input := `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
	assert.Equal(t, 13, part2(input))
}

func TestPart2_3(t *testing.T) {
	input := `012345
123456
234567
345678
4.6789
56789.`
	assert.Equal(t, 227, part2(input))
}

func TestPart2_4(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	assert.Equal(t, 81, part2(input))
}
