package main

import "testing"

func TestCode_1573(t *testing.T) {
	type args struct {
		s string
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
				s: "10101",
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				s: "000000",
			},
			want: 10,
		},
		{
			name: "test3",
			args: args{
				s: "100100010100110",
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1573(tt.args.s); got != tt.want {
				t.Errorf("Code_1573() = %v, want %v", got, tt.want)
			}
		})
	}
}
