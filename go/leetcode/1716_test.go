package main

import "testing"

func TestCode_1716(t *testing.T) {
	type args struct {
		n int
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
				n: 4,
			},
			want: 10,
		},
		{
			name: "test2",
			args: args{
				n: 10,
			},
			want: 37,
		},
		{
			name: "test3",
			args: args{
				n: 20,
			},
			want: 96,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1716(tt.args.n); got != tt.want {
				t.Errorf("Code_1716() = %v, want %v", got, tt.want)
			}
		})
	}
}
