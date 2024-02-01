
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
// RedBlackTree holds elements of the red-black tree.
<原文结束>

# <翻译开始>
// RedBlackTree 用于存储红黑树的元素。
# <翻译结束>


<原文开始>
// RedBlackTreeNode is a single element within the tree.
<原文结束>

# <翻译开始>
// RedBlackTreeNode 是树中的单个元素。
# <翻译结束>


<原文开始>
// NewRedBlackTree instantiates a red-black tree with the custom key comparator.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewRedBlackTree 创建一个带有自定义键比较器的红黑树。
// 参数`safe`用于指定是否在并发安全的情况下使用该树，默认为false。
# <翻译结束>


<原文开始>
// NewRedBlackTreeFrom instantiates a red-black tree with the custom key comparator and `data` map.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewRedBlackTreeFrom 通过自定义键比较器和 `data` 映射创建一个红黑树实例。
// 参数 `safe` 用于指定是否在并发安全的情况下使用该树，默认为 false。
# <翻译结束>


<原文开始>
// SetComparator sets/changes the comparator for sorting.
<原文结束>

# <翻译开始>
// SetComparator 设置/更改用于排序的比较器。
# <翻译结束>


<原文开始>
// Resort the tree if comparator is changed.
<原文结束>

# <翻译开始>
// 如果比较器发生改变，则重新对树进行排序。
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
// Sets batch sets key-values to the tree.
<原文结束>

# <翻译开始>
// Sets批量设置键值对到树中。
# <翻译结束>


<原文开始>
// doSet inserts key-value item into the tree without mutex.
<原文结束>

# <翻译开始>
// doSet 在没有互斥锁的情况下将键值对项插入到树中。
# <翻译结束>


<原文开始>
// Assert key is of comparator's type for initial tree
<原文结束>

# <翻译开始>
// 确保键是初始树中比较器的类型
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
// doRemove removes the node from the tree by `key` without mutex.
<原文结束>

# <翻译开始>
// doRemove 在没有互斥锁的情况下，通过 `key` 从树中移除节点。
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
// IsEmpty returns true if tree does not contain any nodes.
<原文结束>

# <翻译开始>
// IsEmpty 返回 true 如果树中不包含任何节点。
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
// Left returns the left-most (min) node or nil if tree is empty.
<原文结束>

# <翻译开始>
// Left 返回树中最左边（最小值）的节点，如果树为空则返回nil。
# <翻译结束>


<原文开始>
// Right returns the right-most (max) node or nil if tree is empty.
<原文结束>

# <翻译开始>
// Right 返回最右侧（最大）的节点，如果树为空则返回nil。
# <翻译结束>


<原文开始>
// leftNode returns the left-most (min) node or nil if tree is empty.
<原文结束>

# <翻译开始>
// leftNode 返回最左边（最小值）的节点，如果树为空则返回 nil。
# <翻译结束>


<原文开始>
// rightNode returns the right-most (max) node or nil if tree is empty.
<原文结束>

# <翻译开始>
// rightNode 返回最右侧（最大值）的节点，如果树为空则返回 nil。
# <翻译结束>


<原文开始>
// Floor Finds floor node of the input key, return the floor node or nil if no floor node is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that its key is smaller than or equal to the given `key`.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree are larger than the given node.
<原文结束>

# <翻译开始>
// Floor 查找输入键的下界节点，返回下界节点或在未找到时返回 nil。
// 第二个返回参数为布尔值，若找到下界节点则为 true，否则为 false。
//
// 下界节点定义为键小于等于给定 `key` 的最大节点。
// 可能无法找到下界节点，原因可能是树为空，或者树中所有节点都大于给定节点。
# <翻译结束>


<原文开始>
// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling node is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that its key is larger than or equal to the given `key`.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree are smaller than the given node.
<原文结束>

# <翻译开始>
// Ceiling 函数查找大于或等于输入键的最小节点（即上限节点），并返回该上限节点；若未找到上限节点，则返回 nil。
// 第二个返回参数为布尔值，若找到上限节点则为 true，否则为 false。
//
// 上限节点定义为：其键大于或等于给定 `key` 的最小节点。
// 可能找不到上限节点，原因可能是树为空，或者树中所有节点都小于给定节点。
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
// String returns a string representation of container.
<原文结束>

# <翻译开始>
// String 返回 container 的字符串表示形式。
# <翻译结束>


<原文开始>
// Print prints the tree to stdout.
<原文结束>

# <翻译开始>
// Print 将树打印到标准输出（stdout）。
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
// Flip exchanges key-value of the tree to value-key.
// Note that you should guarantee the value is the same type as key,
// or else the comparator would panic.
//
// If the type of value is different with key, you pass the new `comparator`.
<原文结束>

# <翻译开始>
// Flip 将树中的键值对进行交换，即将键和值互换。
// 注意，你需要确保值的类型与键相同，
// 否则比较器将会触发 panic 异常。
//
// 如果值的类型与键不同，你需要传入新的 `comparator`（比较器）。
# <翻译结束>


<原文开始>
// doSearch searches the tree with given `key` without mutex.
// It returns the node if found or otherwise nil.
<原文结束>

# <翻译开始>
// doSearch 在没有互斥锁的情况下，使用给定的 `key` 搜索树。
// 如果找到，则返回节点，否则返回 nil。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for map.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于为 map 设置任意类型的值。
# <翻译结束>


<原文开始>
// getComparator returns the comparator if it's previously set,
// or else it panics.
<原文结束>

# <翻译开始>
// getComparator 返回之前设置的比较器，如果之前未设置，则会引发panic。
# <翻译结束>

