package main

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestGetPrice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().GetNameById(gomock.Eq(1)).Return("mac")
	m.EXPECT().GetGoodsPriceById(gomock.Eq(1)).Return(10000)

	m.EXPECT().GetNameById(gomock.Eq(2)).Return("notebook")
	m.EXPECT().GetGoodsPriceById(gomock.Eq(2)).Return(5000)

	type args struct {
		db DB
		id int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				db: m,
				id: 1,
			},
			want: 20000,
		},
		{
			name: "case2",
			args: args{
				db: m,
				id: 2,
			},
			want: 5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPrice(tt.args.db, tt.args.id); got != tt.want {
				t.Errorf("GetPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
