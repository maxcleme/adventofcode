package utils

import "testing"

func TestTotalOrbits(t *testing.T) {
	type args struct {
		input []string
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
				input: []string{
					"COM)B",
					"B)C",
					"C)D",
					"D)E",
					"E)F",
					"B)G",
					"G)H",
					"D)I",
					"E)J",
					"J)K",
					"K)L",
				},
			},
			want:    42,
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				input: []string{
					"B)C",
					"J)K",
					"D)I",
					"E)F",
					"C)D",
					"COM)B",
					"D)E",
					"B)G",
					"G)H",
					"E)J",
					"K)L",
				},
			},
			want:    42,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TotalOrbits(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalOrbits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TotalOrbits() got = %v, want %v", got, tt.want)
			}
		})
	}
}
