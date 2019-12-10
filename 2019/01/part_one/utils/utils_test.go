package utils

import "testing"

func TestRequiredFuel(t *testing.T) {
	type args struct {
		mass float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{12},
			want: 2,
		},
		{
			name: "2",
			args: args{14},
			want: 2,
		},
		{
			name: "3",
			args: args{1969},
			want: 654,
		},
		{
			name: "4",
			args: args{100756},
			want: 33583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequiredFuel(tt.args.mass); got != tt.want {
				t.Errorf("RequiredFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}