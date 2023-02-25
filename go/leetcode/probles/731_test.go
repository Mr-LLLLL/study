package probles

import (
	"testing"
)

func TestMyCalendarTwo_Book(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				start: 10,
				end:   20,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				start: 50,
				end:   60,
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				start: 10,
				end:   40,
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				start: 5,
				end:   15,
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				start: 5,
				end:   10,
			},
			want: true,
		},
		{
			name: "test6",
			args: args{
				start: 25,
				end:   55,
			},
			want: true,
		},
	}
	m := Constructor731()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MyCalendarTwo.Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
