package _09

import "testing"

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`,
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
