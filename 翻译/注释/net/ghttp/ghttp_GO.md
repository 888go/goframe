
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
// Package ghttp provides powerful http server and simple client implements.
<原文结束>

# <翻译开始>
// Package ghttp 提供了强大的HTTP服务器及简洁的客户端实现。
# <翻译结束>


<原文开始>
// Server wraps the http.Server and provides more rich features.
<原文结束>

# <翻译开始>
// Server 包装了 http.Server，并提供了更多丰富的功能。
# <翻译结束>


<原文开始>
// Instance name of current HTTP server.
<原文结束>

# <翻译开始>
// 当前HTTP服务器的实例名称。
# <翻译结束>







<原文开始>
// Plugin array to extend server functionality.
<原文结束>

# <翻译开始>
// 插件数组，用于扩展服务器功能。
# <翻译结束>


<原文开始>
// Underlying http.Server array.
<原文结束>

# <翻译开始>
// 底层 http.Server 数组
# <翻译结束>


<原文开始>
// Underlying http.Server number for internal usage.
<原文结束>

# <翻译开始>
// 用于内部使用的底层http.Server数字。
# <翻译结束>


<原文开始>
// Used for underlying server closing event notification.
<原文结束>

# <翻译开始>
// 用于底层服务器关闭事件的通知。
# <翻译结束>







<原文开始>
// Server caches for internal usage.
<原文结束>

# <翻译开始>
// Server内部使用的缓存。
# <翻译结束>


<原文开始>
// Route map mainly for route dumps and repeated route checks.
<原文结束>

# <翻译开始>
// 路由映射表主要用于路由转储和重复路由检查。
# <翻译结束>







<原文开始>
// The OpenApi specification management object.
<原文结束>

# <翻译开始>
// OpenApi规范管理对象。
# <翻译结束>


<原文开始>
// Concurrent safety for operations of attribute service.
<原文结束>

# <翻译开始>
// 并发安全：保证属性服务操作的并发安全性。
# <翻译结束>







<原文开始>
// Registrar for service register.
<原文结束>

# <翻译开始>
// Registrar 用于服务注册。
# <翻译结束>


<原文开始>
// Parsed regular expression for route matching.
<原文结束>

# <翻译开始>
// 解析后的用于路由匹配的正则表达式。
# <翻译结束>


<原文开始>
// Parsed router parameter names.
<原文结束>

# <翻译开始>
// 解析后的路由参数名称。
# <翻译结束>







<原文开始>
// RouterItem is just for route dumps.
<原文结束>

# <翻译开始>
// RouterItem 仅用于路由转储。
# <翻译结束>






















<原文开始>
// HandlerFunc is request handler function.
<原文结束>

# <翻译开始>
// HandlerFunc 是请求处理函数。
# <翻译结束>


<原文开始>
// handlerFuncInfo contains the HandlerFunc address and its reflection type.
<原文结束>

# <翻译开始>
// handlerFuncInfo 包含 HandlerFunc 的地址及其反射类型。
# <翻译结束>







<原文开始>
// Reflect type information for current handler, which is used for extensions of the handler feature.
<原文结束>

# <翻译开始>
// 反射当前处理器的类型信息，该信息用于扩展处理器功能。
# <翻译结束>


<原文开始>
// Reflect value information for current handler, which is used for extensions of the handler feature.
<原文结束>

# <翻译开始>
// Reflect当前处理器的值信息，该信息用于扩展处理器功能。
# <翻译结束>


<原文开始>
// Whether strict route matching is enabled.
<原文结束>

# <翻译开始>
// 是否启用严格的路由匹配。
# <翻译结束>







<原文开始>
	// HandlerItem is the registered handler for route handling,
	// including middleware and hook functions.
<原文结束>

# <翻译开始>
// HandlerItem 是注册的路由处理程序，
// 包括中间件和钩子函数。
# <翻译结束>


<原文开始>
		// Unique handler item id mark.
		// Note that the handler function may be registered multiple times as different handler items,
		// which have different handler item id.
