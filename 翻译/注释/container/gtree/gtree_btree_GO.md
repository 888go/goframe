
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
// BTree holds elements of the B-tree.
<原文结束>

# <翻译开始>
// BTree 保存了 B-树 的元素。
# <翻译结束>


<原文开始>
// Total number of keys in the tree
<原文结束>

# <翻译开始>
// 树中键的总数
# <翻译结束>


<原文开始>
// order (maximum number of children)
<原文结束>

# <翻译开始>
// order （最大子节点数）
# <翻译结束>


<原文开始>
// BTreeNode is a single element within the tree.
<原文结束>

# <翻译开始>
// BTreeNode 是树中的单个元素。
# <翻译结束>







<原文开始>
// BTreeEntry represents the key-value pair contained within nodes.
<原文结束>

# <翻译开始>
// BTreeEntry 代表节点中包含的键值对。
# <翻译结束>


<原文开始>
// NewBTree instantiates a B-tree with `m` (maximum number of children) and a custom key comparator.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
// Note that the `m` must be greater or equal than 3, or else it panics.
<原文结束>

# <翻译开始>
// NewBTree 创建一个具有`m`（最大子节点数量）的B树，并使用自定义键比较器。
// 参数`safe`用于指定是否在并发安全环境下使用该树，默认为false。
// 注意，`m`必须大于或等于3，否则会引发panic。
# <翻译结束>


<原文开始>
// NewBTreeFrom instantiates a B-tree with `m` (maximum number of children), a custom key comparator and data map.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewBTreeFrom 通过给定的最大子节点数 `m`，自定义键比较器和数据映射来实例化一个 B-树。
// 参数 `safe` 用于指定是否在并发安全的情况下使用该树，默认为 false。
# <翻译结束>


<原文开始>
// Clone returns a new tree with a copy of current tree.
<原文结束>

# <翻译开始>
// Clone 返回一个新的树，其中包含当前树的副本。
# <翻译结束>


<原文开始>
// Set inserts key-value item into the tree.
<原文结束>

# <翻译开始>
// Set 将键值对项插入到树中。
# <翻译结束>


<原文开始>
// doSet inserts key-value pair node into the tree.
// If key already exists, then its value is updated with the new value.
<原文结束>

# <翻译开始>
// doSet 将键值对节点插入到树中。
// 如果键已存在，则用新值更新其原有值。
# <翻译结束>


<原文开始>
// Sets batch sets key-values to the tree.
<原文结束>

# <翻译开始>
// Sets批量设置键值对到树中。
# <翻译结束>


<原文开始>
// Get searches the node in the tree by `key` and returns its value or nil if key is not found in tree.
<原文结束>

# <翻译开始>
// Get通过`key`在树中搜索节点，并返回其对应的值，如果在树中未找到该键，则返回nil。
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
// doSetWithLockCheck 检查在对 mutex 锁定后，给定 key 对应的值是否存在，
// 如果不存在，则使用给定的 `key` 将 value 设置到映射中；
// 否则，直接返回已存在的 value。
//
// 在设置值的过程中，如果 `value` 的类型为 <func() interface {}>，
// 会在哈希映射的 mutex 锁定下执行该函数，
// 并将函数的返回值以 `key` 为键设置到映射中。
//
// 最终返回给定 `key` 对应的值。
# <翻译结束>


<原文开始>
// GetOrSet returns the value by key,
// or sets value with given `value` if it does not exist and then returns this value.
<原文结束>

# <翻译开始>
// GetOrSet 函数通过 key 返回对应的 value，
// 若该 key 不存在，则使用给定的 `value` 设置并返回这个设置后的值。
# <翻译结束>


<原文开始>
// GetOrSetFunc returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
<原文结束>

# <翻译开始>
// GetOrSetFunc 通过键返回值，如果该键不存在，
// 则使用回调函数 `f` 返回的值进行设置，并随后返回这个设置后的值。
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
// GetOrSetFuncLock 通过键返回值，如果该键不存在，则使用回调函数 `f` 返回的值设置并返回这个新值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
# <翻译结束>


