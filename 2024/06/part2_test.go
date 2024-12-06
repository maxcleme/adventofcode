package _06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2_1(t *testing.T) {
	input := `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	assert.Equal(t, 6, part2(input))
}

func TestPart2_2(t *testing.T) {
	input := `
.#..
#..#
....
^...
#...
.#..`
	assert.Equal(t, 1, part2(input))
}

func TestPart2_3(t *testing.T) {
	input := `
.#....
.....#
...^..
#.....
....#.`
	assert.Equal(t, 1, part2(input))
}

func TestPart2_4(t *testing.T) {
	input := `
....#.
.#.#.#
......
.^.#.#
....#.`
	assert.Equal(t, 1, part2(input))
}

func TestPart2_5(t *testing.T) {
	input := `
....
#..#
.^#,
`
	assert.Equal(t, 1, part2(input))
}
