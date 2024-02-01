
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
// RouterGroup is a group wrapping multiple routes and middleware.
<原文结束>

# <翻译开始>
// RouterGroup 是一个包裹了多个路由和中间件的分组。
# <翻译结束>







<原文开始>
	// preBindItem is item for lazy registering feature of router group. preBindItem is not really registered
	// to server when route function of the group called but is lazily registered when server starts.
<原文结束>

# <翻译开始>
// preBindItem 是 router group 中用于懒加载注册功能的项目。当调用该组的路由函数时，preBindItem 并未真正注册到服务器上，
// 而是在服务器启动时才进行惰性注册。
# <翻译结束>


<原文开始>
// Can be handler, controller or object.
<原文结束>

# <翻译开始>
// 可以是处理器、控制器或对象。
# <翻译结束>


<原文开始>
// Extra parameters for route registering depending on the type.
<原文结束>

# <翻译开始>
// 根据类型为路由注册提供额外的参数。
# <翻译结束>


<原文开始>
// Handler is a register at a certain source file path: line.
<原文结束>

# <翻译开始>
// Handler 是在某个源文件路径：行处注册的处理器
# <翻译结束>


<原文开始>
// Is this item bound to server?
<原文结束>

# <翻译开始>
// 这个条目是否绑定到服务器？
# <翻译结束>


<原文开始>
// handlePreBindItems is called when server starts, which does really route registering to the server.
<原文结束>

# <翻译开始>
// handlePreBindItems 在服务器启动时被调用，它负责向服务器进行实际的路由注册。
# <翻译结束>


<原文开始>
// Handle the items of current server.
<原文结束>

# <翻译开始>
// 处理当前服务器的项目。
# <翻译结束>


<原文开始>
// Group creates and returns a RouterGroup object.
<原文结束>

# <翻译开始>
// Group 创建并返回一个 RouterGroup 对象。
# <翻译结束>


<原文开始>
// Group creates and returns a RouterGroup object, which is bound to a specified domain.
<原文结束>

# <翻译开始>
// Group 创建并返回一个 RouterGroup 对象，该对象与指定的域名绑定。
# <翻译结束>


<原文开始>
// Group creates and returns a subgroup of the current router group.
<原文结束>

# <翻译开始>
// Group 创建并返回当前路由组的一个子组。
# <翻译结束>


<原文开始>
// Clone returns a new router group which is a clone of the current group.
<原文结束>

# <翻译开始>
// Clone 返回一个新的路由组，它是当前组的克隆副本。
# <翻译结束>


<原文开始>
// Bind does batch route registering feature for a router group.
<原文结束>

# <翻译开始>
// Bind 为一个路由组提供批量注册路由的功能。
# <翻译结束>


<原文开始>
// ALL register an http handler to give the route pattern and all http methods.
<原文结束>

# <翻译开始>
// ALL 注册一个HTTP处理器，通过给定路由模式和所有HTTP方法。
# <翻译结束>


<原文开始>
// ALLMap registers http handlers for http methods using map.
<原文结束>

# <翻译开始>
// ALLMap 使用map为HTTP方法注册HTTP处理程序。
# <翻译结束>


<原文开始>
// Map registers http handlers for http methods using map.
<原文结束>

# <翻译开始>
// Map 通过使用映射表注册HTTP方法对应的HTTP处理器。
# <翻译结束>


<原文开始>
// GET registers an http handler to give the route pattern and the http method: GET.
<原文结束>

# <翻译开始>
// GET 注册一个 HTTP 处理器，用于给定的路由模式和 HTTP 方法：GET。
# <翻译结束>


<原文开始>
// PUT registers an http handler to give the route pattern and the http method: PUT.
<原文结束>

# <翻译开始>
// PUT 注册一个 HTTP 处理器，用于给定路由模式和 HTTP 方法：PUT。
# <翻译结束>


<原文开始>
// POST registers an http handler to give the route pattern and the http method: POST.
<原文结束>

# <翻译开始>
// POST 注册一个 HTTP 处理器，用于给定路由模式和 HTTP 方法：POST。
# <翻译结束>


