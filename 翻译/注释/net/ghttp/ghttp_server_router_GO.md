
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// handlerIdGenerator is handler item id generator.
<原文结束>

# <翻译开始>
	// handlerIdGenerator 是处理器项ID生成器。 md5:0a9e55b1609ba4df
# <翻译结束>


<原文开始>
// routerMapKey creates and returns a unique router key for given parameters.
// This key is used for Server.routerMap attribute, which is mainly for checks for
// repeated router registering.
<原文结束>

# <翻译开始>
// routerMapKey 为给定参数创建并返回一个唯一的路由键。这个键用于 `Server.routerMap` 属性，主要用于检查重复的路由注册。 md5:0a5f0d744a55d4ed
# <翻译结束>


<原文开始>
// parsePattern parses the given pattern to domain, method and path variable.
<原文结束>

# <翻译开始>
// parsePattern 将给定的模式解析为域名、方法和路径变量。 md5:9f5177d72b0e5cf6
# <翻译结束>


<原文开始>
// setHandler creates router item with a given handler and pattern and registers the handler to the router tree.
// The router tree can be treated as a multilayer hash table, please refer to the comment in the following codes.
// This function is called during server starts up, which cares little about the performance. What really cares
// is the well-designed router storage structure for router searching when the request is under serving.
<原文结束>

# <翻译开始>
// setHandler 根据给定的处理器和模式创建路由项，并将处理器注册到路由树中。
// 路由树可以看作是一个多层哈希表，请参考下面代码中的注释。
// 此函数在服务器启动时被调用，此时对性能要求不高。真正重要的是设计良好的
// 路由存储结构，以便在处理请求时能够高效地进行路由查找。 md5:325f9f3a1c077ca7
# <翻译结束>


<原文开始>
	// ====================================================================================
	// Change the registered route according to meta info from its request structure.
	// It supports multiple methods that are joined using char `,`.
	// ====================================================================================
<原文结束>

# <翻译开始>
	// ====================================================================================
	// 根据请求结构体中的元信息更改注册的路由。
	// 它支持使用字符 ',' 连接的多个方法。
	// ==================================================================================== md5:1ec1db24aded2c53
# <翻译结束>


<原文开始>
// Multiple methods registering, which are joined using char `,`.
<原文结束>

# <翻译开始>
		// 使用字符`,`连接的多个方法注册。 md5:8edb2f5feed892c9
# <翻译结束>


<原文开始>
// Each method has it own handler.
<原文结束>

# <翻译开始>
				// 每个方法都有自己的处理器。 md5:006ab83dd8178a73
# <翻译结束>


<原文开始>
// Converts `all` to `ALL`.
<原文结束>

# <翻译开始>
		// 将`all`转换为`ALL`。 md5:85c2f9ce5460fdd6
# <翻译结束>


<原文开始>
// Prefix for URI feature.
<原文结束>

# <翻译开始>
	// URI功能的前缀。 md5:7ec2c5614dbd89a6
# <翻译结束>


<原文开始>
// Repeated router checks, this feature can be disabled by server configuration.
<原文结束>

# <翻译开始>
	// 重复的路由检查，这个功能可以通过服务器配置来禁用。 md5:16d9ca5ef5f6ce27
# <翻译结束>


<原文开始>
// Unique id for each handler.
<原文结束>

# <翻译开始>
	// 每个处理器的唯一标识符。 md5:d5cdd6ccf90c625e
# <翻译结束>


<原文开始>
// Create a new router by given parameter.
<原文结束>

# <翻译开始>
	// 根据给定的参数创建一个新的路由。 md5:a6e213f025b1718b
# <翻译结束>


<原文开始>
	// List array, very important for router registering.
	// There may be multiple lists adding into this array when searching from root to leaf.
<原文结束>

# <翻译开始>
	// List数组，对路由器注册非常重要。
	// 在从根到叶的搜索过程中，可能会有多个列表添加到这个数组中。 md5:7ddaff62bcec3109
# <翻译结束>


<原文开始>
	// Multilayer hash table:
	// 1. Each node of the table is separated by URI path which is split by char '/'.
	// 2. The key "*fuzz" specifies this node is a fuzzy node, which has no certain name.
	// 3. The key "*list" is the list item of the node, MOST OF THE NODES HAVE THIS ITEM,
	//    especially the fuzzy node. NOTE THAT the fuzzy node must have the "*list" item,
	//    and the leaf node also has "*list" item. If the node is not a fuzzy node either
	//    a leaf, it neither has "*list" item.
	// 2. The "*list" item is a list containing registered router items ordered by their
	//    priorities from high to low. If it's a fuzzy node, all the sub router items
	//    from this fuzzy node will also be added to its "*list" item.
	// 3. There may be repeated router items in the router lists. The lists' priorities
	//    from root to leaf are from low to high.
