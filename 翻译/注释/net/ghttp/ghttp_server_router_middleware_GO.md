
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
// The default route pattern for global middleware.
<原文结束>

# <翻译开始>
	// 全局中间件的默认路由模式。 md5:104618f487cee8ac
# <翻译结束>


<原文开始>
// BindMiddleware registers one or more global middleware to the server.
// Global middleware can be used standalone without service handler, which intercepts all dynamic requests
// before or after service handler. The parameter `pattern` specifies what route pattern the middleware intercepts,
// which is usually a "fuzzy" pattern like "/:name", "/*any" or "/{field}".
<原文结束>

# <翻译开始>
// BindMiddleware 会在服务器上注册一个或多个全局中间件。全局中间件可以在没有服务处理器的情况下单独使用，它会在服务处理器之前或之后拦截所有动态请求。参数 `pattern` 指定了中间件拦截的路由模式，通常是一个模糊模式，如 "/:name"、"/*any" 或 "/{field}"。 md5:a58488c3f3613ab4
# <翻译结束>


<原文开始>
// BindMiddlewareDefault registers one or more global middleware to the server using default pattern "/*".
// Global middleware can be used standalone without service handler, which intercepts all dynamic requests
// before or after service handler.
<原文结束>

# <翻译开始>
// BindMiddlewareDefault 使用默认模式 "/*" 向服务器注册一个或多个全局中间件。
// 全局中间件可以独立使用，无需服务处理器，它能在所有动态请求的前后拦截处理。 md5:fc212697fcedf39e
# <翻译结束>


<原文开始>
// Use is the alias of BindMiddlewareDefault.
// See BindMiddlewareDefault.
<原文结束>

# <翻译开始>
// Use 是 BindMiddlewareDefault 的别名。
// 参见 BindMiddlewareDefault。 md5:713ca39a398552e8
# <翻译结束>

