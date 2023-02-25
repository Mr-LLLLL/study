package probles

import "testing"

func TestCode_1572(t *testing.T) {
	type args struct {
		mat [][]int
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
				mat: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 25,
		},
		{
			name: "test2",
			args: args{
				mat: [][]int{
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
				},
			},
			want: 8,
		},
		{
			name: "test3",
			args: args{
				mat: [][]int{{5}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1572(tt.args.mat); got != tt.want {
				t.Errorf("Code_1572() = %v, want %v", got, tt.want)
			}
		})
	}
}
