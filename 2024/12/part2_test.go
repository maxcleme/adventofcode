package _12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2_1(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	assert.Equal(t, 80, part2(input))
}

func TestPart2_2(t *testing.T) {
	input := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	assert.Equal(t, 236, part2(input))
}

func TestPart2_3(t *testing.T) {
	input := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	assert.Equal(t, 368, part2(input))
}

func TestPart2_4(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	assert.Equal(t, 1206, part2(input))
}
