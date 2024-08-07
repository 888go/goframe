// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

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

	garray "github.com/888go/goframe/container/garray"
	gset "github.com/888go/goframe/container/gset"
	gtype "github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/debug/gdebug"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/ghttp/internal/swaggerui"
	"github.com/888go/goframe/net/goai"
	"github.com/888go/goframe/net/gsvc"
	gcache "github.com/888go/goframe/os/gcache"
	gctx "github.com/888go/goframe/os/gctx"
	genv "github.com/888go/goframe/os/genv"
	gfile "github.com/888go/goframe/os/gfile"
	glog "github.com/888go/goframe/os/glog"
	gproc "github.com/888go/goframe/os/gproc"
	gsession "github.com/888go/goframe/os/gsession"
	gtimer "github.com/888go/goframe/os/gtimer"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

func init() {
		// 初始化方法映射。 md5:19df407918075ac2
	for _, v := range strings.Split(supportedHttpMethods, ",") {
		methodsMap[v] = struct{}{}
	}
}

// serverProcessInit 初始化一些进程配置，这些只能做一次。 md5:768ec2687e3a0f24
func serverProcessInit() {
	var ctx = context.TODO()
	if !serverProcessInitialized.Cas(false, true) {
		return
	}
	// 这意味着它是一个重启服务器。在开始监听之前，它应该杀死其父进程，以防止两个进程占用同一个端口进行监听。
	// md5:534f767682eae0c3
	if !genv.X取值(adminActionRestartEnvKey).X是否为空() {
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

	// 处理消息的处理器。
	// 仅当启用了优雅特性时，它才可用。
	// md5:b41176de77cc4833
	if gracefulEnabled {
		intlog.Printf(ctx, "pid[%d]: graceful reload feature is enabled", gproc.Pid())
		go handleProcessMessage()
	} else {
		intlog.Printf(ctx, "pid[%d]: graceful reload feature is disabled", gproc.Pid())
	}

	// 这是一种在源代码开发环境中更好地初始化主包路径的不优雅的方式。
	// 它只应在主 goroutine 中使用。
	// 在异步 goroutines 中无法获取主包路径。
	// md5:e0733b122721224e
	gfile.X取main路径()
}

// X取服务对象 使用给定名称和默认配置创建并返回一个服务器实例。
// 注意，参数 `name` 应该在不同的服务器中是唯一的。如果给定的 `name` 已经存在于服务器映射中，它将返回现有的服务器实例。
// md5:ad04664fa9750188
func X取服务对象(名称 ...interface{}) *X服务 {
	serverName := DefaultServerName
	if len(名称) > 0 && 名称[0] != "" {
		serverName = gconv.String(名称[0])
	}
	v := serverMapping.X取值或设置值_函数带锁(serverName, func() interface{} {
		s := &X服务{
			instance:         serverName,
			plugins:          make([]Plugin, 0),
			servers:          make([]*gracefulServer, 0),
			closeChan:        make(chan struct{}, 10000),
			serverCount:      gtype.NewInt(),
			statusHandlerMap: make(map[string][]HandlerFunc),
			serveTree:        make(map[string]interface{}),
			serveCache:       gcache.X创建(),
			routesMap:        make(map[string][]*X路由处理函数),
			openapi:          goai.New(),
			registrar:        gsvc.GetRegistry(),
		}
				// 使用默认配置初始化服务器。 md5:ac8ad35c2e6592fb
		if err := s.X设置配置项(X创建默认配置项()); err != nil {
			panic(gerror.X多层错误码(gcode.CodeInvalidConfiguration, err, ""))
		}
				// 它默认为服务器启用OpenTelemetry。 md5:2a2de55e6612dec7
		s.Use别名(internalMiddlewareServerTracing)
		return s
	})
	return v.(*X服务)
}

// X开始监听 开始在配置的端口上监听。
// 此函数不会阻塞进程，你可以使用 Wait 函数来阻塞进程。
// md5:05c1c66553fa4a61
func (s *X服务) X开始监听() error {
	var ctx = gctx.X取初始化上下文()

	// Swagger UI.
	if s.config.SwaggerPath != "" {
		swaggerui.Init()
		s.X静态文件添加目录映射(s.config.SwaggerPath, swaggerUIPackedPath)
		s.X绑定Hook(s.config.SwaggerPath+"/*", HookBeforeServe, s.swaggerUI)
	}

		// 开放API规范生成JSON处理器。 md5:62cb20bceb4ec15e
	if s.config.OpenApiPath != "" {
		s.X绑定(s.config.OpenApiPath, s.openapiSpec)
	}

	// Register group routes.
	s.handlePreBindItems(ctx)

		// 服务器进程初始化，这只能初始化一次。 md5:f4fd6ab84839bb71
	serverProcessInit()

		// 服务器只能运行一次。 md5:4372da1cc9e271f0
	if s.X取服务状态() == ServerStatusRunning {
		return gerror.X创建错误码(gcode.CodeInvalidOperation, "server is already running")
	}

		// 日志记录路径设置检查。 md5:b1b53a71404f3b28
	if s.config.LogPath != "" && s.config.LogPath != s.config.Logger.X取文件路径() {
		if err := s.config.Logger.X设置文件路径(s.config.LogPath); err != nil {
			return err
		}
	}
		// 默认会话存储。 md5:e5e4d66ee85c002a
	if s.config.SessionStorage == nil {
		sessionStoragePath := ""
		if s.config.SessionPath != "" {
			sessionStoragePath = gfile.X路径生成(s.config.SessionPath, s.config.Name)
			if !gfile.X是否存在(sessionStoragePath) {
				if err := gfile.X创建目录(sessionStoragePath); err != nil {
					return gerror.X多层错误并格式化(err, `mkdir failed for "%s"`, sessionStoragePath)
				}
			}
		}
		s.config.SessionStorage = gsession.NewStorageFile(sessionStoragePath, s.config.SessionMaxAge)
	}
		// 在启动时初始化会话管理器。 md5:060e4cc6f4f8b93e
	s.sessionManager = gsession.New(
		s.config.SessionMaxAge,
		s.config.SessionStorage,
	)

	// PProf feature.
	if s.config.PProfEnabled {
		s.PProf开启(s.config.PProfPattern)
	}

	// Default HTTP handler.
	if s.config.Handler == nil {
		s.config.Handler = s.ServeHTTP
	}

		// 安装外部插件。 md5:5a986e9f0fb84368
	for _, p := range s.plugins {
		if err := p.Install(s); err != nil {
			s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
		}
	}
		// 检查内部注册的路由再次应用于组路由。 md5:7949c3fe59e30c8c
	s.handlePreBindItems(ctx)

	// 如果没有注册路由且没有启用静态服务，它将返回服务器使用无效的错误。
	// md5:d916b25cf4c384d4
	if len(s.routesMap) == 0 && !s.config.FileServerEnabled {
		return gerror.X创建错误码(
			gcode.CodeInvalidOperation,
			`there's no route set or static feature enabled, did you forget import the router?`,
		)
	}
	// ================================================================================================
	// 启动HTTP服务器。
	// ================================================================================================
	// md5:9c551ab8996cea5a
	reloaded := false
	fdMapStr := genv.X取值(adminActionReloadEnvKey).String()
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
		// OpenAPI规范信息。 md5:e99b4557db364598
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

		// 如果该进程是子进程，那么它将通知其父进程退出。 md5:c9a0c66193cbdbfc
	if gproc.IsChild() {
		gtimer.SetTimeout别名(ctx, time.Duration(s.config.GracefulTimeout)*time.Second, func(ctx context.Context) {
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

func (s *X服务) getLocalListenedAddress() string {
	return fmt.Sprintf(`http://127.0.0.1:%d`, s.X取已监听端口())
}

// doRouterMapDump 检查并把路由映射dump到日志中。 md5:658175121fd84066
func (s *X服务) doRouterMapDump() {
	if !s.config.DumpRouterMap {
		return
	}

	var (
		ctx                          = context.TODO()
		routes                       = s.X取路由切片()
		isJustDefaultServerAndDomain = true
		headers                      = []string{
			"SERVER", "DOMAIN", "ADDRESS", "METHOD", "ROUTE", "HANDLER", "MIDDLEWARE",
		}
	)
	for _, item := range routes {
		if item.X服务器名称 != DefaultServerName || item.Domain != DefaultDomainName {
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
				handlerName = gstr.X过滤尾字符(item.Handler.X处理器名称, "-fm")
				middlewares = gstr.X分割并忽略空值(item.X中间件名称, ",")
			)

						// 不打印可能导致混淆的特殊内部中间件。 md5:16290644b5e53f3b
			if gstr.X取右边并倒找(handlerName, ".") == noPrintInternalRoute {
				continue
			}
			for k, v := range middlewares {
				middlewares[k] = gstr.X过滤尾字符(v, "-fm")
			}
			item.X中间件名称 = gstr.X连接(middlewares, "\n")
			if isJustDefaultServerAndDomain {
				data = append(
					data,
					item.X监听地址,
					item.Method,
					item.X路由URI,
					handlerName,
					item.X中间件名称,
				)
			} else {
				data = append(
					data,
					item.X服务器名称,
					item.Domain,
					item.X监听地址,
					item.Method,
					item.X路由URI,
					handlerName,
					item.X中间件名称,
				)
			}
			table.Append(data)
		}
		table.Render()
		s.config.Logger.X是否输出头信息(false).X输出并格式化(ctx, "\n%s", buffer.String())
	}
}

// X取OpenApi对象 返回当前服务器的 OpenAPI 规范管理对象。 md5:8926983c83bc678d
func (s *X服务) X取OpenApi对象() *goai.OpenApiV3 {
	return s.openapi
}

// X取路由切片 获取并返回路由器数组。 md5:4fe4fe015c1fb8e8
func (s *X服务) X取路由切片() []RouterItem {
	var (
		m              = make(map[string]*garray.SortedArray)
		routeFilterSet = gset.X创建文本()
		address        = s.X取已监听地址()
	)
	if s.config.HTTPSAddr != "" {
		if len(address) > 0 {
			address += ","
		}
		address += "tls" + s.config.HTTPSAddr
	}
	for k, handlerItems := range s.routesMap {
		array, _ := gregex.X匹配文本(`(.*?)%([A-Z]+):(.+)@(.+)`, k)
		for index := len(handlerItems) - 1; index >= 0; index-- {
			var (
				handlerItem = handlerItems[index]
				item        = RouterItem{
					X服务器名称:     s.config.Name,
					X监听地址:    address,
					Domain:     array[4],
					Type:       handlerItem.Type,
					X中间件名称: array[1],
					Method:     array[2],
					X路由URI:      array[3],
					Priority:   index,
					Handler:    handlerItem,
				}
			)
			switch item.Handler.Type {
			case HandlerTypeObject, HandlerTypeHandler:
				item.X是否为服务处理器 = true

			case HandlerTypeMiddleware:
				item.X中间件名称 = "GLOBAL MIDDLEWARE"
			}
						// 为dump重复路由过滤。 md5:e9067e19b70a8904
			var setKey = fmt.Sprintf(
				`%s|%s|%s|%s`,
				item.Method, item.X路由URI, item.Domain, item.Type,
			)
			if !routeFilterSet.X加入值并跳过已存在(setKey) {
				continue
			}
			if len(item.Handler.X中间件切片) > 0 {
				for _, v := range item.Handler.X中间件切片 {
					if item.X中间件名称 != "" {
						item.X中间件名称 += ","
					}
					item.X中间件名称 += gdebug.FuncName(v)
				}
			}
			// 如果域名在dump映射中不存在，它会创建该映射。
			// 映射的值是一个自定义排序的数组。
			// md5:b6191883863f4f52
			if _, ok := m[item.Domain]; !ok {
				// Sort in ASC order.
				m[item.Domain] = garray.X创建排序(func(v1, v2 interface{}) int {
					item1 := v1.(RouterItem)
					item2 := v2.(RouterItem)
					r := 0
					if r = strings.Compare(item1.Domain, item2.Domain); r == 0 {
						if r = strings.Compare(item1.X路由URI, item2.X路由URI); r == 0 {
							if r = strings.Compare(item1.Method, item2.Method); r == 0 {
								if item1.Handler.Type == HandlerTypeMiddleware && item2.Handler.Type != HandlerTypeMiddleware {
									return -1
								} else if item1.Handler.Type == HandlerTypeMiddleware && item2.Handler.Type == HandlerTypeMiddleware {
									return 1
								} else if r = strings.Compare(item1.X中间件名称, item2.X中间件名称); r == 0 {
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

// X启动服务 以阻塞方式启动服务器监听。
// 它通常用于单服务器场景。
// md5:4035b4359934ad62
func (s *X服务) X启动服务() {
	var ctx = context.TODO()

	if err := s.X开始监听(); err != nil {
		s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
	}

		// 异步信号处理程序。 md5:7eaa2de84f3b5dae
	go handleProcessSignal()

		// 通过通道进行阻塞以实现优雅重启。 md5:68e2b8bbfb67985a
	<-s.closeChan
	// Remove plugins.
	if len(s.plugins) > 0 {
		for _, p := range s.plugins {
			intlog.Printf(ctx, `remove plugin: %s`, p.X名称())
			if err := p.Remove(); err != nil {
				intlog.Errorf(ctx, "%+v", err)
			}
		}
	}
	s.doServiceDeregister()
	s.Logger别名().X输出并格式化INFO(ctx, "pid[%d]: all servers shutdown", gproc.Pid())
}

// X等待所有服务完成 会阻塞等待所有服务器完成。它通常用于多服务器情况。
// md5:69d8345a5fb12619
func X等待所有服务完成() {
	var ctx = context.TODO()

		// 异步信号处理程序。 md5:7eaa2de84f3b5dae
	go handleProcessSignal()

	<-allShutdownChan

	// Remove plugins.
	serverMapping.X遍历(func(k string, v interface{}) bool {
		s := v.(*X服务)
		if len(s.plugins) > 0 {
			for _, p := range s.plugins {
				intlog.Printf(ctx, `remove plugin: %s`, p.X名称())
				if err := p.Remove(); err != nil {
					intlog.Errorf(ctx, `%+v`, err)
				}
			}
		}
		return true
	})
	glog.X输出并格式化INFO(ctx, "pid[%d]: all servers shutdown", gproc.Pid())
}

// startServer 启动底层服务器并开始监听。 md5:9a21546d820319ef
func (s *X服务) startServer(fdMap listenerFdMap) {
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
				// Windows操作系统不支持从父进程传递套接字文件描述符。
				// md5:ab73e9587a9e540d
				if runtime.GOOS != "windows" {
					fd = gconv.X取整数(addrAndFd[1])
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
		array = gstr.X分割并忽略空值(v, ",")
	} else {
		array = gstr.X分割并忽略空值(s.config.Address, ",")
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
			// Windows操作系统不支持从父进程传递套接字文件描述符。
			// md5:68778a7e9822b36e
			if runtime.GOOS != "windows" {
				fd = gconv.X取整数(addrAndFd[1])
			}
		}
		if fd > 0 {
			s.servers = append(s.servers, s.newGracefulServer(itemFunc, fd))
		} else {
			s.servers = append(s.servers, s.newGracefulServer(itemFunc))
		}
	}
		// 开始异步监听。 md5:9d840d3502f6ae05
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
						// 以阻塞的方式开始监听和服务。 md5:066c0cdda27b7cd1
			err = server.Serve(ctx)
						// 如果服务器在没有错误的情况下关闭，进程将退出。 md5:0a64c12a91a31329
			if err != nil && !strings.EqualFold(http.ErrServerClosed.Error(), err.Error()) {
				s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
			}
						// 如果所有底层服务器都关闭，进程退出。 md5:4b398d3f7ef09228
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

// X取服务状态 获取并返回服务器状态。 md5:2f87146be638ddb6
func (s *X服务) X取服务状态() X服务状态 {
	if serverRunning.X取值() == 0 {
		return ServerStatusStopped
	}
		// 如果任何底层服务器正在运行，那么服务器状态为运行中。 md5:5e0e398c116a1838
	for _, v := range s.servers {
		if v.status.X取值() == ServerStatusRunning {
			return ServerStatusRunning
		}
	}
	return ServerStatusStopped
}

// getListenerFdMap 获取并返回套接字文件描述符的映射。
// 返回映射的键为 "http" 和 "https"。
// md5:970d132151bcc23b
func (s *X服务) getListenerFdMap() map[string]string {
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

// X取已监听端口 获取并返回当前服务器正在监听的其中一个端口。 md5:7e75c1b2c91e6f3e
func (s *X服务) X取已监听端口() int {
	ports := s.X取所有监听已端口()
	if len(ports) > 0 {
		return ports[0]
	}
	return 0
}

// X取所有监听已端口 获取并返回当前服务器监听的端口。 md5:98a46fc6cbcd9703
func (s *X服务) X取所有监听已端口() []int {
	ports := make([]int, 0)
	for _, server := range s.servers {
		ports = append(ports, server.GetListenedPort())
	}
	return ports
}

// X取已监听地址 获取并返回当前服务器所监听的地址字符串。 md5:51d352ffec9dc329
func (s *X服务) X取已监听地址() string {
	if !gstr.X是否包含(s.config.Address, X空闲端口地址) {
		return s.config.Address
	}
	var (
		address       = s.config.Address
		listenedPorts = s.X取所有监听已端口()
	)
	for _, listenedPort := range listenedPorts {
		address = gstr.X替换(address, X空闲端口地址, fmt.Sprintf(`:%d`, listenedPort), 1)
	}
	return address
}
