package utils

import (
	"reflect"
	"testing"
)

func TestClosestManhattanDistance(t *testing.T) {
	type args struct {
		stringWire1 []string
		stringWire2 []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				stringWire1: []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
				stringWire2: []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			},
			want:    159,
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				stringWire1: []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
				stringWire2: []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			},
			want:    135,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClosestManhattanDistance(tt.args.stringWire1, tt.args.stringWire2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClosestManhattanDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClosestManhattanDistance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
