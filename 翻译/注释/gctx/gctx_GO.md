
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
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package gctx wraps context.Context and provides extra context features.
<原文结束>

# <翻译开始>
// 包 gctx 对 context.Context 进行了封装，并提供了额外的上下文功能。
# <翻译结束>


<原文开始>
// Ctx is short name alias for context.Context.
<原文结束>

# <翻译开始>
// Ctx 是 context.Context 的简写别名。
# <翻译结束>


<原文开始>
// StrKey is a type for warps basic type string as context key.
<原文结束>

# <翻译开始>
// StrKey 是一种类型，用于将基本类型 string 包装为上下文键（Context key）。
# <翻译结束>


<原文开始>
// initCtx is the context initialized from process environment.
<原文结束>

# <翻译开始>
// initCtx 是从进程环境初始化的上下文。
# <翻译结束>


<原文开始>
// All environment key-value pairs.
<原文结束>

# <翻译开始>
// 所有环境键值对。
# <翻译结束>


<原文开始>
// OpenTelemetry from environments.
<原文结束>

# <翻译开始>
// 从环境变量中获取OpenTelemetry配置
# <翻译结束>


<原文开始>
// New creates and returns a context which contains context id.
<原文结束>

# <翻译开始>
// New 创建并返回一个包含上下文ID的上下文。
# <翻译结束>


<原文开始>
// WithCtx creates and returns a context containing context id upon given parent context `ctx`.
<原文结束>

# <翻译开始>
// WithCtx 在给定的父级上下文 `ctx` 的基础上创建并返回一个包含上下文 ID 的新上下文。
# <翻译结束>


<原文开始>
// CtxId retrieves and returns the context id from context.
<原文结束>

# <翻译开始>
// CtxId 从 context 中检索并返回上下文 id。
# <翻译结束>


<原文开始>
// SetInitCtx sets custom initialization context.
// Note that this function cannot be called in multiple goroutines.
<原文结束>

# <翻译开始>
// SetInitCtx 设置自定义初始化上下文。
// 注意：该函数不能在多个goroutine中被调用。
# <翻译结束>


<原文开始>
// GetInitCtx returns the initialization context.
// Initialization context is used in `main` or `init` functions.
<原文结束>

# <翻译开始>
// GetInitCtx 返回初始化上下文。
// 初始化上下文用于在 `main` 或 `init` 函数中使用。
# <翻译结束>

