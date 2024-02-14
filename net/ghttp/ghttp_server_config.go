// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/gsvc"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/os/gsession"
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

const (
	defaultHttpAddr  = ":80"  // 默认的HTTP监听端口。
	defaultHttpsAddr = ":443" // 默认的HTTPS监听端口。
	UriTypeDefault   = 0      // 转换方法名称到URI的类型，该类型将名称转换为小写并将单词使用字符'-'连接起来。
	UriTypeFullName  = 1      // MethodNamesToURI 是一个用于将方法名转换为URI的类型，但并不转换回方法名。
	UriTypeAllLower  = 2      // MethodNamesToURI 是一个将方法名转换为 URI 的类型，该类型会将名称转换为其小写形式。
	UriTypeCamel     = 3      // MethodNamesToURI 是一个用于将方法名称转换为 URI 的类型，该类型将名称转换为其驼峰式表示。
)

// ServerConfig 是 HTTP 服务器的配置管理器。
type X服务配置项 struct {
// ======================================================================================================
// 基础部分.
// ======================================================================================================
// 这段注释表明接下来的代码是关于“基础”部分，用于概括或分隔代码的不同模块或功能区块。

	// 服务名称，用于服务注册与发现。
	X服务名称 string `json:"name"`

// Address 指定服务器监听地址，格式如 "端口" 或 ":端口"，
// 多个地址之间使用 ',' 连接。
	X监听地址 string `json:"address"`

	// HTTPSAddr 指定HTTPS服务地址，多个地址之间使用逗号（,）连接。
	HTTPS监听地址 string `json:"httpsAddr"`

	// Listeners 指定自定义监听器。
	X自定义监听器 []net.Listener `json:"listeners"`

	// Endpoints 是服务注册的自定义端点，如果为空则使用 Address。
	Endpoints []string `json:"endpoints"`

	// HTTPSCertPath 指定 HTTPS 服务的证书文件路径。
	HTTPS证书路径 string `json:"httpsCertPath"`

	// HTTPSKeyPath 指定HTTPS服务的密钥文件路径。
	HTTPS密钥路径 string `json:"httpsKeyPath"`

// TLSConfig 提供一个可选的 TLS 配置，用于 ServeTLS 和 ListenAndServeTLS。请注意，
// 此值将被 ServeTLS 和 ListenAndServeTLS 拷贝使用，因此不能通过诸如 tls.Config.SetSessionTicketKeys
// 等方法修改配置。若要使用 SetSessionTicketKeys，请改用 Server.Serve 方法配合一个 TLS 监听器来实现。
	TLS配置 *tls.Config `json:"tlsConfig"`

	// Handler HTTP请求的处理器。
	Handler func(w http.ResponseWriter, r *http.Request) `json:"-"`

// ReadTimeout 是读取整个请求（包括主体）的最大持续时间。
//
// 由于 ReadTimeout 不允许处理程序为每个请求主体设置可接受的截止日期或上传速率，
// 大多数用户可能会更倾向于使用 ReadHeaderTimeout。同时使用它们是有效的。
	X读取超时 time.Duration `json:"readTimeout"`

// WriteTimeout 是在超时前写入响应的最大持续时间。每当读取到新请求的头部时，该时间就会重置。类似于 ReadTimeout，它并不允许 Handlers 根据每个请求自行决定是否超时。
	X写入超时 time.Duration `json:"writeTimeout"`

// IdleTimeout 是在启用 keep-alive 时，等待下一个请求的最大时间间隔。如果 IdleTimeout 设为零，则使用 ReadTimeout 的值。如果两者都为零，则表示没有超时限制。
	X长连接超时 time.Duration `json:"idleTimeout"`

// MaxHeaderBytes 控制服务器在解析请求头（包括请求行）的键和值时，
// 会读取的最大字节数。但请注意，它并不会限制请求体的大小。
//
// 你可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串来配置这个参数。
// 默认情况下，其值为 10240 字节。
	X最大协议头长度 int `json:"maxHeaderBytes"`

	// KeepAlive 启用 HTTP 保持连接（Keep-alive）功能。
	X启用长连接 bool `json:"keepAlive"`

// ServerAgent 指定服务器代理信息，该信息会被写入
// HTTP 响应头中作为 "Server"。
	X服务器代理 string `json:"serverAgent"`

	// View 指定了服务器的默认模板视图对象。
	X模板默认 *模板类.View `json:"view"`

// ======================================================================================================
// 静态部分。
// ======================================================================================================
// 这段注释表明了该处代码是“静态”部分，但具体含义可能需要更多上下文信息来准确翻译。在程序中，“Static”通常是指不变的、非运行时动态改变的数据或函数，或者是初始化后在整个程序生命周期中保持固定的部分。

	// Rewrites 指定了 URI 重写规则映射。
	X路由URI重写规则Map map[string]string `json:"rewrites"`

	// IndexFiles 指定静态文件夹的索引文件。
	X静态文件索引 []string `json:"indexFiles"`

// IndexFolder 指定在请求文件夹时是否列出子文件。
// 如果该值为false，服务器将返回HTTP状态码403。
	X静态文件是否列出子文件 bool `json:"indexFolder"`

	// ServerRoot 指定静态服务的根目录。
	X静态文件根目录 string `json:"serverRoot"`

	// SearchPaths 指定静态服务的额外搜索目录。
	X静态文件额外搜索目录 []string `json:"searchPaths"`

	// StaticPaths 指定了URI到目录映射的数组。
	X静态文件目录映射 []静态文件配置项 `json:"staticPaths"`

// FileServerEnabled 是静态服务的全局开关。
// 如果设置了任何静态路径，它将自动设置为启用状态。
	X静态文件是否开启 bool `json:"fileServerEnabled"`

// ======================================================================================================
// Cookie.
// ======================================================================================================
// 以下是翻译后的中文注释：
// ======================================================================================================
// Cookie（cookies）.
// ======================================================================================================
// 此处的注释表明该部分代码与处理或操作Cookie相关的功能有关。"Cookie"在Web开发中指的是服务器发送到用户浏览器并存储在本地的一小段数据，用于识别不同的用户和跟踪会话状态等信息。

	// CookieMaxAge 指定 cookie 项的最大生存时间（TTL）。
	Cookie最大存活时长 time.Duration `json:"cookieMaxAge"`

// CookiePath 指定cookie路径。
// 同时，它也影响session id的默认存储位置。
	Cookie路径 string `json:"cookiePath"`

// CookieDomain 指定 cookie 域名。
// 同时，它也影响 session id 的默认存储方式。
	Cookie域名 string `json:"cookieDomain"`

// CookieSameSite 指定 cookie 的 SameSite 属性。
// 同时，它也影响会话 ID 的默认存储方式。
	CookieSameSite string `json:"cookieSameSite"`

// CookieSameSite 指定 cookie 的 Secure 属性。
// 同时，它也影响 session id 的默认存储方式。
	Cookie安全 bool `json:"cookieSecure"`

// CookieSameSite 指定 cookie 的 HttpOnly 属性。
// 同时，它也会影响 session id 的默认存储方式。
	Cookie跨站访问控制 bool `json:"cookieHttpOnly"`

// ======================================================================================================
// 会话.
// ======================================================================================================
// 这段代码中的注释表明了接下来要定义或描述的内容是关于“Session”（会话）的，但没有给出具体的代码实现细节。在程序中，"Session"通常用于表示用户与服务器之间交互过程的状态信息，用于维持状态、存储临时数据等。

	// SessionIdName 指定会话 ID 名称。
	SessionID名称 string `json:"sessionIdName"`

	// SessionMaxAge 指定会话项的最大生存时间（TTL）。
	Session最大存活时长 time.Duration `json:"sessionMaxAge"`

// SessionPath 指定用于存储会话文件的会话存储目录路径。
// 只有当会话存储类型为文件存储时，这个配置才有意义。
	Session存储目录路径 string `json:"sessionPath"`

	// SessionStorage 指定会话存储。
	Session存储 session类.Storage `json:"sessionStorage"`

// SessionCookieMaxAge 指定会话 ID 的 cookie 存活时间（TTL）。
// 如果设置为 0，表示它将随浏览器会话一同结束时失效。
	SessionCookie存活时长 time.Duration `json:"sessionCookieMaxAge"`

	// SessionCookieOutput 指定是否自动将会话ID输出到cookie中。
	SessionID输出到Cookie bool `json:"sessionCookieOutput"`

// ======================================================================================================
// 日志记录。
// ======================================================================================================

	X日志记录器           *日志类.Logger `json:"logger"`           // Logger 指定服务器使用的日志记录器。
	X日志存储目录          string       `json:"logPath"`          // LogPath 指定存储日志文件的目录。
	X日志记录等级         string       `json:"logLevel"`         // LogLevel 指定 logger 的日志记录级别。
	X日志开启输出到CMD        bool         `json:"logStdout"`        // LogStdout 指定是否将日志内容输出到标准输出（stdout）中。
	X日志开启错误堆栈记录       bool         `json:"errorStack"`       // ErrorStack 指定在出现错误时是否记录堆栈信息。
	X日志开启错误记录  bool         `json:"errorLogEnabled"`  // ErrorLogEnabled 开启错误日志功能，将错误内容记录到文件中。
	X日志错误文件命名模式  string       `json:"errorLogPattern"`  // ErrorLogPattern 指定错误日志文件的命名模式，例如：error-{Ymd}.log
	X日志开启访客记录 bool         `json:"accessLogEnabled"` // AccessLogEnabled 开启访问日志功能，将访问内容记录到文件中。
	X日志访客文件命名模式 string       `json:"accessLogPattern"` // AccessLogPattern 指定访问日志文件的命名模式，如：access-{Ymd}.log

// ======================================================================================================
// PProf.
// ======================================================================================================
// 此处为Golang代码注释的中文翻译：
// ======================================================================================================
// PProf：这是一个对Go程序进行性能分析的接口包，主要用于生成和解析CPU、内存等资源占用情况的采样数据，
// 以便开发者找出程序中的性能瓶颈。

	PProf开启 bool   `json:"pprofEnabled"` // PProfEnabled 开启PProf功能。
	PProf模式 string `json:"pprofPattern"` // PProfPattern 指定路由器的 PProf 服务模式。

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

	APIOpenApiUI路径 string `json:"openapiPath"` // OpenApiPath 指定OpenApi规范文件的路径。
	APISwaggerUI路径 string `json:"swaggerPath"` // SwaggerPath 指定Swagger UI的路径，用于注册路由。

// ======================================================================================================
// 其他
// ======================================================================================================

// ClientMaxBodySize 指定客户端请求的最大正文大小限制，单位为字节。
// 你可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串进行配置。
// 默认值为 `8MB`。
	X客户端请求最大长度 int64 `json:"clientMaxBodySize"`

// FormParsingMemory 指定用于解析多媒体表单的最大内存缓冲区大小（以字节为单位）。
// 可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串进行配置。
// 默认值为 1MB。
	X表单解析最大缓冲区长度 int64 `json:"formParsingMemory"`

// NameToUriType 指定了在注册路由时，将结构体方法名转换为URI的类型。
	NameToUriType int `json:"nameToUriType"`

	// RouteOverWrite 允许在出现重复路由时进行覆盖。
	X路由允许覆盖 bool `json:"routeOverWrite"`

	// DumpRouterMap 指定在服务器启动时是否自动转储路由映射。
	DumpRouterMap bool `json:"dumpRouterMap"`

	// Graceful 启用进程内所有服务器的优雅重启功能。
	X平滑重启开启 bool `json:"graceful"`

	// GracefulTimeout 设置父进程的最大存活时间（秒）。
	GracefulTimeout uint8 `json:"gracefulTimeout"`

	// GracefulShutdownTimeout 设置在停止服务器之前其最大存活时间（秒）。
	GracefulShutdownTimeout uint8 `json:"gracefulShutdownTimeout"`
}

