package utils

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Sequence
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				reader: strings.NewReader(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid",
			args: args{
				reader: strings.NewReader("123234bc12389713"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid",
			args: args{
				reader: strings.NewReader("12345678"),
			},
			want:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		sequence   Sequence
		phaseCount int
		repeat     int
		offset     int
	}
	tests := []struct {
		name string
		args args
		want Sequence
	}{
		{
			name: "",
			args: args{
				sequence:   []int{0, 3, 0, 3, 6, 7, 3, 2, 5, 7, 7, 2, 1, 2, 9, 4, 4, 0, 6, 3, 4, 9, 1, 5, 6, 5, 4, 7, 4, 6, 6, 4},
				phaseCount: 100,
				repeat:     10000,
				offset:     303673,
			},
			want: []int{8, 4, 4, 6, 2, 0, 2, 6},
		},
		{
			name: "",
			args: args{
				sequence:   []int{0, 2, 9, 3, 5, 1, 0, 9, 6, 9, 9, 9, 4, 0, 8, 0, 7, 4, 0, 7, 5, 8, 5, 4, 4, 7, 0, 3, 4, 3, 2, 3},
				phaseCount: 100,
				repeat:     10000,
				offset:     293510,
			},
			want: []int{7, 8, 7, 2, 5, 2, 7, 0},
		},
		{
			name: "",
			args: args{
				sequence:   []int{0, 3, 0, 8, 1, 7, 7, 0, 8, 8, 4, 9, 2, 1, 9, 5, 9, 7, 3, 1, 1, 6, 5, 4, 4, 6, 8, 5, 0, 5, 1, 7},
				phaseCount: 100,
				repeat:     10000,
				offset:     308177,
			},
			want: []int{5, 3, 5, 5, 3, 7, 3, 1},
		},
		{
			name: "",
			args: args{
				sequence:   []int{1, 2, 3, 4, 5, 6, 7, 8},
				phaseCount: 1,
				repeat:     2,
				offset:     8,
			},
			want: []int{6, 5, 3, 0, 6, 1, 5, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Apply(tt.args.sequence, tt.args.phaseCount, tt.args.repeat, tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		sequence Sequence
		times    int
	}
	tests := []struct {
		name string
		args args
		want Sequence
	}{
		{
			name: "0 times",
			args: args{
				sequence: []int{1, 2, 3, 4, 5},
				times:    0,
			},
			want: []int{},
		},
		{
			name: "1 times",
			args: args{
				sequence: []int{1, 2, 3, 4, 5},
				times:    1,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "3 times",
			args: args{
				sequence: []int{1, 2, 3, 4, 5},
				times:    3,
			},
			want: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.sequence, tt.args.times); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOffset(t *testing.T) {
	type args struct {
		sequence Sequence
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "too short",
			args: args{
				sequence: []int{1, 2, 3, 4},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "valid",
			args: args{
				sequence: []int{0, 0, 0, 0, 0, 1, 7, 1, 2, 3, 4, 5},
			},
			want:    17,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Offset(tt.args.sequence)
			if (err != nil) != tt.wantErr {
				t.Errorf("Offset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Offset() got = %v, want %v", got, tt.want)
			}
		})
	}
}
