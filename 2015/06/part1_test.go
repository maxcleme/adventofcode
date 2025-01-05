package _06

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"turn on 0,0 through 999,999", 1_000_000},
		{"toggle 0,0 through 999,0", 1_000},
		{"turn off 499,499 through 500,500", 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