// NewConfig 创建并返回一个具有默认配置的 ServerConfig 对象。
// 注意，不要将此默认配置定义为本地包变量，因为存在一些指针属性，
// 这些属性可能在不同的服务器中被共享。
func X创建默认配置项() X服务配置项 {
	return X服务配置项{
		X服务名称:                    DefaultServerName,
		X监听地址:                 ":0",
		HTTPS监听地址:               "",
		X自定义监听器:               nil,
		Handler:                 nil,
		X读取超时:             60 * time.Second,
		X写入超时:            0, // No timeout.
		X长连接超时:             60 * time.Second,
		X最大协议头长度:          10240, // 10KB
		X启用长连接:               true,
		X静态文件索引:              []string{"index.html", "index.htm"},
		X静态文件是否列出子文件:             false,
		X服务器代理:             "GoFrame HTTP Server",
		X静态文件根目录:              "",
		X静态文件目录映射:             make([]静态文件配置项, 0),
		X静态文件是否开启:       false,
		Cookie最大存活时长:            time.Hour * 24 * 365,
		Cookie路径:              "/",
		Cookie域名:            "",
		SessionID名称:           "gfsessionid",
		Session存储目录路径:             session类.DefaultStorageFilePath,
		Session最大存活时长:           time.Hour * 24,
		SessionID输出到Cookie:     true,
		SessionCookie存活时长:     time.Hour * 24,
		X日志记录器:                  日志类.X创建(),
		X日志记录等级:                "all",
		X日志开启输出到CMD:               true,
		X日志开启错误堆栈记录:              true,
		X日志开启错误记录:         true,
		X日志错误文件命名模式:         "error-{Ymd}.log",
		X日志开启访客记录:        false,
		X日志访客文件命名模式:        "access-{Ymd}.log",
		DumpRouterMap:           true,
		X客户端请求最大长度:       8 * 1024 * 1024, // 8MB
		X表单解析最大缓冲区长度:       1024 * 1024,     // 1MB
		X路由URI重写规则Map:                make(map[string]string),
		X平滑重启开启:                false,
		GracefulTimeout:         2, // seconds
		GracefulShutdownTimeout: 5, // seconds
	}
}

