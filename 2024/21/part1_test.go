package _21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	input := `029A
980A
179A
456A
379A`
	assert.Equal(t, 126384, part1(input))
}
