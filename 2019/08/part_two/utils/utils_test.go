package utils

import (
	"bufio"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		reader io.Reader
		height int
		width  int
	}
	tests := []struct {
		name    string
		args    args
		want    *Image
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				reader: bufio.NewReader(strings.NewReader("0222112222120000")),
				height: 2,
				width:  2,
			},
			want: &Image{
				Height: 2,
				Width:  2,
				Pixels: [][]*int{
					{p(0), p(1)},
					{p(1), p(0)},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.reader, tt.args.height, tt.args.width)
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

func p(i int) *int {
	return &i
}
