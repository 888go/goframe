
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
// Request is the context object for a request.
<原文结束>

# <翻译开始>
// Request 是一个请求的上下文对象。
# <翻译结束>


<原文开始>
// Corresponding Response of this request.
<原文结束>

# <翻译开始>
// 此请求对应的响应。
# <翻译结束>


<原文开始>
// Request starting time in milliseconds.
<原文结束>

# <翻译开始>
// 请求开始时间（毫秒）
# <翻译结束>


<原文开始>
// Request to end time in milliseconds.
<原文结束>

# <翻译开始>
// 请求结束时的时间（毫秒）。
# <翻译结束>







<原文开始>
// Static file object for static file serving.
<原文结束>

# <翻译开始>
// 静态文件对象，用于静态文件服务。
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
# <翻译结束>


<原文开始>
// All matched handlers containing handler, hook and middleware for this request.
<原文结束>

# <翻译开始>
// 此请求所匹配的所有包含处理器、钩子和中间件的处理程序集合。
# <翻译结束>


<原文开始>
// Real handler serving for this request, not hook or middleware.
<原文结束>

# <翻译开始>
// 此请求的实际处理程序，非钩子或中间件。
# <翻译结束>


<原文开始>
// Handler response object for Request/Response handler.
<原文结束>

# <翻译开始>
// Handler响应对象，用于Request/Response处理器。
# <翻译结束>


<原文开始>
// A bool marking whether the GET parameters parsed.
<原文结束>

# <翻译开始>
// 一个布尔值，标记是否已解析GET参数。
# <翻译结束>


<原文开始>
// A bool marking whether the request body parsed.
<原文结束>

# <翻译开始>
// 一个布尔值，标记请求体是否已解析。
# <翻译结束>


<原文开始>
// A bool marking whether request Form parsed for HTTP method PUT, POST, PATCH.
<原文结束>

# <翻译开始>
// A bool 标记是否已解析HTTP方法 PUT, POST, PATCH 的请求 Form。
# <翻译结束>







<原文开始>
// Router parameters map, which might be nil if there are no router parameters.
<原文结束>

# <翻译开始>
// 路由参数映射表，如果没有路由参数，则可能为nil。
# <翻译结束>


<原文开始>
// Body parameters map, which might be nil if their nobody content.
<原文结束>

# <翻译开始>
// Body参数映射，如果请求体没有内容，则可能为nil。
# <翻译结束>


<原文开始>
// Current executing error of the request.
<原文结束>

# <翻译开始>
// 当前请求正在执行的错误
# <翻译结束>


<原文开始>
// A bool marking whether current request is exited.
<原文结束>

# <翻译开始>
// 一个布尔值，标记当前请求是否已退出。
# <翻译结束>


<原文开始>
// The parsed host name for current host used by GetHost function.
<原文结束>

# <翻译开始>
// GetHost 函数使用的当前主机解析后的主机名。
# <翻译结束>


<原文开始>
// The parsed client ip for current host used by GetClientIp function.
<原文结束>

# <翻译开始>
// GetClientIp 函数当前使用的、已解析的客户端 IP 地址（针对当前主机）。
# <翻译结束>







<原文开始>
// A bool marking whether current request is file serving.
<原文结束>

# <翻译开始>
// 一个布尔值，表示当前请求是否正在提供文件服务。
# <翻译结束>


<原文开始>
// Custom template view engine object for this response.
<原文结束>

# <翻译开始>
// 自定义模板视图引擎对象，用于此次响应。
# <翻译结束>


<原文开始>
// Custom template view variables for this response.
<原文结束>

# <翻译开始>
// 为本次响应自定义模板视图变量。
# <翻译结束>


<原文开始>
// Original URL path that passed from client.
<原文结束>

# <翻译开始>
// 从客户端传递过来的原始URL路径。
# <翻译结束>


<原文开始>
// staticFile is the file struct for static file service.
<原文结束>

# <翻译开始>
// staticFile 是用于静态文件服务的文件结构体。
# <翻译结束>







<原文开始>
// newRequest creates and returns a new request object.
<原文结束>

# <翻译开始>
// newRequest 创建并返回一个新的请求对象。
# <翻译结束>


<原文开始>
// Custom session id creating function.
<原文结束>

# <翻译开始>
// 自定义会话ID生成函数。
# <翻译结束>


<原文开始>
// Remove char '/' in the tail of URI.
<原文结束>

# <翻译开始>
// 删除URI尾部的字符'/'
# <翻译结束>


<原文开始>
// Default URI value if it's empty.
<原文结束>

# <翻译开始>
// 如果URI为空，则使用此默认URI值。
# <翻译结束>


<原文开始>
// WebSocket upgrades current request as a websocket request.
// It returns a new WebSocket object if success, or the error if failure.
// Note that the request should be a websocket request, or it will surely fail upgrading.
<原文结束>

