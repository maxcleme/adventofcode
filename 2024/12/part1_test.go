package _12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	assert.Equal(t, 140, part1(input))
}

func TestPart1_2(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	assert.Equal(t, 772, part1(input))
}

func TestPart1_3(t *testing.T) {
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
	assert.Equal(t, 1930, part1(input))
}
