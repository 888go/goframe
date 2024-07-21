
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Initialize the method map.
<原文结束>

# <翻译开始>
// 初始化方法映射。 md5:19df407918075ac2
# <翻译结束>


<原文开始>
// serverProcessInit initializes some process configurations, which can only be done once.
<原文结束>

# <翻译开始>
// serverProcessInit 初始化一些进程配置，这些只能做一次。 md5:768ec2687e3a0f24
# <翻译结束>


<原文开始>
	// This means it is a restart server. It should kill its parent before starting its listening,
	// to avoid duplicated port listening in two processes.
<原文结束>

# <翻译开始>
	// 这意味着它是一个重启服务器。在开始监听之前，它应该杀死其父进程，以防止两个进程占用同一个端口进行监听。
	// md5:534f767682eae0c3
# <翻译结束>


<原文开始>
	// Process message handler.
	// It enabled only a graceful feature is enabled.
<原文结束>

# <翻译开始>
	// 处理消息的处理器。
	// 仅当启用了优雅特性时，它才可用。
	// md5:b41176de77cc4833
# <翻译结束>


<原文开始>
	// It's an ugly calling for better initializing the main package path
	// in source development environment. It is useful only be used in main goroutine.
	// It fails to retrieve the main package path in asynchronous goroutines.
<原文结束>

# <翻译开始>
	// 这是一种在源代码开发环境中更好地初始化主包路径的不优雅的方式。
	// 它只应在主 goroutine 中使用。
	// 在异步 goroutines 中无法获取主包路径。
	// md5:e0733b122721224e
# <翻译结束>


<原文开始>
// GetServer creates and returns a server instance using given name and default configurations.
// Note that the parameter `name` should be unique for different servers. It returns an existing
// server instance if given `name` is already existing in the server mapping.
<原文结束>

# <翻译开始>
// GetServer 使用给定名称和默认配置创建并返回一个服务器实例。
// 注意，参数 `name` 应该在不同的服务器中是唯一的。如果给定的 `name` 已经存在于服务器映射中，它将返回现有的服务器实例。
// md5:ad04664fa9750188
# <翻译结束>


<原文开始>
// Initialize the server using default configurations.
<原文结束>

# <翻译开始>
// 使用默认配置初始化服务器。 md5:ac8ad35c2e6592fb
# <翻译结束>


<原文开始>
// It enables OpenTelemetry for server in default.
<原文结束>

# <翻译开始>
// 它默认为服务器启用OpenTelemetry。 md5:2a2de55e6612dec7
# <翻译结束>


<原文开始>
// Start starts listening on configured port.
// This function does not block the process, you can use function Wait blocking the process.
<原文结束>

# <翻译开始>
// Start 开始在配置的端口上监听。
// 此函数不会阻塞进程，你可以使用 Wait 函数来阻塞进程。
// md5:05c1c66553fa4a61
# <翻译结束>


<原文开始>
// OpenApi specification json producing handler.
<原文结束>

# <翻译开始>
// 开放API规范生成JSON处理器。 md5:62cb20bceb4ec15e
# <翻译结束>


<原文开始>
// Server process initialization, which can only be initialized once.
<原文结束>

# <翻译开始>
// 服务器进程初始化，这只能初始化一次。 md5:f4fd6ab84839bb71
# <翻译结束>


<原文开始>
// Server can only be run once.
<原文结束>

# <翻译开始>
// 服务器只能运行一次。 md5:4372da1cc9e271f0
# <翻译结束>


<原文开始>
// Logging path setting check.
<原文结束>

# <翻译开始>
// 日志记录路径设置检查。 md5:b1b53a71404f3b28
# <翻译结束>


<原文开始>
// Default session storage.
<原文结束>

# <翻译开始>
// 默认会话存储。 md5:e5e4d66ee85c002a
# <翻译结束>


<原文开始>
// Initialize session manager when start running.
<原文结束>

# <翻译开始>
// 在启动时初始化会话管理器。 md5:060e4cc6f4f8b93e
# <翻译结束>


<原文开始>
// Install external plugins.
<原文结束>

# <翻译开始>
// 安装外部插件。 md5:5a986e9f0fb84368
# <翻译结束>


<原文开始>
// Check the group routes again for internally registered routes.
<原文结束>

# <翻译开始>
// 检查内部注册的路由再次应用于组路由。 md5:7949c3fe59e30c8c
# <翻译结束>


<原文开始>
	// If there's no route registered and no static service enabled,
	// it then returns an error of invalid usage of server.
<原文结束>

# <翻译开始>
	// 如果没有注册路由且没有启用静态服务，它将返回服务器使用无效的错误。
	// md5:d916b25cf4c384d4
# <翻译结束>


<原文开始>
	// ================================================================================================
	// Start the HTTP server.
	// ================================================================================================
<原文结束>

# <翻译开始>
	// ================================================================================================
	// 启动HTTP服务器。
	// ================================================================================================
	// md5:9c551ab8996cea5a
# <翻译结束>


<原文开始>
// OpenApi specification info.
<原文结束>

# <翻译开始>
// OpenAPI规范信息。 md5:e99b4557db364598
# <翻译结束>


<原文开始>
// If this is a child process, it then notifies its parent exit.
<原文结束>

# <翻译开始>
// 如果该进程是子进程，那么它将通知其父进程退出。 md5:c9a0c66193cbdbfc
# <翻译结束>


