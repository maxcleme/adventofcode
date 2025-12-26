package _11

import "testing"

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`,
			want: 2,
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
