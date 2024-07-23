
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
// Default listening port for HTTP.
<原文结束>

# <翻译开始>
// HTTP的默认监听端口。 md5:d368f1a97ab395c2
# <翻译结束>


<原文开始>
// Default listening port for HTTPS.
<原文结束>

# <翻译开始>
// HTTPS的默认监听端口。 md5:72d4583ff29fdc86
# <翻译结束>


<原文开始>
// Method names to the URI converting type, which converts name to its lower case and joins the words using char '-'.
<原文结束>

# <翻译开始>
// 将方法名称转换为URI的类型，该类型将名称转换为小写，并使用字符'-'连接单词。 md5:1fc3cb97052e806f
# <翻译结束>


<原文开始>
// Method names to the URI converting type, which does not convert to the method name.
<原文结束>

# <翻译开始>
// 将方法名称转换为URI转换类型，但不将其转换为方法名称。 md5:75ab2b02cd49ae6b
# <翻译结束>


<原文开始>
// Method names to the URI converting type, which converts name to its lower case.
<原文结束>

# <翻译开始>
// 转换方法名称为URI的类型，将名称转换为小写。 md5:ed7460ac07fbb91a
# <翻译结束>


<原文开始>
// Method names to the URI converting type, which converts name to its camel case.
<原文结束>

# <翻译开始>
// 将方法名称转换为URI的类型，该类型将名称转换为驼峰式命名。 md5:2e028fc00d70d9bf
# <翻译结束>


<原文开始>
// ServerConfig is the HTTP Server configuration manager.
<原文结束>

# <翻译开始>
// ServerConfig 是HTTP服务器的配置管理器。
// 备注: 此配置结构不做名称翻译, 单元测试内的SetConfigWithMap()方法, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// ServerConfig 是HTTP服务器的配置管理器。
// 备注: 此配置结构不做名称翻译, 单元测试内的SetConfigWithMap()方法, 会直接将文本名称转换成配置项名称, 导致找不到原名的配置项. (2024-07-21)
// md5:a2c6c214e9d64d54
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Basic.
	// ======================================================================================================
<原文结束>

# <翻译开始>
	// ======================================================================================================
	// 基础内容.
	// ======================================================================================================
	// md5:accd23363a592bfa
# <翻译结束>


<原文开始>
// Service name, which is for service registry and discovery.
<原文结束>

# <翻译开始>
// 服务名称，用于服务注册和发现。 md5:a0d782209905145d
# <翻译结束>


<原文开始>
	// Address specifies the server listening address like "port" or ":port",
	// multiple addresses joined using ','.
<原文结束>

# <翻译开始>
	// Address 指定服务器监听的地址，如 "port" 或 ":port"，多个地址使用 "," 分隔。
	// md5:ae0da0e5cf0a3e61
# <翻译结束>


<原文开始>
// HTTPSAddr specifies the HTTPS addresses, multiple addresses joined using char ','.
<原文结束>

# <翻译开始>
// HTTPSAddr 指定了HTTPS地址，多个地址使用逗号','连接。 md5:c776678a9eac5a90
# <翻译结束>


<原文开始>
// Listeners specifies the custom listeners.
<原文结束>

# <翻译开始>
// Listeners 指定了自定义的监听器。 md5:3b866f27e7903dac
# <翻译结束>


<原文开始>
// Endpoints are custom endpoints for service register, it uses Address if empty.
<原文结束>

# <翻译开始>
// 如果为空，Endpoints 是服务注册的自定义端点，它使用 Address。 md5:57f627d789f1ea89
# <翻译结束>


<原文开始>
// HTTPSCertPath specifies certification file path for HTTPS service.
<原文结束>

# <翻译开始>
// HTTPSCertPath 指定HTTPS服务的证书文件路径。 md5:3613bb98322987e7
# <翻译结束>


<原文开始>
// HTTPSKeyPath specifies the key file path for HTTPS service.
<原文结束>

