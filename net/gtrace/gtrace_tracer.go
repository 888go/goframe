// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtrace
import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	)
// Tracer 为兼容性和扩展性而封装了 trace.Tracer。
type Tracer struct {
	trace.Tracer
}

// NewTracer 是一个用于获取 Tracer 的便捷函数。
func NewTracer(name ...string) *Tracer {
	tracerName := ""
	if len(name) > 0 {
		tracerName = name[0]
	}
	return &Tracer{
		Tracer: otel.Tracer(tracerName),
	}
}
