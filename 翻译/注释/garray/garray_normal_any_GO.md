
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
// Array is a golang array with rich features.
// It contains a concurrent-safe/unsafe switch, which should be set
// when its initialization and cannot be changed then.
<原文结束>

# <翻译开始>
// Array 是一个具有丰富特性的 Go 语言数组。
// 它包含一个并发安全/不安全切换开关，该开关应在初始化时设置，并且之后不可更改。
# <翻译结束>


<原文开始>
// New creates and returns an empty array.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// New 创建并返回一个空数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
# <翻译结束>


<原文开始>
// NewArray is alias of New, please see New.
<原文结束>

# <翻译开始>
// NewArray 是 New 的别名，请参阅 New。
# <翻译结束>


<原文开始>
// NewArraySize create and returns an array with given size and cap.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewArraySize 根据给定的大小和容量创建并返回一个数组。
// 参数`safe`用于指定是否在并发安全的情况下使用数组，默认为false。
# <翻译结束>


<原文开始>
// NewArrayRange creates and returns an array by a range from `start` to `end`
// with step value `step`.
<原文结束>

# <翻译开始>
// NewArrayRange 根据指定的范围从 `start` 到 `end` 并以步长值 `step` 创建并返回一个数组。
# <翻译结束>


<原文开始>
// NewFrom is alias of NewArrayFrom.
// See NewArrayFrom.
<原文结束>

# <翻译开始>
// NewFrom 是 NewArrayFrom 的别名。
// 请参阅 NewArrayFrom。
# <翻译结束>


<原文开始>
// NewFromCopy is alias of NewArrayFromCopy.
// See NewArrayFromCopy.
<原文结束>

# <翻译开始>
// NewFromCopy 是 NewArrayFromCopy 的别名。
// 请参阅 NewArrayFromCopy。
# <翻译结束>


<原文开始>
// NewArrayFrom creates and returns an array with given slice `array`.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewArrayFrom 通过给定的切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
# <翻译结束>


<原文开始>
// NewArrayFromCopy creates and returns an array from a copy of given slice `array`.
// The parameter `safe` is used to specify whether using array in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewArrayFromCopy 通过复制给定切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
# <翻译结束>


<原文开始>
// At returns the value by the specified index.
// If the given `index` is out of range of the array, it returns `nil`.
<原文结束>

# <翻译开始>
// At通过指定的索引返回值。
// 如果给定的`index`超出数组范围，它将返回`nil`。
# <翻译结束>


<原文开始>
// Get returns the value by the specified index.
// If the given `index` is out of range of the array, the `found` is false.
<原文结束>

# <翻译开始>
// Get 通过指定的索引返回值。
// 如果给定的 `index` 超出了数组的范围，那么 `found` 将为 false。
# <翻译结束>


<原文开始>
// Set sets value to specified index.
<原文结束>

# <翻译开始>
// Set将值设置到指定的索引位置。
# <翻译结束>


<原文开始>
// SetArray sets the underlying slice array with the given `array`.
<原文结束>

# <翻译开始>
// SetArray 将底层的切片数组设置为给定的 `array`。
# <翻译结束>


<原文开始>
// Replace replaces the array items by given `array` from the beginning of array.
<原文结束>

# <翻译开始>
// Replace 从数组起始位置开始，使用给定的`array`替换原有数组元素。
# <翻译结束>


<原文开始>
// Sum returns the sum of values in an array.
<原文结束>

# <翻译开始>
// Sum 返回数组中所有值的和。
# <翻译结束>


<原文开始>
// SortFunc sorts the array by custom function `less`.
<原文结束>

# <翻译开始>
// SortFunc 通过自定义函数 `less` 对数组进行排序。
# <翻译结束>


<原文开始>
// InsertBefore inserts the `values` to the front of `index`.
<原文结束>

# <翻译开始>
// InsertBefore 将 `values` 插入到 `index` 之前的位置。
# <翻译结束>


<原文开始>
// InsertAfter inserts the `values` to the back of `index`.
<原文结束>

