package options

import (
	"context"
	"strconv"

	"time"

	cfctx "git.dustess.com/mk-base/context"
	idempotenceID "git.dustess.com/mk-base/idempotence/id"
	"git.dustess.com/mk-base/util/snowflake"
)

const (
	// IDField id
	IDField = "eventbusID"
	// IdempotenceIDField IdempotenceID
	IdempotenceIDField = "eventbusIdempotenceID"
	// TraceIDField TraceIDField
	TraceIDField = "eventbusTraceIDField"
	// OrderField OrderField
	OrderField = "eventbusOrderField"
	// IDField id
	ValueIDField = "eventbusValueID"
	// PulsarDeliverRemainDurationField RemainDurationField
	PulsarDeliverRemainDurationField = "eventbusPulsarDeliverRemainDurationField"
)

func DefaulHeader() map[string]string {
	return make(map[string]string)
}

// HeaderWithID header新增id
func HeaderWithID(header map[string]string) map[string]string {
	id := snowflake.BaseNumber()
	header[IDField] = strconv.Itoa(int(id))
	header[IdempotenceIDField] = string(idempotenceID.NewGenerator().GenerateID(idempotenceID.Snowflake))
	return header
}

// HeaderMerge header 合并
func HeaderMerge(header map[string]string, customeHeader map[string]string) map[string]string {
	for k, v := range customeHeader {
		header[k] = v
	}
	return header
}

// HeaderWithContext header新增ctx相关信息
func HeaderWithContext(ctx context.Context, header map[string]string) map[string]string {
	ctxValue := CtxValue(ctx)
	common := ctxValue.GetCommonValue()
	traceID, ok := common[cfctx.CtxValueCommonKeyTraceID]
	if ok && len(traceID) > 0 {
		header[TraceIDField] = traceID
	}
	ctxHeader := ctxValue.GetHeaderValue()
	for k, v := range ctxHeader {
		header[k] = v
	}
	return header
}

// HeaderWithPulsarDeliverRemainDuration header新增 Pulsar延迟剩余时间
func HeaderWithPulsarDeliverRemainDuration(header map[string]string, t time.Duration) map[string]string {
	header[PulsarDeliverRemainDurationField] = t.String()
	return header
}

// HeaderPulsarDeliverRemainDuration header新增 Pulsar延迟剩余时间
func HeaderPulsarDeliverRemainDuration(header map[string]string) (time.Duration, bool) {
	var d time.Duration
	t, ok := header[PulsarDeliverRemainDurationField]
	if ok {
		d, _ = time.ParseDuration(t)
	}
	return d, ok
}

// HeaderOrderField 获取OrderField
func HeaderOrderField(header map[string]string) (string, bool) {
	id, ok := header[OrderField]
	return id, ok
}
