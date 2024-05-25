// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gctx封装了context.Context并提供了额外的上下文功能。. md5:edcfb6983b687169
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
	Ctx    = context.Context // Ctx是context.Context的简短名称别名。. md5:2c9c93edc22890c4
	StrKey string            // StrKey 是一个类型，用于将基本类型字符串作为上下文键进行封装。. md5:84db5aa6fb6ea74e
)

var (
	// initCtx是从进程环境初始化的上下文。. md5:7e2eda888a5b2cc9
	initCtx context.Context
)

func init() {
	// 所有环境键值对。. md5:4c0179afb7589376
	m := make(map[string]string)
	i := 0
	for _, s := range os.Environ() {
		i = strings.IndexByte(s, '=')
		if i == -1 {
			continue
		}
		m[s[0:i]] = s[i+1:]
	}
	// 从环境获取OpenTelemetry。. md5:95f284182505db14
	initCtx = otel.GetTextMapPropagator().Extract(
		context.Background(),
		propagation.MapCarrier(m),
	)
	initCtx = WithCtx(initCtx)
}

// New 创建并返回一个包含上下文ID的上下文。. md5:ace97871c3d80d4f
func New() context.Context {
	return WithCtx(context.Background())
}

// WithCtx 根据给定的父上下文`ctx`创建并返回一个包含上下文ID的新上下文。. md5:bea2d0daa280a6eb
func WithCtx(ctx context.Context) context.Context {
	if CtxId(ctx) != "" {
		return ctx
	}
	var span *gtrace.Span
	ctx, span = gtrace.NewSpan(ctx, "gctx.WithCtx")
	defer span.End()
	return ctx
}

// CtxId从上下文中检索并返回context ID。. md5:bd18ae591706e243
func CtxId(ctx context.Context) string {
	return gtrace.GetTraceID(ctx)
}

// SetInitCtx 设置自定义初始化上下文。
// 注意，此函数不能在多个goroutine中调用。
// md5:10830063aafa5df4
func SetInitCtx(ctx context.Context) {
	initCtx = ctx
}

// GetInitCtx 返回初始化上下文。
// 初始化上下文用于在`main`函数或`init`函数中。
// md5:5608d282e442f76c
func GetInitCtx() context.Context {
	return initCtx
}
