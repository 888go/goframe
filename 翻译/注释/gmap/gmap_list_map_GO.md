
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。
# <翻译结束>


<原文开始>
// ListMap is a map that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
<原文结束>

# <翻译开始>
// ListMap 是一个保留插入顺序的映射。
//
// 它底层通过哈希表存储值，并通过双向链表来保存元素的顺序。
//
// 该结构不保证线程安全。
//
// 参考文献：http://en.wikipedia.org/wiki/关联切片
# <翻译结束>


<原文开始>
// NewListMap returns an empty link map.
// ListMap is backed by a hash table to store values and doubly-linked list to store ordering.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewListMap 返回一个空的链式映射。
// ListMap 由哈希表（用于存储值）和双向链表（用于存储顺序）作为底层支持。
// 参数 `safe` 用于指定是否在并发环境中安全地使用映射，默认情况下为 false。
// 以下是更详细的翻译：
// ```go
// NewListMap 函数用于创建并返回一个新的、空的链式映射结构体实例。
// 这个 ListMap 结构体内部结合了哈希表和双向链表两种数据结构：
// 哈希表用于高效地存储和查找键值对，而双向链表则用于记录键值对的插入顺序。
// 
// 参数 `safe` 表示是否需要保证该链式映射在并发环境中的安全性（即线程安全），
// 如果设置为 true，则在多 goroutine 并发访问时会进行相应的同步控制；
// 若不特别指定，其默认值为 false，此时不提供并发安全保证。
# <翻译结束>


<原文开始>
// NewListMapFrom returns a link map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
<原文结束>

# <翻译开始>
// NewListMapFrom 从给定的 `data` map 中创建一个链接映射。
// 注意，参数 `data` 中的映射将被设置为底层数据映射（非深度复制），
// 因此在外部修改该映射时可能会存在一些并发安全问题。
# <翻译结束>


<原文开始>
// Iterator is alias of IteratorAsc.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名。
# <翻译结束>


<原文开始>
// IteratorAsc iterates the map readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAsc 以升序遍历给定回调函数 `f` 的只读映射。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
# <翻译结束>


<原文开始>
// IteratorDesc iterates the map readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 以降序方式遍历给定的只读映射，并使用回调函数 `f` 进行处理。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止迭代。
# <翻译结束>


<原文开始>
// Clone returns a new link map with copy of current map data.
<原文结束>

# <翻译开始>
// Clone 返回一个新的链接映射，其中包含当前映射数据的副本。
# <翻译结束>


<原文开始>
// Clear deletes all data of the map, it will remake a new underlying data map.
<原文结束>

# <翻译开始>
// 清空删除映射中的所有数据，它会重新创建一个新的底层数据映射。
# <翻译结束>


<原文开始>
// Replace the data of the map with given `data`.
<原文结束>

# <翻译开始>
// 用给定的`data`替换map中的数据。
# <翻译结束>


<原文开始>
// Map returns a copy of the underlying data of the map.
<原文结束>

# <翻译开始>
// Map 返回映射底层数据的一个副本。
# <翻译结束>


<原文开始>
// MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.
<原文结束>

# <翻译开始>
// MapStrAny 返回该映射底层数据的一个副本，类型为 map[string]interface{}。
# <翻译结束>


<原文开始>
// FilterEmpty deletes all key-value pair of which the value is empty.
<原文结束>

# <翻译开始>
// FilterEmpty 删除所有值为空的键值对。
# <翻译结束>


<原文开始>
// Set sets key-value to the map.
<原文结束>

# <翻译开始>
// Set 将键值对设置到映射中。
# <翻译结束>


<原文开始>
// Sets batch sets key-values to the map.
<原文结束>

# <翻译开始>
// 设置批量数据：将键值对设置到映射（map）中。
# <翻译结束>


<原文开始>
// Search searches the map with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
<原文结束>

# <翻译开始>
// Search 通过给定的 `key` 在映射中搜索。
// 第二个返回参数 `found` 如果找到了 key，则为 true，否则为 false。
# <翻译结束>


<原文开始>
// Get returns the value by given `key`.
<原文结束>

# <翻译开始>
// Get 通过给定的 `key` 返回对应的值。
# <翻译结束>


<原文开始>
// Pop retrieves and deletes an item from the map.
<原文结束>

# <翻译开始>
// Pop 从映射中检索并删除一个项目。
# <翻译结束>


<原文开始>
// Pops retrieves and deletes `size` items from the map.
// It returns all items if size == -1.
<原文结束>

# <翻译开始>
// Pops 从映射中获取并删除 `size` 个元素。
// 当 size == -1 时，它返回所有元素。
# <翻译结束>


<原文开始>
// doSetWithLockCheck checks whether value of the key exists with mutex.Lock,
// if not exists, set value to the map with given `key`,
// or else just return the existing value.
//
// When setting value, if `value` is type of `func() interface {}`,
// it will be executed with mutex.Lock of the map,
// and its return value will be set to the map with `key`.
//
// It returns value with given `key`.
<原文结束>