// ConfigFromMap 根据给定的映射和默认配置对象创建并返回一个 ServerConfig 对象。
func X创建配置对象Map(配置 map[string]interface{}) (X服务配置项, error) {
	config := X创建默认配置项()
	if err := 转换类.Struct(配置, &config); err != nil {
		return config, err
	}
	return config, nil
}

// SetConfigWithMap 使用map设置服务器的配置
func (s *X服务) X设置配置项Map(配置 map[string]interface{}) error {
// 现在的m是m的一个浅拷贝。
// 对m的任何改动都不会影响原始的那个m。
// 这有点小巧妙，不是吗？
	配置 = 工具类.MapCopy(配置)
// 允许使用类似“1m、100mb、512kb”等字符串形式的大小来设置尺寸配置项：
	if k, v := 工具类.MapPossibleItemByKey(配置, "MaxHeaderBytes"); k != "" {
		配置[k] = 文件类.X易读格式转字节长度(转换类.String(v))
	}
	if k, v := 工具类.MapPossibleItemByKey(配置, "ClientMaxBodySize"); k != "" {
		配置[k] = 文件类.X易读格式转字节长度(转换类.String(v))
	}
	if k, v := 工具类.MapPossibleItemByKey(配置, "FormParsingMemory"); k != "" {
		配置[k] = 文件类.X易读格式转字节长度(转换类.String(v))
	}
// 更新当前配置对象。
// 仅更新已配置的键，而非整个对象。
	if err := 转换类.Struct(配置, &s.config); err != nil {
		return err
	}
	return s.X设置配置项(s.config)
}

