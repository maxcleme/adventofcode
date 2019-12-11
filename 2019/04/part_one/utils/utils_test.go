package utils

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		password int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "six digit number", args: args{password: 123}, want: false},
		{name: "no double", args: args{password: 123789}, want: false},
		{name: "decreasing", args: args{password: 223450}, want: false},
		{name: "valid", args: args{password: 111111}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.password); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
