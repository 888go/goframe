
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
// BTree holds elements of the B-tree.
<原文结束>

# <翻译开始>
// BTree 存储 B 树的元素。 md5:191d2e09c9c918ab
# <翻译结束>


<原文开始>
// Total number of keys in the tree
<原文结束>

# <翻译开始>
//树中的总键数. md5:894ec399ab2f88ea
# <翻译结束>


<原文开始>
// order (maximum number of children)
<原文结束>

# <翻译开始>
// 顺序（最大子节点数量）. md5:9c909788c23fe0a9
# <翻译结束>


<原文开始>
// BTreeNode is a single element within the tree.
<原文结束>

# <翻译开始>
// BTreeNode 是树中的一个单个元素。 md5:f6e73ea6e5510845
# <翻译结束>


<原文开始>
// BTreeEntry represents the key-value pair contained within nodes.
<原文结束>

# <翻译开始>
// BTreeEntry表示节点中包含的键值对。 md5:215d17b4d8c2f829
# <翻译结束>


<原文开始>
// NewBTree instantiates a B-tree with `m` (maximum number of children) and a custom key comparator.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
// Note that the `m` must be greater or equal than 3, or else it panics.
<原文结束>

# <翻译开始>
// NewBTree 创建一个具有 `m`（最大子节点数）和自定义键比较器的 B 树。参数 `safe` 用于指定是否在并发安全模式下使用树，其默认值为 false。
// 注意，`m` 必须大于或等于 3，否则将引发 panic。
// md5:63e15eb274ca4e1d
# <翻译结束>


<原文开始>
// NewBTreeFrom instantiates a B-tree with `m` (maximum number of children), a custom key comparator and data map.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewBTreeFrom 根据给定的参数实例化一个 B-树，包括孩子节点的最大数量 `m`、自定义键比较器和数据映射。
// 参数 `safe` 用于指定是否需要并发安全，默认情况下为 false。
// md5:7a8fbca9b49feb70
# <翻译结束>


<原文开始>
// Clone returns a new tree with a copy of current tree.
<原文结束>

# <翻译开始>
// Clone 返回一个新的树，其中包含当前树的副本。 md5:256477216ae712b7
# <翻译结束>


<原文开始>
// Set inserts key-value item into the tree.
<原文结束>

# <翻译开始>
// Set 将键值对插入到树中。 md5:af4d398e6bf21959
# <翻译结束>


<原文开始>
// doSet inserts key-value pair node into the tree.
// If key already exists, then its value is updated with the new value.
<原文结束>

# <翻译开始>
// doSet 将键值对节点插入到树中。
// 如果键已存在，则用新值更新其值。
// md5:dd34c6d624358b26
# <翻译结束>


<原文开始>
// Sets batch sets key-values to the tree.
<原文结束>

# <翻译开始>
// 设置批处理将键值对添加到树中。 md5:70c6ec85c8b7476c
# <翻译结束>


<原文开始>
// Get searches the node in the tree by `key` and returns its value or nil if key is not found in tree.
<原文结束>

# <翻译开始>
// Get 通过`key`在树中搜索节点，并返回其值，如果`key`在树中未找到，则返回nil。 md5:2e2483db20a69167
# <翻译结束>


<原文开始>
// doSetWithLockCheck checks whether value of the key exists with mutex.Lock,
// if not exists, set value to the map with given `key`,
// or else just return the existing value.
//
// When setting value, if `value` is type of <func() interface {}>,
// it will be executed with mutex.Lock of the hash map,
// and its return value will be set to the map with `key`.
//
// It returns value with given `key`.
<原文结束>

# <翻译开始>
// doSetWithLockCheck 使用互斥锁(mutex.Lock)检查键的值是否存在，
// 如果不存在，则将给定的`key`和`value`设置到映射中，
// 否则直接返回已存在的值。
//
// 在设置值时，如果`value`是<func() interface {}>类型，
// 它将在哈希映射的互斥锁(mutex.Lock)保护下执行，
// 并将其返回值设置到以`key`为键的映射中。
//
// 返回给定`key`对应的值。
// md5:1de9ffab89f3c38a
# <翻译结束>


<原文开始>
// GetOrSet returns the value by key,
// or sets value with given `value` if it does not exist and then returns this value.
<原文结束>

# <翻译开始>
// GetOrSet 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。
// md5:d8f89b6dec47292b
# <翻译结束>


<原文开始>
// GetOrSetFunc returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
<原文结束>

# <翻译开始>
// GetOrSetFunc 通过键获取值，
// 如果键不存在，则使用回调函数`f`的返回值设置值，
// 并返回这个设置的值。
// md5:f584dd7547dfbcc0
# <翻译结束>