# <翻译开始>
// doSetWithLockCheck 在对mutex.Lock进行检查后，确认键对应的值是否存在，
// 如果不存在，则使用给定的`key`将值设置到映射中；
// 否则仅返回已存在的值。
//
// 在设置值时，如果`value`的类型为`func() interface {}`，
// 将在映射的mutex.Lock保护下执行该函数，
// 并将其返回值以`key`为键设置到映射中。
//
// 最终返回带有给定`key`的值。
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
// with mutex.Lock of the map.
<原文结束>

# <翻译开始>
// GetOrSetFuncLock 通过键返回值，如果不存在该键，则使用回调函数`f`返回的值设置该值，
// 并随后返回这个新设置的值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，
// 它在对 map 进行 mutex.Lock 锁定后执行函数 `f`。
// 这段代码注释翻译成中文后的意思是：
// ```markdown
// GetOrSetFuncLock 方法通过给定的键获取值，
// 若键对应的值不存在，则会使用回调函数 `f` 返回的值进行设置，
// 并最终返回这个已设置的值。
//
// GetOrSetFuncLock 方法与 GetOrSetFunc 方法的区别在于，
// 在对映射（map）执行操作前，它会先调用 mutex.Lock 进行锁定。
# <翻译结束>


<原文开始>
// GetVar returns a Var with the value by given `key`.
// The returned Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVar 通过给定的 `key` 返回一个具有相应值的 Var。
// 返回的 Var 不是线程安全的。
# <翻译结束>


<原文开始>
// GetVarOrSet returns a Var with result from GetVarOrSet.
// The returned Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSet 返回一个从 GetVarOrSet 获取结果的 Var。
// 返回的 Var 对象不保证线程安全。
# <翻译结束>


<原文开始>
// GetVarOrSetFunc returns a Var with result from GetOrSetFunc.
// The returned Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFunc 返回一个 Var，其结果来自 GetOrSetFunc 的调用。
// 返回的 Var 不具备并发安全特性。
# <翻译结束>


<原文开始>
// GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock.
// The returned Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFuncLock 返回一个 Var，其结果来自 GetOrSetFuncLock。
// 返回的 Var 不是并发安全的。
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
// it executes function `f` with mutex.Lock of the map.
<原文结束>

# <翻译开始>
// SetIfNotExistFuncLock 函数设置键值对，其值为回调函数 `f` 的返回值，并在设置成功时返回 true。
// 若 `key` 已存在，则返回 false，同时将忽略 `value`。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在操作 map 时使用了 mutex.Lock 进行加锁，确保在执行函数 `f` 期间数据同步安全。
# <翻译结束>


<原文开始>
// Remove deletes value from map by given `key`, and return this deleted value.
<原文结束>

# <翻译开始>
// Remove通过给定的`key`从map中删除值，并返回这个被删除的值。
# <翻译结束>


<原文开始>
// Removes batch deletes values of the map by keys.
<原文结束>

# <翻译开始>
// 删除map中通过keys指定的所有值，进行批量删除。
# <翻译结束>


<原文开始>
// Keys returns all keys of the map as a slice in ascending order.
<原文结束>

# <翻译开始>
// Keys 返回映射的所有键，以升序排列的切片形式。
# <翻译结束>


<原文开始>
// Values returns all values of the map as a slice.
<原文结束>

# <翻译开始>
// Values 返回该映射的所有值作为一个切片。
# <翻译结束>


<原文开始>
// Contains checks whether a key exists.
// It returns true if the `key` exists, or else false.
<原文结束>

# <翻译开始>
// Contains 检查键是否存在。
// 如果 `key` 存在，则返回 true，否则返回 false。
# <翻译结束>


<原文开始>
// Size returns the size of the map.
<原文结束>

# <翻译开始>
// Size 返回映射的大小。
# <翻译结束>


<原文开始>
// IsEmpty checks whether the map is empty.
// It returns true if map is empty, or else false.
<原文结束>

# <翻译开始>
// IsEmpty 检查该映射是否为空。
// 如果映射为空，则返回 true，否则返回 false。
# <翻译结束>


<原文开始>
// Flip exchanges key-value of the map to value-key.
<原文结束>

# <翻译开始>
// Flip 将映射中的键值对进行交换，即把键变成值，值变成键。
# <翻译结束>


<原文开始>
// Merge merges two link maps.
// The `other` map will be merged into the map `m`.
<原文结束>

# <翻译开始>
// Merge 合并两个链表映射。
// `other` 映射将会被合并到映射 `m` 中。
# <翻译结束>


<原文开始>
// String returns the map as a string.
<原文结束>

# <翻译开始>
// String 将映射转换为字符串并返回。
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
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy 实现接口，用于当前类型的深度复制。
# <翻译结束>

