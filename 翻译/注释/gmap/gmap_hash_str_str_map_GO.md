
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证的条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一份。
//
# <翻译结束>


<原文开始>
// StrStrMap implements map[string]string with RWMutex that has switch.
<原文结束>

# <翻译开始>
// StrStrMap 实现了一个带有 RWMutex（读写互斥锁）和 switch 功能的 map[string]string 类型。
// （译注：该结构体或类型提供了一种线程安全的方式存储和操作键值对，其中键和值都是字符串类型。通过使用 RWMutex，可以在多线程环境下进行读写操作时保证数据一致性，switch 可能是指该实现中包含了一些用于控制并发访问模式的开关功能。）
# <翻译结束>


<原文开始>
// NewStrStrMap returns an empty StrStrMap object.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewStrStrMap 返回一个空的 StrStrMap 对象。
// 参数 `safe` 用于指定是否使用线程安全的 map， 默认为 false。
# <翻译结束>


<原文开始>
// NewStrStrMapFrom creates and returns a hash map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
<原文结束>

# <翻译开始>
// NewStrStrMapFrom 通过给定的 `data` 字典创建并返回一个哈希映射。
// 注意，参数 `data` 字典将被直接设置为底层数据字典（非深度复制），
// 因此在外部修改该字典时可能会存在并发安全问题。
# <翻译结束>


<原文开始>
// Iterator iterates the hash map readonly with custom callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// Iterator 使用自定义回调函数 `f` 以只读方式迭代哈希表。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止迭代。
# <翻译结束>


<原文开始>
// Clone returns a new hash map with copy of current map data.
<原文结束>

# <翻译开始>
// Clone 返回一个新的哈希映射，其中包含当前映射数据的副本。
# <翻译结束>


<原文开始>
// Map returns the underlying data map.
// Note that, if it's in concurrent-safe usage, it returns a copy of underlying data,
// or else a pointer to the underlying data.
<原文结束>

# <翻译开始>
// Map 返回底层数据映射。
// 注意，如果它在并发安全的使用场景下，将会返回底层数据的一个副本，
// 否则将返回指向底层数据的指针。
# <翻译结束>


<原文开始>
// MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.
<原文结束>

# <翻译开始>
// MapStrAny 返回该映射底层数据的一个副本，类型为 map[string]interface{}。
# <翻译结束>


<原文开始>
// MapCopy returns a copy of the underlying data of the hash map.
<原文结束>

# <翻译开始>
// MapCopy 返回哈希映射底层数据的一个副本。
# <翻译结束>


<原文开始>
// FilterEmpty deletes all key-value pair of which the value is empty.
// Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.
<原文结束>

# <翻译开始>
// FilterEmpty 删除所有值为空的键值对。
// 以下类型的值被视为空：0, nil, false, "", 切片/映射/通道长度为0。
# <翻译结束>


<原文开始>
// Set sets key-value to the hash map.
<原文结束>

# <翻译开始>
// Set 将键值对设置到哈希映射中。
# <翻译结束>


<原文开始>
// Sets batch sets key-values to the hash map.
<原文结束>

# <翻译开始>
// Sets批量设置键值对到哈希映射中。
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
// It returns value with given `key`.
<原文结束>

# <翻译开始>
// doSetWithLockCheck 在对 mutex.Lock 进行检查后，判断给定 key 的值是否存在，
// 若不存在，则使用给定的 `key` 将 value 设置到 map 中；
// 否则，直接返回已存在的 value。
//
// 它将返回具有给定 `key` 的 value。
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
// Removes batch deletes values of the map by keys.
<原文结束>

# <翻译开始>
// 删除map中通过keys指定的所有值，进行批量删除。
# <翻译结束>


<原文开始>
// Remove deletes value from map by given `key`, and return this deleted value.
<原文结束>

# <翻译开始>
// Remove通过给定的`key`从map中删除值，并返回这个被删除的值。
# <翻译结束>


<原文开始>
// Keys returns all keys of the map as a slice.
<原文结束>

# <翻译开始>
// Keys 返回该映射的所有键作为一个切片。
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
// LockFunc locks writing with given callback function `f` within RWMutex.Lock.
<原文结束>

# <翻译开始>
// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 内锁定写入操作。
# <翻译结束>


<原文开始>
// RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
<原文结束>

# <翻译开始>
// RLockFunc 在 RWMutex.RLock 内使用给定的回调函数 `f` 进行读取锁定。
# <翻译结束>


<原文开始>
// Flip exchanges key-value of the map to value-key.
<原文结束>

# <翻译开始>
// Flip 将映射中的键值对进行交换，即把键变成值，值变成键。
# <翻译结束>


<原文开始>
// Merge merges two hash maps.
// The `other` map will be merged into the map `m`.
<原文结束>

# <翻译开始>
// Merge 合并两个哈希映射。
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


<原文开始>
// IsSubOf checks whether the current map is a sub-map of `other`.
<原文结束>

# <翻译开始>
// IsSubOf 检查当前 map 是否为 `other` 的子集。
# <翻译结束>


<原文开始>
// Diff compares current map `m` with map `other` and returns their different keys.
// The returned `addedKeys` are the keys that are in map `m` but not in map `other`.
// The returned `removedKeys` are the keys that are in map `other` but not in map `m`.
// The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).
<原文结束>

# <翻译开始>
// Diff 函数用于比较当前映射 `m` 与映射 `other`，并返回它们不同的键。
// 返回的 `addedKeys` 是存在于映射 `m` 中但不在映射 `other` 中的键。
// 返回的 `removedKeys` 是存在于映射 `other` 中但不在映射 `m` 中的键。
// 返回的 `updatedKeys` 是同时存在于映射 `m` 和 `other` 中，但其对应值不相等（`!=`）的键。
# <翻译结束>

