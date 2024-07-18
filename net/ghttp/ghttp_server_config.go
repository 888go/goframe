// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

const (
	defaultHttpAddr  = ":80"  // HTTP的默认监听端口。 md5:d368f1a97ab395c2
	defaultHttpsAddr = ":443" // HTTPS的默认监听端口。 md5:72d4583ff29fdc86

)

const (
	UriTypeDefault  = iota // 将方法名称转换为URI的类型，该类型将名称转换为小写，并使用字符'-'连接单词。 md5:1fc3cb97052e806f
	UriTypeFullName        // 将方法名称转换为URI转换类型，但不将其转换为方法名称。 md5:75ab2b02cd49ae6b
	UriTypeAllLower        // 转换方法名称为URI的类型，将名称转换为小写。 md5:ed7460ac07fbb91a
	UriTypeCamel           // 将方法名称转换为URI的类型，该类型将名称转换为驼峰式命名。 md5:2e028fc00d70d9bf
)

// ServerConfig 是HTTP服务器的配置管理器。 md5:a2c6c214e9d64d54
type ServerConfig struct {
// ======================================================================================================
// 基础内容.
// ======================================================================================================
// md5:accd23363a592bfa

	// 服务名称，用于服务注册和发现。 md5:a0d782209905145d
	Name string `json:"name"`//qm:服务名称  cz:Name string `json:"name"`  

// Address 指定服务器监听的地址，如 "port" 或 ":port"，多个地址使用 "," 分隔。
// md5:ae0da0e5cf0a3e61
	Address string `json:"address"`//qm:监听地址  cz:Address string `json:"address"`  

	// HTTPSAddr 指定了HTTPS地址，多个地址使用逗号','连接。 md5:c776678a9eac5a90
	HTTPSAddr string `json:"httpsAddr"`//qm:HTTPS监听地址  cz:HTTPSAddr string `json:"httpsAddr"`  

	// Listeners 指定了自定义的监听器。 md5:3b866f27e7903dac
	Listeners []net.Listener `json:"listeners"`//qm:自定义监听器  cz:Listeners []net.Listener  

	// 如果为空，Endpoints 是服务注册的自定义端点，它使用 Address。 md5:57f627d789f1ea89
	Endpoints []string `json:"endpoints"`

	// HTTPSCertPath 指定HTTPS服务的证书文件路径。 md5:3613bb98322987e7
	HTTPSCertPath string `json:"httpsCertPath"`//qm:HTTPS证书路径  cz:HTTPSCertPath string  

	// HTTPSKeyPath 指定了HTTPS服务的密钥文件路径。 md5:2a9c1b2c382d01b8
	HTTPSKeyPath string `json:"httpsKeyPath"`//qm:HTTPS密钥路径  cz:HTTPSKeyPath string  

// TLSConfig 可选地提供一个用于 ServeTLS 和 ListenAndServeTLS 的 TLS 配置。请注意，此值会被 ServeTLS 和 ListenAndServeTLS 克隆，
// 因此无法通过如 tls.Config.SetSessionTicketKeys 这样的方法修改配置。若要使用 SetSessionTicketKeys，应使用带有 TLS 监听器的 Server.Serve。
// md5:16a9af3e2eb3eabc
	TLSConfig *tls.Config `json:"tlsConfig"`//qm:TLS配置  cz:TLSConfig *tls  

	// 处理HTTP请求的处理器。 md5:bd278835c47f74d4
	Handler func(w http.ResponseWriter, r *http.Request) `json:"-"`

// ReadTimeout是读取整个请求（包括正文）的最大持续时间。
// 
// 由于ReadTimeout不允許Handler根据每个请求体的可接受截止日期或上传速率做出逐个请求的决定，大多数用户更喜欢使用ReadHeaderTimeout。同时使用它们两者也是可以的。
// md5:45add6b4a3777e9a
	ReadTimeout time.Duration `json:"readTimeout"`//qm:读取超时  cz:ReadTimeout time.Duration  

// WriteTimeout 是响应写入操作超时的最长时间。每当读取新请求的头信息时，它都会重置。与 ReadTimeout 相似，它不允许处理程序针对每个请求做出超时时间的决定。
// md5:7cba2c215a8e6f3d
	WriteTimeout time.Duration `json:"writeTimeout"`//qm:写入超时  cz:WriteTimeout time.Duration  

// IdleTimeout 是在保持连接开启状态下，等待下一个请求的最大时间。如果IdleTimeout为零，则使用ReadTimeout的值。如果两者都为零，则没有超时设置。
// md5:340816400bd04176
	IdleTimeout time.Duration `json:"idleTimeout"`//qm:长连接超时  cz:IdleTimeout time.Duration  

// MaxHeaderBytes 控制服务器解析请求头的键值对（包括请求行）时的最大字节数。它不限制请求体的大小。
// 
// 可以在配置文件中通过字符串进行配置，例如：1m、10m、500kb等。默认值为 10240 字节。
// md5:8a61cc04b245e3d0
	MaxHeaderBytes int `json:"maxHeaderBytes"`//qm:最大协议头长度  cz:MaxHeaderBytes int  

	// EnableHTTPKeepAlive 启用 HTTP 保持活动。 md5:316ccd3dea9e6e4e
	KeepAlive bool `json:"keepAlive"`//qm:启用长连接  cz:KeepAlive bool  

// ServerAgent 指定了服务器代理信息，该信息将被写入HTTP响应头中作为 "Server"。
// md5:08ba94b701f34cac
	ServerAgent string `json:"serverAgent"`//qm:服务器代理  cz:ServerAgent string  

	// View 指定了服务器的默认模板视图对象。 md5:1d7f3415e9af9116
	View *gview.View `json:"view"`//qm:模板默认  cz:View *  

// ======================================================================================================
// 静态部分。
// ======================================================================================================
// md5:67b32db36648f355

	// Rewrites 指定了URI重写规则映射。 md5:177b674dcc27b406
	Rewrites map[string]string `json:"rewrites"`//qm:路由URI重写规则Map  cz:Rewrites map  

	// IndexFiles 指定了静态文件夹的索引文件。 md5:2beaa7ae62753536
	IndexFiles []string `json:"indexFiles"`//qm:静态文件索引  cz:IndexFiles []string  

// IndexFolder 指定是否在请求文件夹时列出子文件。
// 如果为 false，服务器将返回 HTTP 状态码 403。
// md5:2d0bfed1b08c9bae
	IndexFolder bool `json:"indexFolder"`//qm:静态文件是否列出子文件  cz:IndexFolder bool  

	// ServerRoot 指定静态服务的根目录。 md5:19d6dd74ec61bd29
	ServerRoot string `json:"serverRoot"`//qm:静态文件根目录  cz:ServerRoot string  

	// SearchPaths 指定静态服务的额外搜索目录。 md5:bd172212e30738c5
	SearchPaths []string `json:"searchPaths"`//qm:静态文件额外搜索目录  cz:SearchPaths []string  

	// StaticPaths 指定了URI到目录的映射数组。 md5:ea14f728a07fbd14
	StaticPaths []staticPathItem `json:"staticPaths"`//qm:静态文件目录映射  cz:StaticPaths []  

// FileServerEnabled 是静态服务的全局开关。
// 如果设置了任何静态路径，它将自动设置为启用。
// md5:c3edec83e9cfd01f
	FileServerEnabled bool `json:"fileServerEnabled"`//qm:静态文件是否开启  cz:FileServerEnabled bool  

// ======================================================================================================
//  cookie。
// ======================================================================================================
// md5:82bbf8b719e4196b

	// CookieMaxAge 指定 cookie 项的最大过期时间（TTL）。 md5:c6712f43beedeb7d
	CookieMaxAge time.Duration `json:"cookieMaxAge"`//qm:Cookie最大存活时长  cz:CookieMaxAge time.Duration  

// CookiePath 指定了Cookie路径。
// 它也影响了会话ID的默认存储位置。
// md5:dad6b405aee54d41
	CookiePath string `json:"cookiePath"`//qm:Cookie路径  cz:CookiePath string  

// CookieDomain 指定Cookie的域名。
// 它也会影响默认的session id存储位置。
// md5:f2d433c3779a94df
	CookieDomain string `json:"cookieDomain"`//qm:Cookie域名  cz:CookieDomain string  

// CookieSameSite 指定 cookie 的 SameSite 属性。它还影响会话 ID 的默认存储。
// md5:bba3b21b29719a23
	CookieSameSite string `json:"cookieSameSite"`

// CookieSameSite 指定 cookie 的 Secure 属性。
// 它还会影响会话 ID 的默认存储。
// md5:e861b293a54a4909
	CookieSecure bool `json:"cookieSecure"`//qm:Cookie安全  cz:CookieSecure bool  

// CookieSameSite 指定了cookie的HttpOnly属性。
// 它也影响会话ID的默认存储方式。
// md5:ffc9065ce43afdd9
	CookieHttpOnly bool `json:"cookieHttpOnly"`//qm:Cookie跨站访问控制  cz:CookieHttpOnly bool  

// ======================================================================================================
// 会话管理。
// ======================================================================================================
// md5:bf2dda055c30c648

	// SessionIdName 指定会话ID的名称。 md5:18c5a80d34f75272
	SessionIdName string `json:"sessionIdName"`//qm:SessionID名称  cz:SessionIdName string  

	// SessionMaxAge 指定会话项的最大超时时间（TTL）。 md5:cc78683f7c70c955
	SessionMaxAge time.Duration `json:"sessionMaxAge"`//qm:Session最大存活时长  cz:SessionMaxAge time.Duration  

// SessionPath 指定了用于存储会话文件的会话存储目录路径。
// 仅当会话存储类型为文件存储时，此设置才有意义。
// md5:ace2dbc78e7f3a04
	SessionPath string `json:"sessionPath"`//qm:Session存储目录路径  cz:SessionPath string  

	// SessionStorage 指定会话存储。 md5:678a55a9e339a25c
	SessionStorage gsession.Storage `json:"sessionStorage"`

// SessionCookieMaxAge 指定会话ID的cookie过期时间。如果设置为0，表示它将随着浏览器会话一起过期。
// md5:7b48b403d924198e
	SessionCookieMaxAge time.Duration `json:"sessionCookieMaxAge"`//qm:SessionCookie存活时长  cz:SessionCookieMaxAge time.Duration  

	// SessionCookieOutput 指定是否自动将会话ID输出到cookie。 md5:040824f71a38e446
	SessionCookieOutput bool `json:"sessionCookieOutput"`//qm:SessionID输出到Cookie  cz:SessionCookieOutput bool  

// ===============================================================================================
// 日志记录。
// ===============================================================================================
// md5:ede34792b995e698

	Logger           *glog.Logger `json:"logger"`//qm:日志记录器  cz:Logger *             // Logger 指定服务器使用的日志记录器。 md5:004b9e605f068eec
	LogPath          string       `json:"logPath"`//qm:日志存储目录  cz:LogPath string            // LogPath 指定日志文件的存储目录。 md5:de94cd356ae22e8a
	LogLevel         string       `json:"logLevel"`//qm:日志记录等级  cz:LogLevel string           // LogLevel 指定了logger的日志级别。 md5:87d747e517ace64c
	LogStdout        bool         `json:"logStdout"`//qm:日志开启输出到CMD  cz:LogStdout bool          // LogStdout 指定是否将日志内容打印到标准输出（stdout）。 md5:f591098f0447f3f8
	ErrorStack       bool         `json:"errorStack"`//qm:日志开启错误堆栈记录  cz:ErrorStack bool         // ErrorStack 指定在发生错误时是否记录堆栈信息。 md5:5cfe84f341788adc
	ErrorLogEnabled  bool         `json:"errorLogEnabled"`//qm:日志开启错误记录  cz:ErrorLogEnabled bool    // ErrorLogEnabled 启用将错误日志内容写入文件。 md5:9065ef46c6d983d0
	ErrorLogPattern  string       `json:"errorLogPattern"`// ErrorLogPattern 指定错误日志文件模式，如：error-YYYYMMDD.log. md5:c59be38d6eeea7aa
	AccessLogEnabled bool         `json:"accessLogEnabled"`//qm:日志开启访客记录  cz:AccessLogEnabled bool   // AccessLogEnabled 启用将访问日志内容记录到文件中。 md5:6867f80f6ec7fb95
	AccessLogPattern string       `json:"accessLogPattern"`// AccessLogPattern 指定错误日志文件的模式，如：access-{Ymd}.log. md5:e01474de0152ebf6

// ======================================================================================================
// PProf。
// ======================================================================================================
// 
// 这段Go代码的注释翻译成中文是：
// 
// ">// ======================================================================================================
// ">// PProf：这是一个用于性能分析和 profiling 的工具，通常在Go语言中用于诊断和优化程序性能。
// >// ======================================================================================================
// md5:94f131ef860cf923

	PProfEnabled bool   `json:"pprofEnabled"`//qm:PProf开启  cz:PProfEnabled bool   // PProfEnabled 启用 PProf 功能。 md5:4847ea23da60be23
	PProfPattern string `json:"pprofPattern"`//qm:PProf模式  cz:PProfPattern string   // PProfPattern 为路由器指定 PProf 服务的模式。 md5:ddf66608babb16a1

// =======================================================================================
// API与Swagger相关代码。
// =======================================================================================
// md5:ec4583c0e3c0aab6

	OpenApiPath       string `json:"openapiPath"`//qm:APIOpenApiUI路径  cz:OpenApiPath string         // OpenApiPath 指定OpenApi规范文件的路径。 md5:a99446ffbab82145
	SwaggerPath       string `json:"swaggerPath"`//qm:APISwaggerUI路径  cz:SwaggerPath string         // SwaggerPath 定义路由注册的Swagger UI路径。 md5:abc9988ac6d860c6
	SwaggerUITemplate string `json:"swaggerUITemplate"` // SwaggerUITemplate 指定 Swagger UI 的自定义模板. md5:0f381e185ab07c43

// ======================================================================================================
// 其他。
// ======================================================================================================
// md5:8a9c4a8ec79cdc30

// ClientMaxBodySize 指定了客户端请求的最大体大小限制（以字节为单位）。它可以在配置文件中通过字符串进行配置，例如：1m、10m、500kb 等。默认值为 `8MB`。
// md5:2ae357b9e73e0ba6
	ClientMaxBodySize int64 `json:"clientMaxBodySize"`//qm:客户端请求最大长度  cz:ClientMaxBodySize int64  

// FormParsingMemory 定义了用于解析多媒体表单的最大内存缓冲区大小（以字节为单位）。可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串进行配置。默认值为 1MB。
// md5:e7808b8ee0d32ae1
	FormParsingMemory int64 `json:"formParsingMemory"`//qm:表单解析最大缓冲区长度  cz:FormParsingMemory int64  

// NameToUriType 用于指定在注册路由时，将结构体方法名转换为URI的类型。
// md5:3853020c5284d13d
	NameToUriType int `json:"nameToUriType"`

	// RouteOverWrite 允许覆盖重复的路由。 md5:11e5811ec1ba25ca
	RouteOverWrite bool `json:"routeOverWrite"`//qm:路由允许覆盖  cz:RouteOverWrite bool  

	// DumpRouterMap 指定服务器启动时是否自动dump路由器映射。 md5:5c37a6000e9858ab
	DumpRouterMap bool `json:"dumpRouterMap"`

	// Graceful启用进程所有服务器的优雅重启功能。 md5:e4f67dc7d507232e
	Graceful bool `json:"graceful"`//qm:平滑重启开启  cz:Graceful bool  

	// GracefulTimeout 设置父进程的最大生存时间（秒）。 md5:09d4293175059ede
	GracefulTimeout uint8 `json:"gracefulTimeout"`

	// GracefulShutdownTimeout 设置在停止服务器之前，允许服务器最大存活的时间（秒数）。 md5:b220917a3a4e4ebf
	GracefulShutdownTimeout uint8 `json:"gracefulShutdownTimeout"`
}

