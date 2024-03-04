
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
// Close the queue in three seconds.
<原文结束>

# <翻译开始>
// 在三秒钟后关闭队列。
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
// 如果队列中没有数据，则会阻塞等待。
// 通过队列对象公开的queue.C属性来读取队列，
// 并使用selectIO多路复用语法进行操作。
// 示例代码如下：
// for {
//     select {
//         case v := <-queue.C: // 从队列channel接收数据
//             if v != nil { // 若接收到的数据不为空
//                 fmt.Println(v) // 输出接收到的数据
//             } else { // 若接收到的数据为空
//                 return // 则退出循环（或程序）
//             }
//     }
// }
# <翻译结束>





 
<原文开始>
// Size is alias of Len.
<原文结束>

# <翻译开始>
// Size 是 Len 的别名。
# <翻译结束>

