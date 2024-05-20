
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b
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
// ListMap 是一个保持插入顺序的映射。
//
// 它使用哈希表存储值，使用双向链表存储顺序。
//
// 该结构不支持多线程安全。
//
// 参考：http://en.wikipedia.org/wiki/Associative_array
// md5:a164418fa5f6798e
# <翻译结束>


<原文开始>
// NewListMap returns an empty link map.
// ListMap is backed by a hash table to store values and doubly-linked list to store ordering.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewListMap 返回一个空的链接映射。
// ListMap 使用哈希表存储值，使用双向链表存储顺序。
// 参数 `safe` 用于指定是否在并发安全模式下使用映射，默认为 false。
// md5:515d74cfd6f50ee5
# <翻译结束>


<原文开始>
// NewListMapFrom returns a link map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
<原文结束>

# <翻译开始>
// NewListMapFrom 从给定的映射（map）`data`创建一个链接映射。
// 注意，参数`data`映射将被设置为底层数据映射（不进行深拷贝），如果在外部修改该映射可能会导致并发安全问题。
// md5:d15c506b7dc77488
# <翻译结束>


<原文开始>
// Iterator is alias of IteratorAsc.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名。. md5:1bfdea306db62845
# <翻译结束>


<原文开始>
// IteratorAsc iterates the map readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAsc 使用给定的回调函数 `f` 以升序遍历地图，并且是只读遍历。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:da20ab57c78da7c6
# <翻译结束>


<原文开始>
// IteratorDesc iterates the map readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 使用给定的回调函数 `f` 以降序遍历只读映射。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止。
// md5:1745f0b396846901
# <翻译结束>


<原文开始>
// Clone returns a new link map with copy of current map data.
<原文结束>

# <翻译开始>
// Clone 返回一个新的链接映射，包含当前映射数据的副本。. md5:c24afe920899b3a6
# <翻译结束>


<原文开始>
// Clear deletes all data of the map, it will remake a new underlying data map.
<原文结束>

# <翻译开始>
// Clear 删除映射中的所有数据，它将重新创建一个新的底层数据映射。. md5:0553a5cd54a22f3c
# <翻译结束>


<原文开始>
// Replace the data of the map with given `data`.
<原文结束>

# <翻译开始>
// 用给定的 `data` 替换映射的数据。. md5:a84ecf2839212d81
# <翻译结束>


<原文开始>
// Map returns a copy of the underlying data of the map.
<原文结束>

# <翻译开始>
// Map返回map底层数据的副本。. md5:1be60fb31c8cf7e9
# <翻译结束>


<原文开始>
// MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.
<原文结束>

# <翻译开始>
// MapStrAny将映射的底层数据复制为map[string]interface{}。. md5:46db5a1110397522
# <翻译结束>


<原文开始>
// FilterEmpty deletes all key-value pair of which the value is empty.
<原文结束>

# <翻译开始>
// FilterEmpty 删除所有值为空的键值对。. md5:77ba324f6e82e0c4
# <翻译结束>


<原文开始>
// Set sets key-value to the map.
<原文结束>

# <翻译开始>
// Set 将键值对设置到映射中。. md5:78bf4bf763bbf6e3
# <翻译结束>


<原文开始>
// Sets batch sets key-values to the map.
<原文结束>

# <翻译开始>
// 将键值对设置到映射中。. md5:863aacdcc54fd6e1
# <翻译结束>


<原文开始>
// Search searches the map with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
<原文结束>

# <翻译开始>
// Search 在给定的`key`下搜索映射。
// 第二个返回参数`found`如果找到键，则为true，否则为false。
// md5:99336de9941a3b02
# <翻译结束>


<原文开始>
// Get returns the value by given `key`.
<原文结束>

# <翻译开始>
// Get 根据给定的 `key` 获取值。. md5:2b744a3e455aadfb
# <翻译结束>


<原文开始>
// Pop retrieves and deletes an item from the map.
<原文结束>

# <翻译开始>
// Pop 从映射中获取并删除一个元素。. md5:2d364ca2b6054111
# <翻译结束>


<原文开始>
// Pops retrieves and deletes `size` items from the map.
// It returns all items if size == -1.
<原文结束>

# <翻译开始>
// Pops 从映射中检索并删除 `size` 个项目。
// 如果 size 等于 -1，则返回所有项目。
// md5:0f2cdbc0238fdc37
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
// doSetWithLockCheck 使用 mutex.Lock 检查给定键的值是否存在。
// 如果不存在，使用给定的 `key` 将值设置到映射中；否则，直接返回现有的值。
// 
// 当设置值时，如果 `value` 是类型为 `func() interface{}`，它将在映射的 mutex.Lock 保护下执行，
// 并将返回值设置为映射中的 `key`。
// 
// 它返回给定 `key` 的值。
// md5:b667e8828a47a6d9
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
// with mutex.Lock of the map.
<原文结束>

