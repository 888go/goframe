
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
// Request is the context object for a request.
<原文结束>

# <翻译开始>
// Request 是请求的上下文对象。. md5:90c1e82eacf87b05
# <翻译结束>


<原文开始>
// Corresponding Response of this request.
<原文结束>

# <翻译开始>
// 对应于此请求的响应。. md5:aca9c7c1b8de7fd3
# <翻译结束>


<原文开始>
// Request starting time in milliseconds.
<原文结束>

# <翻译开始>
// 请求开始时间（以毫秒为单位）。. md5:c8ed1608de735520
# <翻译结束>


<原文开始>
// Request to end time in milliseconds.
<原文结束>

# <翻译开始>
// 结束时间（以毫秒为单位）的请求。. md5:d10eec1013ccf8ce
# <翻译结束>


<原文开始>
// Static file object for static file serving.
<原文结束>

# <翻译开始>
// 用于静态文件服务的静态文件对象。. md5:c9ba568602380c10
# <翻译结束>


<原文开始>
	// =================================================================================================================
	// Private attributes for internal usage purpose.
	// =================================================================================================================
<原文结束>

# <翻译开始>
// =================================================================================================================
// 用于内部使用的私有属性。
// =================================================================================================================
// md5:5edc202a134e48b1
# <翻译结束>


<原文开始>
// All matched handlers containing handler, hook and middleware for this request.
<原文结束>

# <翻译开始>
// 包含此请求的处理器、挂钩和中间件的所有匹配处理程序。. md5:42382b62d60509ae
# <翻译结束>


<原文开始>
// Real handler serving for this request, not hook or middleware.
<原文结束>

# <翻译开始>
// 为这次请求的实际处理程序，不是钩子或中间件。. md5:60c0eea59027a44f
# <翻译结束>


<原文开始>
// Handler response object for Request/Response handler.
<原文结束>

# <翻译开始>
// 用于请求/响应处理程序的处理器响应对象。. md5:adefb98d0b335dd6
# <翻译结束>


<原文开始>
// A bool marking whether the GET parameters parsed.
<原文结束>

# <翻译开始>
// 一个布尔值，表示是否已经解析了GET参数。. md5:08f143c93bb48593
# <翻译结束>


<原文开始>
// A bool marking whether the request body parsed.
<原文结束>

# <翻译开始>
// 一个标记请求体是否已解析的布尔值。. md5:f14ff93f821a596e
# <翻译结束>


<原文开始>
// A bool marking whether request Form parsed for HTTP method PUT, POST, PATCH.
<原文结束>

# <翻译开始>
// 一个布尔值，表示HTTP方法为PUT、POST或PATCH时，请求表单是否已解析。. md5:b372893681fa6b1a
# <翻译结束>


<原文开始>
// Router parameters map, which might be nil if there are no router parameters.
<原文结束>

# <翻译开始>
// 路由参数映射，如果没有路由参数，则可能为nil。. md5:1f669e8d65cc7201
# <翻译结束>


<原文开始>
// Body parameters map, which might be nil if their nobody content.
<原文结束>

# <翻译开始>
// Body参数映射，如果内容为空则可能为nil。. md5:318eb5ce8cbb633b
# <翻译结束>


<原文开始>
// Current executing error of the request.
<原文结束>

# <翻译开始>
// 请求当前执行的错误。. md5:fb1258df50e5180f
# <翻译结束>


<原文开始>
// A bool marking whether current request is exited.
<原文结束>

# <翻译开始>
// 一个标记当前请求是否已退出的布尔值。. md5:df96d943707fa105
# <翻译结束>


<原文开始>
// The parsed host name for current host used by GetHost function.
<原文结束>

# <翻译开始>
// 由GetHost函数使用的当前主机的解析后的主机名。. md5:f8aa264e6a28b762
# <翻译结束>


<原文开始>
// The parsed client ip for current host used by GetClientIp function.
<原文结束>

# <翻译开始>
// 当前主机中由GetClientIp函数使用的解析后的客户端IP。. md5:de8953e109404d72
# <翻译结束>


<原文开始>
// A bool marking whether current request is file serving.
<原文结束>

# <翻译开始>
// 一个bool标记当前请求是否是文件服务。. md5:71668278a657d9c6
# <翻译结束>


<原文开始>
// Custom template view engine object for this response.
<原文结束>

# <翻译开始>
// 为这次响应自定义的模板视图引擎对象。. md5:ddd788a8e05477da
# <翻译结束>


<原文开始>
// Custom template view variables for this response.
<原文结束>

# <翻译开始>
// 为此响应定制的模板视图变量。. md5:4138dc353e5f6967
# <翻译结束>


<原文开始>
// Original URL path that passed from client.
<原文结束>

# <翻译开始>
// 从客户端传递的原始URL路径。. md5:c7368bb1d20fddcd
# <翻译结束>


<原文开始>
// staticFile is the file struct for static file service.
<原文结束>

# <翻译开始>
// staticFile是静态文件服务的文件结构体。. md5:1a45356e421cf5d2
# <翻译结束>


<原文开始>
// newRequest creates and returns a new request object.
<原文结束>

# <翻译开始>
// newRequest 创建并返回一个新的请求对象。. md5:bbe6326af48b82f2
# <翻译结束>


<原文开始>
// Custom session id creating function.
<原文结束>

# <翻译开始>
// 自定义会话ID生成函数。. md5:1530052018d41784
# <翻译结束>


<原文开始>
// Remove char '/' in the tail of URI.
<原文结束>

# <翻译开始>
// 删除URI尾部的'/'字符。. md5:5cd878ba748e3629
# <翻译结束>


<原文开始>
// Default URI value if it's empty.
<原文结束>

