package _09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `2333133121414131402`
	assert.Equal(t, 1928, part1(input))
}