# <翻译开始>
// WebSocket将当前请求升级为websocket请求。
// 如果升级成功，返回一个新的WebSocket对象；如果失败，则返回错误信息。
// 注意，该请求必须是websocket请求，否则升级必定会失败。
# <翻译结束>


<原文开始>
// Exit exits executing of current HTTP handler.
<原文结束>

# <翻译开始>
// Exit 中断当前HTTP处理器的执行。
# <翻译结束>


<原文开始>
// ExitAll exits executing of current and following HTTP handlers.
<原文结束>

# <翻译开始>
// ExitAll 退出当前及后续HTTP处理器的执行。
# <翻译结束>


<原文开始>
// ExitHook exits executing of current and following HTTP HOOK handlers.
<原文结束>

# <翻译开始>
// ExitHook 结束当前及后续 HTTP HOOK 处理器的执行。
# <翻译结束>


<原文开始>
// IsExited checks and returns whether current request is exited.
<原文结束>

# <翻译开始>
// IsExited 检查并返回当前请求是否已退出。
# <翻译结束>


<原文开始>
// GetHeader retrieves and returns the header value with given `key`.
<原文结束>

# <翻译开始>
// GetHeader根据给定的`key`检索并返回头部值。
# <翻译结束>


<原文开始>
// GetHost returns current request host name, which might be a domain or an IP without port.
<原文结束>

# <翻译开始>
// GetHost 返回当前请求的主机名，该主机名可能是域名或不带端口号的IP地址。
# <翻译结束>


<原文开始>
// IsFileRequest checks and returns whether current request is serving file.
<原文结束>

# <翻译开始>
// IsFileRequest 检查并返回当前请求是否正在提供文件服务。
# <翻译结束>


<原文开始>
// IsAjaxRequest checks and returns whether current request is an AJAX request.
<原文结束>

# <翻译开始>
// IsAjaxRequest 检查并返回当前请求是否为 AJAX 请求。
# <翻译结束>


<原文开始>
// GetClientIp returns the client ip of this request without port.
// Note that this ip address might be modified by client header.
<原文结束>

# <翻译开始>
// GetClientIp 返回该请求的客户端IP地址，不包含端口号。
// 注意：此IP地址可能已被客户端头部信息修改。
# <翻译结束>


<原文开始>
// GetRemoteIp returns the ip from RemoteAddr.
<原文结束>

# <翻译开始>
// GetRemoteIp 从 RemoteAddr 返回 IP 地址。
# <翻译结束>


<原文开始>
// GetUrl returns current URL of this request.
<原文结束>

# <翻译开始>
// GetUrl 返回当前请求的URL。
# <翻译结束>


<原文开始>
// GetSessionId retrieves and returns session id from cookie or header.
<原文结束>

# <翻译开始>
// GetSessionId 从cookie或header中检索并返回会话ID。
# <翻译结束>


<原文开始>
// GetReferer returns referer of this request.
<原文结束>

# <翻译开始>
// GetReferer 返回该请求的引用来源。
# <翻译结束>


<原文开始>
// GetError returns the error occurs in the procedure of the request.
// It returns nil if there's no error.
<原文结束>

# <翻译开始>
// GetError 返回在请求过程中发生的错误。
// 如果没有错误，它将返回 nil。
# <翻译结束>


<原文开始>
// SetError sets custom error for current request.
<原文结束>

# <翻译开始>
// SetError为当前请求设置自定义错误。
# <翻译结束>


<原文开始>
// ReloadParam is used for modifying request parameter.
// Sometimes, we want to modify request parameters through middleware, but directly modifying Request.Body
// is invalid, so it clears the parsed* marks of Request to make the parameters reparsed.
<原文结束>

# <翻译开始>
// ReloadParam 用于修改请求参数。
// 有时，我们希望通过中间件来修改请求参数，但直接修改 Request.Body 是无效的，
// 所以它会清除 Request 中已解析标志，以使参数重新解析。
# <翻译结束>


<原文开始>
// GetHandlerResponse retrieves and returns the handler response object and its error.
<原文结束>

# <翻译开始>
// GetHandlerResponse 获取并返回处理器响应对象及其错误信息。
# <翻译结束>


<原文开始>
// GetServeHandler retrieves and returns the user defined handler used to serve this request.
<原文结束>

# <翻译开始>
// GetServeHandler 获取并返回用户自定义的用于处理当前请求的处理器。
# <翻译结束>


<原文开始>
// Middleware manager.
<原文结束>

# <翻译开始>
// 中间件管理器。
# <翻译结束>


<原文开始>
// Custom parameters map.
<原文结束>

# <翻译开始>
// 自定义参数映射。
# <翻译结束>


<原文开始>
// Request body content.
<原文结束>

# <翻译开始>
// 请求体内容。
# <翻译结束>


<原文开始>
// Resource file object.
<原文结束>

# <翻译开始>
// 资源文件对象。
# <翻译结束>