# <翻译开始>
// 如果为空，默认的URI值。. md5:ba9367be3b98edbd
# <翻译结束>


<原文开始>
// WebSocket upgrades current request as a websocket request.
// It returns a new WebSocket object if success, or the error if failure.
// Note that the request should be a websocket request, or it will surely fail upgrading.
//
// Deprecated: will be removed in the future, please use third-party websocket library instead.
<原文结束>

# <翻译开始>
// 将当前请求升级为WebSocket请求。
// 如果成功，返回一个新的WebSocket对象，如果失败，则返回错误。
// 注意，请求必须是WebSocket请求，否则升级肯定会失败。
// 
// 警告：将来将被移除，请使用第三方WebSocket库代替。
// md5:68ad17d4740330b2
# <翻译结束>


<原文开始>
// Exit exits executing of current HTTP handler.
<原文结束>

# <翻译开始>
// Exit 退出当前HTTP处理器的执行。. md5:3a3298adda39cc74
# <翻译结束>


<原文开始>
// ExitAll exits executing of current and following HTTP handlers.
<原文结束>

# <翻译开始>
// ExitAll 退出当前及后续的HTTP处理器执行。. md5:53932e5e1bdd10d5
# <翻译结束>


<原文开始>
// ExitHook exits executing of current and following HTTP HOOK handlers.
<原文结束>

# <翻译开始>
// ExitHook 退出当前及后续HTTP钩子处理器的执行。. md5:ef92857b0e046888
# <翻译结束>


<原文开始>
// IsExited checks and returns whether current request is exited.
<原文结束>

# <翻译开始>
// IsExited 检查并返回当前请求是否已退出。. md5:9198deaaaf14733a
# <翻译结束>


<原文开始>
// GetHeader retrieves and returns the header value with given `key`.
<原文结束>

# <翻译开始>
// GetHeader 获取并返回给定`key`对应的头值。. md5:3088bb7beaf8a754
# <翻译结束>


<原文开始>
// GetHost returns current request host name, which might be a domain or an IP without port.
<原文结束>

# <翻译开始>
// GetHost 返回当前请求的主机名，可能是不带端口的域名或IP。. md5:3a06fa36ddefd149
# <翻译结束>


<原文开始>
// IsFileRequest checks and returns whether current request is serving file.
<原文结束>

# <翻译开始>
// IsFileRequest 检查并返回当前请求是否是为文件服务的。. md5:a849769abec62994
# <翻译结束>


<原文开始>
// IsAjaxRequest checks and returns whether current request is an AJAX request.
<原文结束>

# <翻译开始>
// IsAjaxRequest 检查并返回当前请求是否为AJAX请求。. md5:c4e5c9eb4c13dae7
# <翻译结束>


<原文开始>
// GetClientIp returns the client ip of this request without port.
// Note that this ip address might be modified by client header.
<原文结束>

# <翻译开始>
// GetClientIp 返回此请求的客户端IP（不包含端口号）。
// 注意，此IP地址可能已被客户端头部信息修改。
// md5:54dc4a1d4744646a
# <翻译结束>


<原文开始>
// GetRemoteIp returns the ip from RemoteAddr.
<原文结束>

# <翻译开始>
// GetRemoteIp 从 RemoteAddr 中返回 IP 地址。. md5:fca642ffe8c25d8c
# <翻译结束>


<原文开始>
// GetSchema returns the schema of this request.
<原文结束>

# <翻译开始>
// GetSchema 返回此请求的架构。. md5:7bbb51fb51117978
# <翻译结束>


<原文开始>
// GetUrl returns current URL of this request.
<原文结束>

# <翻译开始>
// GetUrl 返回本次请求的当前URL。. md5:8be855812fe4346f
# <翻译结束>


<原文开始>
// GetSessionId retrieves and returns session id from cookie or header.
<原文结束>

# <翻译开始>
// GetSessionId 从cookie或header中检索并返回会话ID。. md5:06fb7350f9f5597f
# <翻译结束>


<原文开始>
// GetReferer returns referer of this request.
<原文结束>

# <翻译开始>
// GetReferer 返回此请求的引荐来源。. md5:4684519c6f7dc2c0
# <翻译结束>


<原文开始>
// GetError returns the error occurs in the procedure of the request.
// It returns nil if there's no error.
<原文结束>

# <翻译开始>
// GetError 返回请求过程中发生的错误。如果没有错误，它将返回 nil。
// md5:035e49a2662f9c04
# <翻译结束>


<原文开始>
// SetError sets custom error for current request.
<原文结束>

# <翻译开始>
// SetError 为当前请求设置自定义错误。. md5:025f3a0817a4be99
# <翻译结束>


<原文开始>
// ReloadParam is used for modifying request parameter.
// Sometimes, we want to modify request parameters through middleware, but directly modifying Request.Body
// is invalid, so it clears the parsed* marks of Request to make the parameters reparsed.
<原文结束>

# <翻译开始>
// ReloadParam 用于修改请求参数。
// 有时，我们希望通过中间件来修改请求参数，但直接修改 Request.Body 是无效的，
// 因此它会清除 Request 的已解析标记，使得参数能够被重新解析。
// md5:af7ff26c797683ef
# <翻译结束>


<原文开始>
// GetHandlerResponse retrieves and returns the handler response object and its error.
<原文结束>

# <翻译开始>
// GetHandlerResponse 获取并返回处理器响应对象及其错误信息。. md5:d6ef2cb1d4fef297
# <翻译结束>


<原文开始>
// GetServeHandler retrieves and returns the user defined handler used to serve this request.
<原文结束>

# <翻译开始>
// GetServeHandler 获取并返回用于处理此请求的用户定义的处理器。. md5:6aef7d779ee52097
# <翻译结束>

