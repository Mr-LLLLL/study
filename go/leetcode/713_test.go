package main

import "testing"

func TestCode_713(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{10, 5, 2, 6},
				k:    100,
			},
			want: 8,
		},
		{
			name: "test2",
			args: args{
				nums: []int{3, 4, 10},
				k:    80,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_713(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("Code_713() = %v, want %v", got, tt.want)
			}
		})
	}
}
