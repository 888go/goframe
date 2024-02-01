
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
// RequestFromCtx retrieves and returns the Request object from context.
<原文结束>

# <翻译开始>
// RequestFromCtx 从 context 中检索并返回 Request 对象。
# <翻译结束>


<原文开始>
// Context is alias for function GetCtx.
// This function overwrites the http.Request.Context function.
// See GetCtx.
<原文结束>

# <翻译开始>
// Context 是 GetCtx 函数的别名。
// 此函数覆盖了 http.Request.Context 函数的功能。
// 请参阅 GetCtx。
# <翻译结束>


<原文开始>
// Check and inject Request object into context.
<原文结束>

# <翻译开始>
// 检查并把Request对象注入到上下文中。
# <翻译结束>


<原文开始>
// Inject Request object into context.
<原文结束>

# <翻译开始>
// 将Request对象注入到context中。
# <翻译结束>


<原文开始>
// Add default tracing info if using default tracing provider.
<原文结束>

# <翻译开始>
// 如果使用默认追踪提供者，则添加默认追踪信息。
# <翻译结束>


<原文开始>
// Update the values of the original HTTP request.
<原文结束>

# <翻译开始>
// 更新原始HTTP请求的值。
# <翻译结束>


<原文开始>
// GetCtx retrieves and returns the request's context.
// Its alias of function Context,to be relevant with function SetCtx.
<原文结束>

# <翻译开始>
// GetCtx 从请求中检索并返回其上下文。
// 它是函数 Context 的别名，为了与函数 SetCtx 保持关联性。
# <翻译结束>


<原文开始>
// GetNeverDoneCtx creates and returns a never done context object,
// which forbids the context manually done, to make the context can be propagated to asynchronous goroutines,
// which will not be affected by the HTTP request ends.
//
// This change is considered for common usage habits of developers for context propagation
// in multiple goroutines creation in one HTTP request.
<原文结束>

# <翻译开始>
// GetNeverDoneCtx 创建并返回一个永不完成的上下文对象，
// 该对象禁止手动设置完成状态，以保证上下文可以传递到异步goroutine中，
// 因此不会受到HTTP请求结束的影响。
//
// 这个改动是为了适应开发者在处理单个HTTP请求时创建多个goroutine进行上下文传播的常见使用习惯。
# <翻译结束>


<原文开始>
// SetCtx custom context for current request.
<原文结束>

# <翻译开始>
// SetCtx 为当前请求设置自定义上下文。
# <翻译结束>


<原文开始>
// GetCtxVar retrieves and returns a Var with a given key name.
// The optional parameter `def` specifies the default value of the Var if given `key`
// does not exist in the context.
<原文结束>

# <翻译开始>
// GetCtxVar 通过给定的键名检索并返回一个 Var。
// 可选参数 `def` 指定了如果给定 `key` 在上下文中不存在时，Var 的默认值。
# <翻译结束>


<原文开始>
// SetCtxVar sets custom parameter to context with key-value pairs.
<原文结束>

# <翻译开始>
// SetCtxVar 通过键值对设置自定义参数到上下文中。
# <翻译结束>

