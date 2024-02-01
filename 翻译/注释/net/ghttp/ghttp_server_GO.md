
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>







<原文开始>
// serverProcessInit initializes some process configurations, which can only be done once.
<原文结束>

# <翻译开始>
// serverProcessInit 初始化一些只能执行一次的进程配置。
# <翻译结束>


<原文开始>
	// This means it is a restart server. It should kill its parent before starting its listening,
	// to avoid duplicated port listening in two processes.
<原文结束>

# <翻译开始>
// 这意味着它是一个重启服务。在开始监听之前，它应该先终止其父进程，
// 以避免在两个进程中因端口重复监听而产生问题。
# <翻译结束>


<原文开始>
	// Process message handler.
	// It enabled only a graceful feature is enabled.
<原文结束>

# <翻译开始>
// 处理消息处理器。
// 只有在启用了优雅特性时，它才会被启用。
# <翻译结束>


<原文开始>
	// It's an ugly calling for better initializing the main package path
	// in source development environment. It is useful only be used in main goroutine.
	// It fails to retrieve the main package path in asynchronous goroutines.
<原文结束>

# <翻译开始>
// 这是一种用于在源代码开发环境中更好地初始化主包路径的丑陋调用方式。它仅在主goroutine中有用。
// 在异步goroutine中，该方法无法正确获取主包路径。
# <翻译结束>


<原文开始>
// GetServer creates and returns a server instance using given name and default configurations.
// Note that the parameter `name` should be unique for different servers. It returns an existing
// server instance if given `name` is already existing in the server mapping.
<原文结束>

# <翻译开始>
// GetServer 根据给定名称和默认配置创建并返回一个服务器实例。
// 注意，参数`name`对于不同服务器应保持唯一。如果给定的`name`已在服务器映射中存在，
// 则它将返回一个已存在的服务器实例。
# <翻译结束>


<原文开始>
// Initialize the server using default configurations.
<原文结束>

# <翻译开始>
// 使用默认配置初始化服务器。
# <翻译结束>


<原文开始>
// It enables OpenTelemetry for server in default.
<原文结束>

# <翻译开始>
// 它默认为服务器启用OpenTelemetry。
# <翻译结束>


<原文开始>
// Start starts listening on configured port.
// This function does not block the process, you can use function Wait blocking the process.
<原文结束>

# <翻译开始>
// Start 开始监听配置好的端口。
// 该函数不会阻塞进程，你可以使用函数 Wait 来阻塞进程。
# <翻译结束>


<原文开始>
// OpenApi specification json producing handler.
<原文结束>

# <翻译开始>
// OpenApi规范JSON生成处理器。
# <翻译结束>







<原文开始>
// Server process initialization, which can only be initialized once.
<原文结束>

# <翻译开始>
// 服务器进程初始化，只能初始化一次。
# <翻译结束>


<原文开始>
// Server can only be run once.
<原文结束>

# <翻译开始>
// 服务只能运行一次。
# <翻译结束>


<原文开始>
// Logging path setting check.
<原文结束>

# <翻译开始>
// 日志路径设置检查
# <翻译结束>







<原文开始>
// Initialize session manager when start running.
<原文结束>

# <翻译开始>
// 在程序启动运行时初始化会话管理器。
# <翻译结束>












<原文开始>
// Check the group routes again for internally registered routes.
<原文结束>

# <翻译开始>
// 再次检查组路由中内部注册的路由。
# <翻译结束>


<原文开始>
	// If there's no route registered and no static service enabled,
	// it then returns an error of invalid usage of server.
<原文结束>

# <翻译开始>
// 如果没有注册路由且未启用静态服务，
// 则返回服务器使用无效的错误。
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
# <翻译结束>


<原文开始>
// OpenApi specification info.
<原文结束>

# <翻译开始>
// OpenApi规范信息
# <翻译结束>


<原文开始>
// If this is a child process, it then notifies its parent exit.
<原文结束>

# <翻译开始>
// 如果这是一个子进程，那么它会通知其父进程已退出。
# <翻译结束>


<原文开始>
// doRouterMapDump checks and dumps the router map to the log.
<原文结束>

# <翻译开始>
// doRouterMapDump 检查并把路由映射表转储到日志中。
# <翻译结束>


<原文开始>
// GetOpenApi returns the OpenApi specification management object of current server.
<原文结束>

# <翻译开始>
// GetOpenApi 返回当前服务器的OpenApi规范管理对象。
# <翻译结束>


<原文开始>
// GetRoutes retrieves and returns the router array.
<原文结束>

# <翻译开始>
// GetRoutes 获取并返回路由数组。
# <翻译结束>


