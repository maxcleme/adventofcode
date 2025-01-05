package _06

import "testing"

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"turn on 0,0 through 0,0", 1},
		{"toggle 0,0 through 999,999", 2000000},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
