
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
// Default listening port for HTTP.
<原文结束>

# <翻译开始>
// 默认的HTTP监听端口。
# <翻译结束>


<原文开始>
// Default listening port for HTTPS.
<原文结束>

# <翻译开始>
// 默认的HTTPS监听端口。
# <翻译结束>


<原文开始>
// Method names to the URI converting type, which does not convert to the method name.
<原文结束>

# <翻译开始>
// MethodNamesToURI 是一个用于将方法名转换为URI的类型，但并不转换回方法名。
# <翻译结束>


<原文开始>
// Method names to the URI converting type, which converts name to its lower case.
<原文结束>

# <翻译开始>
// MethodNamesToURI 是一个将方法名转换为 URI 的类型，该类型会将名称转换为其小写形式。
# <翻译结束>


<原文开始>
// Method names to the URI converting type, which converts name to its camel case.
<原文结束>

# <翻译开始>
// MethodNamesToURI 是一个用于将方法名称转换为 URI 的类型，该类型将名称转换为其驼峰式表示。
# <翻译结束>


<原文开始>
// ServerConfig is the HTTP Server configuration manager.
<原文结束>

# <翻译开始>
// ServerConfig 是 HTTP 服务器的配置管理器。
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Basic.
	// ======================================================================================================
<原文结束>

# <翻译开始>
// ======================================================================================================
// 基础部分.
// ======================================================================================================
// 这段注释表明接下来的代码是关于“基础”部分，用于概括或分隔代码的不同模块或功能区块。
# <翻译结束>


<原文开始>
// Service name, which is for service registry and discovery.
<原文结束>

# <翻译开始>
// 服务名称，用于服务注册与发现。
# <翻译结束>


<原文开始>
	// Address specifies the server listening address like "port" or ":port",
	// multiple addresses joined using ','.
<原文结束>

# <翻译开始>
// Address 指定服务器监听地址，格式如 "端口" 或 ":端口"，
// 多个地址之间使用 ',' 连接。
# <翻译结束>


<原文开始>
// HTTPSAddr specifies the HTTPS addresses, multiple addresses joined using char ','.
<原文结束>

# <翻译开始>
// HTTPSAddr 指定HTTPS服务地址，多个地址之间使用逗号（,）连接。
# <翻译结束>


<原文开始>
// Listeners specifies the custom listeners.
<原文结束>

# <翻译开始>
// Listeners 指定自定义监听器。
# <翻译结束>


<原文开始>
// Endpoints are custom endpoints for service register, it uses Address if empty.
<原文结束>

# <翻译开始>
// Endpoints 是服务注册的自定义端点，如果为空则使用 Address。
# <翻译结束>


<原文开始>
// HTTPSCertPath specifies certification file path for HTTPS service.
<原文结束>

# <翻译开始>
// HTTPSCertPath 指定 HTTPS 服务的证书文件路径。
# <翻译结束>


<原文开始>
// HTTPSKeyPath specifies the key file path for HTTPS service.
<原文结束>

# <翻译开始>
// HTTPSKeyPath 指定HTTPS服务的密钥文件路径。
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
// TLSConfig 提供一个可选的 TLS 配置，用于 ServeTLS 和 ListenAndServeTLS。请注意，
// 此值将被 ServeTLS 和 ListenAndServeTLS 拷贝使用，因此不能通过诸如 tls.Config.SetSessionTicketKeys
// 等方法修改配置。若要使用 SetSessionTicketKeys，请改用 Server.Serve 方法配合一个 TLS 监听器来实现。
# <翻译结束>


<原文开始>
// Handler the handler for HTTP request.
<原文结束>

# <翻译开始>
// Handler HTTP请求的处理器。
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
// ReadTimeout 是读取整个请求（包括主体）的最大持续时间。
//
// 由于 ReadTimeout 不允许处理程序为每个请求主体设置可接受的截止日期或上传速率，
// 大多数用户可能会更倾向于使用 ReadHeaderTimeout。同时使用它们是有效的。
# <翻译结束>


<原文开始>
	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
<原文结束>

