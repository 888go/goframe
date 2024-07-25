
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
// RequestFromCtx retrieves and returns the Request object from context.
<原文结束>

# <翻译开始>
// RequestFromCtx 从上下文中检索并返回Request对象。 md5:c247eac3d031fb2b
# <翻译结束>


<原文开始>
// Context is alias for function GetCtx.
// This function overwrites the http.Request.Context function.
// See GetCtx.
<原文结束>

# <翻译开始>
// Context 是函数 GetCtx 的别名。
// 这个函数重写了 http.Request.Context 函数。
// 参考 GetCtx 函数。 md5:cc1d0847fad835ae
# <翻译结束>


<原文开始>
// Check and inject Request object into context.
<原文结束>

# <翻译开始>
	// 检查并把Request对象注入到context中。 md5:6db244059baa2f77
# <翻译结束>


<原文开始>
// Inject Request object into context.
<原文结束>

# <翻译开始>
		// 将 Request 对象注入到上下文中。 md5:561b491dd0877f18
# <翻译结束>


<原文开始>
// Add default tracing info if using default tracing provider.
<原文结束>

# <翻译开始>
		// 如果使用默认跟踪提供程序，添加默认跟踪信息。 md5:ffe39a076b5cfe89
# <翻译结束>


<原文开始>
// Update the values of the original HTTP request.
<原文结束>

# <翻译开始>
		// 更新原始HTTP请求的值。 md5:b7f6f8f267437800
# <翻译结束>


<原文开始>
// GetCtx retrieves and returns the request's context.
// Its alias of function Context,to be relevant with function SetCtx.
<原文结束>

# <翻译开始>
// GetCtx 获取并返回请求的上下文。
// 它是函数Context的别名，与SetCtx函数相关联。 md5:67b84a74829ae7e1
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
// 允许在手动完成上下文之前，使得该上下文可以传播到异步goroutine中，
// 不受HTTP请求结束的影响。
//
// 这个更改是考虑到开发人员在单个HTTP请求中创建多个goroutine时，常见的上下文传播习惯。 md5:14245b7febbf75b4
# <翻译结束>


<原文开始>
// SetCtx custom context for current request.
<原文结束>

# <翻译开始>
// SetCtx 为当前请求设置自定义上下文。 md5:447efd465325d822
# <翻译结束>


<原文开始>
// GetCtxVar retrieves and returns a Var with a given key name.
// The optional parameter `def` specifies the default value of the Var if given `key`
// does not exist in the context.
<原文结束>

# <翻译开始>
// GetCtxVar 根据给定的键名获取并返回一个 Var。
// 可选参数 `def` 指定了如果给定的 `key` 在上下文中不存在时，Var 的默认值。 md5:8e874c6ac730ae7b
# <翻译结束>


<原文开始>
// SetCtxVar sets custom parameter to context with key-value pairs.
<原文结束>

# <翻译开始>
// SetCtxVar 为上下文（context）设置自定义参数，使用键值对形式。 md5:f6782a132f936ef1
# <翻译结束>

