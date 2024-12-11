package _11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	input := `125 17`
	assert.Equal(t, 55312, part1(input))
}