# <翻译开始>
// WriteTimeout 是在超时前写入响应的最大持续时间。每当读取到新请求的头部时，该时间就会重置。类似于 ReadTimeout，它并不允许 Handlers 根据每个请求自行决定是否超时。
# <翻译结束>


<原文开始>
	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alive are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
<原文结束>

# <翻译开始>
// IdleTimeout 是在启用 keep-alive 时，等待下一个请求的最大时间间隔。如果 IdleTimeout 设为零，则使用 ReadTimeout 的值。如果两者都为零，则表示没有超时限制。
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
// MaxHeaderBytes 控制服务器在解析请求头（包括请求行）的键和值时，
// 会读取的最大字节数。但请注意，它并不会限制请求体的大小。
//
// 你可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串来配置这个参数。
// 默认情况下，其值为 10240 字节。
# <翻译结束>


<原文开始>
// KeepAlive enables HTTP keep-alive.
<原文结束>

# <翻译开始>
// KeepAlive 启用 HTTP 保持连接（Keep-alive）功能。
# <翻译结束>


<原文开始>
	// ServerAgent specifies the server agent information, which is wrote to
	// HTTP response header as "Server".
<原文结束>

# <翻译开始>
// ServerAgent 指定服务器代理信息，该信息会被写入
// HTTP 响应头中作为 "Server"。
# <翻译结束>


<原文开始>
// View specifies the default template view object for the server.
<原文结束>

# <翻译开始>
// View 指定了服务器的默认模板视图对象。
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
// 这段注释表明了该处代码是“静态”部分，但具体含义可能需要更多上下文信息来准确翻译。在程序中，“Static”通常是指不变的、非运行时动态改变的数据或函数，或者是初始化后在整个程序生命周期中保持固定的部分。
# <翻译结束>


<原文开始>
// Rewrites specifies the URI rewrite rules map.
<原文结束>

# <翻译开始>
// Rewrites 指定了 URI 重写规则映射。
# <翻译结束>


<原文开始>
// IndexFiles specifies the index files for static folder.
<原文结束>

# <翻译开始>
// IndexFiles 指定静态文件夹的索引文件。
# <翻译结束>


<原文开始>
	// IndexFolder specifies if listing sub-files when requesting folder.
	// The server responses HTTP status code 403 if it is false.
<原文结束>

# <翻译开始>
// IndexFolder 指定在请求文件夹时是否列出子文件。
// 如果该值为false，服务器将返回HTTP状态码403。
# <翻译结束>


<原文开始>
// ServerRoot specifies the root directory for static service.
<原文结束>

# <翻译开始>
// ServerRoot 指定静态服务的根目录。
# <翻译结束>


<原文开始>
// SearchPaths specifies additional searching directories for static service.
<原文结束>

# <翻译开始>
// SearchPaths 指定静态服务的额外搜索目录。
# <翻译结束>


<原文开始>
// StaticPaths specifies URI to directory mapping array.
<原文结束>

# <翻译开始>
// StaticPaths 指定了URI到目录映射的数组。
# <翻译结束>


<原文开始>
	// FileServerEnabled is the global switch for static service.
	// It is automatically set enabled if any static path is set.
<原文结束>

# <翻译开始>
// FileServerEnabled 是静态服务的全局开关。
// 如果设置了任何静态路径，它将自动设置为启用状态。
# <翻译结束>







<原文开始>
// CookieMaxAge specifies the max TTL for cookie items.
<原文结束>

# <翻译开始>
// CookieMaxAge 指定 cookie 项的最大生存时间（TTL）。
# <翻译结束>


<原文开始>
	// CookiePath specifies cookie path.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
// CookiePath 指定cookie路径。
// 同时，它也影响session id的默认存储位置。
# <翻译结束>


<原文开始>
	// CookieDomain specifies cookie domain.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
// CookieDomain 指定 cookie 域名。
// 同时，它也影响 session id 的默认存储方式。
# <翻译结束>


<原文开始>
	// CookieSameSite specifies cookie SameSite property.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
// CookieSameSite 指定 cookie 的 SameSite 属性。
// 同时，它也影响会话 ID 的默认存储方式。
# <翻译结束>


