package probles

import "testing"

func TestCode_1189(t *testing.T) {
	type args struct {
		text string
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
				text: "nlaebolko",
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				text: "loonbalxballpoon",
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				text: "leetcode",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1189(tt.args.text); got != tt.want {
				t.Errorf("Code_1189() = %v, want %v", got, tt.want)
			}
		})
	}
}
