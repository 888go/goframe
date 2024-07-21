
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
// RouterGroup is a group wrapping multiple routes and middleware.
<原文结束>

# <翻译开始>
// RouterGroup 是一个包裹多个路由和中间件的分组。 md5:609e7eb75d8a51f0
# <翻译结束>


<原文开始>
	// preBindItem is item for lazy registering feature of router group. preBindItem is not really registered
	// to server when route function of the group called but is lazily registered when server starts.
<原文结束>

# <翻译开始>
	// preBindItem 是用于路由器组延迟注册功能的项目。当路由组的路由函数被调用时，preBindItem 并没有真正注册到服务器，
	// 而是在服务器启动时进行懒惰注册。
	// md5:4255b2f4d61ba05c
# <翻译结束>


<原文开始>
// Can be handler, controller or object.
<原文结束>

# <翻译开始>
// 可以是处理器、控制器或对象。 md5:0c53d2880dc0aafc
# <翻译结束>


<原文开始>
// Extra parameters for route registering depending on the type.
<原文结束>

# <翻译开始>
// 根据类型，为路由注册额外的参数。 md5:9b2faaa25e40fcdb
# <翻译结束>


<原文开始>
// Handler is a register at a certain source file path: line.
<原文结束>

# <翻译开始>
// Handler 是在特定源文件路径和行号处的注册处理器。 md5:0cf92074c14f8d58
# <翻译结束>


<原文开始>
// Is this item bound to server?
<原文结束>

# <翻译开始>
// 是否将此项目绑定到服务器？. md5:7da1001818a8ed44
# <翻译结束>


<原文开始>
// handlePreBindItems is called when server starts, which does really route registering to the server.
<原文结束>

# <翻译开始>
// handlePreBindItems 在服务器启动时被调用，它确实将路由注册到服务器上。 md5:4f0e019b6d905274
# <翻译结束>


<原文开始>
// Handle the items of current server.
<原文结束>

# <翻译开始>
// 处理当前服务器的项目。 md5:5caf897b6c7b073a
# <翻译结束>


<原文开始>
// Group creates and returns a RouterGroup object.
<原文结束>

# <翻译开始>
// Group 创建并返回一个 RouterGroup 对象。 md5:ab811975f9ba0334
# <翻译结束>


<原文开始>
// Group creates and returns a RouterGroup object, which is bound to a specified domain.
<原文结束>

# <翻译开始>
// Group 创建并返回一个 RouterGroup 对象，该对象绑定到指定的域名。 md5:bd60cfbd62234fcd
# <翻译结束>


<原文开始>
// Group creates and returns a subgroup of the current router group.
<原文结束>

# <翻译开始>
// Group 创建并返回当前路由器组的一个子组。 md5:9706484677759d8f
# <翻译结束>


<原文开始>
// Clone returns a new router group which is a clone of the current group.
<原文结束>

# <翻译开始>
// Clone 返回一个新的路由组，它是当前组的克隆。 md5:a3328662d1da7f5f
# <翻译结束>


<原文开始>
// Bind does batch route registering feature for a router group.
<原文结束>

# <翻译开始>
// Bind 为路由器组提供了批量路由注册的功能。 md5:16fbec330e17cafe
# <翻译结束>


<原文开始>
// ALL register an http handler to give the route pattern and all http methods.
<原文结束>

# <翻译开始>
// ALL 注册一个http处理器，用于处理给定路由模式的所有HTTP方法。 md5:06f3f9b3c30b17f0
# <翻译结束>


<原文开始>
// ALLMap registers http handlers for http methods using map.
<原文结束>

# <翻译开始>
// ALLMap 使用映射注册HTTP处理程序，针对HTTP方法。 md5:4baef0383348c469
# <翻译结束>


<原文开始>
// Map registers http handlers for http methods using map.
<原文结束>

# <翻译开始>
// Map使用映射注册HTTP方法的处理器。 md5:234d05d7bb247514
# <翻译结束>


<原文开始>
// GET registers an http handler to give the route pattern and the http method: GET.
<原文结束>

# <翻译开始>
// GET 函数用于注册一个HTTP处理程序，该程序根据给定的路由模式和HTTP方法（GET）进行处理。 md5:28790c458e1b962d
# <翻译结束>


<原文开始>
// PUT registers an http handler to give the route pattern and the http method: PUT.
<原文结束>

# <翻译开始>
// PUT 注册一个 HTTP 处理器，用于处理给定的路由模式和 HTTP 方法：PUT。 md5:28ecbdff64685060
# <翻译结束>


