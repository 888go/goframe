
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//
<原文结束>

# <翻译开始>
// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//
# <翻译结束>


<原文开始>
// NewStrSet create and returns a new set, which contains un-repeated items.
// The parameter `safe` is used to specify whether using set in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewStrSet 创建并返回一个新的不包含重复项的集合。
// 参数`safe`用于指定是否在并发安全的情况下使用集合，默认为false。
// 这里，NewStrSet函数用于创建一个字符串集合，并确保其中的元素互不重复。该函数接受一个可选参数`safe`，它是一个布尔值，表示是否需要保证在并发环境下的安全性。如果不特别指定，那么默认情况下这个集合是不保证并发安全的。
# <翻译结束>


<原文开始>
// NewStrSetFrom returns a new set from `items`.
<原文结束>

# <翻译开始>
// NewStrSetFrom 从 `items` 中创建并返回一个新的集合。
# <翻译结束>


<原文开始>
// Iterator iterates the set readonly with given callback function `f`,
// if `f` returns true then continue iterating; or false to stop.
<原文结束>

# <翻译开始>
// Iterator 使用给定的回调函数`f`对集合进行只读遍历，
// 如果`f`返回true，则继续遍历；若返回false，则停止遍历。
# <翻译结束>


<原文开始>
// Add adds one or multiple items to the set.
<原文结束>

# <翻译开始>
// Add 向集合中添加一个或多个项目。
# <翻译结束>


<原文开始>
// AddIfNotExist checks whether item exists in the set,
// it adds the item to set and returns true if it does not exist in the set,
// or else it does nothing and returns false.
<原文结束>

# <翻译开始>
// AddIfNotExist 检查项目是否已存在于集合中，
// 若该项目不在集合中，则将其添加到集合并返回 true；
// 否则，不做任何操作并返回 false。
# <翻译结束>


<原文开始>
// AddIfNotExistFunc checks whether item exists in the set,
// it adds the item to set and returns true if it does not exists in the set and
// function `f` returns true, or else it does nothing and returns false.
//
// Note that, the function `f` is executed without writing lock.
<原文结束>

# <翻译开始>
// AddIfNotExistFunc 检查项是否已存在于集合中，
// 如果该项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合并返回 true；
// 否则，不执行任何操作并返回 false。
//
// 注意，函数 `f` 在无写入锁的情况下执行。
# <翻译结束>


<原文开始>
// AddIfNotExistFuncLock checks whether item exists in the set,
// it adds the item to set and returns true if it does not exists in the set and
// function `f` returns true, or else it does nothing and returns false.
//
// Note that, the function `f` is executed without writing lock.
<原文结束>

# <翻译开始>
// AddIfNotExistFuncLock 检查项是否存在集合中，
// 如果该项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合并返回 true；
// 否则，不执行任何操作并返回 false。
//
// 注意，函数 `f` 在无写入锁的情况下执行。
# <翻译结束>


<原文开始>
// Contains checks whether the set contains `item`.
<原文结束>

# <翻译开始>
// Contains 检查集合中是否包含 `item`。
# <翻译结束>


<原文开始>
// ContainsI checks whether a value exists in the set with case-insensitively.
// Note that it internally iterates the whole set to do the comparison with case-insensitively.
<原文结束>

# <翻译开始>
// ContainsI 检查某个值是否以不区分大小写的方式存在于集合中。
// 注意：它内部会遍历整个集合，以不区分大小写的方式进行比较。
# <翻译结束>


<原文开始>
// Remove deletes `item` from set.
<原文结束>

# <翻译开始>
// Remove 从集合中删除`item`。
# <翻译结束>


<原文开始>
// Size returns the size of the set.
<原文结束>

# <翻译开始>
// Size 返回集合的大小。
# <翻译结束>


<原文开始>
// Clear deletes all items of the set.
<原文结束>

# <翻译开始>
// 清除删除集合中的所有项。
# <翻译结束>


<原文开始>
// Slice returns the an of items of the set as slice.
<原文结束>

# <翻译开始>
// Slice 返回集合中项目的切片形式。
# <翻译结束>


<原文开始>
// Join joins items with a string `glue`.
<原文结束>

# <翻译开始>
// Join通过字符串`glue`连接items。
# <翻译结束>


