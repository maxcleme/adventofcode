package _22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2_1(t *testing.T) {
	input := `1
2
3
2024`
	assert.Equal(t, 23, part2(input))
}