<原文结束>

# <翻译开始>
// 唯一处理器项标识符标记。
// 注意，处理器函数可能会以不同的处理器项身份注册多次，
// 这些处理器项具有不同的处理器项标识符。
# <翻译结束>


<原文开始>
// Handler name, which is automatically retrieved from runtime stack when registered.
<原文结束>

# <翻译开始>
// 处理器名称，在注册时会自动从运行时堆栈中获取。
# <翻译结束>


<原文开始>
// Handler type: object/handler/middleware/hook.
<原文结束>

# <翻译开始>
// 处理器类型：对象/处理器/中间件/钩子。
# <翻译结束>


<原文开始>
// Handler function information.
<原文结束>

# <翻译开始>
// 处理函数信息
# <翻译结束>


<原文开始>
// Initialization function when request enters the object (only available for object register type).
<原文结束>

# <翻译开始>
// 初始化函数：当请求进入对象时调用（仅适用于对象注册类型）。
# <翻译结束>


<原文开始>
// Shutdown function when request leaves out the object (only available for object register type).
<原文结束>

# <翻译开始>
// Shutdown 函数在请求离开对象时调用（仅适用于对象注册类型）。
# <翻译结束>







<原文开始>
// Hook type name, only available for the hook type.
<原文结束>

# <翻译开始>
// Hook类型名称，仅适用于Hook类型。
# <翻译结束>


<原文开始>
// HandlerItemParsed is the item parsed from URL.Path.
<原文结束>

# <翻译开始>
// HandlerItemParsed是从URL.Path中解析出的项目。
# <翻译结束>












<原文开始>
// ServerStatus is the server status enum type.
<原文结束>

# <翻译开始>
// ServerStatus 是服务器状态枚举类型。
# <翻译结束>


<原文开始>
// HookName is the route hook name enum type.
<原文结束>

# <翻译开始>
// HookName 是路由钩子名称的枚举类型。
# <翻译结束>


<原文开始>
// HandlerType is the route handler enum type.
<原文结束>

# <翻译开始>
// HandlerType 是路由处理器的枚举类型。
# <翻译结束>


<原文开始>
	// Listening file descriptor mapping.
	// The key is either "http" or "https" and the value is its FD.
<原文结束>

# <翻译开始>
// 监听文件描述符映射。
// 键是 "http" 或 "https"，其对应的值为相应的文件描述符（FD）。
# <翻译结束>


<原文开始>
// internalPanic is the custom panic for internal usage.
<原文结束>

# <翻译开始>
// internalPanic 是用于内部使用的自定义恐慌函数。
# <翻译结束>


<原文开始>
// FreePortAddress marks the server listens using random free port.
<原文结束>

# <翻译开始>
// FreePortAddress 表示服务器使用随机空闲端口进行监听。
# <翻译结束>


<原文开始>
// Used for custom route handler, which does not change URL.Path.
<原文结束>

# <翻译开始>
// 用于自定义路由处理器，在此情况下，URL.Path 不会发生变化。
# <翻译结束>


<原文开始>
// Hook handler before route handler/file serving.
<原文结束>

# <翻译开始>
// 在路由处理器/文件服务之前执行的钩子处理器。
# <翻译结束>


<原文开始>
// Hook handler after route handler/file serving.
<原文结束>

# <翻译开始>
// 在路由处理器/文件服务之后的钩子处理器。
# <翻译结束>


<原文开始>
// Hook handler before response output.
<原文结束>

# <翻译开始>
// 在响应输出前的钩子处理器
# <翻译结束>


<原文开始>
// Hook handler after response output.
<原文结束>

# <翻译开始>
// Hook处理器在响应输出之后。
# <翻译结束>


<原文开始>
	// methodsMap stores all supported HTTP method.
	// It is used for quick HTTP method searching using map.
<原文结束>

# <翻译开始>
// methodsMap 存储所有支持的HTTP方法。
// 该映射用于通过map快速搜索HTTP方法。
# <翻译结束>


