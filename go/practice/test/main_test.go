package main

import (
	"testing"
)

func Benchmark_A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 10000)
		for i := range arr {
			v := arr[i]
			_ = v
		}

		nums := [10000]int{1, 2, 3}
		for i := range nums {
			v := nums[i]
			_ = v
		}
	}
}

func Benchmark_B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 10000)
		for i := range arr[:] {
			v := arr[i]
			_ = v
		}

		nums := [10000]int{1, 2, 3}
		for i := range nums[:] {
			v := nums[i]
			_ = v
		}
	}
}

func TestSlice_Append(t *testing.T) {
	type fields struct {
		Len   int
		Paris [][2]Pair
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Slice{
				Len:   tt.fields.Len,
				Paris: tt.fields.Paris,
			}
			s.Append(tt.args.key, tt.args.value)
		})
	}
}