<原文开始>
// String returns items as a string, which implements like json.Marshal does.
<原文结束>

# <翻译开始>
// String 返回 items 作为字符串，其实现方式类似于 json.Marshal。
# <翻译结束>


<原文开始>
// LockFunc locks writing with callback function `f`.
<原文结束>

# <翻译开始>
// LockFunc 使用回调函数`f`进行写入锁定。
# <翻译结束>


<原文开始>
// RLockFunc locks reading with callback function `f`.
<原文结束>

# <翻译开始>
// RLockFunc 通过回调函数 `f` 对读取进行加锁。
# <翻译结束>


<原文开始>
// Equal checks whether the two sets equal.
<原文结束>

# <翻译开始>
// Equal 检查两个集合是否相等。
# <翻译结束>


<原文开始>
// IsSubsetOf checks whether the current set is a sub-set of `other`.
<原文结束>

# <翻译开始>
// IsSubsetOf 检查当前集合是否为 `other` 的子集。
# <翻译结束>


<原文开始>
// Union returns a new set which is the union of `set` and `other`.
// Which means, all the items in `newSet` are in `set` or in `other`.
<原文结束>

# <翻译开始>
// Union 返回一个新的集合，该集合是 `set` 和 `other` 的并集。
// 这意味着，`newSet` 中的所有元素都在 `set` 或者 `other` 中。
# <翻译结束>


<原文开始>
// Diff returns a new set which is the difference set from `set` to `other`.
// Which means, all the items in `newSet` are in `set` but not in `other`.
<原文结束>

# <翻译开始>
// Diff 返回一个新的集合，这个集合是 `set` 与 `other` 的差集。
// 这意味着，新集合 `newSet` 中的所有元素都在 `set` 中，但不在 `other` 中。
# <翻译结束>


<原文开始>
// Intersect returns a new set which is the intersection from `set` to `other`.
// Which means, all the items in `newSet` are in `set` and also in `other`.
<原文结束>

# <翻译开始>
// Intersect 返回一个新的集合，它是从 `set` 到 `other` 的交集。
// 这意味着，`newSet` 中的所有元素都在 `set` 中，并且也在 `other` 中。
# <翻译结束>


<原文开始>
// Complement returns a new set which is the complement from `set` to `full`.
// Which means, all the items in `newSet` are in `full` and not in `set`.
//
// It returns the difference between `full` and `set`
// if the given set `full` is not the full set of `set`.
<原文结束>

# <翻译开始>
// Complement 返回一个新的集合，该集合为 `set` 在 `full` 中的补集。
// 这意味着，新集合 `newSet` 中的所有元素都在 `full` 中但不在 `set` 中。
//
// 如果给定的集合 `full` 不是 `set` 的全集，则返回 `full` 与 `set` 之间的差集。
# <翻译结束>


<原文开始>
// Merge adds items from `others` sets into `set`.
<原文结束>

# <翻译开始>
// Merge 将 `others` 中的元素合并到 `set` 中。
# <翻译结束>


<原文开始>
// Sum sums items.
// Note: The items should be converted to int type,
// or you'd get a result that you unexpected.
<原文结束>

# <翻译开始>
// Sum 计算项的总和。
// 注意：项应转换为 int 类型，
// 否则你将得到一个意想不到的结果。
# <翻译结束>


<原文开始>
// Pop randomly pops an item from set.
<原文结束>

# <翻译开始>
// Pop 随机地从集合中弹出一个元素。
# <翻译结束>


<原文开始>
// Pops randomly pops `size` items from set.
// It returns all items if size == -1.
<原文结束>

# <翻译开始>
// Pops 随机地从集合中弹出 `size` 个元素。
// 如果 size == -1，则返回所有元素。
# <翻译结束>


<原文开始>
// Walk applies a user supplied function `f` to every item of set.
<原文结束>

# <翻译开始>
// Walk 对集合中的每一个元素应用用户提供的函数 `f`。
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
// UnmarshalValue is an interface implement which sets any type of value for set.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于为 set 设置任意类型的值。
# <翻译结束>


<原文开始>
// DeepCopy implements interface for deep copy of current type.
<原文结束>

# <翻译开始>
// DeepCopy 实现接口，用于当前类型的深度复制。
# <翻译结束>

