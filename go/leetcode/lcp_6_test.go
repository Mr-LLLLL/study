package main

import "testing"

func TestCode_Lcp_6(t *testing.T) {
	type args struct {
		coins []int
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
				coins: []int{4, 2, 1},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				coins: []int{2, 3, 10},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_Lcp_6(tt.args.coins); got != tt.want {
				t.Errorf("Code_Lcp_6() = %v, want %v", got, tt.want)
			}
		})
	}
}