# <翻译开始>
// HTTPSKeyPath 指定了HTTPS服务的密钥文件路径。 md5:2a9c1b2c382d01b8
# <翻译结束>


<原文开始>
	// TLSConfig optionally provides a TLS configuration for use
	// by ServeTLS and ListenAndServeTLS. Note that this value is
	// cloned by ServeTLS and ListenAndServeTLS, so it's not
	// possible to modify the configuration with methods like
	// tls.Config.SetSessionTicketKeys. To use
	// SetSessionTicketKeys, use Server.Serve with a TLS Listener
	// instead.
<原文结束>

# <翻译开始>
	// TLSConfig 可选地提供一个用于 ServeTLS 和 ListenAndServeTLS 的 TLS 配置。请注意，此值会被 ServeTLS 和 ListenAndServeTLS 克隆，
	// 因此无法通过如 tls.Config.SetSessionTicketKeys 这样的方法修改配置。若要使用 SetSessionTicketKeys，应使用带有 TLS 监听器的 Server.Serve。
	// md5:16a9af3e2eb3eabc
# <翻译结束>


<原文开始>
// Handler the handler for HTTP request.
<原文结束>

# <翻译开始>
// 处理HTTP请求的处理器。 md5:bd278835c47f74d4
# <翻译结束>


<原文开始>
	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
<原文结束>

# <翻译开始>
	// ReadTimeout是读取整个请求（包括正文）的最大持续时间。
	// 
	// 由于ReadTimeout不允許Handler根据每个请求体的可接受截止日期或上传速率做出逐个请求的决定，大多数用户更喜欢使用ReadHeaderTimeout。同时使用它们两者也是可以的。
	// md5:45add6b4a3777e9a
# <翻译结束>


<原文开始>
	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
<原文结束>

# <翻译开始>
	// WriteTimeout 是响应写入操作超时的最长时间。每当读取新请求的头信息时，它都会重置。与 ReadTimeout 相似，它不允许处理程序针对每个请求做出超时时间的决定。
	// md5:7cba2c215a8e6f3d
# <翻译结束>


<原文开始>
	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alive are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
<原文结束>

# <翻译开始>
	// IdleTimeout 是在保持连接开启状态下，等待下一个请求的最大时间。如果IdleTimeout为零，则使用ReadTimeout的值。如果两者都为零，则没有超时设置。
	// md5:340816400bd04176
# <翻译结束>


<原文开始>
	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	//
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's 10240 bytes in default.
<原文结束>

# <翻译开始>
	// MaxHeaderBytes 控制服务器解析请求头的键值对（包括请求行）时的最大字节数。它不限制请求体的大小。
	// 
	// 可以在配置文件中通过字符串进行配置，例如：1m、10m、500kb等。默认值为 10240 字节。
	// md5:8a61cc04b245e3d0
# <翻译结束>


<原文开始>
// KeepAlive enables HTTP keep-alive.
<原文结束>

# <翻译开始>
// EnableHTTPKeepAlive 启用 HTTP 保持活动。 md5:316ccd3dea9e6e4e
# <翻译结束>


<原文开始>
	// ServerAgent specifies the server agent information, which is wrote to
	// HTTP response header as "Server".
<原文结束>

# <翻译开始>
	// ServerAgent 指定了服务器代理信息，该信息将被写入HTTP响应头中作为 "Server"。
	// md5:08ba94b701f34cac
# <翻译结束>


<原文开始>
// View specifies the default template view object for the server.
<原文结束>

# <翻译开始>
// View 指定了服务器的默认模板视图对象。 md5:1d7f3415e9af9116
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Static.
	// ======================================================================================================
<原文结束>

# <翻译开始>
	// ======================================================================================================
	// 静态部分。
	// ======================================================================================================
	// md5:67b32db36648f355
# <翻译结束>


<原文开始>
// Rewrites specifies the URI rewrite rules map.
<原文结束>

# <翻译开始>
// Rewrites 指定了URI重写规则映射。 md5:177b674dcc27b406
# <翻译结束>


