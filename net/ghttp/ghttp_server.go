// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
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
	if !genv.Get(adminActionRestartEnvKey).IsEmpty() {
		if p, err := os.FindProcess(gproc.PPid()); err == nil {
			if err = p.Kill(); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
			if _, err = p.Wait(); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		} else {
			glog.Error(ctx, err)
		}
	}

// 处理消息处理器。
// 只有在启用了优雅特性时，它才会被启用。
	if gracefulEnabled {
		intlog.Printf(ctx, "pid[%d]: graceful reload feature is enabled", gproc.Pid())
		go handleProcessMessage()
	} else {
		intlog.Printf(ctx, "pid[%d]: graceful reload feature is disabled", gproc.Pid())
	}

// 这是一种用于在源代码开发环境中更好地初始化主包路径的丑陋调用方式。它仅在主goroutine中有用。
// 在异步goroutine中，该方法无法正确获取主包路径。
	gfile.MainPkgPath()
}

// GetServer 根据给定名称和默认配置创建并返回一个服务器实例。
// 注意，参数`name`对于不同服务器应保持唯一。如果给定的`name`已在服务器映射中存在，
// 则它将返回一个已存在的服务器实例。
func GetServer(name ...interface{}) *Server {
	serverName := DefaultServerName
	if len(name) > 0 && name[0] != "" {
		serverName = gconv.String(name[0])
	}
	v := serverMapping.GetOrSetFuncLock(serverName, func() interface{} {
		s := &Server{
			instance:         serverName,
			plugins:          make([]Plugin, 0),
			servers:          make([]*gracefulServer, 0),
			closeChan:        make(chan struct{}, 10000),
			serverCount:      gtype.NewInt(),
			statusHandlerMap: make(map[string][]HandlerFunc),
			serveTree:        make(map[string]interface{}),
			serveCache:       gcache.New(),
			routesMap:        make(map[string][]*HandlerItem),
			openapi:          goai.New(),
			registrar:        gsvc.GetRegistry(),
		}
		// 使用默认配置初始化服务器。
		if err := s.SetConfig(NewConfig()); err != nil {
			panic(gerror.WrapCode(gcode.CodeInvalidConfiguration, err, ""))
		}
		// 它默认为服务器启用OpenTelemetry。
		s.Use(internalMiddlewareServerTracing)
		return s
	})
	return v.(*Server)
}