// NewConfig 创建并返回一个带有默认配置的ServerConfig对象。
// 注意，不要将这些默认配置定义为本地包变量，因为存在一些指针属性可能在不同的服务器中被共享。
// md5:3a8bce955120579e
// ff:创建默认配置项
func NewConfig() ServerConfig {
	return ServerConfig{
		Name:                    DefaultServerName,
		Address:                 ":0",
		HTTPSAddr:               "",
		Listeners:               nil,
		Handler:                 nil,
		ReadTimeout:             60 * time.Second,
		WriteTimeout:            0, // No timeout.
		IdleTimeout:             60 * time.Second,
		MaxHeaderBytes:          10240, // 10KB
		KeepAlive:               true,
		IndexFiles:              []string{"index.html", "index.htm"},
		IndexFolder:             false,
		ServerAgent:             "GoFrame HTTP Server",
		ServerRoot:              "",
		StaticPaths:             make([]staticPathItem, 0),
		FileServerEnabled:       false,
		CookieMaxAge:            time.Hour * 24 * 365,
		CookiePath:              "/",
		CookieDomain:            "",
		SessionIdName:           "gfsessionid",
		SessionPath:             gsession.DefaultStorageFilePath,
		SessionMaxAge:           time.Hour * 24,
		SessionCookieOutput:     true,
		SessionCookieMaxAge:     time.Hour * 24,
		Logger:                  glog.New(),
		LogLevel:                "all",
		LogStdout:               true,
		ErrorStack:              true,
		ErrorLogEnabled:         true,
		ErrorLogPattern:         "error-{Ymd}.log",
		AccessLogEnabled:        false,
		AccessLogPattern:        "access-{Ymd}.log",
		DumpRouterMap:           true,
		ClientMaxBodySize:       8 * 1024 * 1024, // 8MB
		FormParsingMemory:       1024 * 1024,     // 1MB
		Rewrites:                make(map[string]string),
		Graceful:                false,
		GracefulTimeout:         2, // seconds
		GracefulShutdownTimeout: 5, // seconds
	}
}