<原文开始>
// IndexFiles specifies the index files for static folder.
<原文结束>

# <翻译开始>
// IndexFiles 指定了静态文件夹的索引文件。 md5:2beaa7ae62753536
# <翻译结束>


<原文开始>
	// IndexFolder specifies if listing sub-files when requesting folder.
	// The server responses HTTP status code 403 if it is false.
<原文结束>

# <翻译开始>
	// IndexFolder 指定是否在请求文件夹时列出子文件。
	// 如果为 false，服务器将返回 HTTP 状态码 403。
	// md5:2d0bfed1b08c9bae
# <翻译结束>


<原文开始>
// ServerRoot specifies the root directory for static service.
<原文结束>

# <翻译开始>
// ServerRoot 指定静态服务的根目录。 md5:19d6dd74ec61bd29
# <翻译结束>


<原文开始>
// SearchPaths specifies additional searching directories for static service.
<原文结束>

# <翻译开始>
// SearchPaths 指定静态服务的额外搜索目录。 md5:bd172212e30738c5
# <翻译结束>


<原文开始>
// StaticPaths specifies URI to directory mapping array.
<原文结束>

# <翻译开始>
// StaticPaths 指定了URI到目录的映射数组。 md5:ea14f728a07fbd14
# <翻译结束>


<原文开始>
	// FileServerEnabled is the global switch for static service.
	// It is automatically set enabled if any static path is set.
<原文结束>

# <翻译开始>
	// FileServerEnabled 是静态服务的全局开关。
	// 如果设置了任何静态路径，它将自动设置为启用。
	// md5:c3edec83e9cfd01f
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Cookie.
	// ======================================================================================================
<原文结束>

# <翻译开始>
	// ======================================================================================================
	//  cookie。
	// ======================================================================================================
	// md5:82bbf8b719e4196b
# <翻译结束>


<原文开始>
// CookieMaxAge specifies the max TTL for cookie items.
<原文结束>

# <翻译开始>
// CookieMaxAge 指定 cookie 项的最大过期时间（TTL）。 md5:c6712f43beedeb7d
# <翻译结束>


<原文开始>
	// CookiePath specifies cookie path.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
	// CookiePath 指定了Cookie路径。
	// 它也影响了会话ID的默认存储位置。
	// md5:dad6b405aee54d41
# <翻译结束>


<原文开始>
	// CookieDomain specifies cookie domain.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
	// CookieDomain 指定Cookie的域名。
	// 它也会影响默认的session id存储位置。
	// md5:f2d433c3779a94df
# <翻译结束>


<原文开始>
	// CookieSameSite specifies cookie SameSite property.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
	// CookieSameSite 指定 cookie 的 SameSite 属性。它还影响会话 ID 的默认存储。
	// md5:bba3b21b29719a23
# <翻译结束>


<原文开始>
	// CookieSameSite specifies cookie Secure property.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
	// CookieSameSite 指定 cookie 的 Secure 属性。
	// 它还会影响会话 ID 的默认存储。
	// md5:e861b293a54a4909
# <翻译结束>


<原文开始>
	// CookieSameSite specifies cookie HttpOnly property.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
	// CookieSameSite 指定了cookie的HttpOnly属性。
	// 它也影响会话ID的默认存储方式。
	// md5:ffc9065ce43afdd9
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Session.
	// ======================================================================================================
<原文结束>

# <翻译开始>
	// ======================================================================================================
	// 会话管理。
	// ======================================================================================================
	// md5:bf2dda055c30c648
# <翻译结束>


<原文开始>
// SessionIdName specifies the session id name.
<原文结束>

# <翻译开始>
// SessionIdName 指定会话ID的名称。 md5:18c5a80d34f75272
# <翻译结束>


<原文开始>
// SessionMaxAge specifies max TTL for session items.
<原文结束>

# <翻译开始>
// SessionMaxAge 指定会话项的最大超时时间（TTL）。 md5:cc78683f7c70c955
# <翻译结束>


