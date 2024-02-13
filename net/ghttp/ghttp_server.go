// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
	
	"github.com/olekukonko/tablewriter"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/ghttp/internal/swaggerui"
	"github.com/888go/goframe/net/goai"
	"github.com/888go/goframe/net/gsvc"
	"github.com/888go/goframe/os/gcache"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/genv"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gproc"
	"github.com/888go/goframe/os/gsession"
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

func init() {
	// 初始化方法映射。
	for _, v := range strings.Split(supportedHttpMethods, ",") {
		methodsMap[v] = struct{}{}
	}
}

// serverProcessInit 初始化一些只能执行一次的进程配置。
func serverProcessInit() {
	var ctx = context.TODO()
	if !serverProcessInitialized.Cas(false, true) {
		return
	}
// 这意味着它是一个重启服务。在开始监听之前，它应该先终止其父进程，
// 以避免在两个进程中因端口重复监听而产生问题。
	if !环境变量类.X取值(adminActionRestartEnvKey).X是否为空() {
		if p, err := os.FindProcess(进程类.PPid()); err == nil {
			if err = p.Kill(); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
			if _, err = p.Wait(); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		} else {
			日志类.Error(ctx, err)
		}
	}

// 处理消息处理器。
// 只有在启用了优雅特性时，它才会被启用。
	if gracefulEnabled {
		intlog.Printf(ctx, "pid[%d]: graceful reload feature is enabled", 进程类.Pid())
		go handleProcessMessage()
	} else {
		intlog.Printf(ctx, "pid[%d]: graceful reload feature is disabled", 进程类.Pid())
	}

// 这是一种用于在源代码开发环境中更好地初始化主包路径的丑陋调用方式。它仅在主goroutine中有用。
// 在异步goroutine中，该方法无法正确获取主包路径。
	文件类.X取main路径()
}

// GetServer 根据给定名称和默认配置创建并返回一个服务器实例。
// 注意，参数`name`对于不同服务器应保持唯一。如果给定的`name`已在服务器映射中存在，
// 则它将返回一个已存在的服务器实例。
func X取服务对象(名称 ...interface{}) *Server {
	serverName := DefaultServerName
	if len(名称) > 0 && 名称[0] != "" {
		serverName = 转换类.String(名称[0])
	}
	v := serverMapping.X取值或设置值_函数带锁(serverName, func() interface{} {
		s := &Server{
			instance:         serverName,
			plugins:          make([]Plugin, 0),
			servers:          make([]*gracefulServer, 0),
			closeChan:        make(chan struct{}, 10000),
			serverCount:      安全变量类.NewInt(),
			statusHandlerMap: make(map[string][]HandlerFunc),
			serveTree:        make(map[string]interface{}),
			serveCache:       缓存类.X创建(),
			routesMap:        make(map[string][]*HandlerItem),
			openapi:          goai.New(),
			registrar:        gsvc.GetRegistry(),
		}
		// 使用默认配置初始化服务器。
		if err := s.X设置配置项(X创建默认配置项()); err != nil {
			panic(错误类.X多层错误码(错误码类.CodeInvalidConfiguration, err, ""))
		}
		// 它默认为服务器启用OpenTelemetry。
		s.Use别名(internalMiddlewareServerTracing)
		return s
	})
	return v.(*Server)
}

