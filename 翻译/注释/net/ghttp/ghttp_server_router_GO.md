
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
// handlerIdGenerator is handler item id generator.
<原文结束>

# <翻译开始>
// handlerIdGenerator 是处理器项 ID 生成器。
# <翻译结束>


<原文开始>
// routerMapKey creates and returns a unique router key for given parameters.
// This key is used for Server.routerMap attribute, which is mainly for checks for
// repeated router registering.
<原文结束>

# <翻译开始>
// routerMapKey 根据给定的参数创建并返回一个唯一的路由键。
// 此键用于 Server.routerMap 属性，主要用于检查重复的路由注册情况。
# <翻译结束>


<原文开始>
// parsePattern parses the given pattern to domain, method and path variable.
<原文结束>

# <翻译开始>
// parsePattern 将给定的模式解析为域名、方法和路径变量。
# <翻译结束>


<原文开始>
// setHandler creates router item with a given handler and pattern and registers the handler to the router tree.
// The router tree can be treated as a multilayer hash table, please refer to the comment in the following codes.
// This function is called during server starts up, which cares little about the performance. What really cares
// is the well-designed router storage structure for router searching when the request is under serving.
<原文结束>

# <翻译开始>
// setHandler 根据给定的处理器和模式创建路由项，并将处理器注册到路由树中。
// 路由树可以被视为一个多层哈希表，请参考下文中的注释。
// 此函数在服务器启动时调用，对性能要求不高。真正重要的是
// 当请求处于服务状态时，用于路由搜索的良好设计的路由存储结构。
# <翻译结束>


<原文开始>
	// ====================================================================================
	// Change the registered route according to meta info from its request structure.
	// It supports multiple methods that are joined using char `,`.
	// ====================================================================================
<原文结束>

# <翻译开始>
// ====================================================================================
// 根据请求结构中的元信息更改已注册的路由。
// 它支持使用逗号 `,` 连接的多种方法。
// ====================================================================================
# <翻译结束>


<原文开始>
// Multiple methods registering, which are joined using char `,`.
<原文结束>

# <翻译开始>
// 多个方法注册，使用字符 `,` 连接。
# <翻译结束>


<原文开始>
// Each method has it own handler.
<原文结束>

# <翻译开始>
// 每个方法都有自己的处理程序。
# <翻译结束>












<原文开始>
// Repeated router checks, this feature can be disabled by server configuration.
<原文结束>

# <翻译开始>
// 重复路由检查，可以通过服务器配置禁用此功能。
# <翻译结束>


<原文开始>
// Unique id for each handler.
<原文结束>

# <翻译开始>
// 每个处理器的唯一标识 ID
# <翻译结束>


<原文开始>
// Create a new router by given parameter.
<原文结束>

# <翻译开始>
// 根据给定参数创建一个新的路由器。
# <翻译结束>


<原文开始>
	// List array, very important for router registering.
	// There may be multiple lists adding into this array when searching from root to leaf.
<原文结束>

# <翻译开始>
// List 数组，对于路由器注册非常重要。
// 从根节点到叶子节点搜索过程中，可能会有多个列表添加到此数组中。
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
// 多层哈希表:
// 1. 表中的每个节点通过以字符 '/' 分割的URI路径进行区分。
// 2. 键 "*fuzz" 指定该节点是一个模糊节点，没有特定名称。
// 3. 键 "*list" 是节点的列表项，**大多数节点都有此项**，特别是模糊节点。注意，模糊节点必须包含"*list"项，并且叶节点也具有"*list"项。如果节点既不是模糊节点也不是叶节点，则它不包含"*list"项。
// 4. "*list" 项是一个按其优先级从高到低排序的已注册路由项列表。如果是模糊节点，从此模糊节点开始的所有子路由项也将被添加到其"*list"项中。
// 5. 路由列表中可能存在重复的路由项。从根节点到叶节点的列表优先级是从低到高。
# <翻译结束>


<原文开始>
// Ignore empty URI part, like: /user//index
<原文结束>