<原文开始>
	// CookieSameSite specifies cookie Secure property.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
// CookieSameSite 指定 cookie 的 Secure 属性。
// 同时，它也影响 session id 的默认存储方式。
# <翻译结束>


<原文开始>
	// CookieSameSite specifies cookie HttpOnly property.
	// It also affects the default storage for session id.
<原文结束>

# <翻译开始>
// CookieSameSite 指定 cookie 的 HttpOnly 属性。
// 同时，它也会影响 session id 的默认存储方式。
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Session.
	// ======================================================================================================
<原文结束>

# <翻译开始>
// ======================================================================================================
// 会话.
// ======================================================================================================
// 这段代码中的注释表明了接下来要定义或描述的内容是关于“Session”（会话）的，但没有给出具体的代码实现细节。在程序中，"Session"通常用于表示用户与服务器之间交互过程的状态信息，用于维持状态、存储临时数据等。
# <翻译结束>


<原文开始>
// SessionIdName specifies the session id name.
<原文结束>

# <翻译开始>
// SessionIdName 指定会话 ID 名称。
# <翻译结束>


<原文开始>
// SessionMaxAge specifies max TTL for session items.
<原文结束>

# <翻译开始>
// SessionMaxAge 指定会话项的最大生存时间（TTL）。
# <翻译结束>


<原文开始>
	// SessionPath specifies the session storage directory path for storing session files.
	// It only makes sense if the session storage is type of file storage.
<原文结束>

# <翻译开始>
// SessionPath 指定用于存储会话文件的会话存储目录路径。
// 只有当会话存储类型为文件存储时，这个配置才有意义。
# <翻译结束>


<原文开始>
// SessionStorage specifies the session storage.
<原文结束>

# <翻译开始>
// SessionStorage 指定会话存储。
# <翻译结束>


<原文开始>
	// SessionCookieMaxAge specifies the cookie ttl for session id.
	// If it is set 0, it means it expires along with browser session.
<原文结束>

# <翻译开始>
// SessionCookieMaxAge 指定会话 ID 的 cookie 存活时间（TTL）。
// 如果设置为 0，表示它将随浏览器会话一同结束时失效。
# <翻译结束>


<原文开始>
// SessionCookieOutput specifies whether automatic outputting session id to cookie.
<原文结束>

# <翻译开始>
// SessionCookieOutput 指定是否自动将会话ID输出到cookie中。
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Logging.
	// ======================================================================================================
<原文结束>

# <翻译开始>
// ======================================================================================================
// 日志记录。
// ======================================================================================================
# <翻译结束>


<原文开始>
// Logger specifies the logger for server.
<原文结束>

# <翻译开始>
// Logger 指定服务器使用的日志记录器。
# <翻译结束>


<原文开始>
// LogPath specifies the directory for storing logging files.
<原文结束>

# <翻译开始>
// LogPath 指定存储日志文件的目录。
# <翻译结束>


<原文开始>
// LogLevel specifies the logging level for logger.
<原文结束>

# <翻译开始>
// LogLevel 指定 logger 的日志记录级别。
# <翻译结束>


<原文开始>
// LogStdout specifies whether printing logging content to stdout.
<原文结束>

# <翻译开始>
// LogStdout 指定是否将日志内容输出到标准输出（stdout）中。
# <翻译结束>


<原文开始>
// ErrorStack specifies whether logging stack information when error.
<原文结束>

# <翻译开始>
// ErrorStack 指定在出现错误时是否记录堆栈信息。
# <翻译结束>


<原文开始>
// ErrorLogEnabled enables error logging content to files.
<原文结束>

# <翻译开始>
// ErrorLogEnabled 开启错误日志功能，将错误内容记录到文件中。
# <翻译结束>


<原文开始>
// ErrorLogPattern specifies the error log file pattern like: error-{Ymd}.log
<原文结束>

# <翻译开始>
// ErrorLogPattern 指定错误日志文件的命名模式，例如：error-{Ymd}.log
# <翻译结束>


<原文开始>
// AccessLogEnabled enables access logging content to files.
<原文结束>