// SetConfig 设置服务器的配置。
func (s *X服务) X设置配置项(c X服务配置项) error {
	s.config = c
	// 如果地址中缺少':'前缀，则自动添加。
	if s.config.X监听地址 != "" && !文本类.X是否包含(s.config.X监听地址, ":") {
		s.config.X监听地址 = ":" + s.config.X监听地址
	}
	// 静态文件根目录。
	if c.X静态文件根目录 != "" {
		s.X设置静态文件根目录(c.X静态文件根目录)
	}
	if len(c.X静态文件额外搜索目录) > 0 {
		paths := c.X静态文件额外搜索目录
		c.X静态文件额外搜索目录 = []string{}
		for _, v := range paths {
			s.X静态文件添加额外搜索目录(v)
		}
	}
	// HTTPS.
	if c.TLS配置 == nil && c.HTTPS证书路径 != "" {
		s.X启用HTTPS(c.HTTPS证书路径, c.HTTPS密钥路径)
	}
	// Logging.
	if s.config.X日志存储目录 != "" && s.config.X日志存储目录 != s.config.X日志记录器.X取文件路径() {
		if err := s.config.X日志记录器.X设置文件路径(s.config.X日志存储目录); err != nil {
			return err
		}
	}
	if err := s.config.X日志记录器.X设置文本级别(s.config.X日志记录等级); err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	gracefulEnabled = c.X平滑重启开启
	intlog.Printf(context.TODO(), "SetConfig: %+v", s.config)
	return nil
}

