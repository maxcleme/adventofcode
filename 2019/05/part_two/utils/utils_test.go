package utils

import (
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/maxcleme/adventofcode/2019/05/part_two/modes"
	"github.com/maxcleme/adventofcode/2019/05/part_two/opcode"
)

func TestRun(t *testing.T) {
	type args struct {
		instructions []int
		input        []int
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
				[]int{},
			},
			want:    []int{2, 0, 0, 0, 99},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				[]int{2, 3, 0, 3, 99},
				[]int{},
			},
			want:    []int{2, 3, 0, 6, 99},
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				[]int{2, 4, 4, 5, 99, 0},
				[]int{},
			},
			want:    []int{2, 4, 4, 5, 99, 9801},
			wantErr: false,
		},
		{
			name: "4",
			args: args{
				[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
				[]int{},
			},
			want:    []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			wantErr: false,
		},
		{
			name: "unknown opcode",
			args: args{
				[]int{42, 0, 0, 0, 99},
				[]int{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run(tt.args.instructions)

			for _, i := range tt.args.input {
				os.Stdout.Write([]byte(strconv.Itoa(i)))
			}

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

func TestParseInstruction(t *testing.T) {
	type args struct {
		instruction int
	}
	tests := []struct {
		name    string
		args    args
		want    *Instruction
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				instruction: 1002,
			},
			want: &Instruction{
				Opcode: opcode.Mult{},
				Modes:  []modes.Mode{modes.Position{}, modes.Immediate{}, modes.Position{}},
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				instruction: 1102,
			},
			want: &Instruction{
				Opcode: opcode.Mult{},
				Modes:  []modes.Mode{modes.Immediate{}, modes.Immediate{}, modes.Position{}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInstruction(tt.args.instruction)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInstruction() got = %v, want %v", got, tt.want)
			}
		})
	}
}