# <翻译开始>
// AccessLogEnabled 开启访问日志功能，将访问内容记录到文件中。
# <翻译结束>


<原文开始>
// AccessLogPattern specifies the error log file pattern like: access-{Ymd}.log
<原文结束>

# <翻译开始>
// AccessLogPattern 指定访问日志文件的命名模式，如：access-{Ymd}.log
# <翻译结束>







<原文开始>
// PProfEnabled enables PProf feature.
<原文结束>

# <翻译开始>
// PProfEnabled 开启PProf功能。
# <翻译结束>


<原文开始>
// PProfPattern specifies the PProf service pattern for router.
<原文结束>

# <翻译开始>
// PProfPattern 指定路由器的 PProf 服务模式。
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// API & Swagger.
	// ======================================================================================================
<原文结束>

# <翻译开始>
// ======================================================================================================
// API 和 Swagger.
// ======================================================================================================
// 这段注释是对代码段的概括性描述，翻译为中文如下：
// ======================================================================================================
// API 接口与 Swagger 文档相关部分.
// ======================================================================================================
// 其中：
// - API：指的是应用程序接口（Application Programming Interface），是程序之间交互的一种方式。
// - Swagger：是一种用于编写和可视化API的工具，现称为OpenAPI Specification，用于定义RESTful API的标准格式。
# <翻译结束>


<原文开始>
// OpenApiPath specifies the OpenApi specification file path.
<原文结束>

# <翻译开始>
// OpenApiPath 指定OpenApi规范文件的路径。
# <翻译结束>


<原文开始>
// SwaggerPath specifies the swagger UI path for route registering.
<原文结束>

# <翻译开始>
// SwaggerPath 指定Swagger UI的路径，用于注册路由。
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Other.
	// ======================================================================================================
<原文结束>

# <翻译开始>
// ======================================================================================================
// 其他
// ======================================================================================================
# <翻译结束>


<原文开始>
	// ClientMaxBodySize specifies the max body size limit in bytes for client request.
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's `8MB` in default.
<原文结束>

# <翻译开始>
// ClientMaxBodySize 指定客户端请求的最大正文大小限制，单位为字节。
// 你可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串进行配置。
// 默认值为 `8MB`。
# <翻译结束>


<原文开始>
	// FormParsingMemory specifies max memory buffer size in bytes which can be used for
	// parsing multimedia form.
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's 1MB in default.
<原文结束>

# <翻译开始>
// FormParsingMemory 指定用于解析多媒体表单的最大内存缓冲区大小（以字节为单位）。
// 可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串进行配置。
// 默认值为 1MB。
# <翻译结束>


<原文开始>
	// NameToUriType specifies the type for converting struct method name to URI when
	// registering routes.
<原文结束>

# <翻译开始>
// NameToUriType 指定了在注册路由时，将结构体方法名转换为URI的类型。
# <翻译结束>


<原文开始>
// RouteOverWrite allows to overwrite the route if duplicated.
<原文结束>

# <翻译开始>
// RouteOverWrite 允许在出现重复路由时进行覆盖。
# <翻译结束>


<原文开始>
// DumpRouterMap specifies whether automatically dumps router map when server starts.
<原文结束>

# <翻译开始>
// DumpRouterMap 指定在服务器启动时是否自动转储路由映射。
# <翻译结束>


<原文开始>
// Graceful enables graceful reload feature for all servers of the process.
<原文结束>

# <翻译开始>
// Graceful 启用进程内所有服务器的优雅重启功能。
# <翻译结束>


<原文开始>
// GracefulTimeout set the maximum survival time (seconds) of the parent process.
<原文结束>

# <翻译开始>
// GracefulTimeout 设置父进程的最大存活时间（秒）。
# <翻译结束>


<原文开始>
// GracefulShutdownTimeout set the maximum survival time (seconds) before stopping the server.
<原文结束>

# <翻译开始>
// GracefulShutdownTimeout 设置在停止服务器之前其最大存活时间（秒）。
# <翻译结束>


<原文开始>
// NewConfig creates and returns a ServerConfig object with default configurations.
// Note that, do not define this default configuration to local package variable, as there are
// some pointer attributes that may be shared in different servers.
<原文结束>

