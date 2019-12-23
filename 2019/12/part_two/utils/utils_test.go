package utils

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	sample := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []*Object
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				reader: strings.NewReader(sample),
			},
			want: []*Object{
				{
					Position: [3]int{-1, 0, 2},
					Velocity: [3]int{0, 0, 0},
				},
				{
					Position: [3]int{2, -10, -7},
					Velocity: [3]int{0, 0, 0},
				},
				{
					Position: [3]int{4, -8, 8},
					Velocity: [3]int{0, 0, 0},
				},
				{
					Position: [3]int{3, 5, -1},
					Velocity: [3]int{0, 0, 0},
				},
			},
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

func TestSimulate(t *testing.T) {
	type args struct {
		objects []*Object
		nth     int
	}
	tests := []struct {
		name string
		args args
		want []*Object
	}{
		{
			name: "after 0 step",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 0,
			},
			want: []*Object{
				{[3]int{-1, 0, 2}, [3]int{0, 0, 0}},
				{[3]int{2, -10, -7}, [3]int{0, 0, 0}},
				{[3]int{4, -8, 8}, [3]int{0, 0, 0}},
				{[3]int{3, 5, -1}, [3]int{0, 0, 0}},
			},
		},

		{
			name: "after 1 step",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 1,
			},
			want: []*Object{
				{[3]int{2, -1, 1}, [3]int{3, -1, -1}},
				{[3]int{3, -7, -4}, [3]int{1, 3, 3}},
				{[3]int{1, -7, 5}, [3]int{-3, 1, -3}},
				{[3]int{2, 2, 0}, [3]int{-1, -3, 1}},
			},
		}, {
			name: "after 2 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 2,
			},
			want: []*Object{
				{[3]int{5, -3, -1}, [3]int{3, -2, -2}},
				{[3]int{1, -2, 2}, [3]int{-2, 5, 6}},
				{[3]int{1, -4, -1}, [3]int{0, 3, -6}},
				{[3]int{1, -4, 2}, [3]int{-1, -6, 2}},
			},
		},
		{
			name: "after 3 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 3,
			},
			want: []*Object{
				{[3]int{5, -6, -1}, [3]int{0, -3, 0}},
				{[3]int{0, 0, 6}, [3]int{-1, 2, 4}},
				{[3]int{2, 1, -5}, [3]int{1, 5, -4}},
				{[3]int{1, -8, 2}, [3]int{0, -4, 0}},
			},
		},
		{
			name: "after 4 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 4,
			},
			want: []*Object{
				{[3]int{2, -8, 0}, [3]int{-3, -2, 1}},
				{[3]int{2, 1, 7}, [3]int{2, 1, 1}},
				{[3]int{2, 3, -6}, [3]int{0, 2, -1}},
				{[3]int{2, -9, 1}, [3]int{1, -1, -1}},
			},
		},
		{
			name: "after 5 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 5,
			},
			want: []*Object{
				{[3]int{-1, -9, 2}, [3]int{-3, -1, 2}},
				{[3]int{4, 1, 5}, [3]int{2, 0, -2}},
				{[3]int{2, 2, -4}, [3]int{0, -1, 2}},
				{[3]int{3, -7, -1}, [3]int{1, 2, -2}},
			},
		},

		{
			name: "after 6 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 6,
			},
			want: []*Object{
				{[3]int{-1, -7, 3}, [3]int{0, 2, 1}},
				{[3]int{3, 0, 0}, [3]int{-1, -1, -5}},
				{[3]int{3, -2, 1}, [3]int{1, -4, 5}},
				{[3]int{3, -4, -2}, [3]int{0, 3, -1}},
			},
		},

		{
			name: "after 7 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 7,
			},
			want: []*Object{
				{[3]int{2, -2, 1}, [3]int{3, 5, -2}},
				{[3]int{1, -4, -4}, [3]int{-2, -4, -4}},
				{[3]int{3, -7, 5}, [3]int{0, -5, 4}},
				{[3]int{2, 0, 0}, [3]int{-1, 4, 2}},
			},
		},
		{
			name: "after 8 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 8,
			},
			want: []*Object{
				{[3]int{5, 2, -2}, [3]int{3, 4, -3}},
				{[3]int{2, -7, -5}, [3]int{1, -3, -1}},
				{[3]int{0, -9, 6}, [3]int{-3, -2, 1}},
				{[3]int{1, 1, 3}, [3]int{-1, 1, 3}},
			},
		},
		{
			name: "after 9 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 9,
			},
			want: []*Object{
				{[3]int{5, 3, -4}, [3]int{0, 1, -2}},
				{[3]int{2, -9, -3}, [3]int{0, -2, 2}},
				{[3]int{0, -8, 4}, [3]int{0, 1, -2}},
				{[3]int{1, 1, 5}, [3]int{0, 0, 2}},
			},
		},
		{
			name: "after 10 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 10,
			},
			want: []*Object{
				{[3]int{2, 1, -3}, [3]int{-3, -2, 1}},
				{[3]int{1, -8, 0}, [3]int{-1, 1, 3}},
				{[3]int{3, -6, 1}, [3]int{3, 2, -3}},
				{[3]int{2, 0, 4}, [3]int{1, -1, -1}},
			},
		},
		{
			name: "after 0 step",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 0,
			},
			want: []*Object{
				{[3]int{-8, -10, 0}, [3]int{0, 0, 0}},
				{[3]int{5, 5, 10}, [3]int{0, 0, 0}},
				{[3]int{2, -7, 3}, [3]int{0, 0, 0}},
				{[3]int{9, -8, -3}, [3]int{0, 0, 0}},
			},
		},
		{
			name: "after 10 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 10,
			},
			want: []*Object{
				{[3]int{-9, -10, 1}, [3]int{-2, -2, -1}},
				{[3]int{4, 10, 9}, [3]int{-3, 7, -2}},
				{[3]int{8, -10, -3}, [3]int{5, -1, -2}},
				{[3]int{5, -10, 3}, [3]int{0, -4, 5}},
			},
		},
		{
			name: "after 20 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 20,
			},
			want: []*Object{
				{[3]int{-10, 3, -4}, [3]int{-5, 2, 0}},
				{[3]int{5, -25, 6}, [3]int{1, 1, -4}},
				{[3]int{13, 1, 1}, [3]int{5, -2, 2}},
				{[3]int{0, 1, 7}, [3]int{-1, -1, 2}},
			},
		},
		{
			name: "after 30 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 30,
			},
			want: []*Object{
				{[3]int{15, -6, -9}, [3]int{-5, 4, 0}},
				{[3]int{-4, -11, 3}, [3]int{-3, -10, 0}},
				{[3]int{0, -1, 11}, [3]int{7, 4, 3}},
				{[3]int{-3, -2, 5}, [3]int{1, 2, -3}},
			},
		},
		{
			name: "after 40 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 40,
			},
			want: []*Object{
				{[3]int{14, -12, -4}, [3]int{11, 3, 0}},
				{[3]int{-1, 18, 8}, [3]int{-5, 2, 3}},
				{[3]int{-5, -14, 8}, [3]int{1, -2, 0}},
				{[3]int{0, -12, -2}, [3]int{-7, -3, -3}},
			},
		},
		{
			name: "after 50 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 50,
			},
			want: []*Object{
				{[3]int{-23, 4, 1}, [3]int{-7, -1, 2}},
				{[3]int{20, -31, 13}, [3]int{5, 3, 4}},
				{[3]int{-4, 6, 1}, [3]int{-1, 1, -3}},
				{[3]int{15, 1, -5}, [3]int{3, -3, -3}},
			},
		},
		{
			name: "after 60 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 60,
			},
			want: []*Object{
				{[3]int{36, -10, 6}, [3]int{5, 0, 3}},
				{[3]int{-18, 10, 9}, [3]int{-3, -7, 5}},
				{[3]int{8, -12, -3}, [3]int{-2, 1, -7}},
				{[3]int{-18, -8, -2}, [3]int{0, 6, -1}},
			},
		},
		{
			name: "after 70 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 70,
			},
			want: []*Object{
				{[3]int{-33, -6, 5}, [3]int{-5, -4, 7}},
				{[3]int{13, -9, 2}, [3]int{-2, 11, 3}},
				{[3]int{11, -8, 2}, [3]int{8, -6, -7}},
				{[3]int{17, 3, 1}, [3]int{-1, -1, -3}},
			},
		},
		{
			name: "after 80 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 80,
			},
			want: []*Object{
				{[3]int{30, -8, 3}, [3]int{3, 3, 0}},
				{[3]int{-2, -4, 0}, [3]int{4, -13, 2}},
				{[3]int{-18, -7, 15}, [3]int{-8, 2, -2}},
				{[3]int{-2, -1, -8}, [3]int{1, 8, 0}},
			},
		},
		{
			name: "after 90 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 90,
			},
			want: []*Object{
				{[3]int{-25, -1, 4}, [3]int{1, -3, 4}},
				{[3]int{2, -9, 0}, [3]int{-3, 13, -1}},
				{[3]int{32, -8, 14}, [3]int{5, -4, 6}},
				{[3]int{-1, -2, -8}, [3]int{-3, -6, -9}},
			},
		},
		{
			name: "after 100 steps",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
				nth: 100,
			},
			want: []*Object{
				{[3]int{8, -12, -9}, [3]int{-7, 3, 0}},
				{[3]int{13, 16, -3}, [3]int{3, -11, -5}},
				{[3]int{-29, -11, -1}, [3]int{-3, 7, 4}},
				{[3]int{16, -13, 23}, [3]int{7, 1, 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Simulate(tt.args.objects, tt.args.nth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Simulate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObject_Potential(t *testing.T) {
	type fields struct {
		Position [3]int
		Velocity [3]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "1",
			fields: fields{
				Position: [3]int{2, 1, -3},
				Velocity: [3]int{-3, -2, 1},
			},
			want: 6,
		},
		{
			name: "2",
			fields: fields{
				Position: [3]int{1, -8, 0},
				Velocity: [3]int{-1, 1, 3},
			},
			want: 9,
		},
		{
			name: "3",
			fields: fields{
				Position: [3]int{3, -6, 1},
				Velocity: [3]int{3, 2, -3},
			},
			want: 10,
		},
		{
			name: "4",
			fields: fields{
				Position: [3]int{2, 0, 4},
				Velocity: [3]int{1, -1, -1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Object{
				Position: tt.fields.Position,
				Velocity: tt.fields.Velocity,
			}
			if got := o.Potential(); got != tt.want {
				t.Errorf("Potential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObject_Kinetic(t *testing.T) {
	type fields struct {
		Position [3]int
		Velocity [3]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "1",
			fields: fields{
				Position: [3]int{2, 1, -3},
				Velocity: [3]int{-3, -2, 1},
			},
			want: 6,
		},
		{
			name: "2",
			fields: fields{
				Position: [3]int{1, -8, 0},
				Velocity: [3]int{-1, 1, 3},
			},
			want: 5,
		},
		{
			name: "3",
			fields: fields{
				Position: [3]int{3, -6, 1},
				Velocity: [3]int{3, 2, -3},
			},
			want: 8,
		},
		{
			name: "4",
			fields: fields{
				Position: [3]int{2, 0, 4},
				Velocity: [3]int{1, -1, -1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Object{
				Position: tt.fields.Position,
				Velocity: tt.fields.Velocity,
			}
			if got := o.Kinetic(); got != tt.want {
				t.Errorf("Potential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObject_Total(t *testing.T) {
	type fields struct {
		Position [3]int
		Velocity [3]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "1",
			fields: fields{
				Position: [3]int{2, 1, -3},
				Velocity: [3]int{-3, -2, 1},
			},
			want: 36,
		},
		{
			name: "2",
			fields: fields{
				Position: [3]int{1, -8, 0},
				Velocity: [3]int{-1, 1, 3},
			},
			want: 45,
		},
		{
			name: "3",
			fields: fields{
				Position: [3]int{3, -6, 1},
				Velocity: [3]int{3, 2, -3},
			},
			want: 80,
		},
		{
			name: "4",
			fields: fields{
				Position: [3]int{2, 0, 4},
				Velocity: [3]int{1, -1, -1},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Object{
				Position: tt.fields.Position,
				Velocity: tt.fields.Velocity,
			}
			if got := o.Total(); got != tt.want {
				t.Errorf("Potential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextIdentity(t *testing.T) {
	type args struct {
		objects []*Object
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "1",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-1, 0, 2},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -10, -7},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{4, -8, 8},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{3, 5, -1},
						Velocity: [3]int{0, 0, 0},
					},
				},
			},
			want: 2772,
		},
		{
			name: "2",
			args: args{
				objects: []*Object{
					{
						Position: [3]int{-8, -10, 0},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{5, 5, 10},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{2, -7, 3},
						Velocity: [3]int{0, 0, 0},
					},
					{
						Position: [3]int{9, -8, -3},
						Velocity: [3]int{0, 0, 0},
					},
				},
			},
			want: 4686774924,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextIdentity(tt.args.objects); got != tt.want {
				t.Errorf("NextIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