<原文结束>

# <翻译开始>
	// 多层哈希表：
	// 1. 表中的每个节点由以字符 '/' 分割的 URI 路径标识。
	// 2. 键 "*fuzz" 指示这是一个模糊节点，它没有确定的名字。
	// 3. 键 "*list" 是节点的列表项，大多数节点都有这个项，特别是模糊节点。注意：模糊节点必须有 "*list" 项，叶子节点也有 "*list" 项。如果节点既不是模糊节点也不是叶子节点，则不包含 "*list" 项。
	// 4. "*list" 项是一个按优先级从高到低排序的已注册路由项的列表。如果是模糊节点，该模糊节点的所有子路由项也会添加到其 "*list" 项中。
	// 5. 路由列表中可能存在重复的路由项。从根到叶的列表优先级是从低到高。 md5:3b9d86c224bf6153
# <翻译结束>


<原文开始>
// Ignore empty URI part, like: /user//index
<原文结束>

# <翻译开始>
		//index. md5:44ed3114aa11886a
# <翻译结束>


<原文开始>
// Check if it's a fuzzy node.
<原文结束>

# <翻译开始>
		// 检查是否为模糊节点。 md5:ea4491ebe7a6c626
# <翻译结束>


<原文开始>
			// If it's a fuzzy node, it creates a "*list" item - which is a list - in the hash map.
			// All the sub router items from this fuzzy node will also be added to its "*list" item.
<原文结束>

# <翻译开始>
			// 如果它是一个模糊节点，它会在哈希映射中创建一个"*list"项，这实际上是一个列表。
			// 该模糊节点下的所有子路由器项也将被添加到它的"*list"项中。 md5:31e4feee2e295113
# <翻译结束>


<原文开始>
// Make a new bucket for the current node.
<原文结束>

# <翻译开始>
		// 为当前节点创建一个新的桶。 md5:a99064964800f461
# <翻译结束>


<原文开始>
		// The leaf is a hash map and must have an item named "*list", which contains the router item.
		// The leaf can be furthermore extended by adding more ket-value pairs into its map.
		// Note that the `v != "*fuzz"` comparison is required as the list might be added in the former
		// fuzzy checks.
<原文结束>

# <翻译开始>
		// 叶节点是一个哈希映射，必须包含一个名为"*list"的项，其中包含路由项。
		// 叶节点可以通过在其映射中添加更多的键值对来进一步扩展。
		// 请注意，需要进行 `v != "*fuzz"` 的比较，因为列表可能在先前的模糊检查中被添加。 md5:0a1026e07b9b2544
# <翻译结束>


<原文开始>
	// It iterates the list array of `lists`, compares priorities and inserts the new router item in
	// the proper position of each list. The priority of the list is ordered from high to low.
<原文结束>

# <翻译开始>
	// 它遍历`lists`的列表数组，比较优先级，并将新的路由项插入到每个列表的适当位置。
	// 列表的优先级从高到低排序。 md5:f7e3738ec2e01b79
# <翻译结束>


<原文开始>
			// Checks the priority whether inserting the route item before current item,
			// which means it has higher priority.
<原文结束>

# <翻译开始>
			// 检查是否应在当前项之前插入路由项，即它具有更高的优先级。 md5:0e6fc2994f00bc96
# <翻译结束>


<原文开始>
// Just push back in default.
<原文结束>

# <翻译开始>
		// 就默认再推回去。 md5:eba32b1a9fbdfd1f
# <翻译结束>


<原文开始>
// Initialize the route map item.
<原文结束>

# <翻译开始>
	// 初始化路由项。 md5:6ae2dff0c163c17e
# <翻译结束>


<原文开始>
// compareRouterPriority compares the priority between `newItem` and `oldItem`. It returns true
// if `newItem`'s priority is higher than `oldItem`, else it returns false. The higher priority
// item will be inserted into the router list before the other one.
//
// Comparison rules:
// 1. The middleware has the most high priority.
// 2. URI: The deeper, the higher (simply check the count of char '/' in the URI).
// 3. Route type: {xxx} > :xxx > *xxx.
<原文结束>

# <翻译开始>
// compareRouterPriority 比较 `newItem` 和 `oldItem` 之间的优先级。如果 `newItem` 的优先级高于 `oldItem`，则返回 true；否则返回 false。优先级较高的项将被插入到路由列表的前面。
//
// 比较规则：
// 1. 中间件具有最高的优先级。
// 2. URI：深度越深，优先级越高（简单地检查 URI 中字符 '/' 的数量）。
// 3. 路由类型：{xxx} > :xxx > *xxx。 md5:d3f2e1aac7e71a05
# <翻译结束>


<原文开始>
// If they're all types of middleware, the priority is according to their registered sequence.
<原文结束>

