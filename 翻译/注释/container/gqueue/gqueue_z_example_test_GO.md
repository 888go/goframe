
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Close the queue in three seconds.
<原文结束>

# <翻译开始>
	// 三秒后关闭队列。 md5:02742be1c1ceef32
# <翻译结束>


<原文开始>
	// The consumer constantly reads the queue data.
	// If there is no data in the queue, it will block.
	// The queue is read using the queue.C property exposed
	// by the queue object and the selectIO multiplexing syntax
	// example:
	// for {
	//    select {
	//        case v := <-queue.C:
	//            if v != nil {
	//                fmt.Println(v)
	//            } else {
	//                return
	//            }
	//    }
	// }
<原文结束>

# <翻译开始>
	// 消费者持续读取队列中的数据。
	// 如果队列中没有数据，它将阻塞等待。
	// 队列的读取是通过队列对象暴露的
	// queue.C 属性以及使用 selectIO 多路复用语法来实现的。
	// 示例代码如下：
	// for {
	//     select {
	//         case v := <-queue.C:
	//             if v != nil {
	//                 fmt.Println(v)
	//             } else {
	//                 return
	//             }
	//     }
	// }
	// 在这个例子中，`queue.C` 是一个通道，消费者通过它接收队列中的元素。
	// 当有值可读时，`v` 会被赋值并打印；若收到的是 `nil`，则表示某种结束条件，消费者退出循环。 md5:4bb8650995a22499
# <翻译结束>


<原文开始>
	// May Output:
	// 0
	// <nil>
<原文结束>

# <翻译开始>
	// May Output:
	// 0
	// <nil>
# <翻译结束>

