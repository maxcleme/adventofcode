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
		name      string
		args      args
		want      *Coordinate
		wantCount int
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
			want: &Coordinate{
				X: 3,
				Y: 4,
			},
			wantCount: 8,
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
			want: &Coordinate{
				X: 5,
				Y: 8,
			},
			wantCount: 33,
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
			want: &Coordinate{
				X: 1,
				Y: 2,
			},
			wantCount: 35,
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
			want: &Coordinate{
				X: 6,
				Y: 3,
			},
			wantCount: 41,
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
			want: &Coordinate{
				X: 11,
				Y: 13,
			},
			wantCount: 210,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := Parse(tt.args.reader)
			if got, count := Best(*m); !reflect.DeepEqual(got, tt.want) || !reflect.DeepEqual(count, tt.wantCount) {
				t.Errorf("Best() = (%v,%d), want (%v,%d)", got, count, tt.want, tt.wantCount)
			}
		})
	}
}
