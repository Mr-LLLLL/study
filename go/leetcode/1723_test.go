package main

import "testing"

func TestCode_1723(t *testing.T) {
	type args struct {
		jobs []int
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
				jobs: []int{3, 2, 3},
				k:    3,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				jobs: []int{1, 2, 4, 7, 8},
				k:    2,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1723(tt.args.jobs, tt.args.k); got != tt.want {
				t.Errorf("Code_1723() = %v, want %v", got, tt.want)
			}
		})
	}
}