<原文开始>
// GetOrSetFuncLock returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
//
// GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f`
// with mutex.Lock of the hash map.
<原文结束>

# <翻译开始>
// GetOrSetFuncLock 通过键获取值，
// 如果不存在，它将使用回调函数 `f` 的返回值设置该值，然后返回这个值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在执行函数 `f` 时会先锁定哈希映射的 mutex。
// md5:d32fdee586d84dde
# <翻译结束>


<原文开始>
// GetVar returns a gvar.Var with the value by given `key`.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVar 函数通过给定的 `key` 返回一个 gvar.Var，其值为对应的变量。
// 返回的 gvar.Var 不是并发安全的。
// md5:a04747902e4bf242
# <翻译结束>


<原文开始>
// GetVarOrSet returns a gvar.Var with result from GetVarOrSet.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSet 返回一个从 GetVarOrSet 获取的结果的 gvar.Var。返回的 gvar.Var 不是线程安全的。
// md5:089beb08264e18cf
# <翻译结束>


<原文开始>
// GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。
// md5:8c97b145faade5ae
# <翻译结束>


<原文开始>
// GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFuncLock 返回一个gvar.Var，其结果来自GetOrSetFuncLock。
// 返回的gvar.Var是非并发安全的。
// md5:90c22300c2187ce4
# <翻译结束>


<原文开始>
// SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
<原文结束>

# <翻译开始>
// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
# <翻译结束>


<原文开始>
// SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
<原文结束>

# <翻译开始>
// SetIfNotExistFunc 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
# <翻译结束>


<原文开始>
// SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
//
// SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that
// it executes function `f` with mutex.Lock of the hash map.
<原文结束>

# <翻译开始>
// SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。
// 如果 `key` 已存在，则返回 false，`value` 将被忽略。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
// md5:a6ee84b157328f61
# <翻译结束>


<原文开始>
// Contains checks whether `key` exists in the tree.
<原文结束>

# <翻译开始>
// Contains 检查键 `key` 是否存在于树中。 md5:77fd85af8e586867
# <翻译结束>


<原文开始>
// doRemove removes the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
<原文结束>

# <翻译开始>
// doRemove 通过键从树中删除节点。
// 键应符合比较器的类型断言，否则方法将 panic。
// md5:748baf2fba8b968d
# <翻译结束>


<原文开始>
// Remove removes the node from the tree by `key`.
<原文结束>

# <翻译开始>
// Remove 通过 `key` 从树中移除节点。 md5:42fcfa1d28b3945f
# <翻译结束>


<原文开始>
// Removes batch deletes values of the tree by `keys`.
<原文结束>

# <翻译开始>
// 通过`keys`移除树中的批量删除值。 md5:4620c81ac88b2936
# <翻译结束>


<原文开始>
// IsEmpty returns true if tree does not contain any nodes
<原文结束>

# <翻译开始>
// IsEmpty 返回true当树中不包含任何节点时. md5:d43f280c082bb0fd
# <翻译结束>


<原文开始>
// Size returns number of nodes in the tree.
<原文结束>

# <翻译开始>
// Size 返回树中的节点数量。 md5:d437d5852f80de5c
# <翻译结束>


<原文开始>
// Keys returns all keys in asc order.
<原文结束>

# <翻译开始>
// Keys 返回所有键，按升序排列。 md5:c2a692ea3491e160
# <翻译结束>


<原文开始>
// Values returns all values in asc order based on the key.
<原文结束>

# <翻译开始>
// Values返回根据键值升序排列的所有值。 md5:6268d60d7aa20c91
# <翻译结束>


<原文开始>
// Map returns all key-value items as map.
<原文结束>

# <翻译开始>
// Map 返回所有键值对项作为一个映射。 md5:c12ca822a6c71dc1
# <翻译结束>


<原文开始>
// MapStrAny returns all key-value items as map[string]interface{}.
<原文结束>

# <翻译开始>
// MapStrAny 将所有键值对作为 map[string]interface{} 返回。 md5:412456aafc43f7a8
# <翻译结束>


<原文开始>
// Clear removes all nodes from the tree.
<原文结束>

# <翻译开始>
// Clear 从树中移除所有节点。 md5:a7db742922264980
# <翻译结束>


<原文开始>
// Replace the data of the tree with given `data`.
<原文结束>

# <翻译开始>
// 使用给定的`data`替换树中的数据。 md5:ff636c579597f294
# <翻译结束>


<原文开始>
// Height returns the height of the tree.
<原文结束>

