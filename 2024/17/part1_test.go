package _17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", part1(input))
}

func TestPart1_2(t *testing.T) {
	input := `Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4
`
	assert.Equal(t, "0,1,2", part1(input))
}

func TestPart1_3(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`
	assert.Equal(t, "4,2,5,6,7,7,7,7,3,1,0", part1(input))
}

func TestPart1_4(t *testing.T) {
	input := `Register A: 117440
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`
	assert.Equal(t, "0,3,5,4,3,0", part1(input))
}
