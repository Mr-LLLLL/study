package leetcode

import (
	"testing"
)

func Test_hasGroupsSizeX(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				deck: []int{1, 2, 3, 4, 4, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				deck: []int{1, 1, 1, 2, 2, 2, 3, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.args.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				i: 3,
				j: 6,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				i: 8,
				j: 3,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				i: 8,
				j: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
