package _17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2_1(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`
	assert.Equal(t, 117440, part2(input))
}
