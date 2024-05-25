
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。
//
// 本源代码遵循MIT许可证条款。
// 如果gm文件未随附MIT许可证的副本，
// 您可以在https://github.com/gogf/gf获取一个。
// md5:c99fd05f11d37c36
# <翻译结束>


<原文开始>
// StrStrMap implements map[string]string with RWMutex that has switch.
<原文结束>

# <翻译开始>
// StrStrMap实现了具有开关的RWMutex映射[string]string。 md5:39383b0c43f8a057
# <翻译结束>


<原文开始>
// NewStrStrMap returns an empty StrStrMap object.
// The parameter `safe` is used to specify whether using map in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewStrStrMap 返回一个空的 StrStrMap 对象。
// 参数 `safe` 用于指定是否在并发安全模式下使用映射， 默认为 false。
// md5:305a371834d43bdd
# <翻译结束>


<原文开始>
// NewStrStrMapFrom creates and returns a hash map from given map `data`.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
<原文结束>

# <翻译开始>
// NewStrStrMapFrom 根据给定的映射 `data` 创建并返回一个哈希映射。
// 注意，参数 `data` 映射将被设置为底层数据映射（非深度复制），
// 因此，在外部修改该映射时可能会存在并发安全问题。
// md5:00f0f09e1bc308ad
# <翻译结束>


<原文开始>
// Iterator iterates the hash map readonly with custom callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// Iterator 使用自定义回调函数 `f` 读取只读哈希映射。如果 `f` 返回 true，则继续迭代；否则停止。
// md5:52d024b320a69c3b
# <翻译结束>


<原文开始>
// Clone returns a new hash map with copy of current map data.
<原文结束>

# <翻译开始>
// Clone 返回一个新的哈希映射，其中包含当前映射数据的副本。 md5:b9264f3636ead08a
# <翻译结束>


<原文开始>
// Map returns the underlying data map.
// Note that, if it's in concurrent-safe usage, it returns a copy of underlying data,
// or else a pointer to the underlying data.
<原文结束>

# <翻译开始>
// Map 返回底层数据映射。
// 注意，如果它在并发安全的使用场景中，它将返回底层数据的一个副本，
// 否则返回指向底层数据的指针。
// md5:7f8e0898ab3ddb0f
# <翻译结束>


<原文开始>
// MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.
<原文结束>

# <翻译开始>
// MapStrAny将映射的底层数据复制为map[string]interface{}。 md5:46db5a1110397522
# <翻译结束>


<原文开始>
// MapCopy returns a copy of the underlying data of the hash map.
<原文结束>

# <翻译开始>
// MapCopy 返回哈希映射底层数据的一个副本。 md5:46f762167d5821b1
# <翻译结束>


<原文开始>
// FilterEmpty deletes all key-value pair of which the value is empty.
// Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.
<原文结束>

# <翻译开始>
// FilterEmpty 删除所有值为空的键值对。空值包括：0、nil、false、""，以及切片、映射（map）或通道（channel）的长度为0的情况。
// md5:6cdcc470e2c0cab1
# <翻译结束>


<原文开始>
// Set sets key-value to the hash map.
<原文结束>

# <翻译开始>
// Set 将键值对设置到哈希映射中。 md5:07ea2dd1ea28820a
# <翻译结束>


<原文开始>
// Sets batch sets key-values to the hash map.
<原文结束>

# <翻译开始>
// 将键值对设置到哈希映射中。 md5:e3f3f8a1b69eb832
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
// Get 根据给定的 `key` 获取值。 md5:2b744a3e455aadfb
# <翻译结束>


<原文开始>
// Pop retrieves and deletes an item from the map.
<原文结束>

# <翻译开始>
// Pop 从映射中获取并删除一个元素。 md5:2d364ca2b6054111
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
// It returns value with given `key`.
<原文结束>

