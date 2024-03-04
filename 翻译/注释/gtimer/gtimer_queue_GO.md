
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
// priorityQueue is an abstract data type similar to a regular queue or stack data structure in which
// each element additionally has a "priority" associated with it. In a priority queue, an element with
// high priority is served before an element with low priority.
// priorityQueue is based on heap structure.
<原文结束>

# <翻译开始>
// priorityQueue 是一种类似于常规队列或堆栈数据结构的抽象数据类型，
// 其中每个元素都额外关联有一个“优先级”。在优先队列中，具有高优先级的元素会优先于低优先级的元素被处理。
// priorityQueue 是基于堆结构实现的。
# <翻译结束>


<原文开始>
// the underlying queue items manager using heap.
<原文结束>

# <翻译开始>
// 使用堆实现的基础队列项管理器。
# <翻译结束>


<原文开始>
// nextPriority stores the next priority value of the heap, which is used to check if necessary to call the Pop of heap by Timer.
<原文结束>

# <翻译开始>
// nextPriority 存储堆（heap）的下一个优先级值，该值用于通过 Timer 检查是否有必要调用堆的 Pop 方法。
# <翻译结束>


<原文开始>
// priorityQueueHeap is a heap manager, of which the underlying `array` is an array implementing a heap structure.
<原文结束>

# <翻译开始>
// priorityQueueHeap 是一个堆管理器，其底层的 `array` 是一个实现了堆结构的数组。
# <翻译结束>


<原文开始>
// priorityQueueItem stores the queue item which has a `priority` attribute to sort itself in heap.
<原文结束>

# <翻译开始>
// priorityQueueItem 用于存储队列项，其中包含一个 `priority` 属性，以便在堆中进行自我排序。
# <翻译结束>


<原文开始>
// newPriorityQueue creates and returns a priority queue.
<原文结束>

# <翻译开始>
// newPriorityQueue 创建并返回一个优先队列。
# <翻译结束>


<原文开始>
// NextPriority retrieves and returns the minimum and the most priority value of the queue.
<原文结束>

# <翻译开始>
// NextPriority 获取并返回队列中的最小且优先级最高的值。
# <翻译结束>


<原文开始>
// Push pushes a value to the queue.
// The `priority` specifies the priority of the value.
// The lesser the `priority` value the higher priority of the `value`.
<原文结束>

# <翻译开始>
// Push 将一个值推送到队列中。
// `priority` 参数用于指定该值的优先级。
// `priority` 值越小，`value` 的优先级越高。
# <翻译结束>


<原文开始>
// Update the minimum priority using atomic operation.
<原文结束>

# <翻译开始>
// 使用原子操作更新最小优先级。
# <翻译结束>


<原文开始>
// Pop retrieves, removes and returns the most high priority value from the queue.
<原文结束>

# <翻译开始>
// Pop从队列中获取、移除并返回优先级最高的值。
# <翻译结束>