// Start 开始监听配置好的端口。
// 该函数不会阻塞进程，你可以使用函数 Wait 来阻塞进程。
func (s *Server) Start() error {
	var ctx = gctx.GetInitCtx()

	// Swagger UI.
	if s.config.SwaggerPath != "" {
		swaggerui.Init()
		s.AddStaticPath(s.config.SwaggerPath, swaggerUIPackedPath)
		s.BindHookHandler(s.config.SwaggerPath+"/*", HookBeforeServe, s.swaggerUI)
	}

	// OpenApi规范JSON生成处理器。
	if s.config.OpenApiPath != "" {
		s.BindHandler(s.config.OpenApiPath, s.openapiSpec)
	}

	// 注册群组路由。
	s.handlePreBindItems(ctx)

	// 服务器进程初始化，只能初始化一次。
	serverProcessInit()

	// 服务只能运行一次。
	if s.Status() == ServerStatusRunning {
		return gerror.NewCode(gcode.CodeInvalidOperation, "server is already running")
	}

	// 日志路径设置检查
	if s.config.LogPath != "" && s.config.LogPath != s.config.Logger.GetPath() {
		if err := s.config.Logger.SetPath(s.config.LogPath); err != nil {
			return err
		}
	}
	// 默认的会话存储。
	if s.config.SessionStorage == nil {
		sessionStoragePath := ""
		if s.config.SessionPath != "" {
			sessionStoragePath = gfile.Join(s.config.SessionPath, s.config.Name)
			if !gfile.Exists(sessionStoragePath) {
				if err := gfile.Mkdir(sessionStoragePath); err != nil {
					return gerror.Wrapf(err, `mkdir failed for "%s"`, sessionStoragePath)
				}
			}
		}
		s.config.SessionStorage = gsession.NewStorageFile(sessionStoragePath, s.config.SessionMaxAge)
	}
	// 在程序启动运行时初始化会话管理器。
	s.sessionManager = gsession.New(
		s.config.SessionMaxAge,
		s.config.SessionStorage,
	)

	// PProf feature.
	if s.config.PProfEnabled {
		s.EnablePProf(s.config.PProfPattern)
	}

	// 默认HTTP处理器
	if s.config.Handler == nil {
		s.config.Handler = s.ServeHTTP
	}

	// 安装外部插件。
	for _, p := range s.plugins {
		if err := p.Install(s); err != nil {
			s.Logger().Fatalf(ctx, `%+v`, err)
		}
	}
	// 再次检查组路由中内部注册的路由。
	s.handlePreBindItems(ctx)

// 如果没有注册路由且未启用静态服务，
// 则返回服务器使用无效的错误。
	if len(s.routesMap) == 0 && !s.config.FileServerEnabled {
		return gerror.NewCode(
			gcode.CodeInvalidOperation,
			`there's no route set or static feature enabled, did you forget import the router?`,
		)
	}
// ================================================================================================
// 启动HTTP服务器。
// ================================================================================================
	reloaded := false
	fdMapStr := genv.Get(adminActionReloadEnvKey).String()
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
		s.Logger().Infof(
			ctx,
			`swagger ui is serving at address: %s%s/`,
			s.getLocalListenedAddress(),
			s.config.SwaggerPath,
		)
	}
	// OpenApi规范信息
	if s.config.OpenApiPath != "" {
		s.Logger().Infof(
			ctx,
			`openapi specification is serving at address: %s%s`,
			s.getLocalListenedAddress(),
			s.config.OpenApiPath,
		)
	} else {
		if s.config.SwaggerPath != "" {
			s.Logger().Warning(
				ctx,
				`openapi specification is disabled but swagger ui is serving, which might make no sense`,
			)
		} else {
			s.Logger().Info(
				ctx,
				`openapi specification is disabled`,
			)
		}
	}

	// 如果这是一个子进程，那么它会通知其父进程已退出。
	if gproc.IsChild() {
		gtimer.SetTimeout(ctx, time.Duration(s.config.GracefulTimeout)*time.Second, func(ctx context.Context) {
			if err := gproc.Send(gproc.PPid(), []byte("exit"), adminGProcCommGroup); err != nil {
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
	return fmt.Sprintf(`http://127.0.0.1:%d`, s.GetListenedPort())
}

// doRouterMapDump 检查并把路由映射表转储到日志中。
func (s *Server) doRouterMapDump() {
	if !s.config.DumpRouterMap {
		return
	}

	var (
		ctx                          = context.TODO()
		routes                       = s.GetRoutes()
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
				handlerName = gstr.TrimRightStr(item.Handler.Name, "-fm")
				middlewares = gstr.SplitAndTrim(item.Middleware, ",")
			)
			for k, v := range middlewares {
				middlewares[k] = gstr.TrimRightStr(v, "-fm")
			}
			item.Middleware = gstr.Join(middlewares, "\n")
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
		s.config.Logger.Header(false).Printf(ctx, "\n%s", buffer.String())
	}
}

// GetOpenApi 返回当前服务器的OpenApi规范管理对象。
func (s *Server) GetOpenApi() *goai.OpenApiV3 {
	return s.openapi
}

// GetRoutes 获取并返回路由数组。
func (s *Server) GetRoutes() []RouterItem {
	var (
		m              = make(map[string]*garray.SortedArray)
		routeFilterSet = gset.NewStrSet()
		address        = s.GetListenedAddress()
	)
	if s.config.HTTPSAddr != "" {
		if len(address) > 0 {
			address += ","
		}
		address += "tls" + s.config.HTTPSAddr
	}
	for k, handlerItems := range s.routesMap {
		array, _ := gregex.MatchString(`(.*?)%([A-Z]+):(.+)@(.+)`, k)
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
			if !routeFilterSet.AddIfNotExist(setKey) {
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
				m[item.Domain] = garray.NewSortedArray(func(v1, v2 interface{}) int {
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
			m[item.Domain].Add(item)
		}
	}

	routerArray := make([]RouterItem, 0, 128)
	for _, array := range m {
		for _, v := range array.Slice() {
			routerArray = append(routerArray, v.(RouterItem))
		}
	}
	return routerArray
}

// Run 启动服务器并以阻塞方式监听。
// 该方法通常用于单服务器场景。
func (s *Server) Run() {
	var ctx = context.TODO()

	if err := s.Start(); err != nil {
		s.Logger().Fatalf(ctx, `%+v`, err)
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
	s.Logger().Infof(ctx, "pid[%d]: all servers shutdown", gproc.Pid())
}

// Wait 阻塞等待所有服务器完成。
// 在多服务器场景中，它通常被使用。
func Wait() {
	var ctx = context.TODO()

	// 以异步方式处理信号。
	go handleProcessSignal()

	<-allShutdownChan

	// Remove plugins.
	serverMapping.Iterator(func(k string, v interface{}) bool {
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
	glog.Infof(ctx, "pid[%d]: all servers shutdown", gproc.Pid())
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
					fd = gconv.Int(addrAndFd[1])
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
		array = gstr.SplitAndTrim(v, ",")
	} else {
		array = gstr.SplitAndTrim(s.config.Address, ",")
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
				fd = gconv.Int(addrAndFd[1])
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
				s.Logger().Fatalf(ctx, `%+v`, err)
			}
			wg.Done()
			// 开始监听并以阻塞方式提供服务。
			err = server.Serve(ctx)
			// 如果服务器在没有关闭错误的情况下被关闭，进程将退出。
			if err != nil && !strings.EqualFold(http.ErrServerClosed.Error(), err.Error()) {
				s.Logger().Fatalf(ctx, `%+v`, err)
			}
			// 如果所有底层服务器都关闭，进程将退出。
			if s.serverCount.Add(-1) < 1 {
				s.closeChan <- struct{}{}
				if serverRunning.Add(-1) < 1 {
					serverMapping.Remove(s.instance)
					allShutdownChan <- struct{}{}
				}
			}
		}(v)
	}
	wg.Wait()
}

// Status 获取并返回服务器状态。
func (s *Server) Status() ServerStatus {
	if serverRunning.Val() == 0 {
		return ServerStatusStopped
	}
	// 如果任何底层服务器正在运行，则服务器状态为运行中。
	for _, v := range s.servers {
		if v.status.Val() == ServerStatusRunning {
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
		str := v.address + "#" + gconv.String(v.Fd()) + ","
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
func (s *Server) GetListenedPort() int {
	ports := s.GetListenedPorts()
	if len(ports) > 0 {
		return ports[0]
	}
	return 0
}

// GetListenedPorts 获取并返回当前服务器正在监听的所有端口。
func (s *Server) GetListenedPorts() []int {
	ports := make([]int, 0)
	for _, server := range s.servers {
		ports = append(ports, server.GetListenedPort())
	}
	return ports
}

// GetListenedAddress 获取并返回当前服务器监听的地址字符串。
func (s *Server) GetListenedAddress() string {
	if !gstr.Contains(s.config.Address, FreePortAddress) {
		return s.config.Address
	}
	var (
		address       = s.config.Address
		listenedPorts = s.GetListenedPorts()
	)
	for _, listenedPort := range listenedPorts {
		address = gstr.Replace(address, FreePortAddress, fmt.Sprintf(`:%d`, listenedPort), 1)
	}
	return address
}
