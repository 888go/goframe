// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包 gctx 对 context.Context 进行了封装，并提供了额外的上下文功能。
package gctx

import (
	"context"
	"os"
	"strings"
	
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	
	"github.com/gogf/gf/v2/net/gtrace"
)

type (
	Ctx    = context.Context // Ctx 是 context.Context 的简写别名。
	StrKey string            // StrKey 是一种类型，用于将基本类型 string 包装为上下文键（Context key）。
)

var (
	// initCtx 是从进程环境初始化的上下文。
	initCtx context.Context
)

func init() {
	// 所有环境键值对。
	m := make(map[string]string)
	i := 0
	for _, s := range os.Environ() {
		i = strings.IndexByte(s, '=')
		if i == -1 {
			continue
		}
		m[s[0:i]] = s[i+1:]
	}
	// 从环境变量中获取OpenTelemetry配置
	initCtx = otel.GetTextMapPropagator().Extract(
		context.Background(),
		propagation.MapCarrier(m),
	)
	initCtx = WithCtx(initCtx)
}

// New 创建并返回一个包含上下文ID的上下文。
func New() context.Context {
	return WithCtx(context.Background())
}

// WithCtx 在给定的父级上下文 `ctx` 的基础上创建并返回一个包含上下文 ID 的新上下文。
func WithCtx(ctx context.Context) context.Context {
	if CtxId(ctx) != "" {
		return ctx
	}
	var span *gtrace.Span
	ctx, span = gtrace.NewSpan(ctx, "gctx.WithCtx")
	defer span.End()
	return ctx
}

// CtxId 从 context 中检索并返回上下文 id。
func CtxId(ctx context.Context) string {
	return gtrace.GetTraceID(ctx)
}

// SetInitCtx 设置自定义初始化上下文。
// 注意：该函数不能在多个goroutine中被调用。
func SetInitCtx(ctx context.Context) {
	initCtx = ctx
}

// GetInitCtx 返回初始化上下文。
// 初始化上下文用于在 `main` 或 `init` 函数中使用。
func GetInitCtx() context.Context {
	return initCtx
}