# <翻译开始>
// Height 返回树的高度。 md5:c3af563cbe50966a
# <翻译结束>


<原文开始>
// Left returns the left-most (min) entry or nil if tree is empty.
<原文结束>

# <翻译开始>
// Left 返回最左边（最小）的条目，如果树为空则返回 nil。 md5:57cf05edc8d10b88
# <翻译结束>


<原文开始>
// Right returns the right-most (max) entry or nil if tree is empty.
<原文结束>

# <翻译开始>
// Right 返回最右边（最大）的条目，如果树为空则返回 nil。 md5:cd331b29b9cc98f8
# <翻译结束>


<原文开始>
// String returns a string representation of container (for debugging purposes)
<原文结束>

# <翻译开始>
// String 返回一个表示容器的字符串（用于调试目的）. md5:2d28c3cbf692ce78
# <翻译结束>


<原文开始>
// Search searches the tree with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
<原文结束>

# <翻译开始>
// Search 函数使用给定的 `key` 在树中进行查找。
// 第二个返回参数 `found` 为 true 表示找到了键，否则为 false。
// md5:d151c3783cadda2c
# <翻译结束>


<原文开始>
// Search searches the tree with given `key` without mutex.
// It returns the entry if found or otherwise nil.
<原文结束>

# <翻译开始>
// Search 使用给定的 `key` 在不加锁的情况下搜索树。如果找到相应的条目，则返回该条目，否则返回 nil。
// md5:2f4ee3482351a19d
# <翻译结束>


<原文开始>
// Print prints the tree to stdout.
<原文结束>

# <翻译开始>
// Print 将树打印到标准输出。 md5:24fd6288549a501b
# <翻译结束>


<原文开始>
// Iterator is alias of IteratorAsc.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名。 md5:1bfdea306db62845
# <翻译结束>


<原文开始>
// IteratorFrom is alias of IteratorAscFrom.
<原文结束>

# <翻译开始>
// IteratorFrom是IteratorAscFrom的别名。 md5:6d3d506bcb5fe942
# <翻译结束>


<原文开始>
// IteratorAsc iterates the tree readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAsc 使用给定的回调函数 `f` 以升序遍历树（只读）。如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:c13b99ae40add3b0
# <翻译结束>


<原文开始>
// IteratorAscFrom iterates the tree readonly in ascending order with given callback function `f`.
// The parameter `key` specifies the start entry for iterating. The `match` specifies whether
// starting iterating if the `key` is fully matched, or else using index searching iterating.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAscFrom 从给定的回调函数 `f` 以升序遍历树。
// 参数 `key` 指定了遍历的起始条目。`match` 参数指定如果 `key` 完全匹配时是否开始遍历，否则使用索引搜索进行遍历。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:c04855bbd3989808
# <翻译结束>


<原文开始>
// Find current entry position in current node
<原文结束>

# <翻译开始>
// 在当前节点中找到当前条目的位置. md5:0a7b8dbdf0511756
# <翻译结束>


<原文开始>
// Try to go down to the child right of the current entry
<原文结束>

# <翻译开始>
// 尝试进入当前条目右侧的子级. md5:76c7333c8aa6548c
# <翻译结束>


<原文开始>
// Try to go down to the child left of the current node
<原文结束>

# <翻译开始>
// 尝试下降到当前节点的左子节点. md5:500eb3344ae2e9dc
# <翻译结束>


<原文开始>
// Return the left-most entry
<原文结束>

# <翻译开始>
// 返回最左边的条目. md5:e1e0f1ca5a78a394
# <翻译结束>


<原文开始>
// Above assures that we have reached a leaf node, so return the next entry in current node (if any)
<原文结束>

# <翻译开始>
// 上面的代码确保我们已经到达了一个叶节点，所以返回当前节点（如果有）的下一个条目. md5:16db742c6c56694a
# <翻译结束>


<原文开始>
// Reached leaf node and there are no entries to the right of the current entry, so go up to the parent
<原文结束>

# <翻译开始>
// 已到达叶子节点，并且当前条目右侧没有更多条目，因此返回到父节点. md5:2c6528856bee4df1
# <翻译结束>


<原文开始>
// Find next entry position in current node (note: search returns the first equal or bigger than entry)
<原文结束>

# <翻译开始>
// 在当前节点中查找下一个条目位置（注意：搜索返回第一个等于或大于条目的位置）. md5:f184010d524512f9
# <翻译结束>


<原文开始>
// Check that there is a next entry position in current node
<原文结束>

# <翻译开始>
// 检查当前节点是否有下一个条目位置. md5:c0e01af4b4d09d1d
# <翻译结束>