<原文开始>
	// SessionPath specifies the session storage directory path for storing session files.
	// It only makes sense if the session storage is type of file storage.
<原文结束>

# <翻译开始>
	// SessionPath 指定了用于存储会话文件的会话存储目录路径。
	// 仅当会话存储类型为文件存储时，此设置才有意义。
	// md5:ace2dbc78e7f3a04
# <翻译结束>


<原文开始>
// SessionStorage specifies the session storage.
<原文结束>

# <翻译开始>
// SessionStorage 指定会话存储。 md5:678a55a9e339a25c
# <翻译结束>


<原文开始>
	// SessionCookieMaxAge specifies the cookie ttl for session id.
	// If it is set 0, it means it expires along with browser session.
<原文结束>

# <翻译开始>
	// SessionCookieMaxAge 指定会话ID的cookie过期时间。如果设置为0，表示它将随着浏览器会话一起过期。
	// md5:7b48b403d924198e
# <翻译结束>


<原文开始>
// SessionCookieOutput specifies whether automatic outputting session id to cookie.
<原文结束>

# <翻译开始>
// SessionCookieOutput 指定是否自动将会话ID输出到cookie。 md5:040824f71a38e446
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Logging.
	// ======================================================================================================
<原文结束>

# <翻译开始>
	// ===============================================================================================
	// 日志记录。
	// ===============================================================================================
	// md5:ede34792b995e698
# <翻译结束>


<原文开始>
// Logger specifies the logger for server.
<原文结束>

# <翻译开始>
// Logger 指定服务器使用的日志记录器。 md5:004b9e605f068eec
# <翻译结束>


<原文开始>
// LogPath specifies the directory for storing logging files.
<原文结束>

# <翻译开始>
// LogPath 指定日志文件的存储目录。 md5:de94cd356ae22e8a
# <翻译结束>


<原文开始>
// LogLevel specifies the logging level for logger.
<原文结束>

# <翻译开始>
// LogLevel 指定了logger的日志级别。 md5:87d747e517ace64c
# <翻译结束>


<原文开始>
// LogStdout specifies whether printing logging content to stdout.
<原文结束>

# <翻译开始>
// LogStdout 指定是否将日志内容打印到标准输出（stdout）。 md5:f591098f0447f3f8
# <翻译结束>


<原文开始>
// ErrorStack specifies whether logging stack information when error.
<原文结束>

# <翻译开始>
// ErrorStack 指定在发生错误时是否记录堆栈信息。 md5:5cfe84f341788adc
# <翻译结束>


<原文开始>
// ErrorLogEnabled enables error logging content to files.
<原文结束>

# <翻译开始>
// ErrorLogEnabled 启用将错误日志内容写入文件。 md5:9065ef46c6d983d0
# <翻译结束>


<原文开始>
// ErrorLogPattern specifies the error log file pattern like: error-{Ymd}.log
<原文结束>

# <翻译开始>
// ErrorLogPattern 指定错误日志文件模式，如：error-YYYYMMDD.log. md5:c59be38d6eeea7aa
# <翻译结束>


<原文开始>
// AccessLogEnabled enables access logging content to files.
<原文结束>

# <翻译开始>
// AccessLogEnabled 启用将访问日志内容记录到文件中。 md5:6867f80f6ec7fb95
# <翻译结束>


<原文开始>
// AccessLogPattern specifies the error log file pattern like: access-{Ymd}.log
<原文结束>

# <翻译开始>
// AccessLogPattern 指定错误日志文件的模式，如：access-{Ymd}.log. md5:e01474de0152ebf6
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// PProf.
	// ======================================================================================================
<原文结束>

# <翻译开始>
	// ======================================================================================================
	// PProf。
	// ======================================================================================================
	// md5:94f131ef860cf923
# <翻译结束>


<原文开始>
// PProfEnabled enables PProf feature.
<原文结束>

# <翻译开始>
// PProfEnabled 启用 PProf 功能。 md5:4847ea23da60be23
# <翻译结束>


