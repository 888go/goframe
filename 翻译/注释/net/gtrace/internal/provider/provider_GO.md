
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
// New returns a new and configured TracerProvider, which has no SpanProcessor.
//
// In default the returned TracerProvider is configured with:
// - a ParentBased(AlwaysSample) Sampler;
// - a unix nano timestamp and random umber based IDGenerator;
// - the resource.Default() Resource;
// - the default SpanLimits.
//
// The passed opts are used to override these default values and configure the
// returned TracerProvider appropriately.
<原文结束>

# <翻译开始>
// New 返回一个新的并配置好的 TracerProvider，它没有 SpanProcessor。
// 
// 默认情况下，返回的 TracerProvider 配置如下：
// - 一个基于父span（AlwaysSample）的采样器；
// - 一个基于Unix纳秒时间戳和随机数的ID生成器；
// - 资源.Default()资源；
// - 默认的Span限制。
// 
// 传递给 opts 的参数将用于覆盖这些默认值，并适当地配置返回的 TracerProvider。
// md5:92a14af244d0cf0e
# <翻译结束>

