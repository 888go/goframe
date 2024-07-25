
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
// AVLTree holds elements of the AVL tree.
<原文结束>

# <翻译开始>
// AVLTree 存储AVL树的元素。 md5:d108e2d6ca60747c
# <翻译结束>


<原文开始>
// AVLTreeNode is a single element within the tree.
<原文结束>

# <翻译开始>
// AVLTreeNode 是树中的一个元素。 md5:ae1bf04ae171ca4e
# <翻译结束>


<原文开始>
// NewAVLTree instantiates an AVL tree with the custom key comparator.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewAVLTree 创建一个带有自定义键比较器的AVL树。
// 参数 `safe` 用于指定是否需要并发安全，默认为false。 md5:dfb13f71bc07620c
# <翻译结束>


<原文开始>
// NewAVLTreeFrom instantiates an AVL tree with the custom key comparator and data map.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewAVLTreeFrom 使用自定义的键比较器和数据映射初始化一个AVL树。
// 参数`safe`用于指定是否在并发安全模式下使用树，其默认值为false。 md5:856b75ecd9dc1540
# <翻译结束>


<原文开始>
// Clone returns a new tree with a copy of current tree.
<原文结束>

# <翻译开始>
// Clone 返回一个新的树，其中包含当前树的副本。 md5:256477216ae712b7
# <翻译结束>


<原文开始>
// Set inserts node into the tree.
<原文结束>

# <翻译开始>
// Set 将节点插入到树中。 md5:8af837873ed60d6a
# <翻译结束>


<原文开始>
// Sets batch sets key-values to the tree.
<原文结束>

# <翻译开始>
// 设置批处理将键值对添加到树中。 md5:70c6ec85c8b7476c
# <翻译结束>


<原文开始>
// Search searches the tree with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
<原文结束>

# <翻译开始>
// Search 函数使用给定的 `key` 在树中进行查找。
// 第二个返回参数 `found` 为 true 表示找到了键，否则为 false。 md5:d151c3783cadda2c
# <翻译结束>


<原文开始>
// doSearch searches the tree with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
<原文结束>

# <翻译开始>
// doSearch 使用给定的 `key` 在树中进行搜索。
// 第二个返回参数 `found` 如果找到了键，则为 true，否则为 false。 md5:46c1cf26ceea8b4b
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
// 返回给定`key`对应的值。 md5:1de9ffab89f3c38a
# <翻译结束>


<原文开始>
// GetOrSet returns the value by key,
// or sets value with given `value` if it does not exist and then returns this value.
<原文结束>

# <翻译开始>
// GetOrSet 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。 md5:d8f89b6dec47292b
# <翻译结束>


<原文开始>
// GetOrSetFunc returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
<原文结束>

# <翻译开始>
// GetOrSetFunc 通过键获取值，
// 如果键不存在，则使用回调函数`f`的返回值设置值，
// 并返回这个设置的值。 md5:f584dd7547dfbcc0
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
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在执行函数 `f` 时会先锁定哈希映射的 mutex。 md5:d32fdee586d84dde
# <翻译结束>


<原文开始>
// GetVar returns a gvar.Var with the value by given `key`.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVar 函数通过给定的 `key` 返回一个 gvar.Var，其值为对应的变量。
// 返回的 gvar.Var 不是并发安全的。 md5:a04747902e4bf242
# <翻译结束>


<原文开始>
// GetVarOrSet returns a gvar.Var with result from GetVarOrSet.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSet 返回一个从 GetVarOrSet 获取的结果的 gvar.Var。返回的 gvar.Var 不是线程安全的。 md5:089beb08264e18cf
# <翻译结束>


<原文开始>
// GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。 md5:8c97b145faade5ae
# <翻译结束>


<原文开始>
// GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock.
// The returned gvar.Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFuncLock 返回一个gvar.Var，其结果来自GetOrSetFuncLock。
// 返回的gvar.Var是非并发安全的。 md5:90c22300c2187ce4
# <翻译结束>


<原文开始>
// SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
<原文结束>

# <翻译开始>
// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。 md5:f80895920828f03e
# <翻译结束>


<原文开始>
// SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
<原文结束>

# <翻译开始>
// SetIfNotExistFunc 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。 md5:326c0b7c63d813e7
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
// 它在哈希映射的 mutex.Lock 保护下执行函数 `f`。 md5:a6ee84b157328f61
# <翻译结束>


<原文开始>
// Contains checks whether `key` exists in the tree.
<原文结束>

# <翻译开始>
// Contains 检查键 `key` 是否存在于树中。 md5:77fd85af8e586867
# <翻译结束>


<原文开始>
// Remove removes the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
<原文结束>

