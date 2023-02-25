package probles

import "testing"

func TestCode_647(t *testing.T) {
	type args struct {
		s string
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
				s: "abc",
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
		{
			name: "test3",
			args: args{
				s: "abbac",
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_647(tt.args.s); got != tt.want {
				t.Errorf("Code_647() = %v, want %v", got, tt.want)
			}
		})
	}
}
