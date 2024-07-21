
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
// handlerCacheItem is an item just for internal router searching cache.
<原文结束>

# <翻译开始>
// handlerCacheItem 是仅用于内部路由器搜索缓存的项。 md5:bff6700a37e67c6b
# <翻译结束>


<原文开始>
// serveHandlerKey creates and returns a handler key for router.
<原文结束>

# <翻译开始>
// serveHandlerKey 为路由器创建并返回一个处理器键。 md5:a4cf69fa7df9d5ac
# <翻译结束>


<原文开始>
// getHandlersWithCache searches the router item with cache feature for a given request.
<原文结束>

# <翻译开始>
// getHandlersWithCache 为给定的请求搜索具有缓存功能的路由项。 md5:00b96b129dd9a5f8
# <翻译结束>


<原文开始>
	// In case of, eg:
	// Case 1:
	// 		GET /net/http
	// 		r.URL.Path    : /net/http
	// 		r.URL.RawPath : (empty string)
	// Case 2:
	// 		GET /net%2Fhttp
	// 		r.URL.Path    : /net/http
	// 		r.URL.RawPath : /net%2Fhttp
<原文结束>

# <翻译开始>
	// 在以下情况中，例如：
	// 情况1：
	// 		GET /net/http
	// 		r.URL.Path    : /net/http
	// 		r.URL.RawPath : （空字符串）
	// 情况2：
	// 		GET /net%2Fhttp
	// 		r.URL.Path    : /net/http
	// 		r.URL.RawPath : /net%2Fhttp
	// md5:97750eaa6ac9d07d
# <翻译结束>


<原文开始>
	// Special http method OPTIONS handling.
	// It searches the handler with the request method instead of OPTIONS method.
<原文结束>

# <翻译开始>
	// 专门处理 HTTP 方法 OPTIONS。
	// 它会使用请求方法搜索处理器，而不是 OPTIONS 方法。
	// md5:2704b13524189224
# <翻译结束>


<原文开始>
// Search and cache the router handlers.
<原文结束>

# <翻译开始>
// 搜索并缓存路由器处理器。 md5:7263da0c149e9280
# <翻译结束>


<原文开始>
// searchHandlers retrieve and returns the routers with given parameters.
// Note that the returned routers contain serving handler, middleware handlers and hook handlers.
<原文结束>

# <翻译开始>
// searchHandlers 根据给定的参数检索并返回路由器。
// 注意，返回的路由器包含了服务处理程序、中间件处理程序和钩子处理程序。
// md5:c8f076ede0fbe806
# <翻译结束>


<原文开始>
	// In case of double '/' URI, for example:
	// /user//index, //user/index, //user//index//
<原文结束>

# <翻译开始>
	// 对于包含连续'/'的URI，例如：
	// /user	//index, 	//user/index, 	//user	//index	//
	// md5:fb272e4928c6b465
# <翻译结束>


<原文开始>
// Split the URL.path to separate parts.
<原文结束>

# <翻译开始>
// 将URL.path分割成多个部分。 md5:421153d9f0413872
# <翻译结束>


<原文开始>
	// The default domain has the most priority when iteration.
	// Please see doSetHandler if you want to get known about the structure of serveTree.
<原文结束>

# <翻译开始>
	// 当迭代时，默认域具有最高优先级。如果您想了解serveTree的结构，请参阅doSetHandler。
	// md5:8bc20bbd07335cfd
# <翻译结束>


<原文开始>
// Make a list array with a capacity of 16.
<原文结束>

# <翻译开始>
// 创建一个容量为16的列表数组。 md5:9ce7c6d246550dea
# <翻译结束>


<原文开始>
// Add all lists of each node to the list array.
<原文结束>

# <翻译开始>
// 将每个节点的所有列表添加到列表数组中。 md5:82a101859541f52e
# <翻译结束>


<原文开始>
// Loop to the next node by certain key name.
<原文结束>

# <翻译开始>
// 通过指定的关键字名称，循环到下一个节点。 md5:e9bddd4258b62329
# <翻译结束>


<原文开始>
// Loop to the next node by fuzzy node item.
<原文结束>

# <翻译开始>
// 通过模糊节点项遍历到下一个节点。 md5:fe1b87e1d17d2d0f
# <翻译结束>


<原文开始>
				// It here also checks the fuzzy item,
				// for rule case like: "/user/*action" matches to "/user".
<原文结束>

# <翻译开始>
				// 这里同时检查模糊项，
				// 适用于诸如规则情况："/user/*action" 匹配到 "/user"。
				// md5:89d31460cfdbd8e6
