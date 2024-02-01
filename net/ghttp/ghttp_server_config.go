// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

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
type ServerConfig struct {
// ======================================================================================================
// 基础部分.
// ======================================================================================================
// 这段注释表明接下来的代码是关于“基础”部分，用于概括或分隔代码的不同模块或功能区块。

	// 服务名称，用于服务注册与发现。
	Name string `json:"name"`

// Address 指定服务器监听地址，格式如 "端口" 或 ":端口"，
// 多个地址之间使用 ',' 连接。
	Address string `json:"address"`

	// HTTPSAddr 指定HTTPS服务地址，多个地址之间使用逗号（,）连接。
	HTTPSAddr string `json:"httpsAddr"`

	// Listeners 指定自定义监听器。
	Listeners []net.Listener `json:"listeners"`

	// Endpoints 是服务注册的自定义端点，如果为空则使用 Address。
	Endpoints []string `json:"endpoints"`

	// HTTPSCertPath 指定 HTTPS 服务的证书文件路径。
	HTTPSCertPath string `json:"httpsCertPath"`

	// HTTPSKeyPath 指定HTTPS服务的密钥文件路径。
	HTTPSKeyPath string `json:"httpsKeyPath"`

// TLSConfig 提供一个可选的 TLS 配置，用于 ServeTLS 和 ListenAndServeTLS。请注意，
// 此值将被 ServeTLS 和 ListenAndServeTLS 拷贝使用，因此不能通过诸如 tls.Config.SetSessionTicketKeys
// 等方法修改配置。若要使用 SetSessionTicketKeys，请改用 Server.Serve 方法配合一个 TLS 监听器来实现。
	TLSConfig *tls.Config `json:"tlsConfig"`

	// Handler HTTP请求的处理器。
	Handler func(w http.ResponseWriter, r *http.Request) `json:"-"`

// ReadTimeout 是读取整个请求（包括主体）的最大持续时间。
//
// 由于 ReadTimeout 不允许处理程序为每个请求主体设置可接受的截止日期或上传速率，
// 大多数用户可能会更倾向于使用 ReadHeaderTimeout。同时使用它们是有效的。
	ReadTimeout time.Duration `json:"readTimeout"`

// WriteTimeout 是在超时前写入响应的最大持续时间。每当读取到新请求的头部时，该时间就会重置。类似于 ReadTimeout，它并不允许 Handlers 根据每个请求自行决定是否超时。
	WriteTimeout time.Duration `json:"writeTimeout"`

// IdleTimeout 是在启用 keep-alive 时，等待下一个请求的最大时间间隔。如果 IdleTimeout 设为零，则使用 ReadTimeout 的值。如果两者都为零，则表示没有超时限制。
	IdleTimeout time.Duration `json:"idleTimeout"`

// MaxHeaderBytes 控制服务器在解析请求头（包括请求行）的键和值时，
// 会读取的最大字节数。但请注意，它并不会限制请求体的大小。
//
// 你可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串来配置这个参数。
// 默认情况下，其值为 10240 字节。
	MaxHeaderBytes int `json:"maxHeaderBytes"`

	// KeepAlive 启用 HTTP 保持连接（Keep-alive）功能。
	KeepAlive bool `json:"keepAlive"`

// ServerAgent 指定服务器代理信息，该信息会被写入
// HTTP 响应头中作为 "Server"。
	ServerAgent string `json:"serverAgent"`

	// View 指定了服务器的默认模板视图对象。
	View *gview.View `json:"view"`

// ======================================================================================================
// 静态部分。
// ======================================================================================================
// 这段注释表明了该处代码是“静态”部分，但具体含义可能需要更多上下文信息来准确翻译。在程序中，“Static”通常是指不变的、非运行时动态改变的数据或函数，或者是初始化后在整个程序生命周期中保持固定的部分。

	// Rewrites 指定了 URI 重写规则映射。
	Rewrites map[string]string `json:"rewrites"`

	// IndexFiles 指定静态文件夹的索引文件。
	IndexFiles []string `json:"indexFiles"`

// IndexFolder 指定在请求文件夹时是否列出子文件。
// 如果该值为false，服务器将返回HTTP状态码403。
	IndexFolder bool `json:"indexFolder"`

	// ServerRoot 指定静态服务的根目录。
	ServerRoot string `json:"serverRoot"`

	// SearchPaths 指定静态服务的额外搜索目录。
	SearchPaths []string `json:"searchPaths"`

	// StaticPaths 指定了URI到目录映射的数组。
	StaticPaths []staticPathItem `json:"staticPaths"`

// FileServerEnabled 是静态服务的全局开关。
// 如果设置了任何静态路径，它将自动设置为启用状态。
	FileServerEnabled bool `json:"fileServerEnabled"`

// ======================================================================================================
// Cookie.
// ======================================================================================================
// 以下是翻译后的中文注释：
// ======================================================================================================
// Cookie（cookies）.
// ======================================================================================================
// 此处的注释表明该部分代码与处理或操作Cookie相关的功能有关。"Cookie"在Web开发中指的是服务器发送到用户浏览器并存储在本地的一小段数据，用于识别不同的用户和跟踪会话状态等信息。

	// CookieMaxAge 指定 cookie 项的最大生存时间（TTL）。
	CookieMaxAge time.Duration `json:"cookieMaxAge"`

// CookiePath 指定cookie路径。
// 同时，它也影响session id的默认存储位置。
	CookiePath string `json:"cookiePath"`

// CookieDomain 指定 cookie 域名。
// 同时，它也影响 session id 的默认存储方式。
	CookieDomain string `json:"cookieDomain"`

// CookieSameSite 指定 cookie 的 SameSite 属性。
// 同时，它也影响会话 ID 的默认存储方式。
	CookieSameSite string `json:"cookieSameSite"`

// CookieSameSite 指定 cookie 的 Secure 属性。
// 同时，它也影响 session id 的默认存储方式。
	CookieSecure bool `json:"cookieSecure"`

// CookieSameSite 指定 cookie 的 HttpOnly 属性。
// 同时，它也会影响 session id 的默认存储方式。
	CookieHttpOnly bool `json:"cookieHttpOnly"`

// ======================================================================================================
// 会话.
// ======================================================================================================
// 这段代码中的注释表明了接下来要定义或描述的内容是关于“Session”（会话）的，但没有给出具体的代码实现细节。在程序中，"Session"通常用于表示用户与服务器之间交互过程的状态信息，用于维持状态、存储临时数据等。

	// SessionIdName 指定会话 ID 名称。
	SessionIdName string `json:"sessionIdName"`

	// SessionMaxAge 指定会话项的最大生存时间（TTL）。
	SessionMaxAge time.Duration `json:"sessionMaxAge"`

// SessionPath 指定用于存储会话文件的会话存储目录路径。
// 只有当会话存储类型为文件存储时，这个配置才有意义。
	SessionPath string `json:"sessionPath"`

	// SessionStorage 指定会话存储。
	SessionStorage gsession.Storage `json:"sessionStorage"`

// SessionCookieMaxAge 指定会话 ID 的 cookie 存活时间（TTL）。
// 如果设置为 0，表示它将随浏览器会话一同结束时失效。
	SessionCookieMaxAge time.Duration `json:"sessionCookieMaxAge"`

	// SessionCookieOutput 指定是否自动将会话ID输出到cookie中。
	SessionCookieOutput bool `json:"sessionCookieOutput"`

// ======================================================================================================
// 日志记录。
// ======================================================================================================

	Logger           *glog.Logger `json:"logger"`           // Logger 指定服务器使用的日志记录器。
	LogPath          string       `json:"logPath"`          // LogPath 指定存储日志文件的目录。
	LogLevel         string       `json:"logLevel"`         // LogLevel 指定 logger 的日志记录级别。
	LogStdout        bool         `json:"logStdout"`        // LogStdout 指定是否将日志内容输出到标准输出（stdout）中。
	ErrorStack       bool         `json:"errorStack"`       // ErrorStack 指定在出现错误时是否记录堆栈信息。
	ErrorLogEnabled  bool         `json:"errorLogEnabled"`  // ErrorLogEnabled 开启错误日志功能，将错误内容记录到文件中。
	ErrorLogPattern  string       `json:"errorLogPattern"`  // ErrorLogPattern 指定错误日志文件的命名模式，例如：error-{Ymd}.log
	AccessLogEnabled bool         `json:"accessLogEnabled"` // AccessLogEnabled 开启访问日志功能，将访问内容记录到文件中。
	AccessLogPattern string       `json:"accessLogPattern"` // AccessLogPattern 指定访问日志文件的命名模式，如：access-{Ymd}.log

// ======================================================================================================
// PProf.
// ======================================================================================================
// 此处为Golang代码注释的中文翻译：
// ======================================================================================================
// PProf：这是一个对Go程序进行性能分析的接口包，主要用于生成和解析CPU、内存等资源占用情况的采样数据，
// 以便开发者找出程序中的性能瓶颈。

	PProfEnabled bool   `json:"pprofEnabled"` // PProfEnabled 开启PProf功能。
	PProfPattern string `json:"pprofPattern"` // PProfPattern 指定路由器的 PProf 服务模式。

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

	OpenApiPath string `json:"openapiPath"` // OpenApiPath 指定OpenApi规范文件的路径。
	SwaggerPath string `json:"swaggerPath"` // SwaggerPath 指定Swagger UI的路径，用于注册路由。

// ======================================================================================================
// 其他
// ======================================================================================================

// ClientMaxBodySize 指定客户端请求的最大正文大小限制，单位为字节。
// 你可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串进行配置。
// 默认值为 `8MB`。
	ClientMaxBodySize int64 `json:"clientMaxBodySize"`

// FormParsingMemory 指定用于解析多媒体表单的最大内存缓冲区大小（以字节为单位）。
// 可以在配置文件中使用类似 "1m"、"10m"、"500kb" 等字符串进行配置。
// 默认值为 1MB。
	FormParsingMemory int64 `json:"formParsingMemory"`

// NameToUriType 指定了在注册路由时，将结构体方法名转换为URI的类型。
	NameToUriType int `json:"nameToUriType"`

	// RouteOverWrite 允许在出现重复路由时进行覆盖。
	RouteOverWrite bool `json:"routeOverWrite"`

	// DumpRouterMap 指定在服务器启动时是否自动转储路由映射。
	DumpRouterMap bool `json:"dumpRouterMap"`

	// Graceful 启用进程内所有服务器的优雅重启功能。
	Graceful bool `json:"graceful"`

	// GracefulTimeout 设置父进程的最大存活时间（秒）。
	GracefulTimeout uint8 `json:"gracefulTimeout"`

	// GracefulShutdownTimeout 设置在停止服务器之前其最大存活时间（秒）。
	GracefulShutdownTimeout uint8 `json:"gracefulShutdownTimeout"`
}