<原文开始>
// doRouterMapDump checks and dumps the router map to the log.
<原文结束>

# <翻译开始>
// doRouterMapDump 检查并把路由映射dump到日志中。 md5:658175121fd84066
# <翻译结束>


<原文开始>
// No printing special internal middleware that may lead confused.
<原文结束>

# <翻译开始>
// 不打印可能导致混淆的特殊内部中间件。 md5:16290644b5e53f3b
# <翻译结束>


<原文开始>
// GetOpenApi returns the OpenApi specification management object of current server.
<原文结束>

# <翻译开始>
// GetOpenApi 返回当前服务器的 OpenAPI 规范管理对象。 md5:8926983c83bc678d
# <翻译结束>


<原文开始>
// GetRoutes retrieves and returns the router array.
<原文结束>

# <翻译开始>
// GetRoutes 获取并返回路由器数组。 md5:4fe4fe015c1fb8e8
# <翻译结束>


<原文开始>
// Repeated route filtering for dump.
<原文结束>

# <翻译开始>
// 为dump重复路由过滤。 md5:e9067e19b70a8904
# <翻译结束>


<原文开始>
			// If the domain does not exist in the dump map, it creates the map.
			// The value of the map is a custom sorted array.
<原文结束>

# <翻译开始>
			// 如果域名在dump映射中不存在，它会创建该映射。
			// 映射的值是一个自定义排序的数组。
			// md5:b6191883863f4f52
# <翻译结束>


<原文开始>
// Run starts server listening in blocking way.
// It's commonly used for single server situation.
<原文结束>

# <翻译开始>
// Run 以阻塞方式启动服务器监听。
// 它通常用于单服务器场景。
// md5:4035b4359934ad62
# <翻译结束>


<原文开始>
// Signal handler in asynchronous way.
<原文结束>

# <翻译开始>
// 异步信号处理程序。 md5:7eaa2de84f3b5dae
# <翻译结束>


<原文开始>
// Blocking using channel for graceful restart.
<原文结束>

# <翻译开始>
// 通过通道进行阻塞以实现优雅重启。 md5:68e2b8bbfb67985a
# <翻译结束>


<原文开始>
// Wait blocks to wait for all servers done.
// It's commonly used in multiple server situation.
<原文结束>

# <翻译开始>
// Wait 会阻塞等待所有服务器完成。它通常用于多服务器情况。
// md5:69d8345a5fb12619
# <翻译结束>


<原文开始>
// startServer starts the underlying server listening.
<原文结束>

# <翻译开始>
// startServer 启动底层服务器并开始监听。 md5:9a21546d820319ef
# <翻译结束>


<原文开始>
				// The Windows OS does not support socket file descriptor passing
				// from parent process.
<原文结束>

# <翻译开始>
				// Windows操作系统不支持从父进程传递套接字文件描述符。
				// md5:ab73e9587a9e540d
# <翻译结束>


<原文开始>
			// The Window OS does not support socket file descriptor passing
			// from the parent process.
<原文结束>

# <翻译开始>
			// Windows操作系统不支持从父进程传递套接字文件描述符。
			// md5:68778a7e9822b36e
# <翻译结束>


<原文开始>
// Start listening asynchronously.
<原文结束>

# <翻译开始>
// 开始异步监听。 md5:9d840d3502f6ae05
# <翻译结束>


<原文开始>
// Start listening and serving in blocking way.
<原文结束>

# <翻译开始>
// 以阻塞的方式开始监听和服务。 md5:066c0cdda27b7cd1
# <翻译结束>


<原文开始>
// The process exits if the server is closed with none closing error.
<原文结束>

# <翻译开始>
// 如果服务器在没有错误的情况下关闭，进程将退出。 md5:0a64c12a91a31329
# <翻译结束>


<原文开始>
// If all the underlying servers' shutdown, the process exits.
<原文结束>

# <翻译开始>
// 如果所有底层服务器都关闭，进程退出。 md5:4b398d3f7ef09228
# <翻译结束>


<原文开始>
// Status retrieves and returns the server status.
<原文结束>

# <翻译开始>
// Status 获取并返回服务器状态。 md5:2f87146be638ddb6
# <翻译结束>


<原文开始>
// If any underlying server is running, the server status is running.
<原文结束>

# <翻译开始>
// 如果任何底层服务器正在运行，那么服务器状态为运行中。 md5:5e0e398c116a1838
# <翻译结束>


<原文开始>
// getListenerFdMap retrieves and returns the socket file descriptors.
// The key of the returned map is "http" and "https".
<原文结束>

# <翻译开始>
// getListenerFdMap 获取并返回套接字文件描述符的映射。
// 返回映射的键为 "http" 和 "https"。
// md5:970d132151bcc23b
# <翻译结束>


<原文开始>
// GetListenedPort retrieves and returns one port which is listened by current server.
<原文结束>

# <翻译开始>
// GetListenedPort 获取并返回当前服务器正在监听的其中一个端口。 md5:7e75c1b2c91e6f3e
# <翻译结束>


<原文开始>
// GetListenedPorts retrieves and returns the ports which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedPorts 获取并返回当前服务器监听的端口。 md5:98a46fc6cbcd9703
# <翻译结束>


<原文开始>
// GetListenedAddress retrieves and returns the address string which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedAddress 获取并返回当前服务器所监听的地址字符串。 md5:51d352ffec9dc329
# <翻译结束>

