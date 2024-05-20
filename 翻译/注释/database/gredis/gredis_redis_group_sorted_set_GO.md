
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
// IGroupSortedSet manages redis sorted set operations.
// Implements see redis.GroupSortedSet.
<原文结束>

# <翻译开始>
// IGroupSortedSet 管理 redis 排序集合操作。
// 实现了 redis.GroupSortedSet。
// md5:85c86f571889c1f2
# <翻译结束>


<原文开始>
// ZAddOption provides options for function ZAdd.
<原文结束>

# <翻译开始>
// ZAddOption 为 ZAdd 函数提供选项。. md5:b3e234b14d4a1ca8
# <翻译结束>


<原文开始>
	// Only update existing elements if the new score is less than the current score.
	// This flag doesn't prevent adding new elements.
<原文结束>

# <翻译开始>
// 只有当新分数小于当前分数时，才更新已存在的元素。
// 此标志不会阻止添加新元素。
// md5:df3556a5d410e3c9
# <翻译结束>


<原文开始>
	// Only update existing elements if the new score is greater than the current score.
	// This flag doesn't prevent adding new elements.
<原文结束>

# <翻译开始>
// 只有当新分数大于当前分数时，才更新现有元素。此标志不会阻止添加新元素。
// md5:4866b5e44d3c1bec
# <翻译结束>


<原文开始>
	// Modify the return value from the number of new elements added, to the total number of elements changed (CH is an abbreviation of changed).
	// Changed elements are new elements added and elements already existing for which the score was updated.
	// So elements specified in the command line having the same score as they had in the past are not counted.
	// Note: normally the return value of ZAdd only counts the number of new elements added.
<原文结束>

# <翻译开始>
// 将返回值从新添加元素的数量修改为更改的总元素数量（CH代表已更改）。 
// 已更改的元素包括新添加的元素和分数已被更新的现有元素。 
// 因此，命令行中指定的与过去相同的分数的元素不计入总数。 
// 注意：通常情况下，ZAdd的返回值只计算新添加的元素数量。
// md5:f80865660e63c42c
# <翻译结束>


<原文开始>
// When this option is specified ZAdd acts like ZIncrBy. Only one score-element pair can be specified in this mode.
<原文结束>

# <翻译开始>
// 当指定了此选项时，ZAdd 命令的行为类似于 ZIncrBy。在这种模式下，只能指定一个分数-元素对。. md5:bb002fb3eec4eb13
# <翻译结束>


<原文开始>
// ZAddMember is element struct for set.
<原文结束>

# <翻译开始>
// ZAddMember 是集合（set）中的元素结构体。. md5:eb7d172c444324d7
# <翻译结束>


<原文开始>
// ZRangeOption provides extra option for ZRange function.
<原文结束>

# <翻译开始>
// ZRangeOption为ZRange函数提供额外选项。. md5:61532d16fe5a1260
# <翻译结束>


<原文开始>
	// The optional REV argument reverses the ordering, so elements are ordered from highest to lowest score,
	// and score ties are resolved by reverse lexicographical ordering.
<原文结束>

# <翻译开始>
// 可选的REV参数会反转顺序，因此元素按照分数从高到低排序，分数相同时则通过反向字典序进行解析。
// md5:a1c79d75cedbff1b
# <翻译结束>


<原文开始>
// The optional WithScores argument supplements the command's reply with the scores of elements returned.
<原文结束>

# <翻译开始>
// 可选的 WithScores 参数会在命令回复中补充返回元素的分数。. md5:26db0341550d511b
# <翻译结束>


<原文开始>
// ZRangeOptionLimit provides LIMIT argument for ZRange function.
// The optional LIMIT argument can be used to obtain a sub-range from the matching elements
// (similar to SELECT LIMIT offset, count in SQL). A negative `Count` returns all elements from the `Offset`.
<原文结束>

# <翻译开始>
// ZRangeOptionLimit 为 ZRange 函数提供 LIMIT 参数。
// 可选的 LIMIT 参数可用于从匹配的元素中获取子范围（类似于 SQL 中的 SELECT LIMIT 偏移量, 数量）。当 `Count` 为负数时，从 `Offset` 开始返回所有元素。
// md5:a910bb82b51914ef
# <翻译结束>


<原文开始>
// ZRevRangeOption provides options for function ZRevRange.
<原文结束>

# <翻译开始>
// ZRevRangeOption为ZRevRange函数提供了选项。. md5:cd0b627793d48f50
# <翻译结束>

