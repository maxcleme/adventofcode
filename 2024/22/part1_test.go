package _22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	input := `1
10
100
2024`
	assert.Equal(t, 37327623, part1(input))
}