// ConfigFromMap根据给定的映射和默认配置对象，创建并返回一个ServerConfig对象。
// md5:4e735da260d3d596
// ff:创建配置对象Map
// m:配置
// ServerConfig:
func ConfigFromMap(m map[string]interface{}) (ServerConfig, error) {
	config := NewConfig()
	if err := gconv.Struct(m, &config); err != nil {
		return config, err
	}
	return config, nil
}

// SetConfigWithMap 使用映射(map)设置服务器的配置。 md5:896223070f4171c3
// ff:设置配置项Map
// s:
// m:配置
func (s *Server) SetConfigWithMap(m map[string]interface{}) error {
// m 现在是 m 的浅拷贝。
// 对 m 的任何修改都不会影响原始对象。
// 这有点巧妙，不是吗？
// md5:4d1dd38c4db57a79
	m = gutil.MapCopy(m)
// 允许使用字符串大小（如：1m、100mb、512kb等）来设置配置项的大小。
// md5:1afcd879d2e46708
	if k, v := gutil.MapPossibleItemByKey(m, "MaxHeaderBytes"); k != "" {
		m[k] = gfile.StrToSize(gconv.String(v))
	}
	if k, v := gutil.MapPossibleItemByKey(m, "ClientMaxBodySize"); k != "" {
		m[k] = gfile.StrToSize(gconv.String(v))
	}
	if k, v := gutil.MapPossibleItemByKey(m, "FormParsingMemory"); k != "" {
		m[k] = gfile.StrToSize(gconv.String(v))
	}
// 更新当前配置对象。
// 它只会更新已配置的键，而不是整个对象。
// md5:31a4f7b2094c229e
	if err := gconv.Struct(m, &s.config); err != nil {
		return err
	}
	return s.SetConfig(s.config)
}

