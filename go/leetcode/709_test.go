package main

import "testing"

func TestCode_709(t *testing.T) {
	type args struct {
		s string
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
				s: "Hello",
			},
			want: "hello",
		},
		{
			name: "test2",
			args: args{
				s: "here",
			},
			want: "here",
		},
		{
			name: "test3",
			args: args{
				s: "LOVELY",
			},
			want: "lovely",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_709(tt.args.s); got != tt.want {
				t.Errorf("Code_709() = %v, want %v", got, tt.want)
			}
		})
	}
}