// Start 开始监听配置好的端口。
// 该函数不会阻塞进程，你可以使用函数 Wait 来阻塞进程。
func (s *Server) X开始监听() error {
	var ctx = 上下文类.X取初始化上下文()

	// Swagger UI.
	if s.config.SwaggerPath != "" {
		swaggerui.Init()
		s.X静态文件添加目录映射(s.config.SwaggerPath, swaggerUIPackedPath)
		s.X绑定Hook(s.config.SwaggerPath+"/*", HookBeforeServe, s.swaggerUI)
	}

	// OpenApi规范JSON生成处理器。
	if s.config.OpenApiPath != "" {
		s.X绑定(s.config.OpenApiPath, s.openapiSpec)
	}

	// 注册群组路由。
	s.handlePreBindItems(ctx)

	// 服务器进程初始化，只能初始化一次。
	serverProcessInit()

	// 服务只能运行一次。
	if s.X取服务状态() == ServerStatusRunning {
		return 错误类.X创建错误码(错误码类.CodeInvalidOperation, "server is already running")
	}

	// 日志路径设置检查
	if s.config.LogPath != "" && s.config.LogPath != s.config.Logger.X取文件路径() {
		if err := s.config.Logger.X设置文件路径(s.config.LogPath); err != nil {
			return err
		}
	}
	// 默认的会话存储。
	if s.config.SessionStorage == nil {
		sessionStoragePath := ""
		if s.config.SessionPath != "" {
			sessionStoragePath = 文件类.X路径生成(s.config.SessionPath, s.config.Name)
			if !文件类.X是否存在(sessionStoragePath) {
				if err := 文件类.X创建目录(sessionStoragePath); err != nil {
					return 错误类.X多层错误并格式化(err, `mkdir failed for "%s"`, sessionStoragePath)
				}
			}
		}
		s.config.SessionStorage = session类.NewStorageFile(sessionStoragePath, s.config.SessionMaxAge)
	}
	// 在程序启动运行时初始化会话管理器。
	s.sessionManager = session类.New(
		s.config.SessionMaxAge,
		s.config.SessionStorage,
	)

	// PProf feature.
	if s.config.PProfEnabled {
		s.PProf开启(s.config.PProfPattern)
	}

	// 默认HTTP处理器
	if s.config.Handler == nil {
		s.config.Handler = s.ServeHTTP
	}

	// 安装外部插件。
	for _, p := range s.plugins {
		if err := p.Install(s); err != nil {
			s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
		}
	}
	// 再次检查组路由中内部注册的路由。
	s.handlePreBindItems(ctx)

// 如果没有注册路由且未启用静态服务，
// 则返回服务器使用无效的错误。
	if len(s.routesMap) == 0 && !s.config.FileServerEnabled {
		return 错误类.X创建错误码(
			错误码类.CodeInvalidOperation,
			`there's no route set or static feature enabled, did you forget import the router?`,
		)
	}
// ================================================================================================
// 启动HTTP服务器。
// ================================================================================================
	reloaded := false
	fdMapStr := 环境变量类.X取值(adminActionReloadEnvKey).String()
	if len(fdMapStr) > 0 {
		sfm := bufferToServerFdMap([]byte(fdMapStr))
		if v, ok := sfm[s.config.Name]; ok {
			s.startServer(v)
			reloaded = true
		}
	}
	if !reloaded {
		s.startServer(nil)
	}

	// Swagger UI info.
	if s.config.SwaggerPath != "" {
		s.Logger别名().X输出并格式化INFO(
			ctx,
			`swagger ui is serving at address: %s%s/`,
			s.getLocalListenedAddress(),
			s.config.SwaggerPath,
		)
	}
	// OpenApi规范信息
	if s.config.OpenApiPath != "" {
		s.Logger别名().X输出并格式化INFO(
			ctx,
			`openapi specification is serving at address: %s%s`,
			s.getLocalListenedAddress(),
			s.config.OpenApiPath,
		)
	} else {
		if s.config.SwaggerPath != "" {
			s.Logger别名().X输出WARN(
				ctx,
				`openapi specification is disabled but swagger ui is serving, which might make no sense`,
			)
		} else {
			s.Logger别名().X输出INFO(
				ctx,
				`openapi specification is disabled`,
			)
		}
	}

	// 如果这是一个子进程，那么它会通知其父进程已退出。
	if 进程类.IsChild() {
		定时类.SetTimeout别名(ctx, time.Duration(s.config.GracefulTimeout)*time.Second, func(ctx context.Context) {
			if err := 进程类.Send(进程类.PPid(), []byte("exit"), adminGProcCommGroup); err != nil {
				intlog.Errorf(ctx, `server error in process communication: %+v`, err)
			}
		})
	}
	s.initOpenApi()
	s.doServiceRegister()
	s.doRouterMapDump()

	return nil
}

func (s *Server) getLocalListenedAddress() string {
	return fmt.Sprintf(`http://127.0.0.1:%d`, s.X取已监听端口())
}

