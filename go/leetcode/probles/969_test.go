package probles

import (
	"reflect"
	"testing"
)

func TestCode_969(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: []int{3, 2, 4, 1},
			},
			want: []int{4, 2, 4, 3},
		},
		{
			name: "test2",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_969(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Code_969() = %v, want %v", got, tt.want)
			}
		})
	}
}
