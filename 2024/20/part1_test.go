package _20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	input := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
	assert.Equal(t, 1, part1(input, 64))
	assert.Equal(t, 2, part1(input, 40))
	assert.Equal(t, 3, part1(input, 38))
	assert.Equal(t, 4, part1(input, 36))
	assert.Equal(t, 5, part1(input, 20))
	assert.Equal(t, 8, part1(input, 12))
	assert.Equal(t, 10, part1(input, 10))
	assert.Equal(t, 14, part1(input, 8))
	assert.Equal(t, 16, part1(input, 6))
	assert.Equal(t, 30, part1(input, 4))
	assert.Equal(t, 44, part1(input, 2))
}