<原文开始>
// POST registers an http handler to give the route pattern and the http method: POST.
<原文结束>

# <翻译开始>
// POST 注册一个http处理器，用于给路由模式和HTTP方法：POST。 md5:a251027c1c7a1d8c
# <翻译结束>


<原文开始>
// DELETE registers an http handler to give the route pattern and the http method: DELETE.
<原文结束>

# <翻译开始>
// DELETE 注册一个 http 处理器，用于给路由模式（pattern）和 http 方法：DELETE。 md5:b493fe2a753e0422
# <翻译结束>


<原文开始>
// PATCH registers an http handler to give the route pattern and the http method: PATCH.
<原文结束>

# <翻译开始>
// PATCH 注册一个HTTP处理器，给定路由模式和HTTP方法：PATCH。 md5:6662f45a2e57a836
# <翻译结束>


<原文开始>
// HEAD registers an http handler to give the route pattern and the http method: HEAD.
<原文结束>

# <翻译开始>
// HEAD 注册一个http处理器，用于指定路由模式和HTTP方法：HEAD。 md5:c1e170eaa1fe60b7
# <翻译结束>


<原文开始>
// CONNECT registers an http handler to give the route pattern and the http method: CONNECT.
<原文结束>

# <翻译开始>
// CONNECT 注册一个 http 处理器，用于指定路由模式和方法：CONNECT。 md5:01352b24b5b15d84
# <翻译结束>


<原文开始>
// OPTIONS register an http handler to give the route pattern and the http method: OPTIONS.
<原文结束>

# <翻译开始>
// OPTIONS 注册一个 http 处理器，用于指定路由模式和方法：OPTIONS。 md5:7c22cd8904d32b99
# <翻译结束>


<原文开始>
// TRACE registers an http handler to give the route pattern and the http method: TRACE.
<原文结束>

# <翻译开始>
// TRACE 注册一个HTTP处理程序，用于提供路由模式和HTTP方法：TRACE。 md5:530929842b31c7fa
# <翻译结束>


<原文开始>
// REST registers an http handler to give the route pattern according to REST rule.
<原文结束>

# <翻译开始>
// REST 根据 REST 规则注册一个 HTTP 处理器，以提供路由模式。 md5:b89313386e2f52de
# <翻译结束>


<原文开始>
// Hook registers a hook to given route pattern.
<原文结束>

# <翻译开始>
// Hook 注册一个钩子到给定的路由模式。 md5:1b98e351ffc870a2
# <翻译结束>


<原文开始>
// Middleware binds one or more middleware to the router group.
<原文结束>

# <翻译开始>
// Middleware 将一个或多个中间件绑定到路由器组。 md5:ba25a44638f73d20
# <翻译结束>


<原文开始>
// preBindToLocalArray adds the route registering parameters to an internal variable array for lazily registering feature.
<原文结束>

# <翻译开始>
// preBindToLocalArray 将路由注册参数添加到内部变量数组中，以便于惰性注册特性。 md5:0b2a8a31bb20bca1
# <翻译结束>


<原文开始>
// getPrefix returns the route prefix of the group, which recursively retrieves its parent's prefix.
<原文结束>

# <翻译开始>
// getPrefix 返回该组的路由前缀，该方法会递归地获取其父组的前缀。 md5:0d086cd9e63f6697
# <翻译结束>


<原文开始>
// doBindRoutersToServer does really register for the group.
<原文结束>

# <翻译开始>
// doBindRoutersToServer 确实为该组注册。 md5:436447cc3534e54c
# <翻译结束>


<原文开始>
// If there is already a domain, unset the domain field in the pattern.
<原文结束>

# <翻译开始>
// 如果已经有域，那么在模式中清除域字段。 md5:e02751d36da77b97
# <翻译结束>


<原文开始>
// Filter repeated char '/'.
<原文结束>

# <翻译开始>
// 过滤重复的字符 '/'。 md5:9b9a7539f6ae7305
# <翻译结束>


<原文开始>
// Convert params to a string array.
<原文结束>

# <翻译开始>
// 将参数转换为字符串数组。 md5:8388b98c9b261cad
# <翻译结束>


<原文开始>
// Check whether it's a hook handler.
<原文结束>

# <翻译开始>
// 检查它是否是钩子处理器。 md5:f6b816a5e567ae34
# <翻译结束>


<原文开始>
// Finally, it treats the `object` as the Object registering type.
<原文结束>

# <翻译开始>
// 最后，它将`object`视为注册类型的对象。 md5:1175240ff3996b3d
# <翻译结束>

