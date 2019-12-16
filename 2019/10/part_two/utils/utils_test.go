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
		want    *Map
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				reader: strings.NewReader(`.#..#
.....
#####
....#
...##`),
			},
			want: &Map{Asteroids: []*Coordinate{
				{1, 0},
				{4, 0},
				{0, 2},
				{1, 2},
				{2, 2},
				{3, 2},
				{4, 2},
				{4, 3},
				{3, 4},
				{4, 4},
			}},
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

func TestBest(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want *Station
	}{
		{
			name: "1",
			args: args{
				reader: strings.NewReader(`.#..#
.....
#####
....#
...##`),
			},
			want: &Station{
				Position: &Coordinate{
					X: 3,
					Y: 4,
				},
				Size: 8,
			},
		},
		{
			name: "2",
			args: args{
				reader: strings.NewReader(`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`),
			},
			want: &Station{
				Position: &Coordinate{
					X: 5,
					Y: 8,
				},
				Size: 33,
			},
		},
		{
			name: "3",
			args: args{
				reader: strings.NewReader(`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`),
			},
			want: &Station{
				Position: &Coordinate{
					X: 1,
					Y: 2,
				},
				Size: 35,
			},
		},
		{
			name: "4",
			args: args{
				reader: strings.NewReader(`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`),
			},
			want: &Station{
				Position: &Coordinate{
					X: 6,
					Y: 3,
				},
				Size: 41,
			},
		},
		{
			name: "5",
			args: args{
				reader: strings.NewReader(`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`),
			},
			want: &Station{
				Position: &Coordinate{
					X: 11,
					Y: 13,
				},
				Size: 210,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := Parse(tt.args.reader)
			if got := Best(*m); !reflect.DeepEqual(got.Position, tt.want.Position) || !reflect.DeepEqual(got.Size, tt.want.Size) {
				t.Errorf("Best() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestNthVaporized(t *testing.T) {
	type args struct {
		reader io.Reader
		nth    int
	}
	text := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	tests := []struct {
		name string
		args args
		want *Coordinate
	}{
		{
			name: "1st to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    1,
			},
			want: &Coordinate{
				X: 11,
				Y: 12,
			},
		},
		{
			name: "2nd to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    2,
			},
			want: &Coordinate{
				X: 12,
				Y: 1,
			},
		},
		{
			name: "3rd to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    3,
			},
			want: &Coordinate{
				X: 12,
				Y: 2,
			},
		},
		{
			name: "10th to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    10,
			},
			want: &Coordinate{
				X: 12,
				Y: 8,
			},
		},
		{
			name: "20th to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    20,
			},
			want: &Coordinate{
				X: 16,
				Y: 0,
			},
		},
		{
			name: "50th to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    50,
			},
			want: &Coordinate{
				X: 16,
				Y: 9,
			},
		},
		{
			name: "100th to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    100,
			},
			want: &Coordinate{
				X: 10,
				Y: 16,
			},
		},
		{
			name: "199th to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    199,
			},
			want: &Coordinate{
				X: 9,
				Y: 6,
			},
		},
		{
			name: "200th to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    200,
			},
			want: &Coordinate{
				X: 8,
				Y: 2,
			},
		},
		{
			name: "201th to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    201,
			},
			want: &Coordinate{
				X: 10,
				Y: 9,
			},
		},
		{
			name: "299th to be vaporized",
			args: args{
				reader: strings.NewReader(text),
				nth:    299,
			},
			want: &Coordinate{
				X: 11,
				Y: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := Parse(tt.args.reader)
			s := Best(*m)
			if got := NthVaporized(s, tt.args.nth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NthVaporized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle(t *testing.T) {
	type args struct {
		c1 Coordinate
		c2 Coordinate
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 degrees",
			args: args{
				c1: Coordinate{5, 5},
				c2: Coordinate{5, 3},
			},
			want: 0,
		},
		{
			name: "90 degrees",
			args: args{
				c1: Coordinate{5, 5},
				c2: Coordinate{6, 5},
			},
			want: 90,
		},

		{
			name: "180 degrees",
			args: args{
				c1: Coordinate{5, 5},
				c2: Coordinate{5, 6},
			},
			want: 180,
		},

		{
			name: "270 degrees",
			args: args{
				c1: Coordinate{5, 5},
				c2: Coordinate{4, 5},
			},
			want: 270,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Angle(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("Angle() = %v, want %v", got, tt.want)
			}
		})
	}
}