<原文开始>
// GetVar returns a gvar.Var with the value by given `key`.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVar 通过给定的 `key` 返回一个包含其值的 gvar.Var。
// 返回的 gvar.Var 对象不支持并发安全。
# <翻译结束>


<原文开始>
// GetVarOrSet returns a gvar.Var with result from GetVarOrSet.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSet 返回一个从 GetVarOrSet 获取结果的 gvar.Var。
// 返回的 gvar.Var 不是线程安全的。
# <翻译结束>


<原文开始>
// GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。
# <翻译结束>


<原文开始>
// GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFuncLock 返回一个 gvar.Var，其结果来自 GetOrSetFuncLock。
// 返回的 gvar.Var 并未实现并发安全。
# <翻译结束>


<原文开始>
// SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
<原文结束>

# <翻译开始>
// SetIfNotExist 如果`key`不存在，则将`value`设置到map中，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
# <翻译结束>


<原文开始>
// SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
<原文结束>

# <翻译开始>
// SetIfNotExistFunc 使用回调函数`f`的返回值设置键值，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
# <翻译结束>


<原文开始>
// SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
//
// SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that
// it executes function `f` with mutex.Lock of the hash map.
<原文结束>

# <翻译开始>
// SetIfNotExistFuncLock 函数用于设置键值对，其值为回调函数 `f` 的返回值，并在设置成功时返回 true。
// 若 `key` 已存在，则返回 false，并且将忽略 `value` 参数。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在执行回调函数 `f` 时会锁定哈希表的 mutex 锁。
# <翻译结束>


<原文开始>
// Contains checks whether `key` exists in the tree.
<原文结束>

# <翻译开始>
// Contains 检查 `key` 是否存在于树中。
# <翻译结束>


<原文开始>
// doRemove removes the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
<原文结束>

# <翻译开始>
// doRemove 通过键从树中移除节点。
// 键应遵循比较器的类型断言，否则方法将引发恐慌。
# <翻译结束>


<原文开始>
// Remove removes the node from the tree by `key`.
<原文结束>

# <翻译开始>
// Remove 通过 `key` 从树中移除节点。
# <翻译结束>


<原文开始>
// Removes batch deletes values of the tree by `keys`.
<原文结束>

# <翻译开始>
// 删除树中通过`keys`指定的一批值。
# <翻译结束>


<原文开始>
// IsEmpty returns true if tree does not contain any nodes
<原文结束>

# <翻译开始>
// IsEmpty 返回 true 如果树中不包含任何节点
# <翻译结束>


<原文开始>
// Size returns number of nodes in the tree.
<原文结束>

# <翻译开始>
// Size 返回树中节点的数量。
# <翻译结束>


<原文开始>
// Keys returns all keys in asc order.
<原文结束>

# <翻译开始>
// Keys 返回所有按升序排列的键。
# <翻译结束>


<原文开始>
// Values returns all values in asc order based on the key.
<原文结束>

# <翻译开始>
// Values 返回所有基于键升序排列的值。
# <翻译结束>


<原文开始>
// Map returns all key-value items as map.
<原文结束>

# <翻译开始>
// Map 返回所有键值对项作为映射（map）。
# <翻译结束>


<原文开始>
// MapStrAny returns all key-value items as map[string]interface{}.
<原文结束>

# <翻译开始>
// MapStrAny 返回所有键值对项作为 map[string]interface{} 类型。
# <翻译结束>


<原文开始>
// Clear removes all nodes from the tree.
<原文结束>

# <翻译开始>
// Clear 清除树中的所有节点。
# <翻译结束>


<原文开始>
// Replace the data of the tree with given `data`.
<原文结束>

# <翻译开始>
// 用给定的`data`替换树的数据。
# <翻译结束>


<原文开始>
// Height returns the height of the tree.
<原文结束>

# <翻译开始>
// Height 返回树的高度。
# <翻译结束>


<原文开始>
// Left returns the left-most (min) entry or nil if tree is empty.
<原文结束>

# <翻译开始>
// Left 返回最左边（最小）的元素，如果树为空则返回 nil。
# <翻译结束>


<原文开始>
// Right returns the right-most (max) entry or nil if tree is empty.
<原文结束>

