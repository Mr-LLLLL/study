package probles

import "testing"

func Test_giveGem(t *testing.T) {
	type args struct {
		gem        []int
		operations [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				gem:        []int{3, 1, 2},
				operations: [][]int{{0, 2}, {2, 1}, {2, 0}},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				gem:        []int{100, 0, 50, 100},
				operations: [][]int{{0, 2}, {0, 1}, {3, 0}, {3, 0}},
			},
			want: 75,
		},
		{
			name: "test3",
			args: args{
				gem:        []int{0, 0, 0, 0},
				operations: [][]int{{1, 2}, {3, 1}, {1, 2}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := giveGem(tt.args.gem, tt.args.operations); got != tt.want {
				t.Errorf("giveGem() = %v, want %v", got, tt.want)
			}
		})
	}
}