# <翻译开始>
// 忽略空的URI部分，例如：/user//index
# <翻译结束>


<原文开始>
// Check if it's a fuzzy node.
<原文结束>

# <翻译开始>
// 检查是否为模糊节点。
# <翻译结束>


<原文开始>
			// If it's a fuzzy node, it creates a "*list" item - which is a list - in the hash map.
			// All the sub router items from this fuzzy node will also be added to its "*list" item.
<原文结束>

# <翻译开始>
// 如果这是一个模糊节点，它会在哈希映射中创建一个“*list”项——这是一个列表。
// 从此模糊节点派生的所有子路由项也将被添加到其“*list”项中。
# <翻译结束>


<原文开始>
// Make a new bucket for the current node.
<原文结束>

# <翻译开始>
// 为当前节点新建一个桶。
# <翻译结束>







<原文开始>
		// The leaf is a hash map and must have an item named "*list", which contains the router item.
		// The leaf can be furthermore extended by adding more ket-value pairs into its map.
		// Note that the `v != "*fuzz"` comparison is required as the list might be added in the former
		// fuzzy checks.
<原文结束>

# <翻译开始>
// 叶子节点是一个哈希表，且必须包含一个名为"*list"的项目，其中存储着路由项。
// 通过向其映射中添加更多键值对，叶子节点可以进一步扩展。
// 注意，由于在之前的模糊检查中可能已添加了列表，所以需要进行 `v != "*fuzz"` 的比较。
# <翻译结束>


<原文开始>
	// It iterates the list array of `lists`, compares priorities and inserts the new router item in
	// the proper position of each list. The priority of the list is ordered from high to low.
<原文结束>

# <翻译开始>
// 它遍历`lists`列表数组，比较优先级并将新的路由项插入到每个列表中的适当位置。列表的优先级从高到低排序。
# <翻译结束>


<原文开始>
			// Checks the priority whether inserting the route item before current item,
			// which means it has higher priority.
<原文结束>

# <翻译开始>
// 检查优先级，是否在当前项之前插入路由项，
// 这意味着它具有更高的优先级。
# <翻译结束>







<原文开始>
// Initialize the route map item.
<原文结束>

# <翻译开始>
// 初始化路由映射项。
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
// compareRouterPriority 比较 `newItem` 和 `oldItem` 之间的优先级。如果 `newItem` 的优先级高于 `oldItem`，则返回 true，否则返回 false。优先级较高的项将被插入到路由器列表的前面。
//
// 比较规则：
// 1. 中间件具有最高的优先级。
// 2. URI：路径越深，优先级越高（简单地检查 URI 中字符 '/' 的数量）。
// 3. 路由类型：{xxx} > :xxx > *xxx。
# <翻译结束>


<原文开始>
// If they're all types of middleware, the priority is according to their registered sequence.
<原文结束>

# <翻译开始>
// 如果它们都是中间件类型，则优先级根据其注册顺序决定。
# <翻译结束>


<原文开始>
// The middleware has the most high priority.
<原文结束>

# <翻译开始>
// 该中间件具有最高优先级。
# <翻译结束>


<原文开始>
// URI: The deeper, the higher (simply check the count of char '/' in the URI).
<原文结束>

# <翻译开始>
// URI：URI中'/'字符出现次数越多，级别越高。
# <翻译结束>


<原文开始>
	// Compare the length of their URI,
	// but the fuzzy and named parts of the URI are not calculated to the result.
<原文结束>

# <翻译开始>
// 比较它们URI的长度，
// 但URI中的模糊部分和命名部分不计算到结果中。
# <翻译结束>


<原文开始>
	// Example:
	// /admin-goods-{page} > /admin-{page}
	// /{hash}.{type}      > /{hash}
<原文结束>

# <翻译开始>
// 示例：
// /admin-goods-{page} > /admin-{page}
// /{hash}.{type}      > /{hash}
// 上面的注释是用于示例URL路径重写规则：
// 第一条规则表示将 "/admin-goods-任意页码" 重写为 "/admin-任意页码"，其中 {page} 是一个占位符，代表任何数字页码。
// 第二条规则表示将 "/任意哈希值.任意类型" 重写为 "/任意哈希值"，其中 {hash} 和 {type} 分别是占位符，代表任何哈希值和文件类型。
# <翻译结束>


