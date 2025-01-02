package _01

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
