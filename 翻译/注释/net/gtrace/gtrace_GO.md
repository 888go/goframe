
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Package gtrace provides convenience wrapping functionality for tracing feature using OpenTelemetry.
<原文结束>

# <翻译开始>
// 包gtrace提供了使用OpenTelemetry进行跟踪功能的便利封装。 md5:4c5ceb4a418df579
# <翻译结束>


<原文开始>
// To avoid too big tracing content.
<原文结束>

# <翻译开始>
// 为了避免跟踪内容过大。 md5:27b6d53e3e6ced7d
# <翻译结束>


<原文开始>
// For detailed controlling for tracing content.
<原文结束>

# <翻译开始>
// 用于详细控制跟踪内容。 md5:b871fd18f633cc43
# <翻译结束>


<原文开始>
// tracingInternal enables tracing for internal type spans.
<原文结束>

# <翻译开始>
// tracingInternal 为内部类型跨度启用跟踪。 md5:f333cf108af1f2e4
# <翻译结束>


<原文开始>
// Max log size for request and response body, especially for HTTP/RPC request.
<原文结束>

# <翻译开始>
// 请求和响应体的最大日志大小，特别是针对HTTP/RPC请求。 md5:329c69958d9e285c
# <翻译结束>


<原文开始>
// defaultTextMapPropagator is the default propagator for context propagation between peers.
<原文结束>

# <翻译开始>
// defaultTextMapPropagator是用于在对等之间进行context传播的默认传播器。 md5:48e8537b612e7062
# <翻译结束>


<原文开始>
// Default trace provider.
<原文结束>

# <翻译开始>
// 默认的追踪提供者。 md5:61744e697ee81d00
# <翻译结束>


<原文开始>
// IsUsingDefaultProvider checks and return if currently using default trace provider.
<原文结束>

# <翻译开始>
// IsUsingDefaultProvider 检查并返回当前是否正在使用默认的跟踪提供程序。 md5:dd9a8bbd104a14cf
# <翻译结束>


<原文开始>
// IsTracingInternal returns whether tracing spans of internal components.
<原文结束>

# <翻译开始>
// IsTracingInternal 返回是否正在追踪内部组件的跨度。 md5:4439b167674c69e6
# <翻译结束>


<原文开始>
// MaxContentLogSize returns the max log size for request and response body, especially for HTTP/RPC request.
<原文结束>

# <翻译开始>
// MaxContentLogSize 返回请求和响应体的最大日志大小，特别是对于HTTP/RPC请求。 md5:762f425039c664ca
# <翻译结束>


<原文开始>
// CommonLabels returns common used attribute labels:
// ip.intranet, hostname.
<原文结束>

# <翻译开始>
// CommonLabels 返回常用属性标签：
// ip.intranet，hostname。
// md5:8affbee0c43e3bad
# <翻译结束>


<原文开始>
// CheckSetDefaultTextMapPropagator sets the default TextMapPropagator if it is not set previously.
<原文结束>

# <翻译开始>
// CheckSetDefaultTextMapPropagator 如果之前未设置，默认情况下会设置文本映射传播器。 md5:586855119e290f63
# <翻译结束>


<原文开始>
// GetDefaultTextMapPropagator returns the default propagator for context propagation between peers.
<原文结束>

# <翻译开始>
// GetDefaultTextMapPropagator 返回用于在对等体之间传播上下文的默认 propagator。 md5:c053466fb206297d
# <翻译结束>


<原文开始>
// GetTraceID retrieves and returns TraceId from context.
// It returns an empty string is tracing feature is not activated.
<原文结束>

# <翻译开始>
// GetTraceID 从 context 中检索并返回 TraceId。如果跟踪功能未启用，则返回空字符串。
// md5:09e9e014a696e105
# <翻译结束>


<原文开始>
// GetSpanID retrieves and returns SpanId from context.
// It returns an empty string is tracing feature is not activated.
<原文结束>

# <翻译开始>
// GetSpanID 从上下文中检索并返回 SpanId。如果跟踪功能未激活，则返回空字符串。
// md5:1cca885adbc44f92
# <翻译结束>


<原文开始>
// SetBaggageValue is a convenient function for adding one key-value pair to baggage.
// Note that it uses attribute.Any to set the key-value pair.
<原文结束>

# <翻译开始>
// SetBaggageValue 是一个便捷函数，用于向 baggage 中添加一个键值对。
// 注意，它使用 attribute.Any 来设置键值对。
// md5:a0a5e77a036e4b8b
# <翻译结束>


<原文开始>
// SetBaggageMap is a convenient function for adding map key-value pairs to baggage.
// Note that it uses attribute.Any to set the key-value pair.
<原文结束>

# <翻译开始>
// SetBaggageMap 是一个方便的函数，用于向行李中添加映射的键值对。
// 请注意，它使用 attribute.Any 来设置键值对。
// md5:635cc7b15635106d
# <翻译结束>


<原文开始>
// GetBaggageMap retrieves and returns the baggage values as map.
<原文结束>

# <翻译开始>
// GetBaggageMap 获取并返回行李（baggage）的值作为map。 md5:c2fd062493b49cd1
# <翻译结束>


<原文开始>
// GetBaggageVar retrieves value and returns a *gvar.Var for specified key from baggage.
<原文结束>

# <翻译开始>
// GetBaggageVar 从 baggage 中检索值，并为指定的键返回一个 *gvar.Var。 md5:b7635ba9a07703cf
# <翻译结束>


<原文开始>
// WithUUID injects custom trace id with UUID into context to propagate.
<原文结束>

# <翻译开始>
// WithUUID 向上下文注入自定义的基于UUID的追踪ID以进行传播。 md5:b75be6e561eacb0c
# <翻译结束>


<原文开始>
// WithTraceID injects custom trace id into context to propagate.
<原文结束>

# <翻译开始>
// WithTraceID 将自定义的跟踪ID注入上下文以进行传播。 md5:74657c53cd9aeefb
# <翻译结束>