<原文开始>
// PProfPattern specifies the PProf service pattern for router.
<原文结束>

# <翻译开始>
// PProfPattern 为路由器指定 PProf 服务的模式。 md5:ddf66608babb16a1
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// API & Swagger.
	// ======================================================================================================
<原文结束>

# <翻译开始>
	// =======================================================================================
	// API与Swagger相关代码。
	// =======================================================================================
	// md5:ec4583c0e3c0aab6
# <翻译结束>


<原文开始>
// OpenApiPath specifies the OpenApi specification file path.
<原文结束>

# <翻译开始>
// OpenApiPath 指定OpenApi规范文件的路径。 md5:a99446ffbab82145
# <翻译结束>


<原文开始>
// SwaggerPath specifies the swagger UI path for route registering.
<原文结束>

# <翻译开始>
// SwaggerPath 定义路由注册的Swagger UI路径。 md5:abc9988ac6d860c6
# <翻译结束>


<原文开始>
// SwaggerUITemplate specifies the swagger UI custom template
<原文结束>

# <翻译开始>
// SwaggerUITemplate 指定 Swagger UI 的自定义模板. md5:0f381e185ab07c43
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Other.
	// ======================================================================================================
<原文结束>

# <翻译开始>
	// ======================================================================================================
	// 其他。
	// ======================================================================================================
	// md5:8a9c4a8ec79cdc30
# <翻译结束>


<原文开始>
	// ClientMaxBodySize specifies the max body size limit in bytes for client request.
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's `8MB` in default.
<原文结束>

# <翻译开始>
	// ClientMaxBodySize 指定了客户端请求的最大体大小限制（以字节为单位）。它可以在配置文件中通过字符串进行配置，例如：1m、10m、500kb 等。默认值为 `8MB`。
	// md5:2ae357b9e73e0ba6
# <翻译结束>


<原文开始>
	// FormParsingMemory specifies max memory buffer size in bytes which can be used for
	// parsing multimedia form.
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's 1MB in default.
<原文结束>

# <翻译开始>
	// FormParsingMemory 定义了用于解析多媒体表单的最大内存缓冲区大小（以字节为单位）。可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串进行配置。默认值为 1MB。
	// md5:e7808b8ee0d32ae1
# <翻译结束>


<原文开始>
	// NameToUriType specifies the type for converting struct method name to URI when
	// registering routes.
<原文结束>

# <翻译开始>
	// NameToUriType 用于指定在注册路由时，将结构体方法名转换为URI的类型。
	// md5:3853020c5284d13d
# <翻译结束>


<原文开始>
// RouteOverWrite allows to overwrite the route if duplicated.
<原文结束>

# <翻译开始>
// RouteOverWrite 允许覆盖重复的路由。 md5:11e5811ec1ba25ca
# <翻译结束>


<原文开始>
// DumpRouterMap specifies whether automatically dumps router map when server starts.
<原文结束>

# <翻译开始>
// DumpRouterMap 指定服务器启动时是否自动dump路由器映射。 md5:5c37a6000e9858ab
# <翻译结束>


<原文开始>
// Graceful enables graceful reload feature for all servers of the process.
<原文结束>

# <翻译开始>
// Graceful启用进程所有服务器的优雅重启功能。 md5:e4f67dc7d507232e
# <翻译结束>


<原文开始>
// GracefulTimeout set the maximum survival time (seconds) of the parent process.
<原文结束>

# <翻译开始>
// GracefulTimeout 设置父进程的最大生存时间（秒）。 md5:09d4293175059ede
# <翻译结束>


<原文开始>
// GracefulShutdownTimeout set the maximum survival time (seconds) before stopping the server.
<原文结束>

# <翻译开始>
// GracefulShutdownTimeout 设置在停止服务器之前，允许服务器最大存活的时间（秒数）。 md5:b220917a3a4e4ebf
# <翻译结束>


