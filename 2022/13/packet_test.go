package _13

import (
	"testing"
)

func TestPair_IsOrdered(t *testing.T) {
	type fields struct {
		Left  Packet
		Right Packet
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				Left:  MustParsePacket(`[1]`),
				Right: MustParsePacket(`[1]`),
			},
			want: true,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[2]`),
				Right: MustParsePacket(`[1]`),
			},
			want: false,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[1]`),
				Right: MustParsePacket(`[[1]]`),
			},
			want: true,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[[1]]`),
				Right: MustParsePacket(`[1]`),
			},
			want: true,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[2]`),
				Right: MustParsePacket(`[[1]]`),
			},
			want: false,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[[2]]`),
				Right: MustParsePacket(`[1]`),
			},
			want: false,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[[1]]`),
				Right: MustParsePacket(`[[1]]`),
			},
			want: true,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[[2]]`),
				Right: MustParsePacket(`[[1]]`),
			},
			want: false,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[1,1,3,1,1]`),
				Right: MustParsePacket(`[1,1,5,1,1]`),
			},
			want: true,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[[1],[2,3,4]]`),
				Right: MustParsePacket(`[[1],4]`),
			},
			want: true,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[9]`),
				Right: MustParsePacket(`[[8,7,6]]`),
			},
			want: false,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[[4,4],4,4]`),
				Right: MustParsePacket(`[[4,4],4,4,4]`),
			},
			want: true,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[7,7,7,7]`),
				Right: MustParsePacket(`[7,7,7]`),
			},
			want: false,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[]`),
				Right: MustParsePacket(`[3]`),
			},
			want: true,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[[[]]]`),
				Right: MustParsePacket(`[[]]`),
			},
			want: false,
		},
		{
			fields: fields{
				Left:  MustParsePacket(`[1,[2,[3,[4,[5,6,7]]]],8,9]`),
				Right: MustParsePacket(`[1,[2,[3,[4,[5,6,0]]]],8,9]`),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOrdered(tt.fields.Left, tt.fields.Right); got != tt.want {
				t.Errorf("IsOrdered() = %v, want %v", got, tt.want)
			}
		})
	}
}