// SetAddr 设置服务器的监听地址。
// 地址格式类似于 ':80'、'0.0.0.0:80'、'127.0.0.1:80'、'180.18.99.10:80' 等。
func (s *X服务) X设置监听地址(地址 string) {
	s.config.X监听地址 = 地址
}

// SetPort 设置服务器监听端口。
// 监听端口可以设置多个，例如：SetPort(80, 8080)。
func (s *X服务) X设置监听端口(端口 ...int) {
	if len(端口) > 0 {
		s.config.X监听地址 = ""
		for _, v := range 端口 {
			if len(s.config.X监听地址) > 0 {
				s.config.X监听地址 += ","
			}
			s.config.X监听地址 += ":" + strconv.Itoa(v)
		}
	}
}

// SetHTTPSAddr 设置服务器的 HTTPS 监听端口。
func (s *X服务) X设置HTTPS监听地址(地址 string) {
	s.config.HTTPS监听地址 = 地址
}

// SetHTTPSPort 设置服务器的 HTTPS 监听端口。
// 可以设置多个监听端口，例如：SetHTTPSPort(443, 500)。
func (s *X服务) X设置HTTPS监听端口(端口 ...int) {
	if len(端口) > 0 {
		s.config.HTTPS监听地址 = ""
		for _, v := range 端口 {
			if len(s.config.HTTPS监听地址) > 0 {
				s.config.HTTPS监听地址 += ","
			}
			s.config.HTTPS监听地址 += ":" + strconv.Itoa(v)
		}
	}
}

// SetListener 为服务器设置自定义监听器。
func (s *X服务) X设置自定义监听器(监听器 ...net.Listener) error {
	if 监听器 == nil {
		return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "SetListener failed: listener can not be nil")
	}
	if len(监听器) > 0 {
		ports := make([]string, len(监听器))
		for k, v := range 监听器 {
			if v == nil {
				return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "SetListener failed: listener can not be nil")
			}
			ports[k] = fmt.Sprintf(":%d", (v.Addr().(*net.TCPAddr)).Port)
		}
		s.config.X监听地址 = strings.Join(ports, ",")
		s.config.X自定义监听器 = 监听器
	}
	return nil
}

