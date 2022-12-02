package main

import "testing"

func Test_round(t *testing.T) {
	type args struct {
		elveInput Input
		myInput   Input
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				elveInput: Rock,
				myInput:   Rock,
			},
			want: 3 + 1,
		},
		{
			args: args{
				elveInput: Rock,
				myInput:   Paper,
			},
			want: 6 + 2,
		},
		{
			args: args{
				elveInput: Rock,
				myInput:   Scissors,
			},
			want: 0 + 3,
		},
		{
			args: args{
				elveInput: Paper,
				myInput:   Rock,
			},
			want: 0 + 1,
		},
		{
			args: args{
				elveInput: Paper,
				myInput:   Paper,
			},
			want: 3 + 2,
		},
		{
			args: args{
				elveInput: Paper,
				myInput:   Scissors,
			},
			want: 6 + 3,
		},
		{
			args: args{
				elveInput: Scissors,
				myInput:   Rock,
			},
			want: 6 + 1,
		},
		{
			args: args{
				elveInput: Scissors,
				myInput:   Paper,
			},
			want: 0 + 2,
		},
		{
			args: args{
				elveInput: Scissors,
				myInput:   Scissors,
			},
			want: 3 + 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := round(tt.args.elveInput, tt.args.myInput); got != tt.want {
				t.Errorf("round() = %v, want %v", got, tt.want)
			}
		})
	}
}
