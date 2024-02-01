
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package gtrace provides convenience wrapping functionality for tracing feature using OpenTelemetry.
<原文结束>

# <翻译开始>
// gtrace包提供了一种便利的封装功能，用于使用OpenTelemetry进行跟踪特性。
# <翻译结束>


<原文开始>
// To avoid too big tracing content.
<原文结束>

# <翻译开始>
// 为了避免跟踪内容过大。
# <翻译结束>


<原文开始>
// For detailed controlling for tracing content.
<原文结束>

# <翻译开始>
// 用于详细控制跟踪内容。
# <翻译结束>


<原文开始>
// tracingInternal enables tracing for internal type spans.
<原文结束>

# <翻译开始>
// tracingInternal 开启内部类型跨度的追踪功能。
# <翻译结束>


<原文开始>
// Max log size for request and response body, especially for HTTP/RPC request.
<原文结束>

# <翻译开始>
// 请求和响应体的最大日志大小，特别是对于HTTP/RPC请求。
# <翻译结束>


<原文开始>
// defaultTextMapPropagator is the default propagator for context propagation between peers.
<原文结束>

# <翻译开始>
// defaultTextMapPropagator 是默认的上下文传播器，用于在对等节点之间进行上下文传播。
# <翻译结束>


<原文开始>
// IsUsingDefaultProvider checks and return if currently using default trace provider.
<原文结束>

# <翻译开始>
// IsUsingDefaultProvider 检查并返回当前是否正在使用默认跟踪提供程序。
# <翻译结束>


<原文开始>
// IsTracingInternal returns whether tracing spans of internal components.
<原文结束>

# <翻译开始>
// IsTracingInternal 返回是否追踪内部组件的跟踪跨度。
# <翻译结束>


<原文开始>
// MaxContentLogSize returns the max log size for request and response body, especially for HTTP/RPC request.
<原文结束>

# <翻译开始>
// MaxContentLogSize 返回请求和响应正文的最大日志大小，特别是针对HTTP/RPC请求。
# <翻译结束>


<原文开始>
// CommonLabels returns common used attribute labels:
// ip.intranet, hostname.
<原文结束>

# <翻译开始>
// CommonLabels 返回常用属性标签：
// ip.intranet（内网IP），hostname（主机名）
# <翻译结束>


<原文开始>
// CheckSetDefaultTextMapPropagator sets the default TextMapPropagator if it is not set previously.
<原文结束>

# <翻译开始>
// CheckSetDefaultTextMapPropagator 检查并设置默认的 TextMapPropagator。如果此前未设置，则进行设置。
# <翻译结束>


<原文开始>
// GetDefaultTextMapPropagator returns the default propagator for context propagation between peers.
<原文结束>

# <翻译开始>
// GetDefaultTextMapPropagator 返回用于在对等节点间传播上下文的默认载体。
# <翻译结束>


<原文开始>
// GetTraceID retrieves and returns TraceId from context.
// It returns an empty string is tracing feature is not activated.
<原文结束>

# <翻译开始>
// GetTraceID 从上下文中检索并返回 TraceId。
// 如果追踪功能未激活，则返回一个空字符串。
# <翻译结束>


<原文开始>
// GetSpanID retrieves and returns SpanId from context.
// It returns an empty string is tracing feature is not activated.
<原文结束>

# <翻译开始>
// GetSpanID 从上下文中检索并返回 SpanId。
// 如果追踪功能未激活，则返回一个空字符串。
# <翻译结束>


<原文开始>
// SetBaggageValue is a convenient function for adding one key-value pair to baggage.
// Note that it uses attribute.Any to set the key-value pair.
<原文结束>

# <翻译开始>
// SetBaggageValue 是一个方便的函数，用于向 baggage 添加一对键值对。
// 注意，它使用 attribute.Any 来设置键值对。
# <翻译结束>


<原文开始>
// SetBaggageMap is a convenient function for adding map key-value pairs to baggage.
// Note that it uses attribute.Any to set the key-value pair.
<原文结束>

# <翻译开始>
// SetBaggageMap 是一个方便的函数，用于向 baggage 中添加映射键值对。
// 注意，它使用 attribute.Any 来设置键值对。
# <翻译结束>


<原文开始>
// GetBaggageMap retrieves and returns the baggage values as map.
<原文结束>

# <翻译开始>
// GetBaggageMap 获取并以map形式返回 baggage 的值。
# <翻译结束>


<原文开始>
// GetBaggageVar retrieves value and returns a *gvar.Var for specified key from baggage.
<原文结束>

# <翻译开始>
// GetBaggageVar 从 baggage 中按照指定键检索值，并返回一个 *gvar.Var 类型的结果。
# <翻译结束>


<原文开始>
// WithUUID injects custom trace id with UUID into context to propagate.
<原文结束>

# <翻译开始>
// WithUUID 向上下文中注入自定义的 UUID 格式的跟踪 ID，以便进行传播。
# <翻译结束>


<原文开始>
// WithTraceID injects custom trace id into context to propagate.
<原文结束>

# <翻译开始>
// WithTraceID 向上下文中注入自定义跟踪 ID，以便进行传播。
# <翻译结束>


<原文开始>
// Default trace provider.
<原文结束>

# <翻译开始>
// 默认追踪提供者。
# <翻译结束>

