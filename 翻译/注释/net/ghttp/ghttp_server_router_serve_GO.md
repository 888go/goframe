
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
// handlerCacheItem is an item just for internal router searching cache.
<原文结束>

# <翻译开始>
// handlerCacheItem 是一个仅用于内部路由器搜索缓存的项目。
# <翻译结束>


<原文开始>
// serveHandlerKey creates and returns a handler key for router.
<原文结束>

# <翻译开始>
// serveHandlerKey 为路由器创建并返回一个处理程序键。
# <翻译结束>


<原文开始>
// getHandlersWithCache searches the router item with cache feature for a given request.
<原文结束>

# <翻译开始>
// getHandlersWithCache 根据给定的请求，搜索具有缓存功能的路由器项。
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
// 举例说明：
// 情况1：
//   GET /net/http
//   r.URL.Path    : /net/http
//   r.URL.RawPath : （空字符串）
// 情况2：
//   GET /net%2Fhttp
//   r.URL.Path    : /net/http
//   r.URL.RawPath : /net%2Fhttp
// 在上述情况下，r 是一个 *http.Request 类型的对象，
// r.URL.Path 表示已解码的请求路径（在情况1和情况2中都显示为"/net/http"）。
// 而 r.URL.RawPath 则表示原始编码格式的请求路径，
// 在情况1中由于未进行编码，所以是空字符串；
// 在情况2中，原始路径经过URL编码，"%2F"代表"/"，因此RawPath为"/net%2Fhttp"。
# <翻译结束>


<原文开始>
	// Special http method OPTIONS handling.
	// It searches the handler with the request method instead of OPTIONS method.
<原文结束>

# <翻译开始>
// 特殊处理HTTP方法OPTIONS。
// 它会按照请求的方法而非OPTIONS方法去搜索处理器。
# <翻译结束>


<原文开始>
// Search and cache the router handlers.
<原文结束>

# <翻译开始>
// 搜索并缓存路由处理器。
# <翻译结束>


<原文开始>
// searchHandlers retrieve and returns the routers with given parameters.
// Note that the returned routers contain serving handler, middleware handlers and hook handlers.
<原文结束>

# <翻译开始>
// searchHandlers 函数根据给定的参数获取并返回路由器。
// 注意，返回的路由器中包含了服务处理程序、中间件处理程序以及钩子处理程序。
# <翻译结束>


<原文开始>
	// In case of double '/' URI, for example:
	// /user//index, //user/index, //user//index//
<原文结束>

# <翻译开始>
// 在出现双斜线（/）的URI情况下，例如：
// /user//index,/user/index,/user//index//
// （注：这些示例中的双斜线表示连续的路径分隔符，Go语言中会进行规范化处理，将多个连续的斜线视为一个斜线。）
// ```go
// 对于如下的含有双斜线（/）的URI情况：
// 比如 "/user//index", "/user/index", "/user//index//" 等
# <翻译结束>


<原文开始>
// Split the URL.path to separate parts.
<原文结束>

# <翻译开始>
// 将URL.path拆分为多个部分。
# <翻译结束>


<原文开始>
	// The default domain has the most priority when iteration.
	// Please see doSetHandler if you want to get known about the structure of serveTree.
<原文结束>

# <翻译开始>
// 当进行迭代时，默认域具有最高优先级。
// 如果您想了解serveTree的结构，请参阅doSetHandler。
# <翻译结束>


<原文开始>
// Make a list array with a capacity of 16.
<原文结束>

# <翻译开始>
// 创建一个容量为16的list数组
# <翻译结束>


<原文开始>
// Add all lists of each node to the list array.
<原文结束>

# <翻译开始>
// 将每个节点的所有列表添加到list数组中。
# <翻译结束>


<原文开始>
// Loop to the next node by certain key name.
<原文结束>

# <翻译开始>
// 通过指定的键名循环到下一个节点。
# <翻译结束>


<原文开始>
// Loop to the next node by fuzzy node item.
<原文结束>

