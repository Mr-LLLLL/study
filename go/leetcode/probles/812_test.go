package probles

import "testing"

func TestCode_812(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_812(tt.args.points); got != tt.want {
				t.Errorf("Code_812() = %v, want %v", got, tt.want)
			}
		})
	}
}
