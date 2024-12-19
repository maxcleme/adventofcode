package _19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2_1(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`
	assert.Equal(t, 16, part2(input))
}
