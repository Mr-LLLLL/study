package mock

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(dbres{i: 1}, errors.New("not exist"))
	m.EXPECT().Get(gomock.Not("Tom")).Return(dbres{i: 2}, nil)
	m.EXPECT().Print().Return("hello")

	type args struct {
		db  DB
		key string
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
				db:  m,
				key: "Tom",
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				db:  m,
				key: "To",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFromDB(tt.args.db, tt.args.key); got != tt.want {
				t.Errorf("GetFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
