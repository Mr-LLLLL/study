package pkg

import (
	"context"

	"github.com/go-kratos/kratos/pkg/log"
)

func LogPath() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "hello", "work")
	log.Infoc(ctx, "skfjskf")
}

type ITest interface {
	print()
	hello()
}

type T1 struct{}
