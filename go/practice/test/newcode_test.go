package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_test(t *testing.T) {
	fmt.Println(os.Stat("./main.go"))
}

func Test_emptyBottle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: 3,
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				n: 10,
			},
			want: 5,
		},
		{
			name: "test3",
			args: args{
				n: 81,
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := emptyBottle(tt.args.n); got != tt.want {
				t.Errorf("emptyBottle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dupAndOsrt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{2, 2, 1},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dupAndOsrt(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dupAndOsrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trans(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "0xAA",
			},
			want: 170,
		},
		{
			name: "test2",
			args: args{
				s: "0xAAA",
			},
			want: 2730,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trans(tt.args.s); got != tt.want {
				t.Errorf("trans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				k: 3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
		{
			name: "test2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				k: 2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			gotArr := []int{}
			for got != nil {
				gotArr = append(gotArr, got.Val)
				got = got.Next
			}
			if !reflect.DeepEqual(gotArr, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", gotArr, tt.want)
			}
		})
	}
}