// SetConfig 为服务器设置配置。 md5:cf09327be418468a
// ff:设置配置项
// s:
// c:
func (s *Server) SetConfig(c ServerConfig) error {
	s.config = c
	// 如果地址缺少':'前缀，自动添加。 md5:7493e54c133e3353
	if s.config.Address != "" && !gstr.Contains(s.config.Address, ":") {
		s.config.Address = ":" + s.config.Address
	}
	// Static files root.
	if c.ServerRoot != "" {
		s.SetServerRoot(c.ServerRoot)
	}
	if len(c.SearchPaths) > 0 {
		paths := c.SearchPaths
		c.SearchPaths = []string{}
		for _, v := range paths {
			s.AddSearchPath(v)
		}
	}
	// HTTPS.
	if c.TLSConfig == nil && c.HTTPSCertPath != "" {
		s.EnableHTTPS(c.HTTPSCertPath, c.HTTPSKeyPath)
	}
	// Logging.
	if s.config.LogPath != "" && s.config.LogPath != s.config.Logger.GetPath() {
		if err := s.config.Logger.SetPath(s.config.LogPath); err != nil {
			return err
		}
	}
	if err := s.config.Logger.SetLevelStr(s.config.LogLevel); err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	gracefulEnabled = c.Graceful
	intlog.Printf(context.TODO(), "SetConfig: %+v", s.config)
	return nil
}

