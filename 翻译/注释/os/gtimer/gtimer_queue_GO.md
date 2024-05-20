
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
// priorityQueue is an abstract data type similar to a regular queue or stack data structure in which
// each element additionally has a "priority" associated with it. In a priority queue, an element with
// high priority is served before an element with low priority.
// priorityQueue is based on heap structure.
<原文结束>

# <翻译开始>
// priorityQueue 是一种抽象数据类型，类似于常规的队列或堆数据结构，但每个元素都附加了一个“优先级”。
// 在优先队列中，高优先级的元素会在低优先级元素之前被处理。
// priorityQueue 基于堆结构实现。
// md5:83d2128c4f6a7192
# <翻译结束>


<原文开始>
// the underlying queue items manager using heap.
<原文结束>

# <翻译开始>
// 使用堆实现的底层队列项目管理器。. md5:e4cb05ba6e4f562a
# <翻译结束>


<原文开始>
// nextPriority stores the next priority value of the heap, which is used to check if necessary to call the Pop of heap by Timer.
<原文结束>

# <翻译开始>
// nextPriority 存储堆的下一个优先级值，用于检查是否需要通过Timer调用堆的Pop方法。. md5:c0572df9a2e3fa9b
# <翻译结束>


<原文开始>
// priorityQueueHeap is a heap manager, of which the underlying `array` is an array implementing a heap structure.
<原文结束>

# <翻译开始>
// priorityQueueHeap 是一个堆管理器，其底层的 `array` 是一个实现堆结构的数组。. md5:6399593cd56ce49b
# <翻译结束>


<原文开始>
// priorityQueueItem stores the queue item which has a `priority` attribute to sort itself in heap.
<原文结束>

# <翻译开始>
// priorityQueueItem 用于存储具有优先级属性的队列项，以便在堆中进行排序。. md5:78016bdeacd4ec5d
# <翻译结束>


<原文开始>
// newPriorityQueue creates and returns a priority queue.
<原文结束>

# <翻译开始>
// newPriorityQueue 创建并返回一个优先级队列。. md5:e7b97ee41a869ee7
# <翻译结束>


<原文开始>
// NextPriority retrieves and returns the minimum and the most priority value of the queue.
<原文结束>

# <翻译开始>
// NextPriority 从队列中获取并返回最小优先级和最高优先级的值。. md5:deb4631876016380
# <翻译结束>


<原文开始>
// Push pushes a value to the queue.
// The `priority` specifies the priority of the value.
// The lesser the `priority` value the higher priority of the `value`.
<原文结束>

# <翻译开始>
// Push 将一个值推入队列。
// `priority` 参数指定了该值的优先级。
// `priority` 的值越小，表示 `value` 的优先级越高。
// md5:cf9acf4068187c77
# <翻译结束>


<原文开始>
// Update the minimum priority using atomic operation.
<原文结束>

# <翻译开始>
// 使用原子操作更新最小优先级。. md5:c7c29d16bf8470d3
# <翻译结束>


<原文开始>
// Pop retrieves, removes and returns the most high priority value from the queue.
<原文结束>

# <翻译开始>
// Pop 从队列中取出、移除并返回最高优先级的值。. md5:828fb8c6fde3e6a4
# <翻译结束>

