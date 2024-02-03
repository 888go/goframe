// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtimer

import (
	"container/heap"
	"math"
	"sync"
	
	"github.com/888go/goframe/container/gtype"
)

// priorityQueue 是一种类似于常规队列或堆栈数据结构的抽象数据类型，
// 其中每个元素都额外关联有一个“优先级”。在优先队列中，具有高优先级的元素会优先于低优先级的元素被处理。
// priorityQueue 是基于堆结构实现的。
type priorityQueue struct {
	mu           sync.Mutex
	heap         *priorityQueueHeap // 使用堆实现的基础队列项管理器。
	nextPriority *gtype.Int64       // nextPriority 存储堆（heap）的下一个优先级值，该值用于通过 Timer 检查是否有必要调用堆的 Pop 方法。
}

// priorityQueueHeap 是一个堆管理器，其底层的 `array` 是一个实现了堆结构的数组。
type priorityQueueHeap struct {
	array []priorityQueueItem
}

// priorityQueueItem 用于存储队列项，其中包含一个 `priority` 属性，以便在堆中进行自我排序。
type priorityQueueItem struct {
	value    interface{}
	priority int64
}

// newPriorityQueue 创建并返回一个优先队列。
func newPriorityQueue() *priorityQueue {
	queue := &priorityQueue{
		heap:         &priorityQueueHeap{array: make([]priorityQueueItem, 0)},
		nextPriority: gtype.NewInt64(math.MaxInt64),
	}
	heap.Init(queue.heap)
	return queue
}

// NextPriority 获取并返回队列中的最小且优先级最高的值。
func (q *priorityQueue) NextPriority() int64 {
	return q.nextPriority.Val()
}

// Push 将一个值推送到队列中。
// `priority` 参数用于指定该值的优先级。
// `priority` 值越小，`value` 的优先级越高。
func (q *priorityQueue) Push(value interface{}, priority int64) {
	q.mu.Lock()
	defer q.mu.Unlock()
	heap.Push(q.heap, priorityQueueItem{
		value:    value,
		priority: priority,
	})
	// 使用原子操作更新最小优先级。
	nextPriority := q.nextPriority.Val()
	if priority >= nextPriority {
		return
	}
	q.nextPriority.Set(priority)
}

// Pop从队列中获取、移除并返回优先级最高的值。
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
