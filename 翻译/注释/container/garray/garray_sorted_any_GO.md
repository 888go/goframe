
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
// SortedArray is a golang sorted array with rich features.
// It is using increasing order in default, which can be changed by
// setting it a custom comparator.
// It contains a concurrent-safe/unsafe switch, which should be set
// when its initialization and cannot be changed then.
<原文结束>

# <翻译开始>
// SortedArray 是一个功能丰富的 Go 语言排序数组。
// 默认情况下，它使用递增顺序，但可以通过设置自定义比较器进行更改。
// 它包含一个并发安全/不安全的开关，该开关应在初始化时设置并且不能更改。
// md5:48308289d58755e8
# <翻译结束>


<原文开始>
// Whether enable unique feature(false)
<原文结束>

# <翻译开始>
// 是否启用唯一功能（false）. md5:e1a1e6b26151e91d
# <翻译结束>


<原文开始>
// Comparison function(it returns -1: a < b; 0: a == b; 1: a > b)
<原文结束>

# <翻译开始>
// 比较函数（返回值：-1 表示 a < b；0 表示 a == b；1 表示 a > b）. md5:2be44acd57b55d6a
# <翻译结束>


<原文开始>
// NewSortedArray creates and returns an empty sorted array.
// The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.
// The parameter `comparator` used to compare values to sort in array,
// if it returns value < 0, means `a` < `b`; the `a` will be inserted before `b`;
// if it returns value = 0, means `a` = `b`; the `a` will be replaced by     `b`;
// if it returns value > 0, means `a` > `b`; the `a` will be inserted after  `b`;
<原文结束>

# <翻译开始>
// NewSortedArray 创建并返回一个空的排序数组。
// 参数 `safe` 用于指定是否在并发安全环境下使用数组，默认为 false。
// 参数 `comparator` 用于比较数组中值的排序，
// 若返回值小于 0，表示 `a` 小于 `b`；`a` 将插入到 `b` 之前；
// 若返回值等于 0，表示 `a` 等于 `b`；`a` 将被 `b` 替换；
// 若返回值大于 0，表示 `a` 大于 `b`；`a` 将插入到 `b` 之后。
// md5:72443a89d087c135
# <翻译结束>


<原文开始>
// NewSortedArraySize create and returns an sorted array with given size and cap.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewSortedArraySize 创建并返回一个给定大小和容量的排序数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:827c3309faba84ac
# <翻译结束>


<原文开始>
// NewSortedArrayRange creates and returns an array by a range from `start` to `end`
// with step value `step`.
<原文结束>

# <翻译开始>
// NewSortedArrayRange 创建并返回一个从 `start` 到 `end` 的范围，步长为 `step` 的数组。
// md5:93c103a8dc8cf9d7
# <翻译结束>


<原文开始>
// NewSortedArrayFrom creates and returns an sorted array with given slice `array`.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewSortedArrayFrom 根据给定的切片 `array` 创建并返回一个排序数组。
// 参数 `safe` 用于指定是否使用并发安全的数组，默认为 false。
// md5:764ff7e74cab303e
# <翻译结束>


<原文开始>
// NewSortedArrayFromCopy creates and returns an sorted array from a copy of given slice `array`.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewSortedArrayFromCopy 根据给定切片 `array` 的副本创建并返回一个已排序的数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:ec79e11f360050f4
# <翻译结束>


<原文开始>
// At returns the value by the specified index.
// If the given `index` is out of range of the array, it returns `nil`.
<原文结束>

# <翻译开始>
// At 通过指定的索引返回值。
// 如果给定的`index`超出了数组的范围，它将返回`nil`。
// md5:09a7e6585d2eba1a
# <翻译结束>


<原文开始>
// SetArray sets the underlying slice array with the given `array`.
<原文结束>

# <翻译开始>
// SetArray 使用给定的 `array` 设置底层切片数组。 md5:160b43a5c0ec752c
# <翻译结束>


<原文开始>
// SetComparator sets/changes the comparator for sorting.
// It resorts the array as the comparator is changed.
<原文结束>

# <翻译开始>
// SetComparator 设置/更改排序的比较器。
// 当比较器更改时，它会重新对数组进行排序。
// md5:1323d8fba2b97b75
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
// Add adds one or multiple values to sorted array, the array always keeps sorted.
// It's alias of function Append, see Append.
<原文结束>

# <翻译开始>
// Add 将一个或多个值添加到已排序的数组中，数组始终保持排序。它是Append函数的别名，请参阅Append。
// md5:34facedfc7e1b731
# <翻译结束>


