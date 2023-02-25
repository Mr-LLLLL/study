package leetcode

import "testing"

func TestCode_interview_01_05(t *testing.T) {
	type args struct {
		first  string
		second string
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
				first:  "a",
				second: "ab",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				first:  "pales",
				second: "pal",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				first:  "helloworld",
				second: "willbeokay",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				first:  "wolrd",
				second: "worrd",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_interview_01_05(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("Code_interview_01_05() = %v, want %v", got, tt.want)
			}
		})
	}
}