// SetAddr 设置服务器的监听地址。
// 地址格式为 ':80'，'0.0.0.0:80'，'127.0.0.1:80'，'180.18.99.10:80' 等。
// md5:c62ee3ae1a0d6764
// ff:设置监听地址
// s:
// address:地址
func (s *Server) SetAddr(address string) {
	s.config.Address = address
}

// SetPort 设置服务器的监听端口。
// 可以设置多个端口，例如：SetPort(80, 8080)。
// md5:e6fc730d6e8ee17c
// ff:设置监听端口
// s:
// port:端口
func (s *Server) SetPort(port ...int) {
	if len(port) > 0 {
		s.config.Address = ""
		for _, v := range port {
			if len(s.config.Address) > 0 {
				s.config.Address += ","
			}
			s.config.Address += ":" + strconv.Itoa(v)
		}
	}
}

// SetHTTPSAddr 设置服务器的HTTPS监听端口。 md5:2cfcb725865474b8
// ff:设置HTTPS监听地址
// s:
// address:地址
func (s *Server) SetHTTPSAddr(address string) {
	s.config.HTTPSAddr = address
}

// SetHTTPSPort 设置服务器的HTTPS监听端口。
// 可以设置多个监听端口，如：SetHTTPSPort(443, 500)。
// md5:cdbfe394641fff8a
// ff:设置HTTPS监听端口
// s:
// port:端口
func (s *Server) SetHTTPSPort(port ...int) {
	if len(port) > 0 {
		s.config.HTTPSAddr = ""
		for _, v := range port {
			if len(s.config.HTTPSAddr) > 0 {
				s.config.HTTPSAddr += ","
			}
			s.config.HTTPSAddr += ":" + strconv.Itoa(v)
		}
	}
}

