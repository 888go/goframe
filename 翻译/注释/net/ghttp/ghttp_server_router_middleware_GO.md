
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
// The default route pattern for global middleware.
<原文结束>

# <翻译开始>
// 全局中间件的默认路由模式。
# <翻译结束>


<原文开始>
// BindMiddleware registers one or more global middleware to the server.
// Global middleware can be used standalone without service handler, which intercepts all dynamic requests
// before or after service handler. The parameter `pattern` specifies what route pattern the middleware intercepts,
// which is usually a "fuzzy" pattern like "/:name", "/*any" or "/{field}".
<原文结束>

# <翻译开始>
// BindMiddleware 注册一个或多个全局中间件到服务器。
// 全局中间件可以在没有服务处理器的情况下独立使用，它会在服务处理器执行前或执行后拦截所有的动态请求。
// 参数 `pattern` 指定了中间件要拦截的路由模式，通常是一个“模糊”模式，如 "/:name"、"/*any" 或 "/{field}"。
# <翻译结束>


<原文开始>
// BindMiddlewareDefault registers one or more global middleware to the server using default pattern "/*".
// Global middleware can be used standalone without service handler, which intercepts all dynamic requests
// before or after service handler.
<原文结束>

# <翻译开始>
// BindMiddlewareDefault 使用默认模式"/*"将一个或多个全局中间件注册到服务器。
// 全局中间件可以在没有服务处理器的情况下独立使用，它会在服务处理器处理所有动态请求之前或之后进行拦截。
# <翻译结束>


<原文开始>
// Use is the alias of BindMiddlewareDefault.
// See BindMiddlewareDefault.
<原文结束>

# <翻译开始>
// Use 是 BindMiddlewareDefault 的别名。
// 请参阅 BindMiddlewareDefault。
# <翻译结束>