<原文开始>
// NewConfig creates and returns a ServerConfig object with default configurations.
// Note that, do not define this default configuration to local package variable, as there are
// some pointer attributes that may be shared in different servers.
<原文结束>

# <翻译开始>
// NewConfig 创建并返回一个带有默认配置的ServerConfig对象。
// 注意，不要将这些默认配置定义为本地包变量，因为存在一些指针属性可能在不同的服务器中被共享。
// md5:3a8bce955120579e
# <翻译结束>


<原文开始>
// ConfigFromMap creates and returns a ServerConfig object with given map and
// default configuration object.
<原文结束>

# <翻译开始>
// ConfigFromMap根据给定的映射和默认配置对象，创建并返回一个ServerConfig对象。
// md5:4e735da260d3d596
# <翻译结束>


<原文开始>
// SetConfigWithMap sets the configuration for the server using map.
<原文结束>

# <翻译开始>
// SetConfigWithMap 使用映射(map)设置服务器的配置。 md5:896223070f4171c3
# <翻译结束>


<原文开始>
	// The m now is a shallow copy of m.
	// Any changes to m does not affect the original one.
	// A little tricky, isn't it?
<原文结束>

# <翻译开始>
	// m 现在是 m 的浅拷贝。
	// 对 m 的任何修改都不会影响原始对象。
	// 这有点巧妙，不是吗？
	// md5:4d1dd38c4db57a79
# <翻译结束>


<原文开始>
	// Allow setting the size configuration items using string size like:
	// 1m, 100mb, 512kb, etc.
<原文结束>

# <翻译开始>
	// 允许使用字符串大小（如：1m、100mb、512kb等）来设置配置项的大小。
	// md5:1afcd879d2e46708
# <翻译结束>


<原文开始>
	// Update the current configuration object.
	// It only updates the configured keys not all the object.
<原文结束>

# <翻译开始>
	// 更新当前配置对象。
	// 它只会更新已配置的键，而不是整个对象。
	// md5:31a4f7b2094c229e
# <翻译结束>


<原文开始>
// SetConfig sets the configuration for the server.
<原文结束>

# <翻译开始>
// SetConfig 为服务器设置配置。 md5:cf09327be418468a
# <翻译结束>


<原文开始>
// Automatically add ':' prefix for address if it is missed.
<原文结束>

# <翻译开始>
// 如果地址缺少':'前缀，自动添加。 md5:7493e54c133e3353
# <翻译结束>


<原文开始>
// SetAddr sets the listening address for the server.
// The address is like ':80', '0.0.0.0:80', '127.0.0.1:80', '180.18.99.10:80', etc.
<原文结束>

# <翻译开始>
// SetAddr 设置服务器的监听地址。
// 地址格式为 ':80'，'0.0.0.0:80'，'127.0.0.1:80'，'180.18.99.10:80' 等。
// md5:c62ee3ae1a0d6764
# <翻译结束>


<原文开始>
// SetPort sets the listening ports for the server.
// The listening ports can be multiple like: SetPort(80, 8080).
<原文结束>

# <翻译开始>
// SetPort 设置服务器的监听端口。
// 可以设置多个端口，例如：SetPort(80, 8080)。
// md5:e6fc730d6e8ee17c
# <翻译结束>


<原文开始>
// SetHTTPSAddr sets the HTTPS listening ports for the server.
<原文结束>

# <翻译开始>
// SetHTTPSAddr 设置服务器的HTTPS监听端口。 md5:2cfcb725865474b8
# <翻译结束>


<原文开始>
// SetHTTPSPort sets the HTTPS listening ports for the server.
// The listening ports can be multiple like: SetHTTPSPort(443, 500).
<原文结束>

# <翻译开始>
// SetHTTPSPort 设置服务器的HTTPS监听端口。
// 可以设置多个监听端口，如：SetHTTPSPort(443, 500)。
// md5:cdbfe394641fff8a
# <翻译结束>


<原文开始>
// SetListener set the custom listener for the server.
<原文结束>

