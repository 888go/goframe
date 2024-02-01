
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
// IGroupSortedSet manages redis sorted set operations.
// Implements see redis.GroupSortedSet.
<原文结束>

# <翻译开始>
// IGroupSortedSet 管理 Redis 有序集合操作。
// 实现参考 redis.GroupSortedSet。
# <翻译结束>


<原文开始>
// ZAddOption provides options for function ZAdd.
<原文结束>

# <翻译开始>
// ZAddOption 提供了函数 ZAdd 的选项。
# <翻译结束>


<原文开始>
	// Only update existing elements if the new score is less than the current score.
	// This flag doesn't prevent adding new elements.
<原文结束>

# <翻译开始>
// 只有当新分数小于当前分数时，才更新已存在的元素。
// 但请注意，此标志不会阻止添加新的元素。
# <翻译结束>


<原文开始>
	// Only update existing elements if the new score is greater than the current score.
	// This flag doesn't prevent adding new elements.
<原文结束>

# <翻译开始>
// 如果新分数大于当前分数，则仅更新现有元素。
// 此标志不会阻止添加新元素。
# <翻译结束>


<原文开始>
	// Modify the return value from the number of new elements added, to the total number of elements changed (CH is an abbreviation of changed).
	// Changed elements are new elements added and elements already existing for which the score was updated.
	// So elements specified in the command line having the same score as they had in the past are not counted.
	// Note: normally the return value of ZAdd only counts the number of new elements added.
<原文结束>

# <翻译开始>
// 将返回值由新增元素的数量修改为更改过的总元素数量（CH 是“changed”的缩写）。
// 更改过的元素包括新添加的元素以及已存在但分数被更新的元素。
// 因此，命令行中指定且其分数与过去相同的元素不会被计算在内。
// 注意：通常情况下，ZAdd 的返回值仅计算新增元素的数量。
# <翻译结束>


<原文开始>
// When this option is specified ZAdd acts like ZIncrBy. Only one score-element pair can be specified in this mode.
<原文结束>

# <翻译开始>
// 当指定了此选项时，ZAdd 表现得如同 ZIncrBy。在这种模式下，只能指定一个分数-元素对。
# <翻译结束>


<原文开始>
// ZAddMember is element struct for set.
<原文结束>

# <翻译开始>
// ZAddMember 是集合中元素的结构体。
# <翻译结束>


<原文开始>
// ZRangeOption provides extra option for ZRange function.
<原文结束>

# <翻译开始>
// ZRangeOption 为 ZRange 函数提供了额外的选项。
# <翻译结束>


<原文开始>
	// The optional REV argument reverses the ordering, so elements are ordered from highest to lowest score,
	// and score ties are resolved by reverse lexicographical ordering.
<原文结束>

# <翻译开始>
// 可选参数 REV 用于反转排序顺序，因此元素按从高到低的分数进行排序，
// 当分数相同时，采用反字典序进行排序结果的确定。
# <翻译结束>


<原文开始>
// The optional WithScores argument supplements the command's reply with the scores of elements returned.
<原文结束>

# <翻译开始>
// 可选的 WithScores 参数会用返回元素的分数来补充命令的回复。
# <翻译结束>


<原文开始>
// ZRangeOptionLimit provides LIMIT argument for ZRange function.
// The optional LIMIT argument can be used to obtain a sub-range from the matching elements
// (similar to SELECT LIMIT offset, count in SQL). A negative `Count` returns all elements from the `Offset`.
<原文结束>

# <翻译开始>
// ZRangeOptionLimit 为 ZRange 函数提供 LIMIT 参数。
// 可选的 LIMIT 参数可用于从匹配元素中获取一个子范围（类似于 SQL 中的 SELECT LIMIT offset, count）。
// 如果 `Count` 为负数，则返回从 `Offset` 开始的所有元素。
# <翻译结束>


<原文开始>
// ZRevRangeOption provides options for function ZRevRange.
<原文结束>

# <翻译开始>
// ZRevRangeOption 提供了 ZRevRange 函数的选项。
# <翻译结束>