// SetListener 设置服务器的自定义监听器。 md5:5ce2e30da0c2811f
// ff:设置自定义监听器
// s:
// listeners:监听器
func (s *Server) SetListener(listeners ...net.Listener) error {
	if listeners == nil {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "SetListener failed: listener can not be nil")
	}
	if len(listeners) > 0 {
		ports := make([]string, len(listeners))
		for k, v := range listeners {
			if v == nil {
				return gerror.NewCodef(gcode.CodeInvalidParameter, "SetListener failed: listener can not be nil")
			}
			ports[k] = fmt.Sprintf(":%d", (v.Addr().(*net.TCPAddr)).Port)
		}
		s.config.Address = strings.Join(ports, ",")
		s.config.Listeners = listeners
	}
	return nil
}

// EnableHTTPS 为服务器启用HTTPS，使用给定的证书和密钥文件。可选参数`tlsConfig`用于指定自定义TLS配置。
// md5:0e566bf091aacfd8
// ff:启用HTTPS
// s:
// certFile:证书路径
// keyFile:密钥路径
// tlsConfig:tls配置
func (s *Server) EnableHTTPS(certFile, keyFile string, tlsConfig ...*tls.Config) {
	var ctx = context.TODO()
	certFileRealPath := gfile.RealPath(certFile)
	if certFileRealPath == "" {
		certFileRealPath = gfile.RealPath(gfile.Pwd() + gfile.Separator + certFile)
		if certFileRealPath == "" {
			certFileRealPath = gfile.RealPath(gfile.MainPkgPath() + gfile.Separator + certFile)
		}
	}
	// Resource.
	if certFileRealPath == "" && gres.Contains(certFile) {
		certFileRealPath = certFile
	}
	if certFileRealPath == "" {
		s.Logger().Fatalf(ctx, `EnableHTTPS failed: certFile "%s" does not exist`, certFile)
	}
	keyFileRealPath := gfile.RealPath(keyFile)
	if keyFileRealPath == "" {
		keyFileRealPath = gfile.RealPath(gfile.Pwd() + gfile.Separator + keyFile)
		if keyFileRealPath == "" {
			keyFileRealPath = gfile.RealPath(gfile.MainPkgPath() + gfile.Separator + keyFile)
		}
	}
	// Resource.
	if keyFileRealPath == "" && gres.Contains(keyFile) {
		keyFileRealPath = keyFile
	}
	if keyFileRealPath == "" {
		s.Logger().Fatal(ctx, `EnableHTTPS failed: keyFile "%s" does not exist`, keyFile)
	}
	s.config.HTTPSCertPath = certFileRealPath
	s.config.HTTPSKeyPath = keyFileRealPath
	if len(tlsConfig) > 0 {
		s.config.TLSConfig = tlsConfig[0]
	}
}