// doRouterMapDump 检查并把路由映射表转储到日志中。
func (s *Server) doRouterMapDump() {
	if !s.config.DumpRouterMap {
		return
	}

	var (
		ctx                          = context.TODO()
		routes                       = s.X取路由数组()
		isJustDefaultServerAndDomain = true
		headers                      = []string{
			"SERVER", "DOMAIN", "ADDRESS", "METHOD", "ROUTE", "HANDLER", "MIDDLEWARE",
		}
	)
	for _, item := range routes {
		if item.Server != DefaultServerName || item.Domain != DefaultDomainName {
			isJustDefaultServerAndDomain = false
			break
		}
	}
	if isJustDefaultServerAndDomain {
		headers = []string{"ADDRESS", "METHOD", "ROUTE", "HANDLER", "MIDDLEWARE"}
	}
	if len(routes) > 0 {
		buffer := bytes.NewBuffer(nil)
		table := tablewriter.NewWriter(buffer)
		table.SetHeader(headers)
		table.SetRowLine(true)
		table.SetBorder(false)
		table.SetCenterSeparator("|")

		for _, item := range routes {
			var (
				data        = make([]string, 0)
				handlerName = 文本类.X过滤尾字符(item.Handler.Name, "-fm")
				middlewares = 文本类.X分割并忽略空值(item.Middleware, ",")
			)
			for k, v := range middlewares {
				middlewares[k] = 文本类.X过滤尾字符(v, "-fm")
			}
			item.Middleware = 文本类.X连接(middlewares, "\n")
			if isJustDefaultServerAndDomain {
				data = append(
					data,
					item.Address,
					item.Method,
					item.Route,
					handlerName,
					item.Middleware,
				)
			} else {
				data = append(
					data,
					item.Server,
					item.Domain,
					item.Address,
					item.Method,
					item.Route,
					handlerName,
					item.Middleware,
				)
			}
			table.Append(data)
		}
		table.Render()
		s.config.Logger.X是否输出头信息(false).X输出并格式化(ctx, "\n%s", buffer.String())
	}
}

// GetOpenApi 返回当前服务器的OpenApi规范管理对象。
func (s *Server) X取OpenApi对象() *goai.OpenApiV3 {
	return s.openapi
}

// GetRoutes 获取并返回路由数组。
func (s *Server) X取路由数组() []RouterItem {
	var (
		m              = make(map[string]*数组类.SortedArray)
		routeFilterSet = 集合类.X创建文本()
		address        = s.X取已监听地址()
	)
	if s.config.HTTPSAddr != "" {
		if len(address) > 0 {
			address += ","
		}
		address += "tls" + s.config.HTTPSAddr
	}
	for k, handlerItems := range s.routesMap {
		array, _ := 正则类.X匹配文本(`(.*?)%([A-Z]+):(.+)@(.+)`, k)
		for index := len(handlerItems) - 1; index >= 0; index-- {
			var (
				handlerItem = handlerItems[index]
				item        = RouterItem{
					Server:     s.config.Name,
					Address:    address,
					Domain:     array[4],
					Type:       handlerItem.Type,
					Middleware: array[1],
					Method:     array[2],
					Route:      array[3],
					Priority:   index,
					Handler:    handlerItem,
				}
			)
			switch item.Handler.Type {
			case HandlerTypeObject, HandlerTypeHandler:
				item.IsServiceHandler = true

			case HandlerTypeMiddleware:
				item.Middleware = "GLOBAL MIDDLEWARE"
			}
			// 重复路径过滤以供转储
			var setKey = fmt.Sprintf(
				`%s|%s|%s|%s`,
				item.Method, item.Route, item.Domain, item.Type,
			)
			if !routeFilterSet.X加入值并跳过已存在(setKey) {
				continue
			}
			if len(item.Handler.Middleware) > 0 {
				for _, v := range item.Handler.Middleware {
					if item.Middleware != "" {
						item.Middleware += ","
					}
					item.Middleware += gdebug.FuncName(v)
				}
			}
// 如果域名不存在于dump映射中，则创建该映射。
// 映射的值是一个自定义排序的数组。
			if _, ok := m[item.Domain]; !ok {
				// 按升序排序
				m[item.Domain] = 数组类.X创建排序(func(v1, v2 interface{}) int {
					item1 := v1.(RouterItem)
					item2 := v2.(RouterItem)
					r := 0
					if r = strings.Compare(item1.Domain, item2.Domain); r == 0 {
						if r = strings.Compare(item1.Route, item2.Route); r == 0 {
							if r = strings.Compare(item1.Method, item2.Method); r == 0 {
								if item1.Handler.Type == HandlerTypeMiddleware && item2.Handler.Type != HandlerTypeMiddleware {
									return -1
								} else if item1.Handler.Type == HandlerTypeMiddleware && item2.Handler.Type == HandlerTypeMiddleware {
									return 1
								} else if r = strings.Compare(item1.Middleware, item2.Middleware); r == 0 {
									r = item2.Priority - item1.Priority
								}
							}
						}
					}
					return r
				})
			}
			m[item.Domain].X入栈右(item)
		}
	}

	routerArray := make([]RouterItem, 0, 128)
	for _, array := range m {
		for _, v := range array.X取切片() {
			routerArray = append(routerArray, v.(RouterItem))
		}
	}
	return routerArray
}