# <翻译开始>
// doSetWithLockCheck 使用互斥锁(mutex.Lock)检查键的值是否存在，
// 如果不存在，则将给定的`value`设置到映射中指定的`key`处，
// 否则，直接返回已存在的值。
//
// 它返回与给定`key`关联的值。
// md5:3a2d1537d3fe7230
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
// Removes batch deletes values of the map by keys.
<原文结束>

# <翻译开始>
// 通过键删除map中的批删除值。 md5:57081208d84ca7e8
# <翻译结束>


<原文开始>
// Remove deletes value from map by given `key`, and return this deleted value.
<原文结束>

# <翻译开始>
// Remove 通过给定的`key`从map中删除值，并返回被删除的值。 md5:5ee6dc9be17b4ab8
# <翻译结束>


<原文开始>
// Keys returns all keys of the map as a slice.
<原文结束>

# <翻译开始>
// Keys 返回映射中所有键的切片。 md5:425640fff4178659
# <翻译结束>


<原文开始>
// Values returns all values of the map as a slice.
<原文结束>

# <翻译开始>
// Values 将地图中的所有值返回为一个切片。 md5:a89b5b485c966abd
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
// Size返回映射的大小。 md5:da42fb3955847483
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
// Clear deletes all data of the map, it will remake a new underlying data map.
<原文结束>

# <翻译开始>
// Clear 删除映射中的所有数据，它将重新创建一个新的底层数据映射。 md5:0553a5cd54a22f3c
# <翻译结束>


<原文开始>
// Replace the data of the map with given `data`.
<原文结束>

# <翻译开始>
// 用给定的 `data` 替换映射的数据。 md5:a84ecf2839212d81
# <翻译结束>


<原文开始>
// LockFunc locks writing with given callback function `f` within RWMutex.Lock.
<原文结束>

# <翻译开始>
// LockFunc 使用给定的回调函数 `f` 在 RWMutex.Lock 中锁定写操作。 md5:e73dbc0381ebb3dc
# <翻译结束>


<原文开始>
// RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
<原文结束>

# <翻译开始>
// RLockFunc 在 RWMutex.RLock 的范围内使用给定的回调函数 `f` 进行读取锁定。 md5:4ae51d9b7445f043
# <翻译结束>


<原文开始>
// Flip exchanges key-value of the map to value-key.
<原文结束>

# <翻译开始>
// Flip 将映射的键值对交换为值键。 md5:dbcb578f1b30fa01
# <翻译结束>


<原文开始>
// Merge merges two hash maps.
// The `other` map will be merged into the map `m`.
<原文结束>

# <翻译开始>
// Merge 合并两个哈希映射。
// `other` 映射将被合并到映射 `m` 中。
// md5:a90c0d2b1f1fdaaa
# <翻译结束>


<原文开始>
// String returns the map as a string.
<原文结束>

# <翻译开始>
// String 将地图转换为字符串形式并返回。 md5:6473318e71d3dfd0
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for map.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于将任何类型的值设置到映射中。 md5:6f3087a6f7df5477
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
# <翻译结束>


<原文开始>
// IsSubOf checks whether the current map is a sub-map of `other`.
<原文结束>

# <翻译开始>
// IsSubOf 检查当前映射是否是`other`的子映射。 md5:9a6c60859c5a0fbc
# <翻译结束>


<原文开始>
// Diff compares current map `m` with map `other` and returns their different keys.
// The returned `addedKeys` are the keys that are in map `m` but not in map `other`.
// The returned `removedKeys` are the keys that are in map `other` but not in map `m`.
// The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).
<原文结束>

# <翻译开始>
// Diff 函数比较当前地图 `m` 与地图 `other` 并返回它们不同的键。
// 返回的 `addedKeys` 是存在于地图 `m` 中但不在地图 `other` 中的键。
// 返回的 `removedKeys` 是存在于地图 `other` 中但不在地图 `m` 中的键。
// 返回的 `updatedKeys` 是同时存在于地图 `m` 和 `other` 中，但其值不相等（`!=`）的键。
// md5:d3bf0bf8c70e9093
# <翻译结束>

