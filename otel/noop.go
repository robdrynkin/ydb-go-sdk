package otel

import (
	"context"

	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

var (
	_ Config = noopConfig{}
	_ Span   = noopSpan{}
)

type (
	noopConfig struct{}
	noopSpan   struct{}
)

func (noopConfig) Details() trace.Details {
	return 0
}

func (noopConfig) SpanFromContext(context.Context) Span {
	return noopSpan{}
}

func (noopConfig) Start(ctx context.Context, _ string, _ ...KeyValue) (context.Context, Span) {
	return ctx, noopSpan{}
}

func (noopSpan) Link(Span, ...KeyValue) {}

func (noopSpan) TraceID() string {
	return ""
}

func (n noopSpan) Log(string, ...KeyValue)  {}
func (n noopSpan) Warn(error, ...KeyValue)  {}
func (n noopSpan) Error(error, ...KeyValue) {}

func (n noopSpan) End(...KeyValue) {}
