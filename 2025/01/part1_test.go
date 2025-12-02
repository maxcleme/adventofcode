package _01

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	lock := &lock{position: 50}
	lock.turn("L", 68)
	lock.turn("L", 30)
	lock.turn("R", 48)
	lock.turn("L", 5)
	lock.turn("R", 60)
	lock.turn("L", 55)
	lock.turn("L", 1)
	lock.turn("L", 99)
	lock.turn("R", 14)
	lock.turn("L", 82)
	require.Equal(t, 3, lock.part1)
}
