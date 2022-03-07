package main

import "testing"

func TestCode_128(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				nums: []int{0, 3, 4, 5, 7, 8, 1, 2, 0, 6, 8},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_128(tt.args.nums); got != tt.want {
				t.Errorf("Code_128() = %v, want %v", got, tt.want)
			}
		})
	}
}