# <翻译开始>
// Right 返回最右侧（最大）的元素，如果树为空，则返回nil。
# <翻译结束>


<原文开始>
// String returns a string representation of container (for debugging purposes)
<原文结束>

# <翻译开始>
// String 返回容器的字符串表示形式（用于调试目的）
# <翻译结束>


<原文开始>
// Search searches the tree with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
<原文结束>

# <翻译开始>
// Search 使用给定的`key`搜索树。
// 第二个返回参数`found`如果找到key则为真（true），否则为假（false）。
# <翻译结束>


<原文开始>
// Search searches the tree with given `key` without mutex.
// It returns the entry if found or otherwise nil.
<原文结束>

# <翻译开始>
// Search 在没有使用互斥锁的情况下，通过给定的`key`搜索树。
// 如果找到对应的项则返回该条目，否则返回nil。
# <翻译结束>


<原文开始>
// Print prints the tree to stdout.
<原文结束>

# <翻译开始>
// Print 将树打印到标准输出（stdout）。
# <翻译结束>


<原文开始>
// Iterator is alias of IteratorAsc.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名。
# <翻译结束>


<原文开始>
// IteratorFrom is alias of IteratorAscFrom.
<原文结束>

# <翻译开始>
// IteratorFrom 是 IteratorAscFrom 的别名。
# <翻译结束>


<原文开始>
// IteratorAsc iterates the tree readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAsc 以升序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
# <翻译结束>


<原文开始>
// IteratorAscFrom iterates the tree readonly in ascending order with given callback function `f`.
// The parameter `key` specifies the start entry for iterating. The `match` specifies whether
// starting iterating if the `key` is fully matched, or else using index searching iterating.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAscFrom 以升序遍历（只读）给定回调函数 `f` 的树。
// 参数 `key` 指定了遍历的起始项。`match` 指定了当 `key` 完全匹配时是否开始遍历，
// 否则使用索引搜索方式进行遍历。
// 若 `f` 返回 true，则继续遍历；若返回 false，则停止遍历。
# <翻译结束>


<原文开始>
// Find current entry position in current node
<原文结束>

# <翻译开始>
// 在当前节点中查找当前条目的位置
# <翻译结束>


<原文开始>
// Try to go down to the child right of the current entry
<原文结束>

# <翻译开始>
// 尝试转到当前条目右侧的子节点
# <翻译结束>


<原文开始>
// Try to go down to the child left of the current node
<原文结束>

# <翻译开始>
// 尝试移动到当前节点的左孩子节点
# <翻译结束>







<原文开始>
// Above assures that we have reached a leaf node, so return the next entry in current node (if any)
<原文结束>

# <翻译开始>
// 上述代码确保我们已经到达一个叶节点，因此返回当前节点中的下一个条目（如果存在）
# <翻译结束>


<原文开始>
// Reached leaf node and there are no entries to the right of the current entry, so go up to the parent
<原文结束>

# <翻译开始>
// 已到达叶节点，并且当前条目右侧没有其他条目，因此向上返回到父节点
# <翻译结束>


<原文开始>
// Find next entry position in current node (note: search returns the first equal or bigger than entry)
<原文结束>

# <翻译开始>
// 在当前节点中查找下一个条目位置（注意：搜索返回第一个等于或大于给定条目的位置）
# <翻译结束>


<原文开始>
// Check that there is a next entry position in current node
<原文结束>

# <翻译开始>
// 检查当前节点中是否存在下一个条目位置
# <翻译结束>


<原文开始>
// IteratorDesc iterates the tree readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 以降序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
# <翻译结束>


<原文开始>
// IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`.
// The parameter `key` specifies the start entry for iterating. The `match` specifies whether
// starting iterating if the `key` is fully matched, or else using index searching iterating.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDescFrom 从指定的键(key)开始以降序方式遍历树（只读模式），并使用给定的回调函数`f`。
// 参数`key`指定了遍历的起始条目。`match`参数指定了如果`key`完全匹配时是否开始遍历，
// 否则将采用索引搜索方式进行遍历。
// 若`f`返回值为true，则继续遍历；若返回false，则停止遍历。
# <翻译结束>