# <翻译开始>
// Remove 通过键从树中移除节点。
// 键应符合比较器的类型断言，否则方法将 panic。 md5:23794cd4708d8756
# <翻译结束>


<原文开始>
// Removes batch deletes values of the tree by `keys`.
<原文结束>

# <翻译开始>
// 通过`keys`移除树中的批量删除值。 md5:4620c81ac88b2936
# <翻译结束>


<原文开始>
// IsEmpty returns true if tree does not contain any nodes.
<原文结束>

# <翻译开始>
// IsEmpty 如果树中不包含任何节点，则返回true。 md5:8f7ae813360d880b
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
// Left returns the minimum element of the AVL tree
// or nil if the tree is empty.
<原文结束>

# <翻译开始>
// Left 返回 AVL 树中的最小元素
// 如果树为空，则返回 nil。 md5:d6b4c070feb60521
# <翻译结束>


<原文开始>
// Right returns the maximum element of the AVL tree
// or nil if the tree is empty.
<原文结束>

# <翻译开始>
// Right 返回AVL树中的最大元素，如果树为空则返回nil。 md5:7f0d34ae61ed561f
# <翻译结束>


<原文开始>
// Floor Finds floor node of the input key, return the floor node or nil if no floor node is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree is larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
<原文结束>

# <翻译开始>
// Floor 找到输入键的地板节点，如果没有找到地板节点，则返回nil。第二个返回参数表示是否找到了地板，如果找到为true，否则为false。
//
// 地板节点定义为大于或等于给定节点的最大节点。可能找不到地板节点，原因可能是树为空，或者树中的所有节点都大于给定节点。
//
// 键应遵循比较器的类型断言，否则方法会 panic。 md5:720f6000179912eb
# <翻译结束>


<原文开始>
// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling node is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree is smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
<原文结束>

# <翻译开始>
// Ceiling 找到输入键的天花板节点，如果没有找到天花板节点，则返回nil。第二个返回参数表示是否找到了天花板（true）或未找到（false）。
//
// 定义天花板节点为大于或等于给定节点的最小节点。可能找不到天花板节点，因为树为空，或者树中的所有节点都小于给定节点。
//
// 键应遵循比较器的类型断言，否则方法会 panic。 md5:6b92342a03f9f586
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
// String returns a string representation of container
<原文结束>

# <翻译开始>
// String 返回容器的字符串表示形式. md5:7daf925d59987319
# <翻译结束>


<原文开始>
// Print prints the tree to stdout.
<原文结束>

# <翻译开始>
// Print 将树打印到标准输出。 md5:24fd6288549a501b
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
// Flip exchanges key-value of the tree to value-key.
// Note that you should guarantee the value is the same type as key,
// or else the comparator would panic.
//
// If the type of value is different with key, you pass the new `comparator`.
<原文结束>

# <翻译开始>
// Flip 将树中的键值对交换为值键。
// 请注意，你应该确保值的类型与键相同，否则比较器会panic。
//
// 如果值的类型与键不同，你需要传递新的 `comparator`。 md5:e71ceac22aee55f1
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
// IteratorAsc 使用给定的回调函数 `f` 以升序遍历树（只读）。如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。 md5:c13b99ae40add3b0
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
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。 md5:c04855bbd3989808
# <翻译结束>


<原文开始>
// IteratorDesc iterates the tree readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 以降序方式遍历树，使用给定的回调函数 `f`。如果 `f` 返回 true，则继续遍历；否则停止。 md5:f6740ea55dafe4bb
# <翻译结束>


<原文开始>
// IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`.
// The parameter `key` specifies the start entry for iterating. The `match` specifies whether
// starting iterating if the `key` is fully matched, or else using index searching iterating.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDescFrom 以降序方式遍历树，使用给定的回调函数 `f`。参数 `key` 指定开始遍历的条目。`match` 表示是否在 `key` 完全匹配时开始遍历，否则使用索引搜索遍历。如果 `f` 返回 true，则继续遍历；否则停止。 md5:e6bb2f7d12ab34f6
# <翻译结束>


<原文开始>
// Prev returns the previous element in an inorder
// walk of the AVL tree.
<原文结束>

# <翻译开始>
// Prev 返回AVL树中序遍历的上一个元素。 md5:d859f5a91f8afa30
# <翻译结束>


<原文开始>
// Next returns the next element in an inorder
// walk of the AVL tree.
<原文结束>

# <翻译开始>
// Next 返回AVL树中序遍历的下一个元素。 md5:bf33a084df1455d4
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
// getComparator 如果之前已设置比较器，则返回该比较器，否则将引发恐慌。 md5:03eac9fd6d838369
# <翻译结束>

