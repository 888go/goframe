
# <翻译开始>
type ServerConfig
X服务配置项
# <翻译结束>

# <翻译开始>
Name string `json:"name"`
X服务名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Address string `json:"address"`
X监听地址
<跳到行首>
# <翻译结束>

# <翻译开始>
HTTPSAddr string `json:"httpsAddr"`
HTTPS监听地址
<跳到行首>
# <翻译结束>

# <翻译开始>
Listeners []net.Listener
X自定义监听器
<跳到行首>
# <翻译结束>

# <翻译开始>
HTTPSCertPath string
HTTPS证书路径
<跳到行首>
# <翻译结束>

# <翻译开始>
HTTPSKeyPath string
HTTPS密钥路径
<跳到行首>
# <翻译结束>

# <翻译开始>
TLSConfig *tls
TLS配置
<跳到行首>
# <翻译结束>

# <翻译开始>
ReadTimeout time.Duration
X读取超时
<跳到行首>
# <翻译结束>

# <翻译开始>
WriteTimeout time.Duration
X写入超时
<跳到行首>
# <翻译结束>

# <翻译开始>
IdleTimeout time.Duration
X长连接超时
<跳到行首>
# <翻译结束>

# <翻译开始>
MaxHeaderBytes int
X最大协议头长度
<跳到行首>
# <翻译结束>

# <翻译开始>
KeepAlive bool
X启用长连接
<跳到行首>
# <翻译结束>

# <翻译开始>
ServerAgent string
X服务器代理
<跳到行首>
# <翻译结束>

# <翻译开始>
View *
X模板默认
<跳到行首>
# <翻译结束>

# <翻译开始>
Rewrites map
X路由URI重写规则Map
<跳到行首>
# <翻译结束>

# <翻译开始>
IndexFiles []string
X静态文件索引
<跳到行首>
# <翻译结束>

# <翻译开始>
IndexFolder bool
X静态文件是否列出子文件
<跳到行首>
# <翻译结束>

# <翻译开始>
ServerRoot string
X静态文件根目录
<跳到行首>
# <翻译结束>

# <翻译开始>
SearchPaths []string
X静态文件额外搜索目录
<跳到行首>
# <翻译结束>

# <翻译开始>
StaticPaths []staticPathItem
X静态文件目录映射
# <翻译结束>

# <翻译开始>
FileServerEnabled bool
X静态文件是否开启
<跳到行首>
# <翻译结束>

# <翻译开始>
CookieMaxAge time.Duration
Cookie最大存活时长
<跳到行首>
# <翻译结束>

# <翻译开始>
CookiePath string
Cookie路径
<跳到行首>
# <翻译结束>

# <翻译开始>
CookieDomain string
Cookie域名
<跳到行首>
# <翻译结束>

# <翻译开始>
CookieSameSite string
CookieSameSite
<跳到行首>
# <翻译结束>

# <翻译开始>
CookieSecure bool
Cookie安全
<跳到行首>
# <翻译结束>

# <翻译开始>
CookieHttpOnly bool
Cookie跨站访问控制
<跳到行首>
# <翻译结束>

# <翻译开始>
SessionIdName string
SessionID名称
<跳到行首>
# <翻译结束>

# <翻译开始>
SessionMaxAge time.Duration
Session最大存活时长
<跳到行首>
# <翻译结束>

# <翻译开始>
SessionPath string
Session存储目录路径
<跳到行首>
# <翻译结束>

# <翻译开始>
SessionStorage gsession.Storage
Session存储
<跳到行首>
# <翻译结束>

# <翻译开始>
SessionCookieMaxAge time.Duration
SessionCookie存活时长
<跳到行首>
# <翻译结束>

# <翻译开始>
SessionCookieOutput bool
SessionID输出到Cookie
<跳到行首>
# <翻译结束>

# <翻译开始>
Logger           *
X日志记录器
<跳到行首>
# <翻译结束>

# <翻译开始>
LogPath          string
X日志存储目录
<跳到行首>
# <翻译结束>

# <翻译开始>
LogLevel         string
X日志记录等级
<跳到行首>
# <翻译结束>

# <翻译开始>
LogStdout        bool
X日志开启输出到CMD
<跳到行首>
# <翻译结束>

# <翻译开始>
ErrorStack       bool
X日志开启错误堆栈记录
<跳到行首>
# <翻译结束>

# <翻译开始>
ErrorLogEnabled  bool
X日志开启错误记录
<跳到行首>
# <翻译结束>

# <翻译开始>
ErrorLogPattern  string
X日志错误文件命名模式
<跳到行首>
# <翻译结束>

# <翻译开始>
AccessLogEnabled bool
X日志开启访客记录
<跳到行首>
# <翻译结束>

# <翻译开始>
AccessLogPattern string
X日志访客文件命名模式
<跳到行首>
# <翻译结束>

# <翻译开始>
PProfEnabled bool
PProf开启
<跳到行首>
# <翻译结束>

# <翻译开始>
PProfPattern string
PProf模式
<跳到行首>
# <翻译结束>

# <翻译开始>
OpenApiPath string
APIOpenApiUI路径
<跳到行首>
# <翻译结束>

# <翻译开始>
SwaggerPath string
APISwaggerUI路径
<跳到行首>
# <翻译结束>

# <翻译开始>
ClientMaxBodySize int64
X客户端请求最大长度
<跳到行首>
# <翻译结束>

# <翻译开始>
FormParsingMemory int64
X表单解析最大缓冲区长度
<跳到行首>
# <翻译结束>

# <翻译开始>
RouteOverWrite bool
X路由允许覆盖
<跳到行首>
# <翻译结束>

# <翻译开始>
Graceful bool
X平滑重启开启
<跳到行首>
# <翻译结束>