# <翻译开始>
	// 如果它们都是中间件类型，则按照注册的顺序决定优先级。 md5:3a53e273b3f3566f
# <翻译结束>


<原文开始>
// The middleware has the most high priority.
<原文结束>

# <翻译开始>
	// 中间件具有最高的优先级。 md5:0ca14429f8f96e06
# <翻译结束>


<原文开始>
// URI: The deeper, the higher (simply check the count of char '/' in the URI).
<原文结束>

# <翻译开始>
	// URI：深度越深（只需检查URI中'/'字符的数量）。 md5:cf10a2d60b6808df
# <翻译结束>


<原文开始>
	// Compare the length of their URI,
	// but the fuzzy and named parts of the URI are not calculated to the result.
<原文结束>

# <翻译开始>
	// 比较它们的URI长度，
	// 但URI中的模糊部分和命名部分不计入结果中。 md5:55bd5729f8c0352a
# <翻译结束>


<原文开始>
	// Example:
	// /admin-goods-{page} > /admin-{page}
	// /{hash}.{type}      > /{hash}
<原文结束>

# <翻译开始>
	// 示例：
	// /admin-goods-{分页}  > /admin-{分页}
	// /{哈希}.{类型}      > /{哈希} md5:482c38c410b3c591
# <翻译结束>


<原文开始>
// Replace "/*" and "/*any".
<原文结束>

# <翻译开始>
// 替换 "/*" 和 "任何字符串"。 md5:4bbaf5031e185545
# <翻译结束>


<原文开始>
	// Route type checks: {xxx} > :xxx > *xxx.
	// Example:
	// /name/act > /{name}/:act
<原文结束>

# <翻译开始>
	// 路由类型检查：{xxx} > :xxx > *xxx。
	// 例子：
	// /name/act > /{name}/:act
	//
	// 这段注释的意思是，它在描述Go语言中的路由类型检查规则。`{xxx}`、`:xxx`和`*xxx`是路由匹配模式：
	//
	// - `{xxx}` 表示路径中可以包含任意字符的占位符，但需要与实际请求中的某个参数匹配。
	// - `:xxx` 表示路径中的命名参数，这些参数将在路由处理函数中作为变量传递。
	// - `*xxx` 表示零个或多个重复的前面的模式，通常用于处理路径中的可选组件。
	//
	// 举例来说，路由`/name/act` 使用了`{name}`和`:act`，表示请求的URL可以形式为`/具体名称/操作名`，`{name}`会被替换为实际请求中的名称，`:act`则是一个动态的操作标识。 md5:5fc64b4a4a78b2aa
# <翻译结束>


<原文开始>
// If the counts of their fuzzy rules are equal.
<原文结束>

# <翻译开始>
	// 如果它们的模糊规则数量相等。 md5:0a1cd4da270f5da3
# <翻译结束>


<原文开始>
// Eg: /name/{act} > /name/:act
<原文结束>

# <翻译开始>
	// 例如：/name/{act} > /name/:act. md5:14051818a0cea80c
# <翻译结束>


<原文开始>
// Eg: /name/:act > /name/*act
<原文结束>

# <翻译开始>
	// 这段注释的意思是，当URL路径匹配模式"/name/:act"时，它会被转换或重写为"/name/*act"。这里的":act"是一个占位符，表示可以包含任意字符的动态部分，"*act"则表示任何字符序列（包括零个字符）。这是一种路由或路径匹配的规则，在Go或其他支持类似语法的语言中常见于处理URL路由。 md5:6f9e027f06c4b833
# <翻译结束>


<原文开始>
	// It then compares the accuracy of their http method,
	// the more accurate the more priority.
<原文结束>

# <翻译开始>
	// 然后，它会比较它们的HTTP方法的准确性，越准确优先级越高。 md5:19e263d51107b5cb
# <翻译结束>


<原文开始>
	// If they have different router type,
	// the new router item has more priority than the other one.
<原文结束>

# <翻译开始>
	// 如果它们具有不同的路由类型，
	// 那么新的路由项比其他项具有更高的优先级。 md5:63dfba3b91db8cc4
# <翻译结束>


<原文开始>
	// Other situations, like HOOK items,
	// the old router item has more priority than the other one.
<原文结束>

# <翻译开始>
	// 其他情况，如HOOK项目，
	// 旧的路由项具有更高的优先级。 md5:53b3ce09282d12db
# <翻译结束>


<原文开始>
// patternToRegular converts route rule to according to regular expression.
<原文结束>

# <翻译开始>
// patternToRegular 将路由规则转换为相应的正则表达式。 md5:c212402d9fd8cb59
# <翻译结束>


<原文开始>
// Special chars replacement.
<原文结束>

# <翻译开始>
			// 特殊字符替换。 md5:fe1b718da00180dd
# <翻译结束>