// EnableHTTPS 通过给定的证书和密钥文件为服务器启用HTTPS。
// 可选参数 `tlsConfig` 指定了自定义的TLS配置。
func (s *X服务) X启用HTTPS(证书路径, 密钥路径 string, tls配置 ...*tls.Config) {
	var ctx = context.TODO()
	certFileRealPath := 文件类.X取绝对路径且效验(证书路径)
	if certFileRealPath == "" {
		certFileRealPath = 文件类.X取绝对路径且效验(文件类.X取当前工作目录() + 文件类.Separator + 证书路径)
		if certFileRealPath == "" {
			certFileRealPath = 文件类.X取绝对路径且效验(文件类.X取main路径() + 文件类.Separator + 证书路径)
		}
	}
	// Resource.
	if certFileRealPath == "" && 资源类.Contains(证书路径) {
		certFileRealPath = 证书路径
	}
	if certFileRealPath == "" {
		s.Logger别名().X输出并格式化FATA(ctx, `EnableHTTPS failed: certFile "%s" does not exist`, 证书路径)
	}
	keyFileRealPath := 文件类.X取绝对路径且效验(密钥路径)
	if keyFileRealPath == "" {
		keyFileRealPath = 文件类.X取绝对路径且效验(文件类.X取当前工作目录() + 文件类.Separator + 密钥路径)
		if keyFileRealPath == "" {
			keyFileRealPath = 文件类.X取绝对路径且效验(文件类.X取main路径() + 文件类.Separator + 密钥路径)
		}
	}
	// Resource.
	if keyFileRealPath == "" && 资源类.Contains(密钥路径) {
		keyFileRealPath = 密钥路径
	}
	if keyFileRealPath == "" {
		s.Logger别名().X输出FATA(ctx, `EnableHTTPS failed: keyFile "%s" does not exist`, 密钥路径)
	}
	s.config.HTTPS证书路径 = certFileRealPath
	s.config.HTTPS密钥路径 = keyFileRealPath
	if len(tls配置) > 0 {
		s.config.TLS配置 = tls配置[0]
	}
}

// SetTLSConfig 设置自定义 TLS 配置，并为服务器启用 HTTPS 功能。
func (s *X服务) X设置TLS配置(tls配置 *tls.Config) {
	s.config.TLS配置 = tls配置
}

// SetReadTimeout 设置服务器的读取超时时间。
func (s *X服务) X设置读取超时(时长 time.Duration) {
	s.config.X读取超时 = 时长
}

// SetWriteTimeout 为服务器设置写入超时时间。
func (s *X服务) X设置写入超时(时长 time.Duration) {
	s.config.X写入超时 = 时长
}

// SetIdleTimeout 为服务器设置空闲超时时间。
func (s *X服务) X设置长连接超时(时长 time.Duration) {
	s.config.X长连接超时 = 时长
}

// SetMaxHeaderBytes 设置服务器的 MaxHeaderBytes 值。
func (s *X服务) X设置协议头最大长度(最大长度 int) {
	s.config.X最大协议头长度 = 最大长度
}

// SetServerAgent 设置服务器的 ServerAgent。
func (s *X服务) X设置服务器代理标识(代理标识 string) {
	s.config.X服务器代理 = 代理标识
}

// SetKeepAlive 设置服务器的 KeepAlive 参数。
func (s *X服务) X设置开启长连接(开启 bool) {
	s.config.X启用长连接 = 开启
}

// SetView 设置服务器的视图。
func (s *X服务) X设置默认模板对象(模板对象 *模板类.View) {
	s.config.X模板默认 = 模板对象
}

// GetName 返回服务器的名称。
func (s *X服务) X取服务名称() string {
	return s.config.X服务名称
}

// SetName 为服务器设置名称。
func (s *X服务) X设置服务名称(名称 string) {
	s.config.X服务名称 = 名称
}

// SetEndpoints 设置服务器的端点。
func (s *X服务) SetEndpoints(endpoints []string) {
	s.config.Endpoints = endpoints
}

// SetHandler 为服务器设置请求处理器。
func (s *X服务) X设置请求处理器(h func(w http.ResponseWriter, r *http.Request)) {
	s.config.Handler = h
}

// GetHandler 返回服务器的请求处理器。
func (s *X服务) X取请求处理器() func(w http.ResponseWriter, r *http.Request) {
	if s.config.Handler == nil {
		return s.ServeHTTP
	}
	return s.config.Handler
}

// SetRegistrar 设置服务器的 Registrar。
func (s *X服务) X设置注册发现对象(注册发现对象 gsvc.Registrar) {
	s.registrar = 注册发现对象
}

// GetRegistrar 返回服务器的注册器。
func (s *X服务) X取注册发现对象() gsvc.Registrar {
	return s.registrar
}