<原文开始>
// IteratorDesc iterates the tree readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 以降序方式遍历树，使用给定的回调函数 `f`。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:f6740ea55dafe4bb
# <翻译结束>


<原文开始>
// IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`.
// The parameter `key` specifies the start entry for iterating. The `match` specifies whether
// starting iterating if the `key` is fully matched, or else using index searching iterating.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDescFrom 以降序方式遍历树，使用给定的回调函数 `f`。参数 `key` 指定开始遍历的条目。`match` 表示是否在 `key` 完全匹配时开始遍历，否则使用索引搜索遍历。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:e6bb2f7d12ab34f6
# <翻译结束>


<原文开始>
// Try to go down to the child left of the current entry
<原文结束>

# <翻译开始>
// 尝试向下进入当前条目左侧的子项. md5:c5a2056515034dc8
# <翻译结束>


<原文开始>
// Try to go down to the child right of the current node
<原文结束>

# <翻译开始>
// 尝试前往当前节点右侧的子节点。 md5:dc7b90ce22f5e3db
# <翻译结束>


<原文开始>
// Return the right-most entry
<原文结束>

# <翻译开始>
// 返回最右边的条目. md5:d99f4b49cd9c2ea6
# <翻译结束>


<原文开始>
// Above assures that we have reached a leaf node, so return the previous entry in current node (if any)
<原文结束>

# <翻译开始>
// 以上确保我们已经到达叶子节点，因此返回当前节点（如果有）的前一个条目. md5:8c1c1d33dbf6920a
# <翻译结束>


<原文开始>
// Reached leaf node and there are no entries to the left of the current entry, so go up to the parent
<原文结束>

# <翻译开始>
// 到达叶子节点，且当前条目左侧没有更多项，因此向上移到父节点. md5:68f4a5ccc4125b55
# <翻译结束>


<原文开始>
// Find previous entry position in current node (note: search returns the first equal or bigger than entry)
<原文结束>

# <翻译开始>
// 在当前节点中查找前一个条目的位置（注意：搜索返回第一个等于或大于该条目的位置）. md5:cf405018e6f98ac2
# <翻译结束>


<原文开始>
// Check that there is a previous entry position in current node
<原文结束>

# <翻译开始>
// 检查当前节点中是否存在前一个条目的位置. md5:38ebf51611534fa0
# <翻译结束>


<原文开始>
// func (tree *BTree) isFull(node *BTreeNode) bool {
//	return len(node.Entries) == tree.maxEntries()
// }
<原文结束>

# <翻译开始>
// 函数（tree *BTree）isFull（node *BTreeNode）bool：
// 返回当前节点（node）的条目数等于BTree类型的maxEntries方法的返回值，即满载状态。
// md5:3a4924d98a84d807
# <翻译结束>


<原文开始>
// "-1" to favor right nodes to have more keys when splitting
<原文结束>

# <翻译开始>
// 当分割时，倾向于将更多的键分配给右侧节点，使用"-1". md5:589fa3b8c8c0ac7b
# <翻译结束>


<原文开始>
// search does search only within the single node among its entries
<原文结束>

# <翻译开始>
// search 仅在单个节点的条目中进行搜索. md5:708796b6f6c04ad5
# <翻译结束>


<原文开始>
// searchRecursively searches recursively down the tree starting at the startNode
<原文结束>

# <翻译开始>
// searchRecursively 从startNode开始递归地在树中搜索. md5:6c5effca3e12cf15
# <翻译结束>


<原文开始>
// Insert entry's key in the middle of the node
<原文结束>

# <翻译开始>
// 将条目的键插入节点的中间. md5:77832e5f62b079b2
# <翻译结束>


<原文开始>
// Move children from the node to be split into left and right nodes
<原文结束>

# <翻译开始>
// 将节点要被分割的子节点移动到左右子节点中. md5:3e37e30e3dd2cb2c
# <翻译结束>


<原文开始>
// Insert middle key into parent
<原文结束>

# <翻译开始>
// 将中间键插入到父节点中. md5:90770d4875d60061
# <翻译结束>


<原文开始>
// Set child left of inserted key in parent to the created left node
<原文结束>

# <翻译开始>
// 将插入键在父节点的左侧子节点设置为创建的左侧节点. md5:3ef858cf2ae0942a
# <翻译结束>


<原文开始>
// Set child right of inserted key in parent to the created right node
<原文结束>

# <翻译开始>
// 在父节点中将插入键的子节点设置为创建的右节点. md5:3ab277966ef065b7
# <翻译结束>


<原文开始>
// Root is a node with one entry and two children (left and right)
<原文结束>

