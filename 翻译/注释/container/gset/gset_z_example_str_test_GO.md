
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
// NewStrSet create and returns a new set, which contains un-repeated items.
// The parameter `safe` is used to specify whether using set in concurrent-safety,
// which is false in default.
<原文结束>

# <翻译开始>
// NewStrSet 创建并返回一个新集合，其中包含不重复的元素。
// 参数 `safe` 用于指定是否在并发安全环境下使用集合，默认为 false。
// md5:b4b32102d4f1da78
# <翻译结束>


<原文开始>
	// May Output:
	// [str3 str1 str2]
<原文结束>

# <翻译开始>
	// May Output:
	// [str3 str1 str2]
# <翻译结束>


<原文开始>
// NewStrSetFrom returns a new set from `items`.
<原文结束>

# <翻译开始>
// NewStrSetFrom 从`items`创建一个新的集合。 md5:6f9a406a984403d2
# <翻译结束>


<原文开始>
	// May Output:
	// [str1 str2 str3]
<原文结束>

# <翻译开始>
	// May Output:
	// [str1 str2 str3]
# <翻译结束>


<原文开始>
// Add adds one or multiple items to the set.
<原文结束>

# <翻译开始>
// Add 将一个或多个项目添加到集合中。 md5:316141ff7d4b8e45
# <翻译结束>


<原文开始>
	// May Output:
	// [str str1 str2 str3]
	// false
<原文结束>

# <翻译开始>
	// May Output:
	// [str str1 str2 str3]
	// false
# <翻译结束>


<原文开始>
// AddIfNotExist checks whether item exists in the set,
// it adds the item to set and returns true if it does not exist in the set,
// or else it does nothing and returns false.
<原文结束>

# <翻译开始>
// AddIfNotExist 检查项是否存在于集合中，
// 如果项不存在于集合中，它会将项添加到集合中并返回 true，否则什么都不做并返回 false。
// md5:9cff508c42cffd55
# <翻译结束>


<原文开始>
// AddIfNotExistFunc checks whether item exists in the set,
// it adds the item to set and returns true if it does not exist in the set and function `f` returns true,
// or else it does nothing and returns false.
// Note that, the function `f` is executed without writing lock.
<原文结束>

# <翻译开始>
// AddIfNotExistFunc 检查项目是否存在于集合中，
// 如果项目不在集合中且函数 `f` 返回 true，则将其添加到集合并返回 true，
// 否则不做任何操作并返回 false。
// 请注意，函数 `f` 在无写锁的情况下执行。
// md5:0a51b9d79022ae82
# <翻译结束>


<原文开始>
	// May Output:
	// [str1 str2 str3 str]
	// true
<原文结束>

# <翻译开始>
	// May Output:
	// [str1 str2 str3 str]
	// true
# <翻译结束>


<原文开始>
// Clear deletes all items of the set.
<原文结束>

# <翻译开始>
// Clear 删除集合中的所有项。 md5:ce349f0cd3114465
# <翻译结束>


<原文开始>
// Complement returns a new set which is the complement from `set` to `full`.
// Which means, all the items in `newSet` are in `full` and not in `set`.
// It returns the difference between `full` and `set` if the given set `full` is not the full set of `set`.
<原文结束>

# <翻译开始>
// Complement 返回一个新的集合，该集合是相对于`set`到`full`的补集。
// 这意味着，`newSet`中的所有项都包含在`full`中但不包含在`set`中。
// 如果给定的集合`full`并不是`set`的全集，则它返回`full`与`set`之间的差异。
// md5:2116fbb7587db792
# <翻译结束>


<原文开始>
	// May Output:
	// [str4 str5]
<原文结束>

# <翻译开始>
	// May Output:
	// [str4 str5]
# <翻译结束>


<原文开始>
// Contains checks whether the set contains `item`.
<原文结束>

# <翻译开始>
// Contains 检查集合是否包含 `item`。 md5:20a3bdc6aeef1d67
# <翻译结束>


<原文开始>
// ContainsI checks whether a value exists in the set with case-insensitively.
// Note that it internally iterates the whole set to do the comparison with case-insensitively.
<原文结束>

