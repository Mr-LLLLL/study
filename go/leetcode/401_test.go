package leetcode

import (
	"reflect"
	"testing"
)

func TestCode_401(t *testing.T) {
	type args struct {
		turnedOn int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				turnedOn: 1,
			},
			want: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			name: "test2",
			args: args{
				turnedOn: 9,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_401(tt.args.turnedOn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Code_401() = %v, want %v", got, tt.want)
			}
		})
	}
}
