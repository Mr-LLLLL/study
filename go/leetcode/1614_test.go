package leetcode

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "(1+(2*3)+((8)/4))+1",
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				s: "(1)+((2))+(((3)))",
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				s: "(((())((",
			},
			want: 2,
		},
		{
			name: "test4",
			args: args{
				s: "))))((((((",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.s); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
