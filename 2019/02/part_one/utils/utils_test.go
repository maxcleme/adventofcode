package utils

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		instructions []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				[]int{1, 0, 0, 0, 99},
			},
			want:    []int{2, 0, 0, 0, 99},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				[]int{2, 3, 0, 3, 99},
			},
			want:    []int{2, 3, 0, 6, 99},
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				[]int{2, 4, 4, 5, 99, 0},
			},
			want:    []int{2, 4, 4, 5, 99, 9801},
			wantErr: false,
		},
		{
			name: "4",
			args: args{
				[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			},
			want:    []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			wantErr: false,
		},
		{
			name: "unknown opcode",
			args: args{
				[]int{7, 0, 0, 0, 99},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run(tt.args.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
