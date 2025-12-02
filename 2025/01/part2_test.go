package _01

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart2(t *testing.T) {
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
	require.Equal(t, 6, lock.part2)
}

func TestPart2_MoreExample(t *testing.T) {
	l := &lock{position: 50}
	l.turn("L", 50)
	l.turn("R", 50)
	require.Equal(t, 1, l.part2)

	l = &lock{position: 50}
	l.turn("L", 50)
	l.turn("L", 50)
	require.Equal(t, 1, l.part2)

	l = &lock{position: 50}
	l.turn("R", 50)
	l.turn("L", 50)
	require.Equal(t, 1, l.part2)

	l = &lock{position: 50}
	l.turn("R", 50)
	l.turn("R", 50)
	require.Equal(t, 1, l.part2)

	l = &lock{position: 50}
	l.turn("L", 150)
	l.turn("L", 50)
	require.Equal(t, 2, l.part2)

	l = &lock{position: 50}
	l.turn("L", 150)
	l.turn("R", 50)
	require.Equal(t, 2, l.part2)

	l = &lock{position: 50}
	l.turn("R", 150)
	l.turn("L", 50)
	require.Equal(t, 2, l.part2)

	l = &lock{position: 50}
	l.turn("R", 150)
	l.turn("R", 50)
	require.Equal(t, 2, l.part2)

	l = &lock{position: 0}
	l.turn("L", 500)
	require.Equal(t, 5, l.part2)

	l = &lock{position: 50}
	l.turn("L", 650)
	require.Equal(t, 7, l.part2)

	l = &lock{position: 1}
	l.turn("L", 2)
	require.Equal(t, 1, l.part2)

	l = &lock{position: 50}
	l.turn("L", 50)
	l.turn("R", 100)
	require.Equal(t, 2, l.part2)

	l = &lock{position: 0}
	l.turn("R", 100)
	require.Equal(t, 1, l.part2)

	l = &lock{position: 50}
	l.turn("L", 50)
	l.turn("L", 1)
	require.Equal(t, 1, l.part2)

	// Negative modulo check
	l = &lock{position: 50}
	l.turn("L", 500)
	require.Equal(t, 5, l.part2)
	require.Equal(t, 50, l.position)

}
