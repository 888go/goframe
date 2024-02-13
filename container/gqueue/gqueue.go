// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gqueue 提供了动态/静态的并发安全队列。
//
// 特性：
//
// 1. 先进先出（FIFO）队列（数据 -> 链表 -> 通道）；
//
// 2. 快速创建和初始化；
//
// 3. 支持动态队列大小（无限制的队列大小）；
//
// 4. 从队列读取数据时会阻塞等待。
package 队列类

import (
	"math"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/container/gtype"
)

// Queue 是一个基于双链表和通道实现的并发安全队列。
type Queue struct {
	limit  int              // 队列大小的限制。
	list   *链表类.List      // 用于维护数据的基础列表结构
	closed *安全变量类.Bool      // 队列是否已关闭。
	events chan struct{}    // 数据写入事件
	C      chan interface{} // 数据读取的基础通道。
}

const (
	defaultQueueSize = 10000 // 队列缓冲区的大小。
	defaultBatchSize = 10    // 从列表中每次获取的最大批量大小
)

// New 函数返回一个空的队列对象。
// 可选参数 `limit` 用于限制队列的大小，默认情况下不限制大小。
// 当提供了 `limit` 参数时，队列将会是静态且高性能的，其性能可与标准库中的 channel 相媲美。
func X创建(队列长度 ...int) *Queue {
	q := &Queue{
		closed: 安全变量类.NewBool(),
	}
	if len(队列长度) > 0 && 队列长度[0] > 0 {
		q.limit = 队列长度[0]
		q.C = make(chan interface{}, 队列长度[0])
	} else {
		q.list = 链表类.New(true)
		q.events = make(chan struct{}, math.MaxInt32)
		q.C = make(chan interface{}, defaultQueueSize)
		go q.asyncLoopFromListToChannel()
	}
	return q
}

// Push 将数据 `v` 推入队列中。
// 注意，如果在队列关闭后调用 Push，将会导致程序panic。
func (q *Queue) X入栈(值 interface{}) {
	if q.limit > 0 {
		q.C <- 值
	} else {
		q.list.PushBack(值)
		if len(q.events) < defaultQueueSize {
			q.events <- struct{}{}
		}
	}
}

// Pop 按先进先出（FIFO）的方式从队列中弹出一个元素。
// 注意，如果在队列关闭后调用 Pop 方法，它会立即返回空值（nil）。
func (q *Queue) X出栈() interface{} {
	return <-q.C
}

// Close 关闭队列。
// 注意：此操作会立即通知所有因调用 Pop 方法而被阻塞的 goroutine 立即返回。
func (q *Queue) X关闭() {
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
			q.X出栈()
		}
	}
}

// Len 返回队列的长度。
// 注意，如果使用无限大小的队列，结果可能不准确，因为存在一个异步通道一直在不断地读取该列表。
func (q *Queue) X取长度() (长度 int64) {
	bufferedSize := int64(len(q.C))
	if q.limit > 0 {
		return bufferedSize
	}
	return int64(q.list.Size()) + bufferedSize
}

// Size 是 Len 的别名。
// 已弃用：请改用 Len。
func (q *Queue) Size弃用() int64 {
	return q.X取长度()
}

// asyncLoopFromListToChannel 启动一个异步 goroutine，
// 该 goroutine 负责从列表 `q.list` 到通道 `q.C` 的数据同步处理。
func (q *Queue) asyncLoopFromListToChannel() {
	defer func() {
		if q.closed.X取值() {
			_ = recover()
		}
	}()
	for !q.closed.X取值() {
		<-q.events
		for !q.closed.X取值() {
			if bufferLength := q.list.Len(); bufferLength > 0 {
// 当q.C被关闭时，此处将会引发panic，特别是当q.C正被阻塞等待写入时。
// 如果此处发生任何错误，将被recover捕获并忽略。
				for i := 0; i < bufferLength; i++ {
					q.C <- q.list.PopFront()
				}
			} else {
				break
			}
		}
		// 清除q.events，仅保留一个事件用于执行下一次同步检查。
		for i := 0; i < len(q.events)-1; i++ {
			<-q.events
		}
	}
// 如果`q`的大小是无限的，那么关闭`q.C`的操作应该在这里进行。
// 关闭通道的责任在于发送者，在适当的时候应由发送者关闭通道。
	close(q.C)
}
