package leetcode

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				nums: []int{2},
				k:    1,
			},
			want: 2,
		},
		{
			name: "test4",
			args: args{
				nums: []int{2, 1},
				k:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partitino(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := make([]int, 1)
		for i := range a {
			a[i] = rand.Intn(500)
		}
		pilot := rand.Intn(len(a))
		simplePartition := func(a []int, l, r, pilot int) int {
			index := 0
			for i := l; i <= r; i++ {
				if a[i] > a[pilot] {
					index++
				}
			}

			return index
		}

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			want := simplePartition(a, 0, len(a)-1, pilot)
			got := partition(a, 0, len(a)-1, pilot)
			if got != want {
				t.Errorf("partition() = %v, want %v", got, want)
			}
		})
	}
}

func Test_heap_heapify(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				a: []int{3, 4, 5, 1, 2, 7, 6},
			},
		},
		{
			name: "test2",
			args: args{
				a: []int{9, 2, 4, 1, 5, 2, 0, 4},
			},
		},
	}
	isHeap := func(a []int) bool {
		for i := range a {
			l := 2*i + 1
			r := 2 * (i + 1)

			if l < len(a) && a[l] < a[i] {
				return false
			}
			if r < len(a) && a[r] < a[i] {
				return false
			}
		}

		return true
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := new(_heap).heapify(tt.args.a)
			if !isHeap(h.arr) {
				t.Errorf("heap_heapify")
			}
		})
	}
}

func Test_heap_push(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test1",
			fields: fields{
				arr: []int{3, 4, 5, 1, 2, 7, 6},
			},
			args: args{
				n: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := new(_heap).heapify(tt.fields.arr)
			h.push(tt.args.n)
			fmt.Println(h.arr)
		})
	}
}
