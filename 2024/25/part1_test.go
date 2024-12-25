package _25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	input := `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`
	assert.Equal(t, 3, part1(input))
}
