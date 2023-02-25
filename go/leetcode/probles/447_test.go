package probles

import "testing"

func TestCode_447(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				points: [][]int{{0, 0}, {1, 0}, {2, 0}},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				points: [][]int{{1, 1}},
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				points: [][]int{{0, 0}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_447(tt.args.points); got != tt.want {
				t.Errorf("Code_447() = %v, want %v", got, tt.want)
			}
		})
	}
}
