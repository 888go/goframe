// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package ghttp 提供了强大的HTTP服务器及简洁的客户端实现。
package ghttp

import (
	"net/http"
	"reflect"
	"sync"
	"time"
	
	"github.com/gorilla/websocket"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/net/goai"
	"github.com/888go/goframe/net/gsvc"
	"github.com/888go/goframe/os/gcache"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/gsession"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/util/gtag"
)

type (
	// Server 包装了 http.Server，并提供了更多丰富的功能。
	Server struct {
		instance         string                    // 当前HTTP服务器的实例名称。
		config           ServerConfig              // 服务器配置
		plugins          []Plugin                  // 插件数组，用于扩展服务器功能。
		servers          []*gracefulServer         // 底层 http.Server 数组
		serverCount      *gtype.Int                // 用于内部使用的底层http.Server数字。
		closeChan        chan struct{}             // 用于底层服务器关闭事件的通知。
		serveTree        map[string]interface{}    // 路由映射树
		serveCache       *gcache.Cache             // Server内部使用的缓存。
		routesMap        map[string][]*HandlerItem // 路由映射表主要用于路由转储和重复路由检查。
		statusHandlerMap map[string][]HandlerFunc  // 自定义状态处理映射。
		sessionManager   *gsession.Manager         // Session manager.
		openapi          *goai.OpenApiV3           // OpenApi规范管理对象。
		serviceMu        sync.Mutex                // 并发安全：保证属性服务操作的并发安全性。
		service          gsvc.Service              // Registry服务。
		registrar        gsvc.Registrar            // Registrar 用于服务注册。
	}

	// Router object.
	Router struct {
		Uri      string   // URI.
		Method   string   // HTTP method
		Domain   string   // Bound domain.
		RegRule  string   // 解析后的用于路由匹配的正则表达式。
		RegNames []string // 解析后的路由参数名称。
		Priority int      // 仅供参考
	}

	// RouterItem 仅用于路由转储。
	RouterItem struct {
		Handler          *HandlerItem // The handler.
		Server           string       // Server name.
		Address          string       // 监听地址。
		Domain           string       // Bound domain.
		Type             HandlerType  // 路由处理器类型。
		Middleware       string       // Bound middleware.
		Method           string       // 处理器方法名称。
		Route            string       // Route URI.
		Priority         int          // 仅供参考
		IsServiceHandler bool         // 是否为服务处理器
	}

	// HandlerFunc 是请求处理函数。
	HandlerFunc = func(r *Request)

	// handlerFuncInfo 包含 HandlerFunc 的地址及其反射类型。
	handlerFuncInfo struct {
		Func            HandlerFunc      // 处理函数地址。
		Type            reflect.Type     // 反射当前处理器的类型信息，该信息用于扩展处理器功能。
		Value           reflect.Value    // Reflect当前处理器的值信息，该信息用于扩展处理器功能。
		IsStrictRoute   bool             // 是否启用严格的路由匹配。
		ReqStructFields []gstructs.Field // 请求结构体字段。
	}

// HandlerItem 是注册的路由处理程序，
// 包括中间件和钩子函数。
	HandlerItem struct {
// 唯一处理器项标识符标记。
// 注意，处理器函数可能会以不同的处理器项身份注册多次，
// 这些处理器项具有不同的处理器项标识符。
		Id         int
		Name       string          // 处理器名称，在注册时会自动从运行时堆栈中获取。
		Type       HandlerType     // 处理器类型：对象/处理器/中间件/钩子。
		Info       handlerFuncInfo // 处理函数信息
		InitFunc   HandlerFunc     // 初始化函数：当请求进入对象时调用（仅适用于对象注册类型）。
		ShutFunc   HandlerFunc     // Shutdown 函数在请求离开对象时调用（仅适用于对象注册类型）。
		Middleware []HandlerFunc   // 绑定中间件数组。
		HookName   HookName        // Hook类型名称，仅适用于Hook类型。
		Router     *Router         // Router object.
		Source     string          // 注册源文件 `路径:行数`。
	}

	// HandlerItemParsed是从URL.Path中解析出的项目。
	HandlerItemParsed struct {
		Handler *HandlerItem      // 处理器信息。
		Values  map[string]string // Router values 从 URL.Path 解析得到。
	}

	// ServerStatus 是服务器状态枚举类型。
	ServerStatus = int

	// HookName 是路由钩子名称的枚举类型。
	HookName string

	// HandlerType 是路由处理器的枚举类型。
	HandlerType string

// 监听文件描述符映射。
// 键是 "http" 或 "https"，其对应的值为相应的文件描述符（FD）。
	listenerFdMap = map[string]string

	// internalPanic 是用于内部使用的自定义恐慌函数。
	internalPanic string
)

