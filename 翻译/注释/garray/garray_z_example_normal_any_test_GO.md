
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
// Insert item before specified index.
<原文结束>

# <翻译开始>
// 在指定索引之前插入项。
# <翻译结束>


<原文开始>
// Insert item after specified index.
<原文结束>

# <翻译开始>
// 在指定索引之后插入项目。
# <翻译结束>







<原文开始>
// Search item and return its index.
<原文结束>

# <翻译开始>
// 搜索指定项并返回其索引。
# <翻译结束>







<原文开始>
// Empty the array, removes all items of it.
<原文结束>

# <翻译开始>
// 清空切片，移除其所有元素。
# <翻译结束>


<原文开始>
	// Iterator is alias of IteratorAsc, which iterates the array readonly in ascending order
	//  with given callback function `f`.
	// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// Iterator 是 IteratorAsc 的别名，用于以升序方式对切片进行只读遍历，
// 同时调用给定的回调函数 `f`。
// 若 `f` 返回 true，则继续遍历；若返回 false，则停止遍历。
# <翻译结束>


<原文开始>
	// IteratorDesc iterates the array readonly in descending order with given callback function `f`.
	// If `f` returns true, then it continues iterating; or false to stop.
<原文结束>

# <翻译开始>
// IteratorDesc 以降序方式遍历给定回调函数 `f` 的只读切片。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
// 这段Go语言代码注释翻译成中文注释如下：
// ```go
// IteratorDesc 函数以降序顺序对给定的只读切片进行迭代，并使用指定的回调函数 `f` 进行处理。
// 若回调函数 `f` 返回值为 true，则会继续进行迭代；若返回值为 false，则停止迭代过程。
# <翻译结束>


<原文开始>
// Reverse makes array with elements in reverse order.
<原文结束>

# <翻译开始>
// Reverse将切片元素按逆序排列。
# <翻译结束>


<原文开始>
// Shuffle randomly shuffles the array.
<原文结束>

# <翻译开始>
// Shuffle 随机地对切片进行洗牌。
# <翻译结束>


<原文开始>
	// Randomly retrieve and return 2 items from the array.
	// It does not delete the items from array.
<原文结束>

# <翻译开始>
// 随机从切片中获取并返回 2 个元素。
// 不会从切片中删除这些元素。
# <翻译结束>


<原文开始>
	// Randomly pick and return one item from the array.
	// It deletes the picked up item from array.
<原文结束>

# <翻译开始>
// 从切片中随机选取并返回一个元素。
// 它会从切片中删除已选取的元素。
# <翻译结束>


<原文开始>
	// Chunk splits an array into multiple arrays,
	// the size of each array is determined by `size`.
	// The last chunk may contain less than size elements.
<原文结束>

# <翻译开始>
// Chunk 函数将一个切片分割成多个子切片，
// 每个子切片的大小由参数 `size` 确定。
// 最后一个子切片可能包含少于 size 个元素。
# <翻译结束>


<原文开始>
// Any Pop* functions pick, delete and return the item from array.
<原文结束>

# <翻译开始>
// 任何 Pop* 函数都会从切片中挑选、删除并返回一个元素。
# <翻译结束>


<原文开始>
// Print the array length.
<原文结束>

# <翻译开始>
// 打印切片长度。
# <翻译结束>


<原文开始>
// Print the array items.
<原文结束>

# <翻译开始>
// 打印切片元素。
# <翻译结束>


<原文开始>
// Retrieve item by index.
<原文结束>

# <翻译开始>
// 通过索引获取项。
# <翻译结束>


<原文开始>
// Check item existence.
<原文结束>

# <翻译开始>
// 检查项目是否存在。
# <翻译结束>


<原文开始>
// Modify item by index.
<原文结束>

# <翻译开始>
// 通过索引修改项。
# <翻译结束>


<原文开始>
// Remove item by index.
<原文结束>

# <翻译开始>
// 通过索引移除项
# <翻译结束>