<原文开始>
// Try to go down to the child left of the current entry
<原文结束>

# <翻译开始>
// 尝试转到当前条目左侧的子节点
# <翻译结束>


<原文开始>
// Try to go down to the child right of the current node
<原文结束>

# <翻译开始>
// 尝试移动到当前节点右侧的子节点
# <翻译结束>


<原文开始>
// Return the right-most entry
<原文结束>

# <翻译开始>
// 返回最右侧的条目
# <翻译结束>


<原文开始>
// Above assures that we have reached a leaf node, so return the previous entry in current node (if any)
<原文结束>

# <翻译开始>
// 上述代码确保我们已经到达一个叶节点，因此返回当前节点中（如果有的话）的前一个条目
# <翻译结束>


<原文开始>
// Reached leaf node and there are no entries to the left of the current entry, so go up to the parent
<原文结束>

# <翻译开始>
// 已到达叶节点，并且当前条目左侧没有条目，因此向上移动到父节点
# <翻译结束>


<原文开始>
// Find previous entry position in current node (note: search returns the first equal or bigger than entry)
<原文结束>

# <翻译开始>
// 在当前节点中查找前一个条目位置（注意：搜索返回第一个等于或大于给定条目的位置）
# <翻译结束>


<原文开始>
// Check that there is a previous entry position in current node
<原文结束>

# <翻译开始>
// 检查当前节点中是否存在前一个条目位置
# <翻译结束>


<原文开始>
// func (tree *BTree) isFull(node *BTreeNode) bool {
//	return len(node.Entries) == tree.maxEntries()
// }
<原文结束>

# <翻译开始>
// 函数 (tree *BTree) isFull(node *BTreeNode) bool 的作用是：
// 判断给定的 BTreeNode 节点（node）是否已满。
// 当节点中的 Entries 数组长度等于 BTree 的最大条目数（通过 tree.maxEntries() 方法获取）时，返回 true，表示节点已满；否则返回 false。
# <翻译结束>


<原文开始>
// "-1" to favor right nodes to have more keys when splitting
<原文结束>

# <翻译开始>
// 当分裂节点时，使用“-1”倾向于使右侧节点拥有更多的键
# <翻译结束>


<原文开始>
// search does search only within the single node among its entries
<原文结束>

# <翻译开始>
// search 在单个节点的条目中执行搜索
# <翻译结束>


<原文开始>
// searchRecursively searches recursively down the tree starting at the startNode
<原文结束>

# <翻译开始>
// searchRecursively 从起始节点startNode开始，递归地向下搜索整个树结构
# <翻译结束>


<原文开始>
// Insert entry's key in the middle of the node
<原文结束>

# <翻译开始>
// 在节点中间插入entry的键
# <翻译结束>


<原文开始>
// Move children from the node to be split into left and right nodes
<原文结束>

# <翻译开始>
// 将待拆分节点的子节点移动到左右两个新节点中
# <翻译结束>


<原文开始>
// Insert middle key into parent
<原文结束>

# <翻译开始>
// 将中间键插入到父节点中
# <翻译结束>


<原文开始>
// Set child left of inserted key in parent to the created left node
<原文结束>

# <翻译开始>
// 将父节点中插入键左侧的子节点设置为新创建的左节点
# <翻译结束>


<原文开始>
// Set child right of inserted key in parent to the created right node
<原文结束>

# <翻译开始>
// 在父节点中，将插入键的右子节点设置为新创建的右节点
# <翻译结束>


<原文开始>
// Root is a node with one entry and two children (left and right)
<原文结束>

# <翻译开始>
// Root 是一个节点，包含一个键值对，并且拥有两个子节点（左、右子节点）
# <翻译结束>


<原文开始>
// leftSibling returns the node's left sibling and child index (in parent) if it exists, otherwise (nil,-1)
// key is any of keys in node (could even be deleted).
<原文结束>

# <翻译开始>
// leftSibling 返回给定节点的左兄弟节点及其在父节点中的子索引（如果存在的话），否则返回 (nil, -1)
// key 是节点中任意一个键（甚至可能是已删除的键）。
// 这段 Go 语言代码的注释翻译成中文为：
// ```go
// leftSibling 函数会返回该节点的左侧兄弟节点及其在父节点中的子节点位置索引，如果存在这样的兄弟节点，则返回相应的信息；否则返回 (nil, -1)。
// 参数 key 是该节点中的任意一个键值（甚至可能是一个已被删除的键值）。
# <翻译结束>