# <翻译开始>
// ContainsI 检查集合中是否存在某个值（忽略大小写）。
// 注意，它内部会遍历整个集合以进行不区分大小写的比较。
// md5:851e1bbfa6da1bae
# <翻译结束>


<原文开始>
// Diff returns a new set which is the difference set from `set` to `other`.
// Which means, all the items in `newSet` are in `set` but not in `other`.
<原文结束>

# <翻译开始>
// Diff 返回一个新的集合，它是 `set` 与 `other` 之间的差集。
// 这意味着，`newSet` 中的所有项目都在 `set` 中，但不在 `other` 中。
// md5:6779e6e007651b53
# <翻译结束>


<原文开始>
// Equal checks whether the two sets equal.
<原文结束>

# <翻译开始>
// Equal 检查两个集合是否相等。 md5:105ea4dd39b57fe8
# <翻译结束>


<原文开始>
// Intersect returns a new set which is the intersection from `set` to `other`.
// Which means, all the items in `newSet` are in `set` and also in `other`.
<原文结束>

# <翻译开始>
// Intersect 返回一个新的集合，这个集合是 `set` 和 `other` 的交集。
// 这意味着，`newSet` 中的所有元素都既存在于 `set` 中也存在于 `other` 中。
// md5:327d3fcc12f06583
# <翻译结束>


<原文开始>
	// May Output:
	// [c a b]
<原文结束>

# <翻译开始>
	// May Output:
	// [c a b]
# <翻译结束>


<原文开始>
// IsSubsetOf checks whether the current set is a sub-set of `other`
<原文结束>

# <翻译开始>
// IsSubsetOf 检查当前集合是否是 `other` 的子集. md5:70b7ed1e77ec2f80
# <翻译结束>


<原文开始>
// Iterator iterates the set readonly with given callback function `f`,
// if `f` returns true then continue iterating; or false to stop.
<原文结束>

# <翻译开始>
// Iterator 使用给定的回调函数 `f` 遍历只读集合，如果 `f` 返回 true，则继续遍历；否则停止。
// md5:b896360b1cf6fc88
# <翻译结束>


<原文开始>
	// May Output:
	// Iterator a
	// Iterator b
	// Iterator c
	// Iterator d
<原文结束>

# <翻译开始>
	// May Output:
	// Iterator a
	// Iterator b
	// Iterator c
	// Iterator d
# <翻译结束>


<原文开始>
// Join joins items with a string `glue`.
<原文结束>

# <翻译开始>
// Join 使用字符串 `glue` 连接多个项目。 md5:c8699391999ac788
# <翻译结束>


<原文开始>
	// May Output:
	// b,c,d,a
<原文结束>

# <翻译开始>
	// May Output:
	// b,c,d,a
# <翻译结束>


<原文开始>
// LockFunc locks writing with callback function `f`.
<原文结束>

# <翻译开始>
// LockFunc 使用回调函数 `f` 为写入操作加锁。 md5:85d746d8a49edab7
# <翻译结束>


<原文开始>
	// May Output
	// [2 3 1]
<原文结束>

# <翻译开始>
	// 可能的输出
	// [2 3 1]
	// md5:294c6ba36e85ea4c
# <翻译结束>


<原文开始>
// MarshalJSON implements the interface MarshalJSON for json.Marshal.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
# <翻译结束>


<原文开始>
	// May Output:
	// {"Id":1,"Name":"john","Scores":["100","99","98"]}
<原文结束>

# <翻译开始>
	// May Output:
	// {"Id":1,"Name":"john","Scores":["100","99","98"]}
# <翻译结束>


<原文开始>
// Merge adds items from `others` sets into `set`.
<原文结束>

# <翻译开始>
// Merge 将 `others` 集合中的项目合并到 `set` 中。 md5:788b02e300c6f440
# <翻译结束>


<原文开始>
	// May Output:
	// [d a b c]
<原文结束>

# <翻译开始>
	// May Output:
	// [d a b c]
# <翻译结束>


<原文开始>
// Pops randomly pops an item from set.
<原文结束>

