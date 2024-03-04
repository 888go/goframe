
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
// Package gqueue provides dynamic/static concurrent-safe queue.
//
// Features:
//
// 1. FIFO queue(data -> list -> chan);
//
// 2. Fast creation and initialization;
//
// 3. Support dynamic queue size(unlimited queue size);
//
// 4. Blocking when reading data from queue;
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// Queue is a concurrent-safe queue built on doubly linked list and channel.
<原文结束>

# <翻译开始>
// Queue 是一个基于双链表和通道实现的并发安全队列。
# <翻译结束>







<原文开始>
// Underlying list structure for data maintaining.
<原文结束>

# <翻译开始>
// 用于维护数据的基础列表结构
# <翻译结束>












<原文开始>
// Underlying channel for data reading.
<原文结束>

# <翻译开始>
// 数据读取的基础通道。
# <翻译结束>







<原文开始>
// Max batch size per-fetching from list.
<原文结束>

# <翻译开始>
// 从列表中每次获取的最大批量大小
# <翻译结束>


<原文开始>
// New returns an empty queue object.
// Optional parameter `limit` is used to limit the size of the queue, which is unlimited in default.
// When `limit` is given, the queue will be static and high performance which is comparable with stdlib channel.
<原文结束>

# <翻译开始>
// New 函数返回一个空的队列对象。
// 可选参数 `limit` 用于限制队列的大小，默认情况下不限制大小。
// 当提供了 `limit` 参数时，队列将会是静态且高性能的，其性能可与标准库中的 channel 相媲美。
# <翻译结束>


<原文开始>
// Push pushes the data `v` into the queue.
// Note that it would panic if Push is called after the queue is closed.
<原文结束>

# <翻译开始>
// Push 将数据 `v` 推入队列中。
// 注意，如果在队列关闭后调用 Push，将会导致程序panic。
# <翻译结束>


<原文开始>
// Pop pops an item from the queue in FIFO way.
// Note that it would return nil immediately if Pop is called after the queue is closed.
<原文结束>

# <翻译开始>
// Pop 按先进先出（FIFO）的方式从队列中弹出一个元素。
// 注意，如果在队列关闭后调用 Pop 方法，它会立即返回空值（nil）。
# <翻译结束>


<原文开始>
// Close closes the queue.
// Notice: It would notify all goroutines return immediately,
// which are being blocked reading using Pop method.
<原文结束>

# <翻译开始>
// Close 关闭队列。
// 注意：此操作会立即通知所有因调用 Pop 方法而被阻塞的 goroutine 立即返回。
# <翻译结束>


<原文开始>
// Len returns the length of the queue.
// Note that the result might not be accurate if using unlimited queue size as there's an
// asynchronous channel reading the list constantly.
<原文结束>

# <翻译开始>
// Len 返回队列的长度。
// 注意，如果使用无限大小的队列，结果可能不准确，因为存在一个异步通道一直在不断地读取该列表。
# <翻译结束>


<原文开始>
// Size is alias of Len.
// Deprecated: use Len instead.
<原文结束>

# <翻译开始>
// Size 是 Len 的别名。
// 已弃用：请改用 Len。
# <翻译结束>


<原文开始>
// asyncLoopFromListToChannel starts an asynchronous goroutine,
// which handles the data synchronization from list `q.list` to channel `q.C`.
<原文结束>

# <翻译开始>
// asyncLoopFromListToChannel 启动一个异步 goroutine，
// 该 goroutine 负责从列表 `q.list` 到通道 `q.C` 的数据同步处理。
# <翻译结束>


<原文开始>
				// When q.C is closed, it will panic here, especially q.C is being blocked for writing.
				// If any error occurs here, it will be caught by recover and be ignored.
<原文结束>

# <翻译开始>
// 当q.C被关闭时，此处将会引发panic，特别是当q.C正被阻塞等待写入时。
// 如果此处发生任何错误，将被recover捕获并忽略。
# <翻译结束>


<原文开始>
// Clear q.events to remain just one event to do the next synchronization check.
<原文结束>

# <翻译开始>
// 清除q.events，仅保留一个事件用于执行下一次同步检查。
# <翻译结束>


<原文开始>
	// It should be here to close `q.C` if `q` is unlimited size.
	// It's the sender's responsibility to close channel when it should be closed.
<原文结束>

# <翻译开始>
// 如果`q`的大小是无限的，那么关闭`q.C`的操作应该在这里进行。
// 关闭通道的责任在于发送者，在适当的时候应由发送者关闭通道。
# <翻译结束>


<原文开始>
// Limit for queue size.
<原文结束>

# <翻译开始>
// 队列大小的限制。
# <翻译结束>


<原文开始>
// Whether queue is closed.
<原文结束>

# <翻译开始>
// 队列是否已关闭。
# <翻译结束>


<原文开始>
// Events for data writing.
<原文结束>

# <翻译开始>
// 数据写入事件
# <翻译结束>


<原文开始>
// Size for queue buffer.
<原文结束>

# <翻译开始>
// 队列缓冲区的大小。
# <翻译结束>

