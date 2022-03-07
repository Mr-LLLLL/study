package options

import (
	"context"
	"time"

	cfctx "git.dustess.com/mk-base/context"
)

// DetachCtx 只保留值
func DetachCtx(ctx context.Context) context.Context {
	return EventContext{ctx}
}

// EventContext EventContext
type EventContext struct{ ctx context.Context }

func (v EventContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (v EventContext) Done() <-chan struct{}             { return nil }
func (v EventContext) Err() error                        { return nil }
func (v EventContext) Value(key interface{}) interface{} { return v.ctx.Value(key) }

func CtxValue(ctx context.Context) *cfctx.CtxValue {
	return cfctx.GetCtxValue(ctx)
}

