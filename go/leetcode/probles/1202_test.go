package probles

import "testing"

func TestCode_1202(t *testing.T) {
	type args struct {
		s     string
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s:     "dcab",
				pairs: [][]int{{0, 3}, {1, 2}},
			},
			want: "bacd",
		},
		{
			name: "test2",
			args: args{
				s:     "dcab",
				pairs: [][]int{{0, 3}, {1, 2}, {0, 2}},
			},
			want: "abcd",
		},
		{
			name: "test3",
			args: args{
				s:     "cba",
				pairs: [][]int{{0, 1}, {1, 2}},
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1202(tt.args.s, tt.args.pairs); got != tt.want {
				t.Errorf("Code_1202() = %v, want %v", got, tt.want)
			}
		})
	}
}
