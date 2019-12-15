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
				reader: bufio.NewReader(strings.NewReader("123456789012")),
				height: 2,
				width:  3,
			},
			want: &Image{
				Height: 2,
				Width:  3,
				Layers: []Layer{
					{
						Colors: map[int]int{
							1 : 1,
							2 : 1,
							3 : 1,
							4 : 1,
							5 : 1,
							6 : 1,
						},
					},
					{
						Colors: map[int]int{
							7 : 1,
							8 : 1,
							9 : 1,
							0 : 1,
							1 : 1,
							2 : 1,
						},
					},
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
