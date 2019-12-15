package utils

import (
	"os"
	"reflect"
	"testing"

	"github.com/maxcleme/adventofcode/2019/09/part_one/model"
	"github.com/maxcleme/adventofcode/2019/09/part_one/modes"
	"github.com/maxcleme/adventofcode/2019/09/part_one/opcode"
)

func TestRun(t *testing.T) {
	type args struct {
		program model.Program
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
				program: model.Program{
					Statements: []int{
						109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
					},
					In:  os.Stdin,
					Out: os.Stdout,
				},
			},
			want: []int{
				109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run(tt.args.program)

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
		want    *opcode.Instruction
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				instruction: 1002,
			},
			want: &opcode.Instruction{
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
			want: &opcode.Instruction{
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
