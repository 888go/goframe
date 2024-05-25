// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gtrace提供了使用OpenTelemetry进行跟踪功能的便利封装。. md5:4c5ceb4a418df579
package gtrace

import (
	"context"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/command"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtrace/internal/provider"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	tracingCommonKeyIpIntranet        = `ip.intranet`
	tracingCommonKeyIpHostname        = `hostname`
	commandEnvKeyForMaxContentLogSize = "gf.gtrace.max.content.log.size" // 为了避免跟踪内容过大。. md5:27b6d53e3e6ced7d
	commandEnvKeyForTracingInternal   = "gf.gtrace.tracing.internal"     // 用于详细控制跟踪内容。. md5:b871fd18f633cc43
)

var (
	intranetIps, _           = gipv4.GetIntranetIpArray()
	intranetIpStr            = strings.Join(intranetIps, ",")
	hostname, _              = os.Hostname()
	tracingInternal          = true       // tracingInternal 为内部类型跨度启用跟踪。. md5:f333cf108af1f2e4
	tracingMaxContentLogSize = 512 * 1024 // 请求和响应体的最大日志大小，特别是针对HTTP/RPC请求。. md5:329c69958d9e285c
	// defaultTextMapPropagator是用于在对等之间进行context传播的默认传播器。. md5:48e8537b612e7062
	defaultTextMapPropagator = propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
)

func init() {
	tracingInternal = gconv.Bool(command.GetOptWithEnv(commandEnvKeyForTracingInternal, "true"))
	if maxContentLogSize := gconv.Int(command.GetOptWithEnv(commandEnvKeyForMaxContentLogSize)); maxContentLogSize > 0 {
		tracingMaxContentLogSize = maxContentLogSize
	}
	// 默认的追踪提供者。. md5:61744e697ee81d00
	otel.SetTracerProvider(provider.New())
	CheckSetDefaultTextMapPropagator()
}

// IsUsingDefaultProvider 检查并返回当前是否正在使用默认的跟踪提供程序。. md5:dd9a8bbd104a14cf
func IsUsingDefaultProvider() bool {
	_, ok := otel.GetTracerProvider().(*provider.TracerProvider)
	return ok
}

// IsTracingInternal 返回是否正在追踪内部组件的跨度。. md5:4439b167674c69e6
func IsTracingInternal() bool {
	return tracingInternal
}

// MaxContentLogSize 返回请求和响应体的最大日志大小，特别是对于HTTP/RPC请求。. md5:762f425039c664ca
func MaxContentLogSize() int {
	return tracingMaxContentLogSize
}

// CommonLabels 返回常用属性标签：
// ip.intranet，hostname。
// md5:8affbee0c43e3bad
func CommonLabels() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String(tracingCommonKeyIpHostname, hostname),
		attribute.String(tracingCommonKeyIpIntranet, intranetIpStr),
		semconv.HostName(hostname),
	}
}

// CheckSetDefaultTextMapPropagator 如果之前未设置，默认情况下会设置文本映射传播器。. md5:586855119e290f63
func CheckSetDefaultTextMapPropagator() {
	p := otel.GetTextMapPropagator()
	if len(p.Fields()) == 0 {
		otel.SetTextMapPropagator(GetDefaultTextMapPropagator())
	}
}

// GetDefaultTextMapPropagator 返回用于在对等体之间传播上下文的默认 propagator。. md5:c053466fb206297d
func GetDefaultTextMapPropagator() propagation.TextMapPropagator {
	return defaultTextMapPropagator
}

// GetTraceID 从 context 中检索并返回 TraceId。如果跟踪功能未启用，则返回空字符串。
// md5:09e9e014a696e105
func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	traceID := trace.SpanContextFromContext(ctx).TraceID()
	if traceID.IsValid() {
		return traceID.String()
	}
	return ""
}

// GetSpanID 从上下文中检索并返回 SpanId。如果跟踪功能未激活，则返回空字符串。
// md5:1cca885adbc44f92
func GetSpanID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	spanID := trace.SpanContextFromContext(ctx).SpanID()
	if spanID.IsValid() {
		return spanID.String()
	}
	return ""
}

// SetBaggageValue 是一个便捷函数，用于向 baggage 中添加一个键值对。
// 注意，它使用 attribute.Any 来设置键值对。
// md5:a0a5e77a036e4b8b
func SetBaggageValue(ctx context.Context, key string, value interface{}) context.Context {
	return NewBaggage(ctx).SetValue(key, value)
}

// SetBaggageMap 是一个方便的函数，用于向行李中添加映射的键值对。
// 请注意，它使用 attribute.Any 来设置键值对。
// md5:635cc7b15635106d
func SetBaggageMap(ctx context.Context, data map[string]interface{}) context.Context {
	return NewBaggage(ctx).SetMap(data)
}

// GetBaggageMap 获取并返回行李（baggage）的值作为map。. md5:c2fd062493b49cd1
func GetBaggageMap(ctx context.Context) *gmap.StrAnyMap {
	return NewBaggage(ctx).GetMap()
}

// GetBaggageVar 从 baggage 中检索值，并为指定的键返回一个 *gvar.Var。. md5:b7635ba9a07703cf
func GetBaggageVar(ctx context.Context, key string) *gvar.Var {
	return NewBaggage(ctx).GetVar(key)
}

// WithUUID 向上下文注入自定义的基于UUID的追踪ID以进行传播。. md5:b75be6e561eacb0c
func WithUUID(ctx context.Context, uuid string) (context.Context, error) {
	return WithTraceID(ctx, gstr.Replace(uuid, "-", ""))
}

// WithTraceID 将自定义的跟踪ID注入上下文以进行传播。. md5:74657c53cd9aeefb
func WithTraceID(ctx context.Context, traceID string) (context.Context, error) {
	generatedTraceID, err := trace.TraceIDFromHex(traceID)
	if err != nil {
		return ctx, gerror.WrapCodef(
			gcode.CodeInvalidParameter,
			err,
			`invalid custom traceID "%s", a traceID string should be composed with [0-f] and fixed length 32`,
			traceID,
		)
	}
	sc := trace.SpanContextFromContext(ctx)
	if !sc.HasTraceID() {
		var span trace.Span
		ctx, span = NewSpan(ctx, "gtrace.WithTraceID")
		defer span.End()
		sc = trace.SpanContextFromContext(ctx)
	}
	ctx = trace.ContextWithRemoteSpanContext(ctx, trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    generatedTraceID,
		SpanID:     sc.SpanID(),
		TraceFlags: sc.TraceFlags(),
		TraceState: sc.TraceState(),
		Remote:     sc.IsRemote(),
	}))
	return ctx, nil
}