# <翻译开始>
// InsertAfter 在`index`之后插入`values`。
# <翻译结束>


<原文开始>
// Remove removes an item by index.
// If the given `index` is out of range of the array, the `found` is false.
<原文结束>

# <翻译开始>
// Remove 通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
# <翻译结束>


<原文开始>
// doRemoveWithoutLock removes an item by index without lock.
<原文结束>

# <翻译开始>
// doRemoveWithoutLock 在没有加锁的情况下通过索引移除一个项。
# <翻译结束>


<原文开始>
// Determine array boundaries when deleting to improve deletion efficiency.
<原文结束>

# <翻译开始>
// 确定删除时的数组边界以提高删除效率
# <翻译结束>


<原文开始>
	// If it is a non-boundary delete,
	// it will involve the creation of an array,
	// then the deletion is less efficient.
<原文结束>

# <翻译开始>
// 如果这是一个非边界删除，
// 那么它将涉及创建一个数组，
// 因此，删除操作效率较低。
# <翻译结束>


<原文开始>
// RemoveValue removes an item by value.
// It returns true if value is found in the array, or else false if not found.
<原文结束>

# <翻译开始>
// RemoveValue 通过值移除一个元素。
// 若在数组中找到该值，则返回 true，否则（未找到时）返回 false。
# <翻译结束>


<原文开始>
// RemoveValues removes multiple items by `values`.
<原文结束>

# <翻译开始>
// RemoveValues 通过 `values` 移除多个项目。
# <翻译结束>


<原文开始>
// PushLeft pushes one or multiple items to the beginning of array.
<原文结束>

# <翻译开始>
// PushLeft 将一个或多个元素推送到数组的起始位置。
# <翻译结束>


<原文开始>
// PushRight pushes one or multiple items to the end of array.
// It equals to Append.
<原文结束>

# <翻译开始>
// PushRight将一个或多个元素推送到数组的末尾。
// 它等同于Append。
# <翻译结束>


<原文开始>
// PopRand randomly pops and return an item out of array.
// Note that if the array is empty, the `found` is false.
<原文结束>

# <翻译开始>
// PopRand 随机地从数组中弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
# <翻译结束>


<原文开始>
// PopRands randomly pops and returns `size` items out of array.
<原文结束>

# <翻译开始>
// PopRands 随机地从数组中弹出并返回 `size` 个元素。
# <翻译结束>


<原文开始>
// PopLeft pops and returns an item from the beginning of array.
// Note that if the array is empty, the `found` is false.
<原文结束>

# <翻译开始>
// PopLeft 从数组开头弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
# <翻译结束>


<原文开始>
// PopRight pops and returns an item from the end of array.
// Note that if the array is empty, the `found` is false.
<原文结束>

# <翻译开始>
// PopRight从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则`found`为false。
# <翻译结束>


<原文开始>
// PopLefts pops and returns `size` items from the beginning of array.
<原文结束>

# <翻译开始>
// PopLefts 从数组开头弹出并返回 `size` 个元素。
# <翻译结束>


<原文开始>
// PopRights pops and returns `size` items from the end of array.
<原文结束>

# <翻译开始>
// PopRights 从数组末尾弹出并返回 `size` 个元素。
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
// Range 函数通过范围选择并返回数组中的元素，类似于 array[start:end]。
// 注意：在并发安全的使用场景下，它会返回一个原数据的副本；否则，返回的是底层数据的指针。
//
// 如果 `end` 为负数，则偏移量将从数组末尾开始计算。
// 如果省略了 `end`，则序列将包含从 start 开始直到数组末尾的所有元素。
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
// SubSlice 返回数组中由 `offset` 和 `size` 参数指定的元素子序列，并将其作为切片。
// 若在并发安全场景下使用，返回该切片的副本；否则返回指向切片的指针。
//
// 如果 offset 非负，则序列从数组该偏移位置开始。
// 如果 offset 为负，则序列从数组末尾向前偏移该距离的位置开始。
//
// 如果提供了 length 并且为正数，则序列将包含最多该数量的元素。
// 若数组长度小于 length，则序列仅包含数组中可获得的元素。
// 如果 length 为负数，则序列将在数组末尾向前停在该距离的位置。
// 如果未提供 length，则序列包含从 offset 开始直到数组末尾的所有元素。
//
// 若有任何可能穿越数组左边界的情况，函数将失败。
# <翻译结束>