const (
	// FreePortAddress 表示服务器使用随机空闲端口进行监听。
	FreePortAddress = ":0"
)

const (
	HeaderXUrlPath                     = "x-url-path"         // 用于自定义路由处理器，在此情况下，URL.Path 不会发生变化。
	HookBeforeServe       HookName     = "HOOK_BEFORE_SERVE"  // 在路由处理器/文件服务之前执行的钩子处理器。
	HookAfterServe        HookName     = "HOOK_AFTER_SERVE"   // 在路由处理器/文件服务之后的钩子处理器。
	HookBeforeOutput      HookName     = "HOOK_BEFORE_OUTPUT" // 在响应输出前的钩子处理器
	HookAfterOutput       HookName     = "HOOK_AFTER_OUTPUT"  // Hook处理器在响应输出之后。
	ServerStatusStopped   ServerStatus = 0
	ServerStatusRunning   ServerStatus = 1
	DefaultServerName                  = "default"
	DefaultDomainName                  = "default"
	HandlerTypeHandler    HandlerType  = "handler"
	HandlerTypeObject     HandlerType  = "object"
	HandlerTypeMiddleware HandlerType  = "middleware"
	HandlerTypeHook       HandlerType  = "hook"
)

const (
	supportedHttpMethods                    = "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE"
	defaultMethod                           = "ALL"
	routeCacheDuration                      = time.Hour
	ctxKeyForRequest            gctx.StrKey = "gHttpRequestObject"
	contentTypeXml                          = "text/xml"
	contentTypeHtml                         = "text/html"
	contentTypeJson                         = "application/json"
	swaggerUIPackedPath                     = "/goframe/swaggerui"
	responseHeaderTraceID                   = "Trace-ID"
	responseHeaderContentLength             = "Content-Length"
	specialMethodNameInit                   = "Init"
	specialMethodNameShut                   = "Shut"
	specialMethodNameIndex                  = "Index"
	defaultEndpointPort                     = 80
)

const (
	exceptionExit     internalPanic = "exit"
	exceptionExitAll  internalPanic = "exit_all"
	exceptionExitHook internalPanic = "exit_hook"
)

var (
// methodsMap 存储所有支持的HTTP方法。
// 该映射用于通过map快速搜索HTTP方法。
	methodsMap = make(map[string]struct{})

// serverMapping 用于存储当前进程中多个服务器实例。
// 键是服务器的名称，值是其对应的实例。
	serverMapping = gmap.NewStrAnyMap(true)

// serverRunning 标记正在运行的服务器数量。
// 如果没有运行成功的服务器，或者所有服务器都已关闭，则该值为0。
	serverRunning = gtype.NewInt()

	// wsUpGrader 是用于websocket的默认升级配置。
	wsUpGrader = websocket.Upgrader{
		// 默认情况下，它不检查源，应用程序可以自行进行检查。
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
// allShutdownChan 是所有服务器完成服务并退出的事件。
// 它用于进程阻塞目的。
	allShutdownChan = make(chan struct{}, 1000)

// serverProcessInitialized 用于服务端的延迟初始化。
// 这个过程只能被初始化一次。
	serverProcessInitialized = gtype.NewBool()

	// gracefulEnabled 用于实现优雅重启功能，默认情况下为 false。
	gracefulEnabled = false

	// defaultValueTags 是用于存储默认值的结构体标签名称。
	defaultValueTags = []string{gtag.DefaultShort, gtag.Default}
)

var (
	ErrNeedJsonBody = gerror.NewWithOption(gerror.Option{
		Text: "the request body content should be JSON format",
		Code: gcode.CodeInvalidRequest,
	})
)