// NewConfig 创建并返回一个具有默认配置的 ServerConfig 对象。
// 注意，不要将此默认配置定义为本地包变量，因为存在一些指针属性，
// 这些属性可能在不同的服务器中被共享。
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

// ConfigFromMap 根据给定的映射和默认配置对象创建并返回一个 ServerConfig 对象。
func ConfigFromMap(m map[string]interface{}) (ServerConfig, error) {
	config := NewConfig()
	if err := gconv.Struct(m, &config); err != nil {
		return config, err
	}
	return config, nil
}

// SetConfigWithMap 使用map设置服务器的配置
func (s *Server) SetConfigWithMap(m map[string]interface{}) error {
// 现在的m是m的一个浅拷贝。
// 对m的任何改动都不会影响原始的那个m。
// 这有点小巧妙，不是吗？
	m = gutil.MapCopy(m)
// 允许使用类似“1m、100mb、512kb”等字符串形式的大小来设置尺寸配置项：
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
// 仅更新已配置的键，而非整个对象。
	if err := gconv.Struct(m, &s.config); err != nil {
		return err
	}
	return s.SetConfig(s.config)
}

// SetConfig 设置服务器的配置。
func (s *Server) SetConfig(c ServerConfig) error {
	s.config = c
	// 如果地址中缺少':'前缀，则自动添加。
	if s.config.Address != "" && !gstr.Contains(s.config.Address, ":") {
		s.config.Address = ":" + s.config.Address
	}
	// 静态文件根目录。
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
// 地址格式类似于 ':80'、'0.0.0.0:80'、'127.0.0.1:80'、'180.18.99.10:80' 等。
func (s *Server) SetAddr(address string) {
	s.config.Address = address
}

// SetPort 设置服务器监听端口。
// 监听端口可以设置多个，例如：SetPort(80, 8080)。
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

// SetHTTPSAddr 设置服务器的 HTTPS 监听端口。
func (s *Server) SetHTTPSAddr(address string) {
	s.config.HTTPSAddr = address
}

// SetHTTPSPort 设置服务器的 HTTPS 监听端口。
// 可以设置多个监听端口，例如：SetHTTPSPort(443, 500)。
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

// SetListener 为服务器设置自定义监听器。
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

// EnableHTTPS 通过给定的证书和密钥文件为服务器启用HTTPS。
// 可选参数 `tlsConfig` 指定了自定义的TLS配置。
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

// SetTLSConfig 设置自定义 TLS 配置，并为服务器启用 HTTPS 功能。
func (s *Server) SetTLSConfig(tlsConfig *tls.Config) {
	s.config.TLSConfig = tlsConfig
}

// SetReadTimeout 设置服务器的读取超时时间。
func (s *Server) SetReadTimeout(t time.Duration) {
	s.config.ReadTimeout = t
}

// SetWriteTimeout 为服务器设置写入超时时间。
func (s *Server) SetWriteTimeout(t time.Duration) {
	s.config.WriteTimeout = t
}

// SetIdleTimeout 为服务器设置空闲超时时间。
func (s *Server) SetIdleTimeout(t time.Duration) {
	s.config.IdleTimeout = t
}

// SetMaxHeaderBytes 设置服务器的 MaxHeaderBytes 值。
func (s *Server) SetMaxHeaderBytes(b int) {
	s.config.MaxHeaderBytes = b
}

// SetServerAgent 设置服务器的 ServerAgent。
func (s *Server) SetServerAgent(agent string) {
	s.config.ServerAgent = agent
}

// SetKeepAlive 设置服务器的 KeepAlive 参数。
func (s *Server) SetKeepAlive(enabled bool) {
	s.config.KeepAlive = enabled
}

// SetView 设置服务器的视图。
func (s *Server) SetView(view *gview.View) {
	s.config.View = view
}

// GetName 返回服务器的名称。
func (s *Server) GetName() string {
	return s.config.Name
}

// SetName 为服务器设置名称。
func (s *Server) SetName(name string) {
	s.config.Name = name
}

// SetEndpoints 设置服务器的端点。
func (s *Server) SetEndpoints(endpoints []string) {
	s.config.Endpoints = endpoints
}

// SetHandler 为服务器设置请求处理器。
func (s *Server) SetHandler(h func(w http.ResponseWriter, r *http.Request)) {
	s.config.Handler = h
}

// GetHandler 返回服务器的请求处理器。
func (s *Server) GetHandler() func(w http.ResponseWriter, r *http.Request) {
	if s.config.Handler == nil {
		return s.ServeHTTP
	}
	return s.config.Handler
}

// SetRegistrar 设置服务器的 Registrar。
func (s *Server) SetRegistrar(registrar gsvc.Registrar) {
	s.registrar = registrar
}

// GetRegistrar 返回服务器的注册器。
func (s *Server) GetRegistrar() gsvc.Registrar {
	return s.registrar
}
