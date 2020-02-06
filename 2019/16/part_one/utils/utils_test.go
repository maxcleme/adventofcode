package utils

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func p(sequence Sequence) *Sequence {
	return &sequence
}

func TestParse(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *Sequence
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
			want:    p("12345678"),
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
		pattern    Pattern
		phaseCount int
	}
	tests := []struct {
		name string
		args args
		want *Sequence
	}{
		{
			name: "0 -> 1",
			args: args{
				sequence:   "12345678",
				pattern:    []int{0, 1, 0, -1},
				phaseCount: 1,
			},
			want: p("48226158"),
		},
		{
			name: "1 -> 2",
			args: args{
				sequence:   "48226158",
				pattern:    []int{0, 1, 0, -1},
				phaseCount: 1,
			},
			want: p("34040438"),
		},
		{
			name: "2 -> 3",
			args: args{
				sequence:   "34040438",
				pattern:    []int{0, 1, 0, -1},
				phaseCount: 1,
			},
			want: p("03415518"),
		},
		{
			name: "3 -> 4",
			args: args{
				sequence:   "03415518",
				pattern:    []int{0, 1, 0, -1},
				phaseCount: 1,
			},
			want: p("01029498"),
		},
		{
			name: "1 -> 4",
			args: args{
				sequence:   "12345678",
				pattern:    []int{0, 1, 0, -1},
				phaseCount: 4,
			},
			want: p("01029498"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Apply(tt.args.sequence, tt.args.pattern, tt.args.phaseCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