<原文开始>
// Append is alias of PushRight, please See PushRight.
<原文结束>

# <翻译开始>
// Append 是 PushRight 的别名，请参阅 PushRight。
# <翻译结束>


<原文开始>
// Len returns the length of array.
<原文结束>

# <翻译开始>
// Len 返回数组的长度。
# <翻译结束>


<原文开始>
// Slice returns the underlying data of array.
// Note that, if it's in concurrent-safe usage, it returns a copy of underlying data,
// or else a pointer to the underlying data.
<原文结束>

# <翻译开始>
// Slice 返回数组的基础数据。
// 注意，如果它在并发安全的使用场景下，会返回基础数据的一个副本，
// 否则，则返回指向基础数据的指针。
# <翻译结束>


<原文开始>
// Interfaces returns current array as []interface{}.
<原文结束>

# <翻译开始>
// Interfaces 函数将当前数组转换为 []interface{} 类型并返回。
# <翻译结束>


<原文开始>
// Clone returns a new array, which is a copy of current array.
<原文结束>

# <翻译开始>
// Clone 返回一个新的数组，它是当前数组的一个副本。
# <翻译结束>


<原文开始>
// Clear deletes all items of current array.
<原文结束>

# <翻译开始>
// 清空删除当前数组中的所有元素。
# <翻译结束>


<原文开始>
// Contains checks whether a value exists in the array.
<原文结束>

# <翻译开始>
// Contains 检查某个值是否存在于数组中。
# <翻译结束>


<原文开始>
// Search searches array by `value`, returns the index of `value`,
// or returns -1 if not exists.
<原文结束>

# <翻译开始>
// Search 在数组中通过 `value` 进行搜索，返回 `value` 的索引，
// 若不存在，则返回 -1。
# <翻译结束>


<原文开始>
// Unique uniques the array, clear repeated items.
// Example: [1,1,2,3,2] -> [1,2,3]
<原文结束>

# <翻译开始>
// Unique 函数用于对数组去重，清除重复的元素。
// 示例：[1,1,2,3,2] -> [1,2,3]
# <翻译结束>


<原文开始>
// LockFunc locks writing by callback function `f`.
<原文结束>

# <翻译开始>
// LockFunc 通过回调函数`f`进行写入锁定。
# <翻译结束>


<原文开始>
// RLockFunc locks reading by callback function `f`.
<原文结束>

# <翻译开始>
// RLockFunc 通过回调函数`f`锁定读取操作。
# <翻译结束>


<原文开始>
// Merge merges `array` into current array.
// The parameter `array` can be any garray or slice type.
// The difference between Merge and Append is Append supports only specified slice type,
// but Merge supports more parameter types.
<原文结束>

# <翻译开始>
// Merge 将`array`合并到当前数组中。
// 参数`array`可以是任何garray类型或切片类型。
// Merge 和 Append 的区别在于，Append 仅支持特定类型的切片作为参数，
// 而 Merge 支持更多类型的参数。
# <翻译结束>


<原文开始>
// Fill fills an array with num entries of the value `value`,
// keys starting at the `startIndex` parameter.
<原文结束>

# <翻译开始>
// Fill 用 `value` 值填充数组，填充 num 个条目，
// 键（索引）从 `startIndex` 参数开始。
# <翻译结束>


<原文开始>
// Chunk splits an array into multiple arrays,
// the size of each array is determined by `size`.
// The last chunk may contain less than size elements.
<原文结束>

# <翻译开始>
// Chunk 函数将一个数组分割成多个子数组，
// 每个子数组的大小由参数 `size` 确定。
// 最后一个子数组可能包含少于 size 个元素。
# <翻译结束>


