# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[Name string `json:"name"`]
qm=服务名称
cz=Name string `json:"name"`

[Address string `json:"address"`]
qm=监听地址
cz=Address string `json:"address"`

[HTTPSAddr string `json:"httpsAddr"`]
qm=HTTPS监听地址
cz=HTTPSAddr string `json:"httpsAddr"`

[Listeners #左中括号##右中括号#net.Listener `json:"listeners"`]
qm=自定义监听器
cz=Listeners []net.Listener

[HTTPSCertPath string `json:"httpsCertPath"`]
qm=HTTPS证书路径
cz=HTTPSCertPath string

[HTTPSKeyPath string `json:"httpsKeyPath"`]
qm=HTTPS密钥路径
cz=HTTPSKeyPath string

[TLSConfig *tls.Config `json:"tlsConfig"`]
qm=TLS配置
cz=TLSConfig *tls

[ReadTimeout time.Duration `json:"readTimeout"`]
qm=读取超时
cz=ReadTimeout time.Duration

[WriteTimeout time.Duration `json:"writeTimeout"`]
qm=写入超时
cz=WriteTimeout time.Duration

[IdleTimeout time.Duration `json:"idleTimeout"`]
qm=长连接超时
cz=IdleTimeout time.Duration

[MaxHeaderBytes int `json:"maxHeaderBytes"`]
qm=最大协议头长度
cz=MaxHeaderBytes int

[KeepAlive bool `json:"keepAlive"`]
qm=启用长连接
cz=KeepAlive bool

[ServerAgent string `json:"serverAgent"`]
qm=服务器代理
cz=ServerAgent string

[View *gview.View `json:"view"`]
qm=模板默认
cz=View *

[Rewrites map#左中括号#string#右中括号#string `json:"rewrites"`]
qm=路由URI重写规则Map
cz=Rewrites map

[IndexFiles #左中括号##右中括号#string `json:"indexFiles"`]
qm=静态文件索引
cz=IndexFiles []string

[IndexFolder bool `json:"indexFolder"`]
qm=静态文件是否列出子文件
cz=IndexFolder bool

[ServerRoot string `json:"serverRoot"`]
qm=静态文件根目录
cz=ServerRoot string

[SearchPaths #左中括号##右中括号#string `json:"searchPaths"`]
qm=静态文件额外搜索目录
cz=SearchPaths []string

[StaticPaths #左中括号##右中括号#staticPathItem `json:"staticPaths"`]
qm=静态文件目录映射
cz=StaticPaths []

[FileServerEnabled bool `json:"fileServerEnabled"`]
qm=静态文件是否开启
cz=FileServerEnabled bool

[CookieMaxAge time.Duration `json:"cookieMaxAge"`]
qm=Cookie最大存活时长
cz=CookieMaxAge time.Duration

[CookiePath string `json:"cookiePath"`]
qm=Cookie路径
cz=CookiePath string

[CookieDomain string `json:"cookieDomain"`]
qm=Cookie域名
cz=CookieDomain string

[CookieSecure bool `json:"cookieSecure"`]
qm=Cookie安全
cz=CookieSecure bool

[CookieHttpOnly bool `json:"cookieHttpOnly"`]
qm=Cookie跨站访问控制
cz=CookieHttpOnly bool

[SessionIdName string `json:"sessionIdName"`]
qm=SessionID名称
cz=SessionIdName string

[SessionMaxAge time.Duration `json:"sessionMaxAge"`]
qm=Session最大存活时长
cz=SessionMaxAge time.Duration

[SessionPath string `json:"sessionPath"`]
qm=Session存储目录路径
cz=SessionPath string

[SessionCookieMaxAge time.Duration `json:"sessionCookieMaxAge"`]
qm=SessionCookie存活时长
cz=SessionCookieMaxAge time.Duration

[SessionCookieOutput bool `json:"sessionCookieOutput"`]
qm=SessionID输出到Cookie
cz=SessionCookieOutput bool

[Logger *glog.Logger `json:"logger"`]
qm=日志记录器
cz=Logger *

[LogPath string `json:"logPath"`]
qm=日志存储目录
cz=LogPath string

[LogLevel string `json:"logLevel"`]
qm=日志记录等级
cz=LogLevel string

[LogStdout bool `json:"logStdout"`]
qm=日志开启输出到CMD
cz=LogStdout bool

[ErrorStack bool `json:"errorStack"`]
qm=日志开启错误堆栈记录
cz=ErrorStack bool

[ErrorLogEnabled bool `json:"errorLogEnabled"`]
qm=日志开启错误记录
cz=ErrorLogEnabled bool

[ErrorLogPattern string `json:"errorLogPattern"`]
qm=日志错误文件命名模式
cz=ErrorLogPattern string

[AccessLogEnabled bool `json:"accessLogEnabled"`]
qm=日志开启访客记录
cz=AccessLogEnabled bool

[AccessLogPattern string `json:"accessLogPattern"`]
qm=日志访客文件命名模式
cz=AccessLogPattern string

[PProfEnabled bool `json:"pprofEnabled"`]
qm=PProf开启
cz=PProfEnabled bool

[PProfPattern string `json:"pprofPattern"`]
qm=PProf模式
cz=PProfPattern string

[OpenApiPath string `json:"openapiPath"`]
qm=APIOpenApiUI路径
cz=OpenApiPath string

[SwaggerPath string `json:"swaggerPath"`]
qm=APISwaggerUI路径
cz=SwaggerPath string

[ClientMaxBodySize int64 `json:"clientMaxBodySize"`]
qm=客户端请求最大长度
cz=ClientMaxBodySize int64

[FormParsingMemory int64 `json:"formParsingMemory"`]
qm=表单解析最大缓冲区长度
cz=FormParsingMemory int64

[RouteOverWrite bool `json:"routeOverWrite"`]
qm=路由允许覆盖
cz=RouteOverWrite bool

[Graceful bool `json:"graceful"`]
qm=平滑重启开启
cz=Graceful bool