// Run 启动服务器并以阻塞方式监听。
// 该方法通常用于单服务器场景。
func (s *Server) X启动服务() {
	var ctx = context.TODO()

	if err := s.X开始监听(); err != nil {
		s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
	}

	// 以异步方式处理信号。
	go handleProcessSignal()

	// 使用通道进行阻塞，实现优雅重启。
	<-s.closeChan
	// Remove plugins.
	if len(s.plugins) > 0 {
		for _, p := range s.plugins {
			intlog.Printf(ctx, `remove plugin: %s`, p.Name())
			if err := p.Remove(); err != nil {
				intlog.Errorf(ctx, "%+v", err)
			}
		}
	}
	s.doServiceDeregister()
	s.Logger别名().X输出并格式化INFO(ctx, "pid[%d]: all servers shutdown", 进程类.Pid())
}

// Wait 阻塞等待所有服务器完成。
// 在多服务器场景中，它通常被使用。
func X等待所有服务完成() {
	var ctx = context.TODO()

	// 以异步方式处理信号。
	go handleProcessSignal()

	<-allShutdownChan

	// Remove plugins.
	serverMapping.X遍历(func(k string, v interface{}) bool {
		s := v.(*Server)
		if len(s.plugins) > 0 {
			for _, p := range s.plugins {
				intlog.Printf(ctx, `remove plugin: %s`, p.Name())
				if err := p.Remove(); err != nil {
					intlog.Errorf(ctx, `%+v`, err)
				}
			}
		}
		return true
	})
	日志类.X输出并格式化INFO(ctx, "pid[%d]: all servers shutdown", 进程类.Pid())
}

