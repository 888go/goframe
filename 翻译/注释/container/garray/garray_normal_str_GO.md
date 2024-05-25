
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
// StrArray is a golang string array with rich features.
// It contains a concurrent-safe/unsafe switch, which should be set
// when its initialization and cannot be changed then.
<原文结束>

# <翻译开始>
// StrArray 是一个具有丰富功能的 Go 语言字符串数组。
// 它包含一个并发安全/不安全的开关，该开关应在初始化时设置，并且之后不能更改。
// md5:60bf9d0fe402df8a
# <翻译结束>


<原文开始>
// NewStrArray creates and returns an empty array.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewStrArray 创建并返回一个空数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:1a16d6b7fa6dc90d
# <翻译结束>


<原文开始>
// NewStrArraySize create and returns an array with given size and cap.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewStrArraySize 创建并返回一个给定大小和容量的数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:d419c5b3ffb2a682
# <翻译结束>


<原文开始>
// NewStrArrayFrom creates and returns an array with given slice `array`.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewStrArrayFrom 根据给定的切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否使用并发安全的数组，默认为 false。
// md5:719d22a529b420db
# <翻译结束>


<原文开始>
// NewStrArrayFromCopy creates and returns an array from a copy of given slice `array`.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewStrArrayFromCopy 根据给定切片 `array` 的副本创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全环境下使用数组，默认为 false。
// md5:71bd55b1c0df65be
# <翻译结束>


<原文开始>
// At returns the value by the specified index.
// If the given `index` is out of range of the array, it returns an empty string.
<原文结束>

# <翻译开始>
// At通过指定的索引返回值。
// 如果给定的`index`超出了数组的范围，它将返回一个空字符串。
// md5:2465f6b1e3ac2863
# <翻译结束>


<原文开始>
// Get returns the value by the specified index.
// If the given `index` is out of range of the array, the `found` is false.
<原文结束>

# <翻译开始>
// Get 函数通过指定的索引返回值。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:ab300cfc0d6dd8ee
# <翻译结束>


<原文开始>
// Set sets value to specified index.
<原文结束>

# <翻译开始>
// Set 设置指定索引的值。 md5:7c1d7ea9df0b722c
# <翻译结束>


<原文开始>
// SetArray sets the underlying slice array with the given `array`.
<原文结束>

# <翻译开始>
// SetArray 使用给定的 `array` 设置底层切片数组。 md5:160b43a5c0ec752c
# <翻译结束>


<原文开始>
// Replace replaces the array items by given `array` from the beginning of array.
<原文结束>

# <翻译开始>
// Replace 从数组的起始位置开始，使用给定的 `array` 替换数组中的元素。 md5:5acead2fd9ec0761
# <翻译结束>


<原文开始>
// Sum returns the sum of values in an array.
<原文结束>

# <翻译开始>
// Sum 返回数组中所有值的和。 md5:b2148175a749b162
# <翻译结束>


<原文开始>
// Sort sorts the array in increasing order.
// The parameter `reverse` controls whether sort
// in increasing order(default) or decreasing order
<原文结束>

# <翻译开始>
// Sort 对数组进行升序排序。
// 参数 `reverse` 控制排序方式，如果为 true，则降序排列（默认为升序）。
// md5:35d4650a0f563ccf
# <翻译结束>


<原文开始>
// SortFunc sorts the array by custom function `less`.
<原文结束>

# <翻译开始>
// SortFunc 使用自定义函数 `less` 对数组进行排序。 md5:8da07d09bbd08513
# <翻译结束>


<原文开始>
// InsertBefore inserts the `values` to the front of `index`.
<原文结束>

# <翻译开始>
// InsertBefore 将`values`插入到`index`的前面。 md5:f5f3b46cd17ba885
# <翻译结束>


<原文开始>
// InsertAfter inserts the `values` to the back of `index`.
<原文结束>

# <翻译开始>
// InsertAfter 将 `values` 插入到 `index` 后面。 md5:b90b80fa75b6b6e0
# <翻译结束>


<原文开始>
// Remove removes an item by index.
// If the given `index` is out of range of the array, the `found` is false.
<原文结束>

# <翻译开始>
// Remove 函数通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:feaf958654838c25
# <翻译结束>


<原文开始>
// doRemoveWithoutLock removes an item by index without lock.
<原文结束>

# <翻译开始>
// doRemoveWithoutLock 不使用锁移除一个项目。 md5:a6a1746903fd131c
# <翻译结束>


<原文开始>
// Determine array boundaries when deleting to improve deletion efficiency.
<原文结束>

# <翻译开始>
// 在删除时确定数组边界，以提高删除效率。 md5:bc969ee880edf699
# <翻译结束>