# <翻译开始>
// 随机从集合中弹出一个元素。 md5:56ac5a59d1852551
# <翻译结束>


<原文开始>
// Pops randomly pops `size` items from set.
// It returns all items if size == -1.
<原文结束>

# <翻译开始>
// Pops 从集合中随机弹出 `size` 个元素。
// 如果 size == -1，它将返回所有元素。
// md5:c687f88e0a2df8f2
# <翻译结束>


<原文开始>
	// May Output:
	// a
	// b
<原文结束>

# <翻译开始>
	// May Output:
	// a
	// b
# <翻译结束>


<原文开始>
// RLockFunc locks reading with callback function `f`.
<原文结束>

# <翻译开始>
// RLockFunc 使用回调函数 `f` 进行读取锁定。 md5:5fe2bf1a85ce319e
# <翻译结束>


<原文开始>
// Remove deletes `item` from set.
<原文结束>

# <翻译开始>
// Remove 从集合中删除 `item`。 md5:ab30c696cc44d190
# <翻译结束>


<原文开始>
	// May Output:
	// [b c d]
<原文结束>

# <翻译开始>
	// May Output:
	// [b c d]
# <翻译结束>


<原文开始>
// Size returns the size of the set.
<原文结束>

# <翻译开始>
// Size 返回集合的大小。 md5:0d55ac576b7779ee
# <翻译结束>


<原文开始>
// Slice returns the an of items of the set as slice.
<原文结束>

# <翻译开始>
// Slice 返回集合中的元素作为切片。 md5:f5bc80ac01ae812b
# <翻译结束>


<原文开始>
	// May Output:
	// [a,b,c,d]
<原文结束>

# <翻译开始>
	// May Output:
	// [a,b,c,d]
# <翻译结束>


<原文开始>
// String returns items as a string, which implements like json.Marshal does.
<原文结束>

# <翻译开始>
// String 将 items 转换为字符串，其实现方式类似于 json.Marshal。 md5:cedb10711c2e5dac
# <翻译结束>


<原文开始>
	// May Output:
	// "a","b","c","d"
<原文结束>

# <翻译开始>
	// May Output:
	// "a","b","c","d"
# <翻译结束>


<原文开始>
// Sum sums items. Note: The items should be converted to int type,
// or you'd get a result that you unexpected.
<原文结束>

# <翻译开始>
// Sum 对项目求和。注意：项目应转换为整型，
// 否则你可能会得到意想不到的结果。
// md5:7cca75708fbf4ffc
# <翻译结束>


<原文开始>
// Union returns a new set which is the union of `set` and `other`.
// Which means, all the items in `newSet` are in `set` or in `other`.
<原文结束>

# <翻译开始>
// Union 返回一个新集合，它是`set`和`other`的并集。
// 意味着，`newSet`中的所有项目都在`set`中或在`other`中。
// md5:420e241c3c12e8e6
# <翻译结束>


<原文开始>
	// May Output:
	// [a b c d]
<原文结束>

# <翻译开始>
	// May Output:
	// [a b c d]
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
# <翻译结束>


<原文开始>
	// May Output:
	// {1 john "99","98","100"}
<原文结束>

# <翻译开始>
	// May Output:
	// {1 john "99","98","100"}
# <翻译结束>


<原文开始>
// UnmarshalValue is an interface implement which sets any type of value for set.
<原文结束>

# <翻译开始>
// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置为集合。 md5:b119247f684920ad
# <翻译结束>


<原文开始>
// Walk applies a user supplied function `f` to every item of set.
<原文结束>

# <翻译开始>
// Walk应用用户提供的函数`f`到集合中的每一项。 md5:d6ceaae555e8a9e6
# <翻译结束>


<原文开始>
// Add prefix for given table names.
<原文结束>

# <翻译开始>
	// 为给定的表名添加前缀。 md5:dea7405f272e0c9e
# <翻译结束>


<原文开始>
	// May Output:
	// [gf_user gf_user_detail]
<原文结束>

# <翻译开始>
	// May Output:
	// [gf_user gf_user_detail]
# <翻译结束>

