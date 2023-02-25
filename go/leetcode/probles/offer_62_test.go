package probles

import "testing"

func TestCode_Offer_62(t *testing.T) {
	type args struct {
		n int
		m int
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
				n: 5,
				m: 3,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				n: 10,
				m: 17,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_Offer_62(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("Code_Offer_62() = %v, want %v", got, tt.want)
			}
		})
	}
}
