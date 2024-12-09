package _09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	input := `2333133121414131402`
	assert.Equal(t, 2858, part2(input))
}
