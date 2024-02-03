// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// gtrace包提供了一种便利的封装功能，用于使用OpenTelemetry进行跟踪特性。
package gtrace

import (
	"context"
	"os"
	"strings"
	
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/net/gipv4"
	"github.com/888go/goframe/net/gtrace/internal/provider"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

const (
	tracingCommonKeyIpIntranet        = `ip.intranet`
	tracingCommonKeyIpHostname        = `hostname`
	commandEnvKeyForMaxContentLogSize = "gf.gtrace.max.content.log.size" // 为了避免跟踪内容过大。
	commandEnvKeyForTracingInternal   = "gf.gtrace.tracing.internal"     // 用于详细控制跟踪内容。
)

var (
	intranetIps, _           = gipv4.GetIntranetIpArray()
	intranetIpStr            = strings.Join(intranetIps, ",")
	hostname, _              = os.Hostname()
	tracingInternal          = true       // tracingInternal 开启内部类型跨度的追踪功能。
	tracingMaxContentLogSize = 512 * 1024 // 请求和响应体的最大日志大小，特别是对于HTTP/RPC请求。
	// defaultTextMapPropagator 是默认的上下文传播器，用于在对等节点之间进行上下文传播。
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
	// 默认追踪提供者。
	otel.SetTracerProvider(provider.New())
	CheckSetDefaultTextMapPropagator()
}

// IsUsingDefaultProvider 检查并返回当前是否正在使用默认跟踪提供程序。
func IsUsingDefaultProvider() bool {
	_, ok := otel.GetTracerProvider().(*provider.TracerProvider)
	return ok
}

// IsTracingInternal 返回是否追踪内部组件的跟踪跨度。
func IsTracingInternal() bool {
	return tracingInternal
}

// MaxContentLogSize 返回请求和响应正文的最大日志大小，特别是针对HTTP/RPC请求。
func MaxContentLogSize() int {
	return tracingMaxContentLogSize
}

// CommonLabels 返回常用属性标签：
// ip.intranet（内网IP），hostname（主机名）
func CommonLabels() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String(tracingCommonKeyIpHostname, hostname),
		attribute.String(tracingCommonKeyIpIntranet, intranetIpStr),
		semconv.HostNameKey.String(hostname),
	}
}

// CheckSetDefaultTextMapPropagator 检查并设置默认的 TextMapPropagator。如果此前未设置，则进行设置。
func CheckSetDefaultTextMapPropagator() {
	p := otel.GetTextMapPropagator()
	if len(p.Fields()) == 0 {
		otel.SetTextMapPropagator(GetDefaultTextMapPropagator())
	}
}

// GetDefaultTextMapPropagator 返回用于在对等节点间传播上下文的默认载体。
func GetDefaultTextMapPropagator() propagation.TextMapPropagator {
	return defaultTextMapPropagator
}

// GetTraceID 从上下文中检索并返回 TraceId。
// 如果追踪功能未激活，则返回一个空字符串。
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

// GetSpanID 从上下文中检索并返回 SpanId。
// 如果追踪功能未激活，则返回一个空字符串。
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

// SetBaggageValue 是一个方便的函数，用于向 baggage 添加一对键值对。
// 注意，它使用 attribute.Any 来设置键值对。
func SetBaggageValue(ctx context.Context, key string, value interface{}) context.Context {
	return NewBaggage(ctx).SetValue(key, value)
}

// SetBaggageMap 是一个方便的函数，用于向 baggage 中添加映射键值对。
// 注意，它使用 attribute.Any 来设置键值对。
func SetBaggageMap(ctx context.Context, data map[string]interface{}) context.Context {
	return NewBaggage(ctx).SetMap(data)
}

// GetBaggageMap 获取并以map形式返回 baggage 的值。
func GetBaggageMap(ctx context.Context) *gmap.StrAnyMap {
	return NewBaggage(ctx).GetMap()
}

// GetBaggageVar 从 baggage 中按照指定键检索值，并返回一个 *gvar.Var 类型的结果。
func GetBaggageVar(ctx context.Context, key string) *gvar.Var {
	return NewBaggage(ctx).GetVar(key)
}

// WithUUID 向上下文中注入自定义的 UUID 格式的跟踪 ID，以便进行传播。
func WithUUID(ctx context.Context, uuid string) (context.Context, error) {
	return WithTraceID(ctx, gstr.Replace(uuid, "-", ""))
}

// WithTraceID 向上下文中注入自定义跟踪 ID，以便进行传播。
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
