// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包ghttp提供了强大的HTTP服务器和简单的客户端实现。 md5:0b6b4fc969b3488d
package ghttp

import (
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gtag"
)

type (
		// Server 包装了 http.Server 并提供了更多丰富的功能。 md5:0f435fc6994521cc
	Server struct {
		instance         string                    // 当前HTTP服务器的实例名称。 md5:9bf2787b3978a65a
		config           ServerConfig              // Server configuration.
		plugins          []Plugin                  // 用于扩展服务器功能的插件数组。 md5:9911152f56cd3480
		servers          []*gracefulServer         // 基础的 http.Server 数组。 md5:02c4898b1f5f6ef6
		serverCount      *gtype.Int                // 用于内部使用的底层http.Server编号。 md5:fa74896195039a36
		closeChan        chan struct{}             // 用于底层服务器关闭事件的通知。 md5:e8245a5fccddd4d8
		serveTree        map[string]interface{}    // The route maps tree.
		serveCache       *gcache.Cache             // 用于内部使用的服务器缓存。 md5:4af11401e9378d69
		routesMap        map[string][]*HandlerItem // 路由映射主要用于路由输出和重复路由检查。 md5:3a6dad4b5b6a5437
		statusHandlerMap map[string][]HandlerFunc  // 定义自定义状态处理器映射。 md5:f9a3aa7ca98708d3
		sessionManager   *gsession.Manager         // Session manager.
		openapi          *goai.OpenApiV3           // OpenApi规范管理对象。 md5:a28d3edf4dc5e22c
		serviceMu        sync.Mutex                // 属性服务的操作的并发安全性。 md5:3b9d36234acd3da2
		service          gsvc.Service              // 仓库服务。 md5:355c5f1d12149e89
		registrar        gsvc.Registrar            // 服务注册器。 md5:b10bd58ca0e98276
	}

	// Router object.
	Router struct {
		Uri      string   // URI.
		Method   string   // HTTP method
		Domain   string   // Bound domain.
		RegRule  string   // 用于路由匹配的解析正则表达式。 md5:8892e0f87233f591
		RegNames []string // 已解析的路由器参数名称。 md5:cb14202c5a1319f3
		Priority int      // Just for reference.
	}

		// RouterItem 仅仅用于路由输出。 md5:50edede8ea2ffda9
	RouterItem struct {
		Handler          *HandlerItem // The handler.
		Server           string       // Server name.
		Address          string       // Listening address.
		Domain           string       // Bound domain.
		Type             HandlerType  // Route handler type.
		Middleware       string       // Bound middleware.
		Method           string       // Handler method name.
		Route            string       // Route URI.
		Priority         int          // Just for reference.
		IsServiceHandler bool         // Is service handler.
	}

		// HandlerFunc是一个请求处理函数。 md5:9d90773cd75863ca
	HandlerFunc = func(r *Request)

		// handlerFuncInfo 包含了 HandlerFunc 的地址及其反射类型。 md5:32f3e1bf7321a5a1
	handlerFuncInfo struct {
		Func            HandlerFunc      // 处理器函数地址。 md5:a1fa398ec433900f
		Type            reflect.Type     // 用于处理程序特性的扩展，它反映了当前处理器的信息类型。 md5:4b1de13beafb74cf
		Value           reflect.Value    // 反射当前处理器的值信息，用于处理器功能的扩展。 md5:d3c9f0b7ed9ab350
		IsStrictRoute   bool             // 是否启用严格路由匹配。 md5:3b0165798f5d3455
		ReqStructFields []gstructs.Field // Request struct fields.
	}

	// HandlerItem是注册的路由处理程序，包括中间件和钩子函数。
	// md5:78b676e6e09329bb
	HandlerItem struct {
		// 唯一的处理器项标识。
		// 注意，同一个处理器函数可能会注册为不同的处理器项，它们具有不同的处理器项ID。
		// md5:7b474802a6e17d79
		Id         int
		Name       string          // 处理器名称，当注册时会自动从运行时堆栈中获取。 md5:72fae2285a3c4c69
		Type       HandlerType     // 处理器类型：对象/处理器/中间件/钩子。 md5:3f54f4463e6e5dc5
		Info       handlerFuncInfo // 处理器函数信息。 md5:471fa38cfb5ee901
		InitFunc   HandlerFunc     // 当请求进入对象时的初始化函数（仅适用于对象注册类型）。 md5:e9c9e4d3f3d42414
		ShutFunc   HandlerFunc     // 当请求删除对象时（仅适用于对象注册类型），调用的退出函数。 md5:b919a3d55a43043b
		Middleware []HandlerFunc   //绑定中间件数组。 md5:8fba709766af338f
		HookName   HookName        // 挂钩类型名称，仅对挂钩类型可用。 md5:13d0e45e1f8d3e9f
		Router     *Router         // Router object.
		Source     string          // 注册源文件`路径:行号`。 md5:681405429ed39e78
	}

		// HandlerItemParsed是从URL.Path解析出的项目。 md5:6e18c9d6fea2d3d1
	HandlerItemParsed struct {
		Handler *HandlerItem      // Handler information.
		Values  map[string]string // 从URL.Path中解析得到的路由值。 md5:f6d02a0cfbdeb7d3
	}

		// ServerStatus 是服务器状态的枚举类型。 md5:6de5e4d7f5fc52a6
	ServerStatus = int

		// HookName是路由钩子名称枚举类型。 md5:9e0295b925c0c40f
	HookName string

		// HandlerType 是路由处理器枚举类型。 md5:940090bf597744cc
	HandlerType string

	// 文件描述符映射的监听。
	// 键可以是 "http" 或 "https"，值则是对应的文件描述符（FD）。
	// md5:203cb0295faad7ff
	listenerFdMap = map[string]string

		// internalPanic 是用于内部使用的自定义恐慌。 md5:287806e552654f1d
	internalPanic string
)

