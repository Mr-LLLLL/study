package main

import "testing"

func TestCode_Lcp_33(t *testing.T) {
	type args struct {
		bucket []int
		vat    []int
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
				bucket: []int{1, 3},
				vat:    []int{6, 8},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				bucket: []int{9, 0, 1},
				vat:    []int{0, 2, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_Lcp_33(tt.args.bucket, tt.args.vat); got != tt.want {
				t.Errorf("Code_Lcp_33() = %v, want %v", got, tt.want)
			}
		})
	}
}
