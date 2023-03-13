package leetcode

import (
	"testing"
)

func TestT_GetSum(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test1",
		},
	}
	simpleSum := func(s []int) int {
		sum := 0
		for _, v := range s {
			sum += v
		}
		return sum
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmp := NewT()
			got := tmp.GetSum()
			want := simpleSum(tmp.GetSli())
			if got != want {
				t.Errorf("T.GetSum() = %v, want %v", got, want)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test1",
			want: "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Print(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