# <翻译结束>


<原文开始>
// The leaf must have a list item. It adds the list to the list array.
<原文结束>

# <翻译开始>
// 叶子节点必须有一个列表项。它将这个列表添加到列表数组中。 md5:dbc074656c73362f
# <翻译结束>


<原文开始>
		// OK, let's loop the result list array, adding the handler item to the result handler result array.
		// As the tail of the list array has the most priority, it iterates the list array from its tail to head.
<原文结束>

# <翻译开始>
		// 好的，让我们遍历结果列表数组，将处理项添加到结果处理器结果数组中。由于列表数组的尾部优先级最高，所以我们从数组尾部开始向前遍历。
		// md5:1f7f116128551404
# <翻译结束>


<原文开始>
				// Filter repeated handler items, especially the middleware and hook handlers.
				// It is necessary, do not remove this checks logic unless you really know how it is necessary.
				//
				// The `repeatHandlerCheckMap` is used for repeat handler filtering during handler searching.
				// As there are fuzzy nodes, and the fuzzy nodes have both sub-nodes and sub-list nodes, there
				// may be repeated handler items in both sub-nodes and sub-list nodes. It here uses handler item id to
				// identify the same handler item that registered.
				//
				// The same handler item is the one that is registered in the same function doSetHandler.
				// Note that, one handler function(middleware or hook function) may be registered multiple times as
				// different handler items using function doSetHandler, and they have different handler item id.
				//
				// Note that twice, the handler function may be registered multiple times as different handler items.
<原文结束>

# <翻译开始>
				// 过滤重复的处理器项，特别是中间件和钩子处理器。
				// 这是必要的，除非你非常清楚为什么需要移除这个检查逻辑，否则请不要删除。
				//
				// `repeatHandlerCheckMap` 用于在搜索处理器时进行重复处理器过滤。由于存在模糊节点，这些模糊节点既有子节点也有子列表节点，因此可能会在子节点和子列表节点中出现重复的处理器项。
				//
				// 同一个处理器项是指使用 `doSetHandler` 函数在同一函数中注册的处理器。需要注意的是，一个处理器函数（中间件或钩子函数）可能通过 `doSetHandler` 函数以不同的处理器项多次注册，并且它们有不同的处理器项 ID。
				//
				// 另外需要注意，同一种处理器函数可能由于不同的处理目的而被多次注册为不同的处理器项。
				// md5:6e4536c4e013b86a
# <翻译结束>


<原文开始>
				// Serving handler can only be added to the handler array just once.
				// The first route item in the list has the most priority than the rest.
				// This ignoring can implement route overwritten feature.
<原文结束>

# <翻译开始>
				// 服务处理程序只能添加到处理器数组中一次。
				// 列表中的第一个路由项比其余项具有更高的优先级。
				// 此忽略功能可以实现路由覆盖功能。
				// md5:6e93290e1cdad8d9
# <翻译结束>


<原文开始>
// Note the rule having no fuzzy rules: len(match) == 1
<原文结束>

# <翻译开始>
// 注意没有模糊规则的规则：match 的长度等于 1. md5:c26d1818ce3f384e
# <翻译结束>


<原文开始>
						// If the rule contains fuzzy names,
						// it needs paring the URL to retrieve the values for the names.
<原文结束>

# <翻译开始>
						// 如果规则包含模糊名称（fuzzy names），
						// 需要对URL进行切分以获取名称的值。
						// md5:022aca8d52d2dc1f
# <翻译结束>


<原文开始>
// It there repeated names, it just overwrites the same one.
<原文结束>

# <翻译开始>
//如果有重复的名称，它就会覆盖相同的名称。 md5:afb894e9dbad1062
# <翻译结束>


<原文开始>
// The serving handler can be added just once.
<原文结束>

# <翻译开始>
// 服务处理程序只能添加一次。 md5:fef3170c186d44cb
# <翻译结束>


<原文开始>
						// The middleware is inserted before the serving handler.
						// If there are multiple middleware, they're inserted into the result list by their registering order.
						// The middleware is also executed by their registered order.
<原文结束>

# <翻译开始>
						// 中间件在服务处理器之前插入。
						// 如果有多个中间件，它们会按照注册的顺序插入到结果列表中。
						// 中间件也会按照注册的顺序执行。
						// md5:3ae9ef9a044965f3
# <翻译结束>


<原文开始>
// HOOK handler, just push it back to the list.
<原文结束>

# <翻译开始>
// HOOK处理器，只是将其推回列表。 md5:5c3afbdb8ce6826c
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
# <翻译结束>

