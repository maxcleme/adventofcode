package _05

import "testing"

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"qjhvhtzxzqqjkmpb", 1},
		{"xxyxx", 1},
		{"uurcxstgmygtbstg", 0},
		{"ieodomkazucvgmuy", 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
