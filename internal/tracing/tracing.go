// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package tracing 提供了一些用于追踪功能的实用函数。 md5:e35e84e99377e9ae
package tracing

import (
	"math"
	"time"

	"go.opentelemetry.io/otel/trace"

	gtype "github.com/888go/goframe/container/gtype"
	gbinary "github.com/888go/goframe/encoding/gbinary"
	grand "github.com/888go/goframe/util/grand"
)

var (
	randomInitSequence = int32(grand.X整数(math.MaxInt32))
	sequence           = gtype.NewInt32(randomInitSequence)
)

// NewIDs 生成并返回一个新的 trace 和 span ID。 md5:1b5d0dc93c6f38ff
func NewIDs() (traceID trace.TraceID, spanID trace.SpanID) {
	return NewTraceID(), NewSpanID()
}

// NewTraceID 创建并返回一个追踪ID。 md5:6832c150aaff566d
func NewTraceID() (traceID trace.TraceID) {
	var (
		timestampNanoBytes = gbinary.EncodeInt64(time.Now().UnixNano())
		sequenceBytes      = gbinary.EncodeInt32(sequence.Add(1))
		randomBytes        = grand.X字节集(4)
	)
	copy(traceID[:], timestampNanoBytes)
	copy(traceID[8:], sequenceBytes)
	copy(traceID[12:], randomBytes)
	return
}

// NewSpanID 创建并返回一个新的span ID。 md5:02d4095273219037
func NewSpanID() (spanID trace.SpanID) {
	copy(spanID[:], gbinary.EncodeInt64(time.Now().UnixNano()/1e3))
	copy(spanID[4:], grand.X字节集(4))
	return
}