# <翻译开始>
// 通过模糊节点项循环到下一个节点。
# <翻译结束>


<原文开始>
				// It here also checks the fuzzy item,
				// for rule case like: "/user/*action" matches to "/user".
<原文结束>

# <翻译开始>
// 此处同时检查模糊匹配项，
// 例如：规则 "/user/*action" 匹配到 "/user" 的情况。
# <翻译结束>


<原文开始>
// The leaf must have a list item. It adds the list to the list array.
<原文结束>

# <翻译开始>
// 叶子节点必须包含一个列表项。它将该列表添加到列表数组中。
# <翻译结束>


<原文开始>
		// OK, let's loop the result list array, adding the handler item to the result handler result array.
		// As the tail of the list array has the most priority, it iterates the list array from its tail to head.
<原文结束>

# <翻译开始>
// 好的，让我们循环遍历结果列表数组，
// 将处理器项添加到结果处理器结果数组中。
// 由于列表数组的尾部具有最高的优先级，
// 因此从尾部到头部遍历列表数组。
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
// 这是必要的，请不要移除这个检查逻辑，除非你确实了解其必要性。
//
// `repeatHandlerCheckMap` 用于在搜索处理器时进行重复处理器项过滤。
// 由于存在模糊节点，且模糊节点同时包含子节点和子列表节点，在子节点和子列表节点中可能存在重复的处理器项。这里使用处理器项ID来识别注册时相同的处理器项。
//
// 同样的处理器项是指在同一个 doSetHandler 函数中注册的项。
// 注意，一个处理器函数（中间件或钩子函数）可能通过 function doSetHandler 以不同的处理器项形式被多次注册，它们具有不同的处理器项ID。
//
// 特别注意，处理器函数可能作为不同的处理器项被多次注册。
# <翻译结束>


<原文开始>
				// Serving handler can only be added to the handler array just once.
				// The first route item in the list has the most priority than the rest.
				// This ignoring can implement route overwritten feature.
<原文结束>

# <翻译开始>
// 服务处理程序只能一次性添加到处理器数组中。
// 列表中的第一条路由项比其余项具有更高的优先级。
// 这种忽略方式可以实现路由覆盖功能。
# <翻译结束>


<原文开始>
// Note the rule having no fuzzy rules: len(match) == 1
<原文结束>

# <翻译开始>
// 注意这个规则：没有模糊规则时，匹配项的长度为1
# <翻译结束>


<原文开始>
						// If the rule contains fuzzy names,
						// it needs paring the URL to retrieve the values for the names.
<原文结束>

# <翻译开始>
// 如果规则中包含模糊名称，
// 则需要解析URL以获取这些名称的值。
# <翻译结束>


<原文开始>
// It there repeated names, it just overwrites the same one.
<原文结束>

# <翻译开始>
// 如果存在重复的名字，它只会覆盖相同的那个。
# <翻译结束>


<原文开始>
// The serving handler can be added just once.
<原文结束>

# <翻译开始>
// 服务处理程序只能添加一次。
# <翻译结束>


<原文开始>
						// The middleware is inserted before the serving handler.
						// If there are multiple middleware, they're inserted into the result list by their registering order.
						// The middleware is also executed by their registered order.
<原文结束>

# <翻译开始>
// 中文注释：
// 中间件在服务处理程序之前插入。
// 如果存在多个中间件，它们会按照注册顺序插入到结果列表中。
// 中间件也会按照注册时的顺序执行。
// 这段Go语言代码注释翻译成中文后为：
// ```markdown
// 中间件在实际服务处理程序之前被插入。
// 若存在多个中间件，它们将根据注册顺序依次插入到结果列表中。
// 同样地，这些中间件也是按照其注册时的顺序来执行的。
# <翻译结束>


<原文开始>
// HOOK handler, just push it back to the list.
<原文结束>

# <翻译开始>
// HOOK 处理函数，只需将其推回列表中。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
# <翻译结束>

