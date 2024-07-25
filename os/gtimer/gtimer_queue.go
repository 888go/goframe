// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gtimer

import (
	"container/heap"
	"math"
	"sync"

	"github.com/gogf/gf/v2/container/gtype"
)

// priorityQueue 是一种抽象数据类型，类似于常规的队列或堆数据结构，但每个元素都附加了一个“优先级”。
// 在优先队列中，高优先级的元素会在低优先级元素之前被处理。
// priorityQueue 基于堆结构实现。 md5:83d2128c4f6a7192
type priorityQueue struct {
	mu           sync.Mutex
	heap         *priorityQueueHeap // 使用堆实现的底层队列项目管理器。 md5:e4cb05ba6e4f562a
	nextPriority *gtype.Int64       // nextPriority 存储堆的下一个优先级值，用于检查是否需要通过Timer调用堆的Pop方法。 md5:c0572df9a2e3fa9b
}

// priorityQueueHeap 是一个堆管理器，其底层的 `array` 是一个实现堆结构的数组。 md5:6399593cd56ce49b
type priorityQueueHeap struct {
	array []priorityQueueItem
}

// priorityQueueItem 用于存储具有优先级属性的队列项，以便在堆中进行排序。 md5:78016bdeacd4ec5d
type priorityQueueItem struct {
	value    interface{}
	priority int64
}

// newPriorityQueue 创建并返回一个优先级队列。 md5:e7b97ee41a869ee7
func newPriorityQueue() *priorityQueue {
	queue := &priorityQueue{
		heap:         &priorityQueueHeap{array: make([]priorityQueueItem, 0)},
		nextPriority: gtype.NewInt64(math.MaxInt64),
	}
	heap.Init(queue.heap)
	return queue
}

// NextPriority 从队列中获取并返回最小优先级和最高优先级的值。 md5:deb4631876016380
func (q *priorityQueue) NextPriority() int64 {
	return q.nextPriority.Val()
}

// Push 将一个值推入队列。
// `priority` 参数指定了该值的优先级。
// `priority` 的值越小，表示 `value` 的优先级越高。 md5:cf9acf4068187c77
func (q *priorityQueue) Push(value interface{}, priority int64) {
	q.mu.Lock()
	defer q.mu.Unlock()
	heap.Push(q.heap, priorityQueueItem{
		value:    value,
		priority: priority,
	})
	// 使用原子操作更新最小优先级。 md5:c7c29d16bf8470d3
	nextPriority := q.nextPriority.Val()
	if priority >= nextPriority {
		return
	}
	q.nextPriority.Set(priority)
}

// Pop 从队列中取出、移除并返回最高优先级的值。 md5:828fb8c6fde3e6a4
func (q *priorityQueue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if v := heap.Pop(q.heap); v != nil {
		var nextPriority int64 = math.MaxInt64
		if len(q.heap.array) > 0 {
			nextPriority = q.heap.array[0].priority
		}
		q.nextPriority.Set(nextPriority)
		return v.(priorityQueueItem).value
	}
	return nil
}