<原文开始>
	// serverMapping stores more than one server instances for current processes.
	// The key is the name of the server, and the value is its instance.
<原文结束>

# <翻译开始>
// serverMapping 用于存储当前进程中多个服务器实例。
// 键是服务器的名称，值是其对应的实例。
# <翻译结束>


<原文开始>
	// serverRunning marks the running server counts.
	// If there is no successful server running or all servers' shutdown, this value is 0.
<原文结束>

# <翻译开始>
// serverRunning 标记正在运行的服务器数量。
// 如果没有运行成功的服务器，或者所有服务器都已关闭，则该值为0。
# <翻译结束>


<原文开始>
// wsUpGrader is the default up-grader configuration for websocket.
<原文结束>

# <翻译开始>
// wsUpGrader 是用于websocket的默认升级配置。
# <翻译结束>


<原文开始>
// It does not check the origin in default, the application can do it itself.
<原文结束>

# <翻译开始>
// 默认情况下，它不检查源，应用程序可以自行进行检查。
# <翻译结束>


<原文开始>
	// allShutdownChan is the event for all servers have done its serving and exit.
	// It is used for process blocking purpose.
<原文结束>

# <翻译开始>
// allShutdownChan 是所有服务器完成服务并退出的事件。
// 它用于进程阻塞目的。
# <翻译结束>


<原文开始>
	// serverProcessInitialized is used for lazy initialization for server.
	// The process can only be initialized once.
<原文结束>

# <翻译开始>
// serverProcessInitialized 用于服务端的延迟初始化。
// 这个过程只能被初始化一次。
# <翻译结束>


<原文开始>
// gracefulEnabled is used for a graceful reload feature, which is false in default.
<原文结束>

# <翻译开始>
// gracefulEnabled 用于实现优雅重启功能，默认情况下为 false。
# <翻译结束>


<原文开始>
// defaultValueTags are the struct tag names for default value storing.
<原文结束>

# <翻译开始>
// defaultValueTags 是用于存储默认值的结构体标签名称。
# <翻译结束>







<原文开始>
// Registering source file `path:line`.
<原文结束>

# <翻译开始>
// 注册源文件 `路径:行数`。
# <翻译结束>


<原文开始>
// Router values parsed from URL.Path.
<原文结束>

# <翻译开始>
// Router values 从 URL.Path 解析得到。
# <翻译结束>


<原文开始>
// Server configuration.
<原文结束>

# <翻译开始>
// 服务器配置
# <翻译结束>


<原文开始>
// The route maps tree.
<原文结束>

# <翻译开始>
// 路由映射树
# <翻译结束>


<原文开始>
// Custom status handler map.
<原文结束>

# <翻译开始>
// 自定义状态处理映射。
# <翻译结束>


<原文开始>
// The service for Registry.
<原文结束>

# <翻译开始>
// Registry服务。
# <翻译结束>


<原文开始>
// Just for reference.
<原文结束>

# <翻译开始>
// 仅供参考
# <翻译结束>


<原文开始>
// Listening address.
<原文结束>

# <翻译开始>
// 监听地址。
# <翻译结束>


<原文开始>
// Route handler type.
<原文结束>

# <翻译开始>
// 路由处理器类型。
# <翻译结束>


<原文开始>
// Handler method name.
<原文结束>

# <翻译开始>
// 处理器方法名称。
# <翻译结束>


<原文开始>
// Is service handler.
<原文结束>

# <翻译开始>
// 是否为服务处理器
# <翻译结束>


<原文开始>
// Handler function address.
<原文结束>

# <翻译开始>
// 处理函数地址。
# <翻译结束>


<原文开始>
// Request struct fields.
<原文结束>

# <翻译开始>
// 请求结构体字段。
# <翻译结束>


<原文开始>
// Bound middleware array.
<原文结束>

# <翻译开始>
// 绑定中间件数组。
# <翻译结束>


<原文开始>
// Handler information.
<原文结束>

# <翻译开始>
// 处理器信息。
# <翻译结束>