<原文开始>
// Append adds one or multiple values to sorted array, the array always keeps sorted.
<原文结束>

# <翻译开始>
// Append 向已排序的数组中添加一个或多个值，数组将始终保持排序状态。 md5:f839b377c2c77f6b
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
// RemoveValues removes an item by `values`.
<原文结束>

# <翻译开始>
// RemoveValues 通过 `values` 删除一个项目。 md5:05e01eb00e998269
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
<原文结束>

# <翻译开始>
// PopRands 随机地从数组中弹出并返回 `size` 个元素。 md5:3e1b1cbd52abd4cf
# <翻译结束>


<原文开始>
// PopLefts pops and returns `size` items from the beginning of array.
<原文结束>

# <翻译开始>
// PopLefts 从数组开头弹出并返回 `size` 个元素。 md5:4a903258f1fe1dd4
# <翻译结束>


<原文开始>
// PopRights pops and returns `size` items from the end of array.
<原文结束>

# <翻译开始>
// PopRights 从数组末尾移除并返回 `size` 个元素。 md5:0b04e6ad99e5349b
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
// Sum returns the sum of values in an array.
<原文结束>

# <翻译开始>
// Sum 返回数组中所有值的和。 md5:b2148175a749b162
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
// Contains checks whether a value exists in the array.
<原文结束>

# <翻译开始>
// Contains 检查值是否存在于数组中。 md5:f209e1f30dd53cb2
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
// Binary search.
// It returns the last compared index and the result.
// If `result` equals to 0, it means the value at `index` is equals to `value`.
// If `result` lesser than 0, it means the value at `index` is lesser than `value`.
// If `result` greater than 0, it means the value at `index` is greater than `value`.
<原文结束>

# <翻译开始>
// 二分查找。
// 它返回最后比较的索引和结果。
// 如果 `result` 等于 0，表示索引处的值等于 `value`。
// 如果 `result` 小于 0，表示索引处的值小于 `value`。
// 如果 `result` 大于 0，表示索引处的值大于 `value`。
// md5:869c6a1ccba79c7a
# <翻译结束>


<原文开始>
// SetUnique sets unique mark to the array,
// which means it does not contain any repeated items.
// It also does unique check, remove all repeated items.
<原文结束>

# <翻译开始>
// SetUnique 将唯一标记设置到数组中，
// 这意味着它不包含任何重复的项目。
// 它还会进行唯一性检查，删除所有重复项。
// md5:acbac75bf944670c
# <翻译结束>


<原文开始>
// Unique uniques the array, clear repeated items.
<原文结束>

# <翻译开始>
// Unique 函数用于清除非唯一元素，确保数组中的每个元素都是唯一的。 md5:6dfd767cdbb67ed2
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
// LockFunc locks writing by callback function `f`.
<原文结束>

# <翻译开始>
// LockFunc 通过回调函数 `f` 实现写入锁定。 md5:d45a130fa9aa0af2
# <翻译结束>


<原文开始>
// Keep the array always sorted.
<原文结束>

# <翻译开始>
	// 保持数组始终排序。 md5:b2ef189f10478e96
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
// Chunk splits an array into multiple arrays,
// the size of each array is determined by `size`.
// The last chunk may contain less than size elements.
<原文结束>

# <翻译开始>
// Chunk 将一个数组分割成多个子数组，每个子数组的大小由 `size` 决定。最后一个子数组可能包含少于 `size` 个元素。
// md5:0f1f74ff34633d24
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
// Note that the comparator is set as string comparator in default.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。
// 注意，默认情况下，比较器被设置为字符串比较器。
// md5:8af2d4d6f742bb31
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for array.
// Note that the comparator is set as string comparator in default.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于为数组设置任何类型的价值。
// 注意，比较器默认设置为字符串比较器。
// md5:5c9d5d1af1e97ec8
# <翻译结束>


<原文开始>
// FilterNil removes all nil value of the array.
<原文结束>

# <翻译开始>
// FilterNil 删除数组中的所有空值（nil）。 md5:df6d66c2056b4815
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
// FilterEmpty removes all empty value of the array.
// Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.
<原文结束>

# <翻译开始>
// FilterEmpty 移除数组中的所有空值。
// 被认为是空的值包括：0，nil，false，""，切片、映射（map）或通道（channel）的长度为0。
// md5:da01f627cd0962db
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
// getComparator returns the comparator if it's previously set,
// or else it panics.
<原文结束>

# <翻译开始>
// getComparator 如果之前已设置比较器，则返回该比较器，否则将引发恐慌。
// md5:03eac9fd6d838369
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
# <翻译结束>