<原文开始>
	// If it is a non-boundary delete,
	// it will involve the creation of an array,
	// then the deletion is less efficient.
<原文结束>

# <翻译开始>
// 如果是一个非边界删除，
// 它将涉及创建一个数组，
// 那么删除操作效率较低。
// md5:6a664196d66bc968
# <翻译结束>


<原文开始>
// RemoveValue removes an item by value.
// It returns true if value is found in the array, or else false if not found.
<原文结束>

# <翻译开始>
// RemoveValue 函数根据值删除一个元素。
// 如果值在数组中找到，它将返回 true，否则如果未找到则返回 false。
// md5:c49c7706ce703d00
# <翻译结束>


<原文开始>
// RemoveValues removes multiple items by `values`.
<原文结束>

# <翻译开始>
// RemoveValues 根据`values`移除多个项目。 md5:fbdf68fa6a8cdd26
# <翻译结束>


<原文开始>
// PushLeft pushes one or multiple items to the beginning of array.
<原文结束>

# <翻译开始>
// PushLeft 将一个或多个项目推送到数组的开头。 md5:9062afab48970bed
# <翻译结束>


<原文开始>
// PushRight pushes one or multiple items to the end of array.
// It equals to Append.
<原文结束>

# <翻译开始>
// PushRight 将一个或多个元素添加到数组的末尾。
// 它等同于 Append。
// md5:bb33f2edfdfd9896
# <翻译结束>


<原文开始>
// PopLeft pops and returns an item from the beginning of array.
// Note that if the array is empty, the `found` is false.
<原文结束>

# <翻译开始>
// PopLeft 从数组的开头弹出并返回一个项目。
// 注意，如果数组为空，`found` 为 false。
// md5:68f14002d84594a4
# <翻译结束>


<原文开始>
// PopRight pops and returns an item from the end of array.
// Note that if the array is empty, the `found` is false.
<原文结束>

# <翻译开始>
// PopRight 从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
// md5:207fa7c7c4a04a10
# <翻译结束>


<原文开始>
// PopRand randomly pops and return an item out of array.
// Note that if the array is empty, the `found` is false.
<原文结束>

# <翻译开始>
// PopRand 从数组中随机弹出并返回一个元素。
// 注意，如果数组为空，`found` 将为 false。
// md5:29338267db400401
# <翻译结束>


<原文开始>
// PopRands randomly pops and returns `size` items out of array.
// If the given `size` is greater than size of the array, it returns all elements of the array.
// Note that if given `size` <= 0 or the array is empty, it returns nil.
<原文结束>

# <翻译开始>
// PopRands 随机地从数组中弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，它将返回数组的所有元素。
// 注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:9fd270d3d3021d32
# <翻译结束>


<原文开始>
// PopLefts pops and returns `size` items from the beginning of array.
// If the given `size` is greater than size of the array, it returns all elements of the array.
// Note that if given `size` <= 0 or the array is empty, it returns nil.
<原文结束>

# <翻译开始>
// PopLefts 从数组开始处弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的长度，它将返回数组中的所有元素。
// 请注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:3ecbe066336a9849
# <翻译结束>


<原文开始>
// PopRights pops and returns `size` items from the end of array.
// If the given `size` is greater than size of the array, it returns all elements of the array.
// Note that if given `size` <= 0 or the array is empty, it returns nil.
<原文结束>

# <翻译开始>
// PopRights 从数组末尾弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，它将返回数组中的所有元素。
// 注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:4f44f32fbb68fb50
# <翻译结束>


<原文开始>
// Range picks and returns items by range, like array[start:end].
// Notice, if in concurrent-safe usage, it returns a copy of slice;
// else a pointer to the underlying data.
//
// If `end` is negative, then the offset will start from the end of array.
// If `end` is omitted, then the sequence will have everything from start up
// until the end of the array.
<原文结束>

# <翻译开始>
// Range通过范围选择并返回项目，就像数组[start:end]一样。
// 请注意，如果在并发安全使用中，它将返回切片的副本；否则返回底层数据的指针。
// 
// 如果`end`为负数，则偏移量将从数组末尾开始。
// 如果省略`end`，则序列将包含从`start`到数组结尾的所有内容。
// md5:8b71690536bb9ec5
# <翻译结束>


<原文开始>
// SubSlice returns a slice of elements from the array as specified
// by the `offset` and `size` parameters.
// If in concurrent safe usage, it returns a copy of the slice; else a pointer.
//
// If offset is non-negative, the sequence will start at that offset in the array.
// If offset is negative, the sequence will start that far from the end of the array.
//
// If length is given and is positive, then the sequence will have up to that many elements in it.
// If the array is shorter than the length, then only the available array elements will be present.
// If length is given and is negative then the sequence will stop that many elements from the end of the array.
// If it is omitted, then the sequence will have everything from offset up until the end of the array.
//
// Any possibility crossing the left border of array, it will fail.
<原文结束>

