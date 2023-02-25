package leetcode

import "testing"

func TestCode_1869(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s: "1101",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "111000",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				s: "110100010",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				s: "11010111111010000",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1869(tt.args.s); got != tt.want {
				t.Errorf("Code_1869() = %v, want %v", got, tt.want)
			}
		})
	}
}