// SetTLSConfig 设置自定义的TLS配置，并为服务器启用HTTPS功能。 md5:e8fae606e7c9daa6
// ff:设置TLS配置
// s:
// tlsConfig:tls配置
func (s *Server) SetTLSConfig(tlsConfig *tls.Config) {
	s.config.TLSConfig = tlsConfig
}

// SetReadTimeout 设置服务器的读取超时时间。 md5:731a0457ad074a1e
// ff:设置读取超时
// s:
// t:时长
func (s *Server) SetReadTimeout(t time.Duration) {
	s.config.ReadTimeout = t
}

// SetWriteTimeout 设置服务器的写超时。 md5:60f9efbd1b42a85a
// ff:设置写入超时
// s:
// t:时长
func (s *Server) SetWriteTimeout(t time.Duration) {
	s.config.WriteTimeout = t
}

// SetIdleTimeout 设置服务器的空闲超时时间。 md5:f18e89634fa33c02
// ff:设置长连接超时
// s:
// t:时长
func (s *Server) SetIdleTimeout(t time.Duration) {
	s.config.IdleTimeout = t
}

// SetMaxHeaderBytes 为服务器设置 MaxHeaderBytes。 md5:2e7198560eedacbb
// ff:设置协议头最大长度
// s:
// b:最大长度
func (s *Server) SetMaxHeaderBytes(b int) {
	s.config.MaxHeaderBytes = b
}

// SetServerAgent 设置服务器的ServerAgent。 md5:ac1c65804355cc50
// ff:设置服务器代理标识
// s:
// agent:代理标识
func (s *Server) SetServerAgent(agent string) {
	s.config.ServerAgent = agent
}

// SetKeepAlive 设置服务器的KeepAlive。 md5:54c342f49d9fa171
// ff:设置开启长连接
// s:
// enabled:开启
func (s *Server) SetKeepAlive(enabled bool) {
	s.config.KeepAlive = enabled
}

// SetView 设置服务器的视图。 md5:ec7bba6db1e3a9cf
// ff:设置默认模板对象
// s:
// view:模板对象
func (s *Server) SetView(view *gview.View) {
	s.config.View = view
}

// GetName 返回服务器的名称。 md5:1662443760c53efe
// ff:取服务名称
// s:
func (s *Server) GetName() string {
	return s.config.Name
}

// SetName 设置服务器的名称。 md5:242f311a4c185514
// ff:设置服务名称
// s:
// name:名称
func (s *Server) SetName(name string) {
	s.config.Name = name
}

// SetEndpoints 设置服务器的端点。 md5:b75987e400904902
// ff:
// s:
// endpoints:
func (s *Server) SetEndpoints(endpoints []string) {
	s.config.Endpoints = endpoints
}

// SetHandler 设置服务器的请求处理器。 md5:c2ce7c70be19e1d5
// ff:设置请求处理器
// s:
// h:
// w:
// r:
func (s *Server) SetHandler(h func(w http.ResponseWriter, r *http.Request)) {
	s.config.Handler = h
}

// GetHandler 返回服务器的请求处理程序。 md5:97d22a3db48dd77d
// ff:取请求处理器
// s:
// w:
// r:
func (s *Server) GetHandler() func(w http.ResponseWriter, r *http.Request) {
	if s.config.Handler == nil {
		return s.ServeHTTP
	}
	return s.config.Handler
}

// SetRegistrar 设置服务器的注册器。 md5:59baf7cae4845598
// ff:设置注册发现对象
// s:
// registrar:注册发现对象
func (s *Server) SetRegistrar(registrar gsvc.Registrar) {
	s.registrar = registrar
}

// GetRegistrar 返回服务器的注册商。 md5:d5a67dbd4e6ac976
// ff:取注册发现对象
// s:
func (s *Server) GetRegistrar() gsvc.Registrar {
	return s.registrar
}