# <翻译开始>
// SetListener 设置服务器的自定义监听器。 md5:5ce2e30da0c2811f
# <翻译结束>


<原文开始>
// EnableHTTPS enables HTTPS with given certification and key files for the server.
// The optional parameter `tlsConfig` specifies custom TLS configuration.
<原文结束>

# <翻译开始>
// EnableHTTPS 为服务器启用HTTPS，使用给定的证书和密钥文件。可选参数`tlsConfig`用于指定自定义TLS配置。
// md5:0e566bf091aacfd8
# <翻译结束>


<原文开始>
// SetTLSConfig sets custom TLS configuration and enables HTTPS feature for the server.
<原文结束>

# <翻译开始>
// SetTLSConfig 设置自定义的TLS配置，并为服务器启用HTTPS功能。 md5:e8fae606e7c9daa6
# <翻译结束>


<原文开始>
// SetReadTimeout sets the ReadTimeout for the server.
<原文结束>

# <翻译开始>
// SetReadTimeout 设置服务器的读取超时时间。 md5:731a0457ad074a1e
# <翻译结束>


<原文开始>
// SetWriteTimeout sets the WriteTimeout for the server.
<原文结束>

# <翻译开始>
// SetWriteTimeout 设置服务器的写超时。 md5:60f9efbd1b42a85a
# <翻译结束>


<原文开始>
// SetIdleTimeout sets the IdleTimeout for the server.
<原文结束>

# <翻译开始>
// SetIdleTimeout 设置服务器的空闲超时时间。 md5:f18e89634fa33c02
# <翻译结束>


<原文开始>
// SetMaxHeaderBytes sets the MaxHeaderBytes for the server.
<原文结束>

# <翻译开始>
// SetMaxHeaderBytes 为服务器设置 MaxHeaderBytes。 md5:2e7198560eedacbb
# <翻译结束>


<原文开始>
// SetServerAgent sets the ServerAgent for the server.
<原文结束>

# <翻译开始>
// SetServerAgent 设置服务器的ServerAgent。 md5:ac1c65804355cc50
# <翻译结束>


<原文开始>
// SetKeepAlive sets the KeepAlive for the server.
<原文结束>

# <翻译开始>
// SetKeepAlive 设置服务器的KeepAlive。 md5:54c342f49d9fa171
# <翻译结束>


<原文开始>
// SetView sets the View for the server.
<原文结束>

# <翻译开始>
// SetView 设置服务器的视图。 md5:ec7bba6db1e3a9cf
# <翻译结束>


<原文开始>
// GetName returns the name of the server.
<原文结束>

# <翻译开始>
// GetName 返回服务器的名称。 md5:1662443760c53efe
# <翻译结束>


<原文开始>
// SetName sets the name for the server.
<原文结束>

# <翻译开始>
// SetName 设置服务器的名称。 md5:242f311a4c185514
# <翻译结束>


<原文开始>
// SetEndpoints sets the Endpoints for the server.
<原文结束>

# <翻译开始>
// SetEndpoints 设置服务器的端点。 md5:b75987e400904902
# <翻译结束>


<原文开始>
// SetHandler sets the request handler for server.
<原文结束>

# <翻译开始>
// SetHandler 设置服务器的请求处理器。 md5:c2ce7c70be19e1d5
# <翻译结束>


<原文开始>
// GetHandler returns the request handler of the server.
<原文结束>

# <翻译开始>
// GetHandler 返回服务器的请求处理程序。 md5:97d22a3db48dd77d
# <翻译结束>


<原文开始>
// SetRegistrar sets the Registrar for server.
<原文结束>

# <翻译开始>
// SetRegistrar 设置服务器的注册器。 md5:59baf7cae4845598
# <翻译结束>


<原文开始>
// GetRegistrar returns the Registrar of server.
<原文结束>

# <翻译开始>
// GetRegistrar 返回服务器的注册商。 md5:d5a67dbd4e6ac976
# <翻译结束>

