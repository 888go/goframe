
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
// 包 gqueue 提供动态/静态并发安全队列。
//
// 功能：
//
// 1.先进先出队列（数据 -> 列表 -> 通道）；
//
// 2.快速创建和初始化；
//
// 3.支持动态队列大小（无限制的队列大小）；
//
// 4.从队列中读取数据时会阻塞。
// md5:ff40490071065bb6
# <翻译结束>


<原文开始>
// Queue is a concurrent-safe queue built on doubly linked list and channel.
<原文结束>

# <翻译开始>
// Queue是一个基于双向链表和通道的并发安全队列。. md5:dc3dd26386e4acfb
# <翻译结束>


<原文开始>
// Underlying list structure for data maintaining.
<原文结束>

# <翻译开始>
// 用于数据维护的基础列表结构。. md5:c41c31abb5b3b9e9
# <翻译结束>


<原文开始>
// Whether queue is closed.
<原文结束>

# <翻译开始>
// 队列是否关闭。. md5:e327adc6a3cf2327
# <翻译结束>


<原文开始>
// Events for data writing.
<原文结束>

# <翻译开始>
// 数据写入事件。. md5:e179f70066a0a70c
# <翻译结束>


<原文开始>
// Underlying channel for data reading.
<原文结束>

# <翻译开始>
// 数据读取的底层通道。. md5:1bf7b6e23c35ba5f
# <翻译结束>


<原文开始>
// Max batch size per-fetching from list.
<原文结束>

# <翻译开始>
// 从列表中每次预取的最大批处理大小。. md5:d8aca34be43fb879
# <翻译结束>


<原文开始>
// New returns an empty queue object.
// Optional parameter `limit` is used to limit the size of the queue, which is unlimited in default.
// When `limit` is given, the queue will be static and high performance which is comparable with stdlib channel.
<原文结束>

# <翻译开始>
// New 返回一个空的队列对象。
// 可选参数 `limit` 用于限制队列的大小，默认情况下无限制。
// 当提供 `limit` 时，队列将变为静态且高性能，其性能可与标准库中的通道相媲美。
// md5:9fbd45b8d84f665e
# <翻译结束>


<原文开始>
// Push pushes the data `v` into the queue.
// Note that it would panic if Push is called after the queue is closed.
<原文结束>

# <翻译开始>
// Push 将数据 `v` 推入队列。
// 注意，如果在关闭队列后调用 Push，它将引发 panic。
// md5:ace317b42ed78776
# <翻译结束>


<原文开始>
// Pop pops an item from the queue in FIFO way.
// Note that it would return nil immediately if Pop is called after the queue is closed.
<原文结束>

# <翻译开始>
// Pop 从队列中按先进先出（FIFO）方式弹出一个项目。
// 如果在关闭队列后调用 Pop，它会立即返回 nil。
// md5:f632ecf6d87ed4c5
# <翻译结束>


<原文开始>
// Close closes the queue.
// Notice: It would notify all goroutines return immediately,
// which are being blocked reading using Pop method.
<原文结束>

# <翻译开始>
// Close 关闭队列。
// 注意：它会通知所有因调用Pop方法而阻塞的goroutine立即返回。
// md5:bd22bcaaebaed5dc
# <翻译结束>


<原文开始>
// Len returns the length of the queue.
// Note that the result might not be accurate if using unlimited queue size as there's an
// asynchronous channel reading the list constantly.
<原文结束>

# <翻译开始>
// Len 返回队列的长度。
// 请注意，如果使用无限大的队列大小，结果可能不准确，因为有一个异步通道持续读取列表。
// md5:b2b860a611742a51
# <翻译结束>


<原文开始>
// Size is alias of Len.
// Deprecated: use Len instead.
<原文结束>

# <翻译开始>
// Size是Len的别名。
// 警告：请改用Len。
// md5:25acbbc5f8f37a14
# <翻译结束>


<原文开始>
// asyncLoopFromListToChannel starts an asynchronous goroutine,
// which handles the data synchronization from list `q.list` to channel `q.C`.
<原文结束>

# <翻译开始>
// asyncLoopFromListToChannel 启动一个异步goroutine，
// 它负责从列表`q.list`到通道`q.C`的数据同步处理。
// md5:fd4f8b385cd5a6ba
# <翻译结束>


<原文开始>
				// When q.C is closed, it will panic here, especially q.C is being blocked for writing.
				// If any error occurs here, it will be caught by recover and be ignored.
<原文结束>

# <翻译开始>
// 当q.C被关闭时，这里将会发生恐慌，尤其是当q.C因写入操作而被阻塞时。
// 如果这里发生任何错误，它将被recover捕获并被忽略。
// md5:eaf48f57d3e8e5be
# <翻译结束>


<原文开始>
// Clear q.events to remain just one event to do the next synchronization check.
<原文结束>

# <翻译开始>
// 清除q.events，只保留一个事件以进行下次同步检查。. md5:925f0acc845d8b6d
# <翻译结束>


<原文开始>
	// It should be here to close `q.C` if `q` is unlimited size.
	// It's the sender's responsibility to close channel when it should be closed.
<原文结束>

# <翻译开始>
// 如果队列 `q` 的大小是无限的，它应该在这里关闭 `q.C`。
// 当需要关闭通道时，发送者有责任关闭通道。
// md5:bd37819839de5b3c
# <翻译结束>

