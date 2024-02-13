// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package tracing 提供了一些用于追踪功能的实用工具函数。
package tracing

import (
	"math"
	"time"
	
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/encoding/gbinary"
	"github.com/888go/goframe/util/grand"
)

var (
	randomInitSequence = int32(随机类.X整数(math.MaxInt32))
	sequence           = 安全变量类.NewInt32(randomInitSequence)
)

// NewIDs 创建并返回一个新的跟踪ID和跨度ID。
func NewIDs() (traceID trace.TraceID, spanID trace.SpanID) {
	return NewTraceID(), NewSpanID()
}

// NewTraceID 创建并返回一个追踪ID。
func NewTraceID() (traceID trace.TraceID) {
	var (
		timestampNanoBytes = 字节集类.EncodeInt64(time.Now().UnixNano())
		sequenceBytes      = 字节集类.EncodeInt32(sequence.Add(1))
		randomBytes        = 随机类.X字节集(4)
	)
	copy(traceID[:], timestampNanoBytes)
	copy(traceID[8:], sequenceBytes)
	copy(traceID[12:], randomBytes)
	return
}

// NewSpanID 创建并返回一个跨度ID。
func NewSpanID() (spanID trace.SpanID) {
	copy(spanID[:], 字节集类.EncodeInt64(time.Now().UnixNano()/1e3))
	copy(spanID[4:], 随机类.X字节集(4))
	return
}
