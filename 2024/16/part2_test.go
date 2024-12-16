package _16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2_DUMMY(t *testing.T) {
	input := `#####
#E..#
#...#
#..S#
#####`
	assert.Equal(t, 8, part2(input))
}

func TestPart2_1(t *testing.T) {
	input := `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`
	assert.Equal(t, 45, part2(input))
}

func TestPart2_2(t *testing.T) {
	input := `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`
	assert.Equal(t, 64, part2(input))
}

// https://www.reddit.com/r/adventofcode/comments/1hfupai/comment/m2edhyn
func TestPart2_3(t *testing.T) {
	input := `#######
#.....#
#.###.#
#....E#
##.#.##
#S....#
#.###.#
#.....#
#######`
	assert.Equal(t, 10, part2(input))
}