<原文开始>
// DELETE registers an http handler to give the route pattern and the http method: DELETE.
<原文结束>

# <翻译开始>
// DELETE 注册一个 HTTP 处理器，根据给定的路由模式和 HTTP 方法（DELETE）进行处理。
# <翻译结束>


<原文开始>
// PATCH registers an http handler to give the route pattern and the http method: PATCH.
<原文结束>

# <翻译开始>
// PATCH 注册一个 HTTP 处理器，用于给定的路由模式和 HTTP 方法：PATCH。
# <翻译结束>


<原文开始>
// HEAD registers an http handler to give the route pattern and the http method: HEAD.
<原文结束>

# <翻译开始>
// HEAD 注册一个 HTTP 处理器，用于指定路由模式和 HTTP 方法：HEAD。
# <翻译结束>


<原文开始>
// CONNECT registers an http handler to give the route pattern and the http method: CONNECT.
<原文结束>

# <翻译开始>
// CONNECT 注册一个 HTTP 处理器，用于给定路由模式和 HTTP 方法：CONNECT。
# <翻译结束>


<原文开始>
// OPTIONS register an http handler to give the route pattern and the http method: OPTIONS.
<原文结束>

# <翻译开始>
// OPTIONS 注册一个HTTP处理器，用于指定路由模式和HTTP方法：OPTIONS。
# <翻译结束>


<原文开始>
// TRACE registers an http handler to give the route pattern and the http method: TRACE.
<原文结束>

# <翻译开始>
// TRACE 注册一个 HTTP 处理器，用于提供路由模式和 HTTP 方法：TRACE。
# <翻译结束>


<原文开始>
// REST registers an http handler to give the route pattern according to REST rule.
<原文结束>

# <翻译开始>
// REST 根据REST规则注册一个HTTP处理器，以便给定路由模式。
# <翻译结束>


<原文开始>
// Hook registers a hook to given route pattern.
<原文结束>

# <翻译开始>
// Hook 将钩子注册到给定的路由模式。
# <翻译结束>


<原文开始>
// Middleware binds one or more middleware to the router group.
<原文结束>

# <翻译开始>
// Middleware 将一个或多个中间件绑定到路由组。
# <翻译结束>


<原文开始>
// preBindToLocalArray adds the route registering parameters to an internal variable array for lazily registering feature.
<原文结束>

# <翻译开始>
// preBindToLocalArray 将路由注册参数预先添加到内部变量数组中，以便于进行惰性注册功能。
# <翻译结束>


<原文开始>
// getPrefix returns the route prefix of the group, which recursively retrieves its parent's prefix.
<原文结束>

# <翻译开始>
// getPrefix 返回当前组的路由前缀，该方法会递归获取其父级的前缀。
# <翻译结束>


<原文开始>
// doBindRoutersToServer does really register for the group.
<原文结束>

# <翻译开始>
// doBindRoutersToServer 是真正执行将路由器绑定到服务器的注册操作。
# <翻译结束>


<原文开始>
// If there is already a domain, unset the domain field in the pattern.
<原文结束>

# <翻译开始>
// 如果已经有域名，那么在模式中取消设置域名字段。
# <翻译结束>







<原文开始>
// Convert params to a string array.
<原文结束>

# <翻译开始>
// 将参数转换为字符串数组。
# <翻译结束>


<原文开始>
// Check whether it's a hook handler.
<原文结束>

# <翻译开始>
// 检查是否为钩子处理器。
# <翻译结束>


<原文开始>
// Finally, it treats the `object` as the Object registering type.
<原文结束>

# <翻译开始>
// 最后，它将`object`视为正在注册的类型对象。
# <翻译结束>







<原文开始>
// Prefix for sub-route.
<原文结束>

# <翻译开始>
// 子路由前缀。
# <翻译结束>


<原文开始>
// Filter repeated char '/'.
<原文结束>

# <翻译开始>
// 过滤重复的字符 '/'。
# <翻译结束>