# <翻译开始>
// NewConfig 创建并返回一个具有默认配置的 ServerConfig 对象。
// 注意，不要将此默认配置定义为本地包变量，因为存在一些指针属性，
// 这些属性可能在不同的服务器中被共享。
# <翻译结束>


<原文开始>
// ConfigFromMap creates and returns a ServerConfig object with given map and
// default configuration object.
<原文结束>

# <翻译开始>
// ConfigFromMap 根据给定的映射和默认配置对象创建并返回一个 ServerConfig 对象。
# <翻译结束>


<原文开始>
// SetConfigWithMap sets the configuration for the server using map.
<原文结束>

# <翻译开始>
// SetConfigWithMap 使用map设置服务器的配置
# <翻译结束>


<原文开始>
	// The m now is a shallow copy of m.
	// Any changes to m does not affect the original one.
	// A little tricky, isn't it?
<原文结束>

# <翻译开始>
// 现在的m是m的一个浅拷贝。
// 对m的任何改动都不会影响原始的那个m。
// 这有点小巧妙，不是吗？
# <翻译结束>


<原文开始>
	// Allow setting the size configuration items using string size like:
	// 1m, 100mb, 512kb, etc.
<原文结束>

# <翻译开始>
// 允许使用类似“1m、100mb、512kb”等字符串形式的大小来设置尺寸配置项：
# <翻译结束>


<原文开始>
	// Update the current configuration object.
	// It only updates the configured keys not all the object.
<原文结束>

# <翻译开始>
// 更新当前配置对象。
// 仅更新已配置的键，而非整个对象。
# <翻译结束>


<原文开始>
// SetConfig sets the configuration for the server.
<原文结束>

# <翻译开始>
// SetConfig 设置服务器的配置。
# <翻译结束>


<原文开始>
// Automatically add ':' prefix for address if it is missed.
<原文结束>

# <翻译开始>
// 如果地址中缺少':'前缀，则自动添加。
# <翻译结束>







<原文开始>
// SetAddr sets the listening address for the server.
// The address is like ':80', '0.0.0.0:80', '127.0.0.1:80', '180.18.99.10:80', etc.
<原文结束>

# <翻译开始>
// SetAddr 设置服务器的监听地址。
// 地址格式类似于 ':80'、'0.0.0.0:80'、'127.0.0.1:80'、'180.18.99.10:80' 等。
# <翻译结束>


<原文开始>
// SetPort sets the listening ports for the server.
// The listening ports can be multiple like: SetPort(80, 8080).
<原文结束>

# <翻译开始>
// SetPort 设置服务器监听端口。
// 监听端口可以设置多个，例如：SetPort(80, 8080)。
# <翻译结束>


<原文开始>
// SetHTTPSAddr sets the HTTPS listening ports for the server.
<原文结束>

# <翻译开始>
// SetHTTPSAddr 设置服务器的 HTTPS 监听端口。
# <翻译结束>


<原文开始>
// SetHTTPSPort sets the HTTPS listening ports for the server.
// The listening ports can be multiple like: SetHTTPSPort(443, 500).
<原文结束>

# <翻译开始>
// SetHTTPSPort 设置服务器的 HTTPS 监听端口。
// 可以设置多个监听端口，例如：SetHTTPSPort(443, 500)。
# <翻译结束>


<原文开始>
// SetListener set the custom listener for the server.
<原文结束>

# <翻译开始>
// SetListener 为服务器设置自定义监听器。
# <翻译结束>


<原文开始>
// EnableHTTPS enables HTTPS with given certification and key files for the server.
// The optional parameter `tlsConfig` specifies custom TLS configuration.
<原文结束>

# <翻译开始>
// EnableHTTPS 通过给定的证书和密钥文件为服务器启用HTTPS。
// 可选参数 `tlsConfig` 指定了自定义的TLS配置。
# <翻译结束>


<原文开始>
// SetTLSConfig sets custom TLS configuration and enables HTTPS feature for the server.
<原文结束>

