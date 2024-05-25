
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
// ServeHTTP is the default handler for http request.
// It should not create new goroutine handling the request as
// it's called by am already created new goroutine from http.Server.
//
// This function also makes serve implementing the interface of http.Handler.
<原文结束>

# <翻译开始>
// ServeHTTP 是 http 请求的默认处理器。
// 它不应该为处理请求创建新的goroutine，因为http.Server已经为此创建了一个新的goroutine。
//
// 这个函数还实现了http.Handler接口。
// md5:82dd5f4475c291db
# <翻译结束>


<原文开始>
// Rewrite feature checks.
<原文结束>

# <翻译开始>
// 重写特征检查。 md5:9dab4befbfff965b
# <翻译结束>


<原文开始>
// Create a new request object.
<原文结束>

# <翻译开始>
// 创建一个新的请求对象。 md5:e4e00eb82696932c
# <翻译结束>


<原文开始>
// Get sessionId before user handler
<原文结束>

# <翻译开始>
// 在用户处理器之前获取sessionId. md5:d1a5359c34fed9f9
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
// md5:9514a47b66a76f01
# <翻译结束>


<原文开始>
	// Search the static file with most high priority,
	// which also handle the index files feature.
<原文结束>

# <翻译开始>
// 搜索具有最高优先级的静态文件，同时也处理索引文件功能。
// md5:f618b1fa06ea7acb
# <翻译结束>


<原文开始>
// Search the dynamic service handler.
<原文结束>

# <翻译开始>
// 搜索动态服务处理器。 md5:0cbcd4f2d3569e55
# <翻译结束>


<原文开始>
// Check the service type static or dynamic for current request.
<原文结束>

# <翻译开始>
// 检查当前请求的服务类型是静态还是动态。 md5:642ac02f364c85bc
# <翻译结束>


<原文开始>
// Call custom status handler.
<原文结束>

# <翻译开始>
// 调用自定义状态处理器。 md5:8a7c4e0df133e717
# <翻译结束>


<原文开始>
	// Automatically set the session id to cookie
	// if it creates a new session id in this request
	// and SessionCookieOutput is enabled.
<原文结束>

# <翻译开始>
// 如果在这个请求中生成了新的会话ID，并且启用了SessionCookieOutput，自动将会话ID设置为cookie。
// md5:2c6864797c5d809f
# <翻译结束>


<原文开始>
		// Can change by r.Session.SetId("") before init session
		// Can change by r.Cookie.SetSessionId("")
<原文结束>

# <翻译开始>
// 初始化会话前，可以通过 r.Session.SetId("") 来更改
// 也可以通过 r.Cookie.SetSessionId("") 来更改
// md5:7175563db73b9a50
# <翻译结束>


<原文开始>
// Output the cookie content to the client.
<原文结束>

# <翻译开始>
// 将cookie内容输出到客户端。 md5:b9694a9aa06119db
# <翻译结束>


<原文开始>
// Output the buffer content to the client.
<原文结束>

# <翻译开始>
// 将缓冲区内容输出到客户端。 md5:fe2997ba592b17ad
# <翻译结束>


<原文开始>
	// Close the session, which automatically update the TTL
	// of the session if it exists.
<原文结束>

# <翻译开始>
// 关闭会话，如果会话存在，这将自动更新其TTL（超时时间）。
// md5:a86a4db886c94158
# <翻译结束>


<原文开始>
	// Close the request and response body
	// to release the file descriptor in time.
<原文结束>

# <翻译开始>
// 关闭请求和响应体以及时释放文件描述符。
// md5:aea97d230b2451b0
# <翻译结束>


<原文开始>
// searchStaticFile searches the file with given URI.
// It returns a file struct specifying the file information.
<原文结束>

# <翻译开始>
// searchStaticFile 根据给定的URI搜索文件。
// 它返回一个file结构体，其中包含文件信息。
// md5:e5b76cc2b6c98a07
# <翻译结束>


<原文开始>
// Firstly search the StaticPaths mapping.
<原文结束>

# <翻译开始>
// 首先搜索StaticPaths映射。 md5:4f9c5afa25bf93dd
# <翻译结束>


<原文开始>
// To avoid case like: /static/style -> /static/style.css
<原文结束>

# <翻译开始>
// 为了避免像这样的情况：/static/style -> /static/style.css. md5:74ccef8cd597d359
# <翻译结束>


<原文开始>
// Secondly search the root and searching paths.
<原文结束>

# <翻译开始>
// 其次，搜索根目录和搜索路径。 md5:9b1b9aadf8478052
# <翻译结束>


<原文开始>
// Lastly search the resource manager.
<原文结束>

# <翻译开始>
// 最后搜索资源管理器。 md5:1ccc8123528fc4a4
# <翻译结束>


<原文开始>
// serveFile serves the static file for the client.
// The optional parameter `allowIndex` specifies if allowing directory listing if `f` is a directory.
<原文结束>

# <翻译开始>
// serveFile 为客户端服务静态文件。
// 可选参数 `allowIndex` 指定如果 `f` 是一个目录时是否允许目录列表。
// md5:1741c137e9fcf4cd
# <翻译结束>


<原文开始>
// Use resource file from memory.
<原文结束>

# <翻译开始>
// 从内存中使用资源文件。 md5:eb37e3d39231ad74
# <翻译结束>


<原文开始>
	// Clear the response buffer before file serving.
	// It ignores all custom buffer content and uses the file content.
<原文结束>

# <翻译开始>
// 在服务文件之前清空响应缓冲区。
// 它忽略所有自定义的缓冲内容，转而使用文件内容。
// md5:b7ae0cf8ef13c29c
# <翻译结束>


<原文开始>
// listDir lists the sub files of specified directory as HTML content to the client.
<原文结束>

# <翻译开始>
// listDir 将指定目录下的子文件以HTML内容的形式列出来发送给客户端。 md5:1648438b6fcd2bd5
# <翻译结束>


<原文开始>
// The folder type has the most priority than file.
<原文结束>

# <翻译开始>
// 文件夹类型优先于文件。 md5:f5cc5a85f701d6c1
# <翻译结束>

