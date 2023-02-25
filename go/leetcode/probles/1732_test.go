package probles

import "testing"

func TestCode_1732(t *testing.T) {
	type args struct {
		gain []int
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
				gain: []int{-5, 1, 5, 0, -7},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				gain: []int{-4, -3, -2, -1, 4, 3, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1732(tt.args.gain); got != tt.want {
				t.Errorf("Code_1732() = %v, want %v", got, tt.want)
			}
		})
	}
}
