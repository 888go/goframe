
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Package gctx wraps context.Context and provides extra context features.
<原文结束>

# <翻译开始>
// 包gctx封装了context.Context并提供了额外的上下文功能。 md5:edcfb6983b687169
# <翻译结束>


<原文开始>
// Ctx is short name alias for context.Context.
<原文结束>

# <翻译开始>
// Ctx是context.Context的简短名称别名。 md5:2c9c93edc22890c4
# <翻译结束>


<原文开始>
// StrKey is a type for warps basic type string as context key.
<原文结束>

# <翻译开始>
// StrKey 是一个类型，用于将基本类型字符串作为上下文键进行封装。 md5:84db5aa6fb6ea74e
# <翻译结束>


<原文开始>
// initCtx is the context initialized from process environment.
<原文结束>

# <翻译开始>
	// initCtx是从进程环境初始化的上下文。 md5:7e2eda888a5b2cc9
# <翻译结束>


<原文开始>
// All environment key-value pairs.
<原文结束>

# <翻译开始>
	// 所有环境键值对。 md5:4c0179afb7589376
# <翻译结束>


<原文开始>
// OpenTelemetry from environments.
<原文结束>

# <翻译开始>
	// 从环境获取OpenTelemetry。 md5:95f284182505db14
# <翻译结束>


<原文开始>
// New creates and returns a context which contains context id.
<原文结束>

# <翻译开始>
// New 创建并返回一个包含上下文ID的上下文。 md5:ace97871c3d80d4f
# <翻译结束>


<原文开始>
// WithCtx creates and returns a context containing context id upon given parent context `ctx`.
<原文结束>

# <翻译开始>
// WithCtx 根据给定的父上下文`ctx`创建并返回一个包含上下文ID的新上下文。 md5:bea2d0daa280a6eb
# <翻译结束>


<原文开始>
// CtxId retrieves and returns the context id from context.
<原文结束>

# <翻译开始>
// CtxId从上下文中检索并返回context ID。 md5:bd18ae591706e243
# <翻译结束>


<原文开始>
// SetInitCtx sets custom initialization context.
// Note that this function cannot be called in multiple goroutines.
<原文结束>

# <翻译开始>
// SetInitCtx 设置自定义初始化上下文。
// 注意，此函数不能在多个goroutine中调用。 md5:10830063aafa5df4
# <翻译结束>


<原文开始>
// GetInitCtx returns the initialization context.
// Initialization context is used in `main` or `init` functions.
<原文结束>

# <翻译开始>
// GetInitCtx 返回初始化上下文。
// 初始化上下文用于在`main`函数或`init`函数中。 md5:5608d282e442f76c
# <翻译结束>