<原文开始>
// Repeated route filtering for dump.
<原文结束>

# <翻译开始>
// 重复路径过滤以供转储
# <翻译结束>


<原文开始>
			// If the domain does not exist in the dump map, it creates the map.
			// The value of the map is a custom sorted array.
<原文结束>

# <翻译开始>
// 如果域名不存在于dump映射中，则创建该映射。
// 映射的值是一个自定义排序的数组。
# <翻译结束>







<原文开始>
// Run starts server listening in blocking way.
// It's commonly used for single server situation.
<原文结束>

# <翻译开始>
// Run 启动服务器并以阻塞方式监听。
// 该方法通常用于单服务器场景。
# <翻译结束>


<原文开始>
// Signal handler in asynchronous way.
<原文结束>

# <翻译开始>
// 以异步方式处理信号。
# <翻译结束>


<原文开始>
// Blocking using channel for graceful restart.
<原文结束>

# <翻译开始>
// 使用通道进行阻塞，实现优雅重启。
# <翻译结束>


<原文开始>
// Wait blocks to wait for all servers done.
// It's commonly used in multiple server situation.
<原文结束>

# <翻译开始>
// Wait 阻塞等待所有服务器完成。
// 在多服务器场景中，它通常被使用。
# <翻译结束>


<原文开始>
// startServer starts the underlying server listening.
<原文结束>

# <翻译开始>
// startServer 启动底层服务器并开始监听。
# <翻译结束>


<原文开始>
				// The Windows OS does not support socket file descriptor passing
				// from parent process.
<原文结束>

# <翻译开始>
// Windows操作系统不支持从父进程传递socket文件描述符。
# <翻译结束>


<原文开始>
			// The Window OS does not support socket file descriptor passing
			// from the parent process.
<原文结束>

# <翻译开始>
// Windows 操作系统不支持从父进程传递套接字文件描述符。
# <翻译结束>


<原文开始>
// Start listening asynchronously.
<原文结束>

# <翻译开始>
// 开始异步监听。
# <翻译结束>







<原文开始>
// Start listening and serving in blocking way.
<原文结束>

# <翻译开始>
// 开始监听并以阻塞方式提供服务。
# <翻译结束>


<原文开始>
// The process exits if the server is closed with none closing error.
<原文结束>

# <翻译开始>
// 如果服务器在没有关闭错误的情况下被关闭，进程将退出。
# <翻译结束>


<原文开始>
// If all the underlying servers' shutdown, the process exits.
<原文结束>

# <翻译开始>
// 如果所有底层服务器都关闭，进程将退出。
# <翻译结束>


<原文开始>
// Status retrieves and returns the server status.
<原文结束>

# <翻译开始>
// Status 获取并返回服务器状态。
# <翻译结束>


<原文开始>
// If any underlying server is running, the server status is running.
<原文结束>

# <翻译开始>
// 如果任何底层服务器正在运行，则服务器状态为运行中。
# <翻译结束>


<原文开始>
// getListenerFdMap retrieves and returns the socket file descriptors.
// The key of the returned map is "http" and "https".
<原文结束>

# <翻译开始>
// getListenerFdMap 获取并返回套接字文件描述符的映射。
// 返回映射中的键为 "http" 和 "https"。
# <翻译结束>


<原文开始>
// GetListenedPort retrieves and returns one port which is listened by current server.
<原文结束>

# <翻译开始>
// GetListenedPort 获取并返回当前服务器正在监听的一个端口。
# <翻译结束>


<原文开始>
// GetListenedPorts retrieves and returns the ports which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedPorts 获取并返回当前服务器正在监听的所有端口。
# <翻译结束>


<原文开始>
// GetListenedAddress retrieves and returns the address string which are listened by current server.
<原文结束>

# <翻译开始>
// GetListenedAddress 获取并返回当前服务器监听的地址字符串。
# <翻译结束>












<原文开始>
// Initialize the method map.
<原文结束>

# <翻译开始>
// 初始化方法映射。
# <翻译结束>


<原文开始>
// Register group routes.
<原文结束>

# <翻译开始>
// 注册群组路由。
# <翻译结束>


<原文开始>
// Default session storage.
<原文结束>

# <翻译开始>
// 默认的会话存储。
# <翻译结束>


<原文开始>
// Default HTTP handler.
<原文结束>

# <翻译开始>
// 默认HTTP处理器
# <翻译结束>


<原文开始>
// Install external plugins.
<原文结束>

# <翻译开始>
// 安装外部插件。
# <翻译结束>


<原文开始>
// Sort in ASC order.
<原文结束>

# <翻译开始>
// 按升序排序
# <翻译结束>