# <翻译开始>
// SubSlice 返回数组中指定的一段元素切片。
// 如果在并发安全的使用场景下，它将返回切片的一个副本；否则返回切片的指针。
//
// 如果偏移量（offset）为非负数，序列将从数组的该位置开始。
// 如果偏移量为负数，序列将从数组末尾向前该距离的位置开始。
//
// 如果提供了长度（size）且为正数，那么序列将包含最多这么多元素。
// 如果数组比指定的长度短，则序列只包含可用的数组元素。
// 如果长度为负数，则序列将在距离数组末尾该数量的元素处停止。
// 如果省略长度参数，那么序列将从偏移量开始直到数组末尾的所有元素。
//
// 如果切片范围的起始位置超出数组左侧边界，操作将失败。
// md5:f87ecd35d1dd7ac8
# <翻译结束>


<原文开始>
// Append is alias of PushRight,please See PushRight.
<原文结束>

# <翻译开始>
// Append 是 PushRight 的别名，详情请参阅 PushRight。 md5:2f083a022f7fd9c3
# <翻译结束>


<原文开始>
// Len returns the length of array.
<原文结束>

# <翻译开始>
// Len 返回数组的长度。 md5:593b37501e98da95
# <翻译结束>


<原文开始>
// Slice returns the underlying data of array.
// Note that, if it's in concurrent-safe usage, it returns a copy of underlying data,
// or else a pointer to the underlying data.
<原文结束>

# <翻译开始>
// Slice 返回数组的底层数据。
// 注意，如果在并发安全的使用情况下，它会返回底层数据的副本，否则返回底层数据的指针。
// md5:111cbee45795a58b
# <翻译结束>


<原文开始>
// Interfaces returns current array as []interface{}.
<原文结束>

# <翻译开始>
// Interfaces 将当前数组作为 []interface{} 返回。 md5:f7a2e3459e185314
# <翻译结束>


<原文开始>
// Clone returns a new array, which is a copy of current array.
<原文结束>

# <翻译开始>
// Clone 返回一个新的数组，它是当前数组的副本。 md5:52ada4030c562295
# <翻译结束>


<原文开始>
// Clear deletes all items of current array.
<原文结束>

# <翻译开始>
// Clear 删除当前数组中的所有项目。 md5:3d9c6d68a5719979
# <翻译结束>


<原文开始>
// Contains checks whether a value exists in the array.
<原文结束>

# <翻译开始>
// Contains 检查值是否存在于数组中。 md5:f209e1f30dd53cb2
# <翻译结束>


<原文开始>
// ContainsI checks whether a value exists in the array with case-insensitively.
// Note that it internally iterates the whole array to do the comparison with case-insensitively.
<原文结束>

# <翻译开始>
// ContainsI 检查数组中是否存在某个值（忽略大小写）。
// 注意，它内部会遍历整个数组以进行不区分大小写的比较。
// md5:faf76a65365aa0ac
# <翻译结束>


<原文开始>
// Search searches array by `value`, returns the index of `value`,
// or returns -1 if not exists.
<原文结束>

# <翻译开始>
// Search 在数组中搜索 `value`，返回 `value` 的索引，
// 如果不存在则返回 -1。
// md5:787617bfeade8f93
# <翻译结束>


<原文开始>
// Unique uniques the array, clear repeated items.
// Example: [1,1,2,3,2] -> [1,2,3]
<原文结束>

# <翻译开始>
// Unique 去除数组中的重复元素。
// 例如：[1,1,2,3,2] -> [1,2,3]
// md5:5083aa414231fd30
# <翻译结束>


<原文开始>
// LockFunc locks writing by callback function `f`.
<原文结束>

# <翻译开始>
// LockFunc 通过回调函数 `f` 实现写入锁定。 md5:d45a130fa9aa0af2
# <翻译结束>


<原文开始>
// RLockFunc locks reading by callback function `f`.
<原文结束>

# <翻译开始>
// RLockFunc 通过回调函数 `f` 实现读取锁定。 md5:a45deee1e6f17c88
# <翻译结束>


<原文开始>
// Merge merges `array` into current array.
// The parameter `array` can be any garray or slice type.
// The difference between Merge and Append is Append supports only specified slice type,
// but Merge supports more parameter types.
<原文结束>

# <翻译开始>
// Merge 将 `array` 合并到当前数组中。
// 参数 `array` 可以是任何 garray 或切片类型。
// Merge 和 Append 的区别在于，Append 只支持特定的切片类型，
// 而 Merge 支持更多种类的参数类型。
// md5:465caccda38e84f8
# <翻译结束>