# <翻译开始>
// SetTLSConfig 设置自定义 TLS 配置，并为服务器启用 HTTPS 功能。
# <翻译结束>


<原文开始>
// SetReadTimeout sets the ReadTimeout for the server.
<原文结束>

# <翻译开始>
// SetReadTimeout 设置服务器的读取超时时间。
# <翻译结束>


<原文开始>
// SetWriteTimeout sets the WriteTimeout for the server.
<原文结束>

# <翻译开始>
// SetWriteTimeout 为服务器设置写入超时时间。
# <翻译结束>


<原文开始>
// SetIdleTimeout sets the IdleTimeout for the server.
<原文结束>

# <翻译开始>
// SetIdleTimeout 为服务器设置空闲超时时间。
# <翻译结束>


<原文开始>
// SetMaxHeaderBytes sets the MaxHeaderBytes for the server.
<原文结束>

# <翻译开始>
// SetMaxHeaderBytes 设置服务器的 MaxHeaderBytes 值。
# <翻译结束>


<原文开始>
// SetServerAgent sets the ServerAgent for the server.
<原文结束>

# <翻译开始>
// SetServerAgent 设置服务器的 ServerAgent。
# <翻译结束>


<原文开始>
// SetKeepAlive sets the KeepAlive for the server.
<原文结束>

# <翻译开始>
// SetKeepAlive 设置服务器的 KeepAlive 参数。
# <翻译结束>


<原文开始>
// SetView sets the View for the server.
<原文结束>

# <翻译开始>
// SetView 设置服务器的视图。
# <翻译结束>


<原文开始>
// GetName returns the name of the server.
<原文结束>

# <翻译开始>
// GetName 返回服务器的名称。
# <翻译结束>


<原文开始>
// SetName sets the name for the server.
<原文结束>

# <翻译开始>
// SetName 为服务器设置名称。
# <翻译结束>


<原文开始>
// SetEndpoints sets the Endpoints for the server.
<原文结束>

# <翻译开始>
// SetEndpoints 设置服务器的端点。
# <翻译结束>


<原文开始>
// SetHandler sets the request handler for server.
<原文结束>

# <翻译开始>
// SetHandler 为服务器设置请求处理器。
# <翻译结束>


<原文开始>
// GetHandler returns the request handler of the server.
<原文结束>

# <翻译开始>
// GetHandler 返回服务器的请求处理器。
# <翻译结束>


<原文开始>
// SetRegistrar sets the Registrar for server.
<原文结束>

# <翻译开始>
// SetRegistrar 设置服务器的 Registrar。
# <翻译结束>


<原文开始>
// GetRegistrar returns the Registrar of server.
<原文结束>

# <翻译开始>
// GetRegistrar 返回服务器的注册器。
# <翻译结束>


<原文开始>
// Method names to the URI converting type, which converts name to its lower case and joins the words using char '-'.
<原文结束>

# <翻译开始>
// 转换方法名称到URI的类型，该类型将名称转换为小写并将单词使用字符'-'连接起来。
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// Cookie.
	// ======================================================================================================
<原文结束>

# <翻译开始>
// ======================================================================================================
// Cookie.
// ======================================================================================================
// 以下是翻译后的中文注释：
// ======================================================================================================
// Cookie（cookies）.
// ======================================================================================================
// 此处的注释表明该部分代码与处理或操作Cookie相关的功能有关。"Cookie"在Web开发中指的是服务器发送到用户浏览器并存储在本地的一小段数据，用于识别不同的用户和跟踪会话状态等信息。
# <翻译结束>


<原文开始>
	// ======================================================================================================
	// PProf.
	// ======================================================================================================
<原文结束>

# <翻译开始>
// ======================================================================================================
// PProf.
// ======================================================================================================
// 此处为Golang代码注释的中文翻译：
// ======================================================================================================
// PProf：这是一个对Go程序进行性能分析的接口包，主要用于生成和解析CPU、内存等资源占用情况的采样数据，
// 以便开发者找出程序中的性能瓶颈。
# <翻译结束>


<原文开始>
// Static files root.
<原文结束>

# <翻译开始>
// 静态文件根目录。
# <翻译结束>

