package leetcode

import "testing"

func Test_toHex(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				num: 2,
			},
			want: "2",
		},
		{
			name: "test2",
			args: args{
				num: 26,
			},
			want: "1a",
		},
		{
			name: "test3",
			args: args{
				num: -1,
			},
			want: "ffffffff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