# <翻译开始>
// GetOrSetFuncLock 通过键获取值，
// 如果该值不存在，则使用回调函数 `f` 的返回值进行设置，
// 然后返回这个值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它在执行映射的 mutex.Lock 保护下执行函数 `f`。
// md5:f5e408a3393171bc
# <翻译结束>


<原文开始>
// GetVar returns a Var with the value by given `key`.
// The returned Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVar通过给定的`key`返回一个Var。返回的Var是非并发安全的。
// md5:debfb1b2bd13312b
# <翻译结束>


<原文开始>
// GetVarOrSet returns a Var with result from GetVarOrSet.
// The returned Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSet 返回一个 Var，其中包含从 GetVarOrSet 获取的结果。
// 返回的 Var 是非并发安全的。
// md5:c3730f368b7f62b5
# <翻译结束>


<原文开始>
// GetVarOrSetFunc returns a Var with result from GetOrSetFunc.
// The returned Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFunc 返回一个Var，其结果来自GetOrSetFunc。
// 返回的Var不具备并发安全性。
// md5:7d7674129b73ead1
# <翻译结束>


<原文开始>
// GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock.
// The returned Var is un-concurrent safe.
<原文结束>

# <翻译开始>
// GetVarOrSetFuncLock 返回一个从 GetOrSetFuncLock 获得结果的 Var。返回的 Var 不是线程安全的。
// md5:bdab644d14c89234
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
// it executes function `f` with mutex.Lock of the map.
<原文结束>

# <翻译开始>
// SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。
// 如果 `key` 已存在，它将返回 false，`value` 将被忽略。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在执行函数 `f` 时会获取映射的 mutex.Lock。
// md5:12e78d7edb4c4c12
# <翻译结束>


<原文开始>
// Remove deletes value from map by given `key`, and return this deleted value.
<原文结束>

# <翻译开始>
// Remove 通过给定的`key`从map中删除值，并返回被删除的值。. md5:5ee6dc9be17b4ab8
# <翻译结束>


<原文开始>
// Removes batch deletes values of the map by keys.
<原文结束>

# <翻译开始>
// 通过键删除map中的批删除值。. md5:57081208d84ca7e8
# <翻译结束>


<原文开始>
// Keys returns all keys of the map as a slice in ascending order.
<原文结束>

# <翻译开始>
// Keys返回映射中所有键作为升序排列的切片。. md5:140d43c5cccae9d9
# <翻译结束>


<原文开始>
// Values returns all values of the map as a slice.
<原文结束>

# <翻译开始>
// Values 将地图中的所有值返回为一个切片。. md5:a89b5b485c966abd
# <翻译结束>


<原文开始>
// Contains checks whether a key exists.
// It returns true if the `key` exists, or else false.
<原文结束>

# <翻译开始>
// Contains 检查键是否存在。
// 如果键存在，它返回 true，否则返回 false。
// md5:d8fb22313aadd65f
# <翻译结束>


<原文开始>
// Size returns the size of the map.
<原文结束>

# <翻译开始>
// Size返回映射的大小。. md5:da42fb3955847483
# <翻译结束>


<原文开始>
// IsEmpty checks whether the map is empty.
// It returns true if map is empty, or else false.
<原文结束>

# <翻译开始>
// IsEmpty 检查映射是否为空。
// 如果映射为空，则返回true，否则返回false。
// md5:ad4bd5c796f79266
# <翻译结束>


<原文开始>
// Flip exchanges key-value of the map to value-key.
<原文结束>

# <翻译开始>
// Flip 将映射的键值对交换为值键。. md5:dbcb578f1b30fa01
# <翻译结束>


<原文开始>
// Merge merges two link maps.
// The `other` map will be merged into the map `m`.
<原文结束>

# <翻译开始>
// Merge 合并两个链接映射。
// 将将`other`映射合并到`m`映射中。
// md5:2ec13ae7c16e16f8
# <翻译结束>


<原文开始>
// String returns the map as a string.
<原文结束>

# <翻译开始>
// String 将地图转换为字符串形式并返回。. md5:6473318e71d3dfd0
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for map.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于将任何类型的值设置到映射中。. md5:6f3087a6f7df5477
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
# <翻译结束>

