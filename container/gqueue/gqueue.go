// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package gqueue provides dynamic/static concurrent-safe queue.
//
//
// 1. FIFO queue(data -> list -> chan);
//
// 2. Fast creation and initialization;
//
// 3. Support dynamic queue size(unlimited queue size);
//
// 4. Blocking when reading data from queue;
package gqueue//bm:队列类

import (
	"math"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/container/gtype"
)

// Queue是一个基于双向链表和通道的并发安全队列。 md5:dc3dd26386e4acfb
type Queue struct {
	limit  int              // Limit for queue size.
	list   *glist.List      // 用于数据维护的基础列表结构。 md5:c41c31abb5b3b9e9
	closed *gtype.Bool      // 队列是否关闭。 md5:e327adc6a3cf2327
	events chan struct{}    // 数据写入事件。 md5:e179f70066a0a70c
	C      chan interface{} // 数据读取的底层通道。 md5:1bf7b6e23c35ba5f
}

const (
	defaultQueueSize = 10000 // Size for queue buffer.
	defaultBatchSize = 10    // 从列表中每次预取的最大批处理大小。 md5:d8aca34be43fb879
)

// New 返回一个空的队列对象。
// 可选参数 `limit` 用于限制队列的大小，默认情况下无限制。
// 当提供 `limit` 时，队列将变为静态且高性能，其性能可与标准库中的通道相媲美。
// md5:9fbd45b8d84f665e
// ff:创建
// limit:队列长度
func New(limit ...int) *Queue {
	q := &Queue{
		closed: gtype.NewBool(),
	}
	if len(limit) > 0 && limit[0] > 0 {
		q.limit = limit[0]
		q.C = make(chan interface{}, limit[0])
	} else {
		q.list = glist.New(true)
		q.events = make(chan struct{}, math.MaxInt32)
		q.C = make(chan interface{}, defaultQueueSize)
		go q.asyncLoopFromListToChannel()
	}
	return q
}

// Push 将数据 `v` 推入队列。
// 注意，如果在关闭队列后调用 Push，它将引发 panic。
// md5:ace317b42ed78776
// ff:入栈
// q:
// v:值
func (q *Queue) Push(v interface{}) {
	if q.limit > 0 {
		q.C <- v
	} else {
		q.list.PushBack(v)
		if len(q.events) < defaultQueueSize {
			q.events <- struct{}{}
		}
	}
}

// Pop 从队列中按先进先出（FIFO）方式弹出一个项目。
// 如果在关闭队列后调用 Pop，它会立即返回 nil。
// md5:f632ecf6d87ed4c5
// ff:出栈
// q:
func (q *Queue) Pop() interface{} {
	return <-q.C
}

// Close closes the queue.
// which are being blocked reading using Pop method.
// ff:关闭
// q:
func (q *Queue) Close() {
	if !q.closed.Cas(false, true) {
		return
	}
	if q.events != nil {
		close(q.events)
	}
	if q.limit > 0 {
		close(q.C)
	} else {
		for i := 0; i < defaultBatchSize; i++ {
			q.Pop()
		}
	}
}

// Len 返回队列的长度。
// 请注意，如果使用无限大的队列大小，结果可能不准确，因为有一个异步通道持续读取列表。
// md5:b2b860a611742a51
// ff:取长度
// q:
// length:长度
func (q *Queue) Len() (length int64) {
	bufferedSize := int64(len(q.C))
	if q.limit > 0 {
		return bufferedSize
	}
	return int64(q.list.Size()) + bufferedSize
}

// Size is alias of Len.
// ff:Size弃用
// q:
func (q *Queue) Size() int64 {
	return q.Len()
}

// asyncLoopFromListToChannel 启动一个异步goroutine，
// 它负责从列表`q.list`到通道`q.C`的数据同步处理。
// md5:fd4f8b385cd5a6ba
func (q *Queue) asyncLoopFromListToChannel() {
	defer func() {
		if q.closed.Val() {
			_ = recover()
		}
	}()
	for !q.closed.Val() {
		<-q.events
		for !q.closed.Val() {
			if bufferLength := q.list.Len(); bufferLength > 0 {
// 当q.C被关闭时，这里将会发生恐慌，尤其是当q.C因写入操作而被阻塞时。
// 如果这里发生任何错误，它将被recover捕获并被忽略。
// md5:eaf48f57d3e8e5be
				for i := 0; i < bufferLength; i++ {
					q.C <- q.list.PopFront()
				}
			} else {
				break
			}
		}
		// 清除q.events，只保留一个事件以进行下次同步检查。 md5:925f0acc845d8b6d
		for i := 0; i < len(q.events)-1; i++ {
			<-q.events
		}
	}
// 如果队列 `q` 的大小是无限的，它应该在这里关闭 `q.C`。
// 当需要关闭通道时，发送者有责任关闭通道。
// md5:bd37819839de5b3c
	close(q.C)
}
