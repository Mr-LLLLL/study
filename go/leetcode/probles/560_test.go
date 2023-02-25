package probles

import "testing"

func TestCode_560(t *testing.T) {
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_560(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("Code_560() = %v, want %v", got, tt.want)
			}
		})
	}
}