<原文开始>
// rightSibling returns the node's right sibling and child index (in parent) if it exists, otherwise (nil,-1)
// key is any of keys in node (could even be deleted).
<原文结束>

# <翻译开始>
// rightSibling 返回给定节点的右兄弟节点及其在父节点中的子节点索引（如果存在的话），否则返回 (nil, -1)
// key 是节点中任意一个键（甚至可能是已被删除的键）。
# <翻译结束>


<原文开始>
// delete deletes an entry in node at entries' index
// ref.: https://en.wikipedia.org/wiki/B-tree#Deletion
<原文结束>

# <翻译开始>
// delete 删除在节点中entries指定索引处的条目
// 参考文献: https://en.wikipedia.org/wiki/B-tree#删除
# <翻译结束>







<原文开始>
// deleting from an internal node
<原文结束>

# <翻译开始>
// 从内部节点进行删除
# <翻译结束>


<原文开始>
// reBalance reBalances the tree after deletion if necessary and returns true, otherwise false.
// Note that we first delete the entry and then call reBalance, thus the passed deleted key as reference.
<原文结束>

# <翻译开始>
// reBalance 在必要时在删除操作后重新平衡树并返回 true，否则返回 false。
// 注意，我们首先删除条目，然后调用 reBalance，因此将已删除键作为参考传入。
# <翻译结束>


<原文开始>
// check if re-balancing is needed
<原文结束>

# <翻译开始>
// 检查是否需要进行平衡调整
# <翻译结束>


<原文开始>
// try to borrow from left sibling
<原文结束>

# <翻译开始>
// 尝试从左兄弟节点借用
# <翻译结束>


<原文开始>
// try to borrow from right sibling
<原文结束>

# <翻译开始>
// 尝试从右兄弟节点借用
# <翻译结束>

















<原文开始>
// make the merged node the root if its parent was the root and the root is empty
<原文结束>

# <翻译开始>
// 如果合并节点的父节点是根节点，并且根节点为空，则将合并节点设置为根节点
# <翻译结束>


<原文开始>
// parent might be underflow, so try to reBalance if necessary
<原文结束>

# <翻译开始>
// 父节点可能发生了下溢（即子节点数量不平衡），因此如果有必要，尝试进行重新平衡
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
# <翻译结束>


<原文开始>
// getComparator returns the comparator if it's previously set,
// or else it panics.
<原文结束>

# <翻译开始>
// getComparator 返回之前设置的比较器，如果之前未设置，则会引发panic。
# <翻译结束>


<原文开始>
// largest node in the left sub-tree (assumed to exist)
<原文结束>

# <翻译开始>
// 左子树中最大的节点（假设存在）
# <翻译结束>


<原文开始>
// prepend parent's separator entry to node's entries
<原文结束>

# <翻译开始>
// 在节点的条目前添加父级的分隔符条目
# <翻译结束>


<原文开始>
// append parent's separator entry to node's entries
<原文结束>

# <翻译开始>
// 将父节点的分隔符条目添加到当前节点的条目中
# <翻译结束>


<原文开始>
// Contained keys in node
<原文结束>

# <翻译开始>
// 节点中包含的键
# <翻译结束>


<原文开始>
// Return the left-most entry
<原文结束>

# <翻译开始>
// 返回最左边的条目
# <翻译结束>


<原文开始>
// deleting from a leaf node
<原文结束>

# <翻译开始>
// 从叶子节点进行删除
# <翻译结束>


<原文开始>
// merge with siblings
<原文结束>

# <翻译开始>
// 与兄弟节点合并
# <翻译结束>


<原文开始>
// merge with right sibling
<原文结束>

# <翻译开始>
// 与右侧兄弟节点合并
# <翻译结束>


<原文开始>
// merge with left sibling
<原文结束>

# <翻译开始>
// 与左兄弟节点合并
# <翻译结束>

