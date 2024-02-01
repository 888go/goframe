
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
// ServeHTTP is the default handler for http request.
// It should not create new goroutine handling the request as
// it's called by am already created new goroutine from http.Server.
//
// This function also makes serve implementing the interface of http.Handler.
<原文结束>

# <翻译开始>
// ServeHTTP 是处理 HTTP 请求的默认处理器。
// 由于它是由 http.Server 已经创建的新 goroutine 调用的，所以不应在此函数中创建新的处理请求的 goroutine。
//
// 此外，这个函数实现了 http.Handler 接口。
# <翻译结束>












<原文开始>
// Create a new request object.
<原文结束>

# <翻译开始>
// 创建一个新的请求对象。
# <翻译结束>


<原文开始>
// Get sessionId before user handler
<原文结束>

# <翻译开始>
// 在用户处理器之前获取sessionId
# <翻译结束>












<原文开始>
		// Close the session, which automatically update the TTL
		// of the session if it exists.
<原文结束>

# <翻译开始>
// 关闭会话，如果会话存在，会自动更新其TTL（生存时间）
# <翻译结束>


<原文开始>
		// Close the request and response body
		// to release the file descriptor in time.
<原文结束>

# <翻译开始>
// 关闭请求和响应体
// 以便及时释放文件描述符。
# <翻译结束>


<原文开始>
	// ============================================================
	// Priority:
	// Static File > Dynamic Service > Static Directory
	// ============================================================
<原文结束>

# <翻译开始>
// ============================================================
// 优先级：
// 静态文件 > 动态服务 > 静态目录
// ============================================================
# <翻译结束>


<原文开始>
	// Search the static file with most high priority,
	// which also handle the index files feature.
<原文结束>

# <翻译开始>
// 搜索优先级最高的静态文件，并处理索引文件特性。
# <翻译结束>


<原文开始>
// Search the dynamic service handler.
<原文结束>

# <翻译开始>
// 搜索动态服务处理器。
# <翻译结束>


<原文开始>
// Check the service type static or dynamic for current request.
<原文结束>

# <翻译开始>
// 检查当前请求的服务类型是静态还是动态。
# <翻译结束>















































<原文开始>
// Call custom status handler.
<原文结束>

# <翻译开始>
// 调用自定义状态处理器。
# <翻译结束>


<原文开始>
	// Automatically set the session id to cookie
	// if it creates a new session id in this request
	// and SessionCookieOutput is enabled.
<原文结束>

# <翻译开始>
// 如果在本次请求中创建了新的会话ID，并且启用了SessionCookieOutput，则自动将会话ID设置到cookie中。
# <翻译结束>


<原文开始>
		// Can change by r.Session.SetId("") before init session
		// Can change by r.Cookie.SetSessionId("")
<原文结束>

# <翻译开始>
// 在初始化session之前，可以通过r.Session.SetId("")来改变
// 也可以通过r.Cookie.SetSessionId("")来改变
# <翻译结束>


<原文开始>
// Output the cookie content to the client.
<原文结束>

# <翻译开始>
// 将cookie内容输出到客户端。
# <翻译结束>


<原文开始>
// Output the buffer content to the client.
<原文结束>

# <翻译开始>
// 将缓冲区内容输出到客户端。
# <翻译结束>







<原文开始>
// searchStaticFile searches the file with given URI.
// It returns a file struct specifying the file information.
<原文结束>

# <翻译开始>
// searchStaticFile 通过给定的 URI 搜索文件。
// 它返回一个文件结构体，该结构体指定了文件信息。
# <翻译结束>


<原文开始>
// Firstly search the StaticPaths mapping.
<原文结束>

# <翻译开始>
// 首先搜索 StaticPaths 映射。
# <翻译结束>


<原文开始>
// To avoid case like: /static/style -> /static/style.css
<原文结束>

# <翻译开始>
// 为避免出现类似这种情况：/static/style -> /static/style.css 的情况
# <翻译结束>


<原文开始>
// Secondly search the root and searching paths.
<原文结束>

# <翻译开始>
// 其次，搜索根目录和搜索路径。
# <翻译结束>


<原文开始>
// Lastly search the resource manager.
<原文结束>

# <翻译开始>
// 最后在资源管理器中进行搜索。
# <翻译结束>


<原文开始>
// serveFile serves the static file for the client.
// The optional parameter `allowIndex` specifies if allowing directory listing if `f` is a directory.
<原文结束>

# <翻译开始>
// serveFile 为客户端提供静态文件服务。
// 可选参数 `allowIndex` 指定当 `f` 是目录时，是否允许目录列表展示。
# <翻译结束>


<原文开始>
// Use resource file from memory.
<原文结束>

# <翻译开始>
// 从内存中使用资源文件。
# <翻译结束>







<原文开始>
	// Clear the response buffer before file serving.
	// It ignores all custom buffer content and uses the file content.
<原文结束>

# <翻译开始>
// 在文件服务之前清空响应缓冲区。
// 它会忽略所有自定义缓冲区内容，并使用文件内容。
# <翻译结束>


<原文开始>
// listDir lists the sub files of specified directory as HTML content to the client.
<原文结束>

# <翻译开始>
// listDir 将指定目录下的子文件列表以HTML内容的形式发送给客户端。
# <翻译结束>


<原文开始>
// The folder type has the most priority than file.
<原文结束>

# <翻译开始>
// 文件夹类型比文件类型具有更高的优先级。
# <翻译结束>


<原文开始>
// Max body size limit.
<原文结束>

# <翻译开始>
// 最大请求体大小限制。
# <翻译结束>


<原文开始>
// Rewrite feature checks.
<原文结束>

# <翻译开始>
// 重写特性检查。
# <翻译结束>


<原文开始>
// error log handling.
<原文结束>

# <翻译开始>
// 错误日志处理。
# <翻译结束>


<原文开始>
// access log handling.
<原文结束>

# <翻译开始>
// 访问日志处理。
# <翻译结束>


<原文开始>
// HOOK - BeforeServe
<原文结束>

# <翻译开始>
// HOOK - 服务启动前
# <翻译结束>


<原文开始>
// Core serving handling.
<原文结束>

# <翻译开始>
// 核心服务处理。
# <翻译结束>


<原文开始>
// Static file service.
<原文结束>

# <翻译开始>
// 静态文件服务。
# <翻译结束>


<原文开始>
// Serve the directory.
<原文结束>

# <翻译开始>
// 服务目录（提供目录内容）
# <翻译结束>


<原文开始>
// HOOK - BeforeOutput
<原文结束>

# <翻译开始>
// HOOK - 输出之前
# <翻译结束>


<原文开始>
// HTTP status checking.
<原文结束>

# <翻译开始>
// HTTP状态检查。
# <翻译结束>


<原文开始>
// HTTP status handler.
<原文结束>

# <翻译开始>
// HTTP状态处理器。
# <翻译结束>


<原文开始>
// HOOK - AfterOutput
<原文结束>

# <翻译开始>
// HOOK - 输出后
# <翻译结束>


<原文开始>
// Use file from dist.
<原文结束>

# <翻译开始>
// 使用来自dist目录的文件。
# <翻译结束>

