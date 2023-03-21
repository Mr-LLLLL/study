package leetcode

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "test3",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "test4",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
		{
			name: "test5",
			args: args{
				m: 51,
				n: 9,
			},
			want: 1916797311,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