<原文开始>
	// Route type checks: {xxx} > :xxx > *xxx.
	// Example:
	// /name/act > /{name}/:act
<原文结束>

# <翻译开始>
// 路由类型检查规则：{xxx} > :xxx > *xxx。
// 示例：
// /name/act 对应于 /{name}/:act
// （注释翻译：这段Go语言代码的注释描述了路由路径匹配的优先级规则，其中花括号 `{}`、冒号 `:` 以及星号 `*` 分别用于表示路径参数的不同格式。按照优先级从高到低排列为：`{xxx}`（命名路径参数）、`:xxx`（动态路径参数）和 `*xxx`（任意长度路径参数）。示例说明了这种对应关系，在路由 `/name/act` 中，`name` 可以映射为 `{name}` 形式的命名路径参数，而 `act` 可以映射为 `:act` 形式的动态路径参数。）
# <翻译结束>


<原文开始>
// If the counts of their fuzzy rules are equal.
<原文结束>

# <翻译开始>
// 如果它们的模糊规则计数相等。
# <翻译结束>


<原文开始>
// Eg: /name/{act} > /name/:act
<原文结束>

# <翻译开始>
// 示例：/name/{act} 转换为 /name/:act
# <翻译结束>


<原文开始>
// Eg: /name/:act > /name/*act
<原文结束>

# <翻译开始>
// 示例：/name/:act > /name/*act
// 注释翻译：该注释用于表示一个路由映射规则的示例，其中"/name/:act"是一个动态路由模板，":act"是一个参数占位符，可以匹配任何非空字符串。在实际应用中，它将被映射到类似"/name/任意动作名"的实际路由路径。例如，如果":act"为"edit"，则此规则可匹配路径"/name/edit"。而"*act"则通常用于匹配任意后缀，包括"/"字符及其后面的所有内容。但在给定的代码片段中，并没有明确指出"*act"的行为，可能需要更多上下文信息来准确解释。
# <翻译结束>


<原文开始>
	// It then compares the accuracy of their http method,
	// the more accurate the more priority.
<原文结束>

# <翻译开始>
// 然后比较它们的HTTP方法的准确性，
// 越准确则优先级越高。
# <翻译结束>


<原文开始>
	// If they have different router type,
	// the new router item has more priority than the other one.
<原文结束>

# <翻译开始>
// 如果它们具有不同的路由类型，
// 那么新的路由项比另一个具有更高的优先级。
# <翻译结束>


<原文开始>
	// Other situations, like HOOK items,
	// the old router item has more priority than the other one.
<原文结束>

# <翻译开始>
// 其他情况，如 HOOK 项，
// 则旧的路由项比其他项具有更高的优先级。
# <翻译结束>


<原文开始>
// patternToRegular converts route rule to according to regular expression.
<原文结束>

# <翻译开始>
// patternToRegular 将路由规则转换为相应的正则表达式。
# <翻译结束>












<原文开始>
// Converts `all` to `ALL`.
<原文结束>

# <翻译开始>
// 将`all`转换为`ALL`。
# <翻译结束>


<原文开始>
// Prefix for URI feature.
<原文结束>

# <翻译开始>
// URI特征的前缀。
# <翻译结束>


<原文开始>
// Loop to next bucket.
<原文结束>

# <翻译开始>
// 循环到下一个桶。
# <翻译结束>


<原文开始>
// Just push back in default.
<原文结束>

# <翻译开始>
// 默认情况下，仅向后推入
# <翻译结束>


<原文开始>
// Replace "/*" and "/*any".
<原文结束>

# <翻译开始>
// 将 "/*" 和 "/*any" 进行替换。
# <翻译结束>


<原文开始>
// Special chars replacement.
<原文结束>

# <翻译开始>
// 特殊字符替换。
# <翻译结束>

