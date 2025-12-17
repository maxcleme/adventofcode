package _05

import "testing"

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
