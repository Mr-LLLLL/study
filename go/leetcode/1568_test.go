package main

import (
	"fmt"
	"testing"
)

func TestCode_1568(t *testing.T) {
	type args struct {
		grid [][]int
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
				grid: [][]int{
					{0, 1, 1, 0},
					{0, 1, 1, 0},
					{0, 0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				grid: [][]int{
					{1, 1},
				},
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				grid: [][]int{
					{1, 0, 1, 0},
				},
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				grid: [][]int{
					{1, 1, 0, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 0, 1, 1},
					{1, 1, 0, 1, 1},
				},
			},
			want: 1,
		},
		{
			name: "test5",
			args: args{
				grid: [][]int{
					{1, 1, 0, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 0, 1, 1},
					{1, 1, 1, 1, 1},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1568(tt.args.grid); got != tt.want {
				t.Errorf("Code_1568() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_dfs(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				x: 0,
				y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid([][]int{{1, 1}, {1, 1}})
			fmt.Println(g.grid)
			g.dfs(tt.args.x, tt.args.y)
			fmt.Println(g.grid)
		})
	}
}