<原文开始>
// Fill fills an array with num entries of the value `value`,
// keys starting at the `startIndex` parameter.
<原文结束>

# <翻译开始>
// Fill 使用`value`值填充数组，从`startIndex`参数开始的num个条目。
// md5:0a7d3daa806b72ca
# <翻译结束>


<原文开始>
// Chunk splits an array into multiple arrays,
// the size of each array is determined by `size`.
// The last chunk may contain less than size elements.
<原文结束>

# <翻译开始>
// Chunk 将一个数组分割成多个子数组，每个子数组的大小由 `size` 决定。最后一个子数组可能包含少于 `size` 个元素。
// md5:0f1f74ff34633d24
# <翻译结束>


<原文开始>
// Pad pads array to the specified length with `value`.
// If size is positive then the array is padded on the right, or negative on the left.
// If the absolute value of `size` is less than or equal to the length of the array
// then no padding takes place.
<原文结束>

# <翻译开始>
// Pad 用`value`将数组填充到指定的长度。
// 如果大小为正数，则在右侧填充数组，如果为负数，则在左侧填充。
// 如果`size`的绝对值小于或等于数组的长度，则不进行填充。
// md5:fbe08b371c540418
# <翻译结束>


<原文开始>
// Rand randomly returns one item from array(no deleting).
<原文结束>

# <翻译开始>
// Rand 随机从数组中返回一个元素（不进行删除）。 md5:e152d2c5bc15ecd7
# <翻译结束>


<原文开始>
// Rands randomly returns `size` items from array(no deleting).
<原文结束>

# <翻译开始>
// Rands 随机从数组中返回 `size` 个元素（不删除）。 md5:09ad7802f8190e3c
# <翻译结束>


<原文开始>
// Shuffle randomly shuffles the array.
<原文结束>

# <翻译开始>
// 随机打乱数组。 md5:5897797461d9f11a
# <翻译结束>


<原文开始>
// Reverse makes array with elements in reverse order.
<原文结束>

# <翻译开始>
// Reverse 函数将数组元素反转顺序。 md5:cc34cd0a2fa08e1c
# <翻译结束>


<原文开始>
// Join joins array elements with a string `glue`.
<原文结束>

# <翻译开始>
// Join 使用字符串 `glue` 连接数组元素。 md5:ec3894b049af1251
# <翻译结束>


<原文开始>
// CountValues counts the number of occurrences of all values in the array.
<原文结束>

# <翻译开始>
// CountValues 计算数组中所有值出现的次数。 md5:95b4772dcb002365
# <翻译结束>


<原文开始>
// Iterator is alias of IteratorAsc.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名。 md5:1bfdea306db62845
# <翻译结束>


<原文开始>
// IteratorAsc iterates the array readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAsc 遍历数组，按照给定的回调函数 `f` 以升序进行只读访问。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:8a125e2dd8982d48
# <翻译结束>


<原文开始>
// IteratorDesc iterates the array readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 以降序遍历数组，并使用给定的回调函数`f`进行只读迭代。
// 如果`f`返回true，则继续遍历；如果返回false，则停止遍历。
// md5:ea0a3805bccce0f7
# <翻译结束>


<原文开始>
// String returns current array as a string, which implements like json.Marshal does.
<原文结束>

# <翻译开始>
// String 将当前数组转换为字符串，其实现方式类似于 json.Marshal。 md5:feda8f29233cde8d
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note that do not use pointer as its receiver here.
<原文结束>

# <翻译开始>
// MarshalJSON实现了json.Marshal接口的MarshalJSON方法。
// 注意，这里不要使用指针作为接收者。
// md5:b4f76062b07a5263
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for array.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于为数组设置任何类型的数据值。 md5:35211e747ab939ab
# <翻译结束>


<原文开始>
// Filter iterates array and filters elements using custom callback function.
// It removes the element from array if callback function `filter` returns true,
// it or else does nothing and continues iterating.
<原文结束>

# <翻译开始>
// Filter 遍历数组，并使用自定义回调函数过滤元素。
// 如果回调函数`filter`返回true，它将从数组中移除该元素，否则不做任何操作并继续遍历。
// md5:d33873cfb9f1bb38
# <翻译结束>


<原文开始>
// FilterEmpty removes all empty string value of the array.
<原文结束>

# <翻译开始>
// FilterEmpty 函数移除数组中的所有空字符串值。 md5:2b2e8cd6c844936a
# <翻译结束>


<原文开始>
// Walk applies a user supplied function `f` to every item of array.
<原文结束>

# <翻译开始>
// Walk 将用户提供的函数 `f` 应用到数组的每个元素上。 md5:51e35ea7c2c6525c
# <翻译结束>


<原文开始>
// IsEmpty checks whether the array is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查数组是否为空。 md5:fb6684351506a02d
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
# <翻译结束>