# <翻译开始>
// Root 是一个具有一个入口和两个子节点（左和右）的节点. md5:5c7ab1f314ee5149
# <翻译结束>


<原文开始>
// leftSibling returns the node's left sibling and child index (in parent) if it exists, otherwise (nil,-1)
// key is any of keys in node (could even be deleted).
<原文结束>

# <翻译开始>
// leftSibling 函数返回节点的左兄弟节点以及该节点在父节点中的索引（如果存在的话），否则返回（nil, -1）。
// key 是节点中的任意一个键（即使是已被删除的键）。
// md5:5df6d39676db1b43
# <翻译结束>


<原文开始>
// rightSibling returns the node's right sibling and child index (in parent) if it exists, otherwise (nil,-1)
// key is any of keys in node (could even be deleted).
<原文结束>

# <翻译开始>
// rightSibling 返回节点的右兄弟节点及其在父节点中的子索引，如果存在的话，否则返回 (nil,-1)。
// key 可以是节点中的任意键（甚至可能是已删除的键）。
// md5:d987c8284e77dafa
# <翻译结束>


<原文开始>
// delete deletes an entry in node at entries' index
// ref.: https://en.wikipedia.org/wiki/B-tree#Deletion
<原文结束>

# <翻译开始>
// delete 删除node中entries索引处的条目
// 参考：https://en.wikipedia.org/wiki/B-tree#Deletion
// md5:b876a095ea679730
# <翻译结束>


<原文开始>
// deleting from a leaf node
<原文结束>

# <翻译开始>
// 从叶节点删除. md5:7876d56e8045e7f9
# <翻译结束>


<原文开始>
// deleting from an internal node
<原文结束>

# <翻译开始>
// 从内部节点删除. md5:4bd2fbac4d732f59
# <翻译结束>


<原文开始>
// largest node in the left sub-tree (assumed to exist)
<原文结束>

# <翻译开始>
// 左子树中最大的节点（假设存在）. md5:58dc1797a99c50fe
# <翻译结束>


<原文开始>
// reBalance reBalances the tree after deletion if necessary and returns true, otherwise false.
// Note that we first delete the entry and then call reBalance, thus the passed deleted key as reference.
<原文结束>

# <翻译开始>
// reBalance 在必要时重新平衡树并返回true，否则返回false。
// 注意，我们首先删除条目，然后调用reBalance，因此将传递已删除的键作为引用。
// md5:c3feadb6a7f38094
# <翻译结束>


<原文开始>
// check if re-balancing is needed
<原文结束>

# <翻译开始>
// 检查是否需要重新平衡. md5:1a872a2636208ac3
# <翻译结束>


<原文开始>
// try to borrow from left sibling
<原文结束>

# <翻译开始>
// 尝试从左侧兄弟节点借用. md5:93535f4b1bfcf27f
# <翻译结束>


<原文开始>
// prepend parent's separator entry to node's entries
<原文结束>

# <翻译开始>
// 将父节点的分隔符条目添加到节点的条目中. md5:aa1e7a85adad7bb6
# <翻译结束>


<原文开始>
// try to borrow from right sibling
<原文结束>

# <翻译开始>
// 尝试从右侧兄弟节点借入. md5:e0e1cfcfc7caad95
# <翻译结束>


<原文开始>
// append parent's separator entry to node's entries
<原文结束>

# <翻译开始>
// 将父节点的分隔符条目追加到节点的条目中. md5:00ee3de89c558897
# <翻译结束>


<原文开始>
// merge with right sibling
<原文结束>

# <翻译开始>
// 与右侧兄弟节点合并. md5:cf809f1e8e2d90dc
# <翻译结束>


<原文开始>
// merge with left sibling
<原文结束>

# <翻译开始>
// 与左兄弟节点合并. md5:411c166f3c82c9dc
# <翻译结束>


<原文开始>
// make the merged node the root if its parent was the root and the root is empty
<原文结束>

# <翻译开始>
// 如果合并节点的父节点是根节点并且根节点为空，将合并节点设为新的根. md5:304d193d00ef0afc
# <翻译结束>


<原文开始>
// parent might be underflow, so try to reBalance if necessary
<原文结束>

# <翻译开始>
// 父元素可能会下溢，因此如果必要的话尝试重新平衡. md5:157406767c643099
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
# <翻译结束>


<原文开始>
// getComparator returns the comparator if it's previously set,
// or else it panics.
<原文结束>

# <翻译开始>
// getComparator 如果之前已设置比较器，则返回该比较器，否则将引发恐慌。
// md5:03eac9fd6d838369
# <翻译结束>