// startServer 启动底层服务器并开始监听。
func (s *Server) startServer(fdMap listenerFdMap) {
	var (
		ctx          = context.TODO()
		httpsEnabled bool
	)
	// HTTPS
	if s.config.TLSConfig != nil || (s.config.HTTPSCertPath != "" && s.config.HTTPSKeyPath != "") {
		if len(s.config.HTTPSAddr) == 0 {
			if len(s.config.Address) > 0 {
				s.config.HTTPSAddr = s.config.Address
				s.config.Address = ""
			} else {
				s.config.HTTPSAddr = defaultHttpsAddr
			}
		}
		httpsEnabled = len(s.config.HTTPSAddr) > 0
		var array []string
		if v, ok := fdMap["https"]; ok && len(v) > 0 {
			array = strings.Split(v, ",")
		} else {
			array = strings.Split(s.config.HTTPSAddr, ",")
		}
		for _, v := range array {
			if len(v) == 0 {
				continue
			}
			var (
				fd        = 0
				itemFunc  = v
				addrAndFd = strings.Split(v, "#")
			)
			if len(addrAndFd) > 1 {
				itemFunc = addrAndFd[0]
// Windows操作系统不支持从父进程传递socket文件描述符。
				if runtime.GOOS != "windows" {
					fd = 转换类.X取整数(addrAndFd[1])
				}
			}
			if fd > 0 {
				s.servers = append(s.servers, s.newGracefulServer(itemFunc, fd))
			} else {
				s.servers = append(s.servers, s.newGracefulServer(itemFunc))
			}
			s.servers[len(s.servers)-1].isHttps = true
		}
	}
	// HTTP
	if !httpsEnabled && len(s.config.Address) == 0 {
		s.config.Address = defaultHttpAddr
	}
	var array []string
	if v, ok := fdMap["http"]; ok && len(v) > 0 {
		array = 文本类.X分割并忽略空值(v, ",")
	} else {
		array = 文本类.X分割并忽略空值(s.config.Address, ",")
	}
	for _, v := range array {
		if len(v) == 0 {
			continue
		}
		var (
			fd        = 0
			itemFunc  = v
			addrAndFd = strings.Split(v, "#")
		)
		if len(addrAndFd) > 1 {
			itemFunc = addrAndFd[0]
// Windows 操作系统不支持从父进程传递套接字文件描述符。
			if runtime.GOOS != "windows" {
				fd = 转换类.X取整数(addrAndFd[1])
			}
		}
		if fd > 0 {
			s.servers = append(s.servers, s.newGracefulServer(itemFunc, fd))
		} else {
			s.servers = append(s.servers, s.newGracefulServer(itemFunc))
		}
	}
	// 开始异步监听。
	serverRunning.Add(1)
	var wg = sync.WaitGroup{}
	for _, v := range s.servers {
		wg.Add(1)
		go func(server *gracefulServer) {
			s.serverCount.Add(1)
			var err error
			// Create listener.
			if server.isHttps {
				err = server.CreateListenerTLS(
					s.config.HTTPSCertPath, s.config.HTTPSKeyPath, s.config.TLSConfig,
				)
			} else {
				err = server.CreateListener()
			}
			if err != nil {
				s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
			}
			wg.Done()
			// 开始监听并以阻塞方式提供服务。
			err = server.Serve(ctx)
			// 如果服务器在没有关闭错误的情况下被关闭，进程将退出。
			if err != nil && !strings.EqualFold(http.ErrServerClosed.Error(), err.Error()) {
				s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
			}
			// 如果所有底层服务器都关闭，进程将退出。
			if s.serverCount.Add(-1) < 1 {
				s.closeChan <- struct{}{}
				if serverRunning.Add(-1) < 1 {
					serverMapping.X删除(s.instance)
					allShutdownChan <- struct{}{}
				}
			}
		}(v)
	}
	wg.Wait()
}

// Status 获取并返回服务器状态。
func (s *Server) X取服务状态() ServerStatus {
	if serverRunning.X取值() == 0 {
		return ServerStatusStopped
	}
	// 如果任何底层服务器正在运行，则服务器状态为运行中。
	for _, v := range s.servers {
		if v.status.X取值() == ServerStatusRunning {
			return ServerStatusRunning
		}
	}
	return ServerStatusStopped
}

// getListenerFdMap 获取并返回套接字文件描述符的映射。
// 返回映射中的键为 "http" 和 "https"。
func (s *Server) getListenerFdMap() map[string]string {
	m := map[string]string{
		"https": "",
		"http":  "",
	}
	for _, v := range s.servers {
		str := v.address + "#" + 转换类.String(v.Fd()) + ","
		if v.isHttps {
			if len(m["https"]) > 0 {
				m["https"] += ","
			}
			m["https"] += str
		} else {
			if len(m["http"]) > 0 {
				m["http"] += ","
			}
			m["http"] += str
		}
	}
	return m
}

// GetListenedPort 获取并返回当前服务器正在监听的一个端口。
func (s *Server) X取已监听端口() int {
	ports := s.X取所有监听已端口()
	if len(ports) > 0 {
		return ports[0]
	}
	return 0
}

// GetListenedPorts 获取并返回当前服务器正在监听的所有端口。
func (s *Server) X取所有监听已端口() []int {
	ports := make([]int, 0)
	for _, server := range s.servers {
		ports = append(ports, server.GetListenedPort())
	}
	return ports
}

// GetListenedAddress 获取并返回当前服务器监听的地址字符串。
func (s *Server) X取已监听地址() string {
	if !文本类.X是否包含(s.config.Address, FreePortAddress) {
		return s.config.Address
	}
	var (
		address       = s.config.Address
		listenedPorts = s.X取所有监听已端口()
	)
	for _, listenedPort := range listenedPorts {
		address = 文本类.X替换(address, FreePortAddress, fmt.Sprintf(`:%d`, listenedPort), 1)
	}
	return address
}