<原文开始>
// Pad pads array to the specified length with `value`.
// If size is positive then the array is padded on the right, or negative on the left.
// If the absolute value of `size` is less than or equal to the length of the array
// then no padding takes place.
<原文结束>

# <翻译开始>
// Pad 通过 `value` 值对数组进行填充，以达到指定长度。
// 如果 size 为正数，则在数组右侧进行填充；若为负数，则在左侧填充。
// 若 `size` 的绝对值小于或等于数组的长度，则不进行填充操作。
# <翻译结束>


<原文开始>
// Rand randomly returns one item from array(no deleting).
<原文结束>

# <翻译开始>
// Rand 随机地从数组中返回一个元素（不删除）。
# <翻译结束>


<原文开始>
// Rands randomly returns `size` items from array(no deleting).
<原文结束>

# <翻译开始>
// Rands 随机返回数组中的 `size` 个元素（不删除）。
# <翻译结束>


<原文开始>
// Shuffle randomly shuffles the array.
<原文结束>

# <翻译开始>
// Shuffle 随机地对数组进行洗牌。
# <翻译结束>


<原文开始>
// Reverse makes array with elements in reverse order.
<原文结束>

# <翻译开始>
// Reverse 将数组元素按逆序排列生成新数组。
# <翻译结束>


<原文开始>
// Join joins array elements with a string `glue`.
<原文结束>

# <翻译开始>
// Join 通过字符串 `glue` 连接数组元素。
# <翻译结束>


<原文开始>
// CountValues counts the number of occurrences of all values in the array.
<原文结束>

# <翻译开始>
// CountValues 计算数组中所有值出现的次数。
# <翻译结束>


<原文开始>
// Iterator is alias of IteratorAsc.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名。
# <翻译结束>


<原文开始>
// IteratorAsc iterates the array readonly in ascending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorAsc 以升序遍历给定数组，并使用回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
# <翻译结束>


<原文开始>
// IteratorDesc iterates the array readonly in descending order with given callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 函数以降序遍历给定的数组，并使用指定回调函数 `f` 进行只读操作。
// 若 `f` 返回 true，则继续迭代；若返回 false，则停止迭代。
# <翻译结束>


<原文开始>
// String returns current array as a string, which implements like json.Marshal does.
<原文结束>

# <翻译开始>
// String 方法将当前数组以字符串形式返回，其实现方式类似于 json.Marshal。
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note that do not use pointer as its receiver here.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
// 注意：此处接收者不使用指针。
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for array.
<原文结束>

# <翻译开始>
// UnmarshalValue 实现了一个接口，该接口用于为数组设置任意类型的值。
# <翻译结束>


<原文开始>
// Filter iterates array and filters elements using custom callback function.
// It removes the element from array if callback function `filter` returns true,
// it or else does nothing and continues iterating.
<原文结束>

# <翻译开始>
// Filter 对数组进行迭代，并通过自定义回调函数进行元素过滤。
// 如果回调函数 `filter` 返回 true，则从数组中移除该元素；
// 否则不做任何处理并继续迭代。
# <翻译结束>


<原文开始>
// FilterNil removes all nil value of the array.
<原文结束>

# <翻译开始>
// FilterNil 移除数组中所有的 nil 值。
# <翻译结束>


<原文开始>
// FilterEmpty removes all empty value of the array.
// Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.
<原文结束>

# <翻译开始>
// FilterEmpty 用于移除数组中所有空值。
// 下列值被认为是空值：0, nil, false, "", 以及长度为0的slice、map或chan。
# <翻译结束>


<原文开始>
// Walk applies a user supplied function `f` to every item of array.
<原文结束>

# <翻译开始>
// Walk 对数组中的每一项应用用户提供的函数 `f`。
# <翻译结束>


<原文开始>
// IsEmpty checks whether the array is empty.
<原文结束>

# <翻译开始>
// IsEmpty 检查数组是否为空。
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy 实现接口，用于当前类型的深度复制。
# <翻译结束>

