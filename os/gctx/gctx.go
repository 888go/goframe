// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gctx wraps context.Context and provides extra context features.
package gctx//bm:上下文类

import (
	"context"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/gogf/gf/v2/net/gtrace"
)

type (
	Ctx    = context.Context // Ctx is short name alias for context.Context.
	StrKey string            // StrKey is a type for warps basic type string as context key.
)

var (
	// initCtx is the context initialized from process environment.
	initCtx context.Context
)

func init() {
	// All environment key-value pairs.
	m := make(map[string]string)
	i := 0
	for _, s := range os.Environ() {
		i = strings.IndexByte(s, '=')
		if i == -1 {
			continue
		}
		m[s[0:i]] = s[i+1:]
	}
	// OpenTelemetry from environments.
	initCtx = otel.GetTextMapPropagator().Extract(
		context.Background(),
		propagation.MapCarrier(m),
	)
	initCtx = WithCtx(initCtx)
}

// New creates and returns a context which contains context id.
// ff:创建
func New() context.Context {
	return WithCtx(context.Background())
}

// WithCtx creates and returns a context containing context id upon given parent context `ctx`.
// ff:创建并从上下文
// ctx:上下文
func WithCtx(ctx context.Context) context.Context {
	if CtxId(ctx) != "" {
		return ctx
	}
	var span *gtrace.Span
	ctx, span = gtrace.NewSpan(ctx, "gctx.WithCtx")
	defer span.End()
	return ctx
}

// CtxId retrieves and returns the context id from context.
// ff:取上下文id
// ctx:上下文
func CtxId(ctx context.Context) string {
	return gtrace.GetTraceID(ctx)
}

// SetInitCtx sets custom initialization context.
// Note that this function cannot be called in multiple goroutines.
// ff:设置初始化上下文
// ctx:上下文
func SetInitCtx(ctx context.Context) {
	initCtx = ctx
}

// GetInitCtx returns the initialization context.
// Initialization context is used in `main` or `init` functions.
// ff:取初始化上下文
func GetInitCtx() context.Context {
	return initCtx
}
