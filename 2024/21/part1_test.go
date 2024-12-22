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

func Test_n2d(t *testing.T) {
	tests := []struct {
		from string
		to   string
		want string
	}{
		{
			from: "",
			to:   "0",
			want: "<A",
		},
		{
			from: "0",
			to:   "2",
			want: "^A",
		},
		{
			from: "2",
			to:   "9",
			want: "^^>A",
		},
		{
			from: "9",
			to:   "A",
			want: "vvvA",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, tt.want, n2d(tt.from, tt.to), "n2d(%v, %v)", tt.from, tt.to)
		})
	}
}

func Test_d2d(t *testing.T) {
	tests := []struct {
		from string
		to   string
		want string
	}{
		{
			from: "",
			to:   "<",
			want: "v<<A",
		},
		{
			from: "<",
			to:   "A",
			want: ">>^A",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, tt.want, d2d(tt.from, tt.to), "d2d(%v, %v)", tt.from, tt.to)
		})
	}
}