const (
		// FreePortAddress 标记服务器使用随机的空闲端口进行监听。 md5:16e8ca0633c4a135
	FreePortAddress = ":0"
)

const (
	HeaderXUrlPath                     = "x-url-path"         // 用于自定义路由处理器，它不会改变URL.Path。 md5:50133279ebd50b30
	HookBeforeServe       HookName     = "HOOK_BEFORE_SERVE"  // 路由处理程序或文件服务之前的钩子处理器。 md5:cf35e94a36837514
	HookAfterServe        HookName     = "HOOK_AFTER_SERVE"   // 路由处理器或文件服务后的钩子处理程序。 md5:d236c9c1ff9f9847
	HookBeforeOutput      HookName     = "HOOK_BEFORE_OUTPUT" // 在输出响应之前处理挂钩。 md5:4d8f72553e739005
	HookAfterOutput       HookName     = "HOOK_AFTER_OUTPUT"  // 在响应输出后处理钩子。 md5:4a1980ddfff8dda6
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
	noPrintInternalRoute                    = "internalMiddlewareServerTracing"
)

const (
	exceptionExit     internalPanic = "exit"
	exceptionExitAll  internalPanic = "exit_all"
	exceptionExitHook internalPanic = "exit_hook"
)

var (
	// methodsMap 存储所有支持的HTTP方法。
	// 它用于通过映射快速搜索HTTP方法。
	// md5:d1f472198ffb7422
	methodsMap = make(map[string]struct{})

	// serverMapping 为当前进程存储多个服务器实例。
	// 键是服务器的名称，值是其对应的实例。
	// md5:ebea00ddd34b8a0e
	serverMapping = gmap.NewStrAnyMap(true)

	// serverRunning 标记运行中的服务器数量。
	// 如果没有成功运行的服务器或所有服务器都已关闭，该值为 0。
	// md5:5521930133e9bda1
	serverRunning = gtype.NewInt()

		// wsUpGrader是websocket的默认升级器配置。 md5:dcb5f656ab8a1a3a
	wsUpGrader = websocket.Upgrader{
				// 默认情况下，它不检查来源，应用程序可以自行执行这一操作。 md5:a508dbea94a89f04
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// allShutdownChan 是所有服务器完成服务并退出的事件。
	// 它用于进程阻塞的目的。
	// md5:7575f7dd8d471d7f
	allShutdownChan = make(chan struct{}, 1000)

	// serverProcessInitialized 用于服务器的懒惰初始化。进程只能初始化一次。
	// md5:34bc30c229d05a4a
	serverProcessInitialized = gtype.NewBool()

		// gracefulEnabled 用于启用优雅重启功能，该功能默认为 false。 md5:ff6dd18537c796d9
	gracefulEnabled = false

		// defaultValueTags 用于存储默认值的结构体标签名称。 md5:9d4fc272f2f20701
	defaultValueTags = []string{gtag.DefaultShort, gtag.Default}
)

var (
	ErrNeedJsonBody = gerror.NewWithOption(gerror.Option{
		Text: "the request body content should be JSON format",
		Code: gcode.CodeInvalidRequest,
	})
)
