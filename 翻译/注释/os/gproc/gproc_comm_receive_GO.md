
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
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// tcpListened marks whether the receiving listening service started.
<原文结束>

# <翻译开始>
// tcpListened 标记接收端监听服务是否已启动。
# <翻译结束>


<原文开始>
// Receive blocks and receives message from other process using local TCP listening.
// Note that, it only enables the TCP listening service when this function called.
<原文结束>

# <翻译开始>
// 接收区块并通过本地TCP监听从其他进程接收消息。
// 注意，只有当调用此函数时，才会启用TCP监听服务。
# <翻译结束>


<原文开始>
// Use atomic operations to guarantee only one receiver goroutine listening.
<原文结束>

# <翻译开始>
// 使用原子操作以确保只有一个接收goroutine在监听。
# <翻译结束>


<原文开始>
// receiveTcpListening scans local for available port and starts listening.
<原文结束>

# <翻译开始>
// receiveTcpListening 在本地扫描可用端口并开始监听。
# <翻译结束>


<原文开始>
// Save the port to the pid file.
<原文结束>

# <翻译开始>
// 将端口保存到pid文件中。
# <翻译结束>


<原文开始>
// receiveTcpHandler is the connection handler for receiving data.
<原文结束>

# <翻译开始>
// receiveTcpHandler 是用于接收数据的连接处理器。
# <翻译结束>


<原文开始>
// Just close the connection if any error occurs.
<原文结束>

# <翻译开始>
// 如果发生任何错误，仅关闭连接即可。
# <翻译结束>


<原文开始>
// Blocking receiving.
<原文结束>

# <翻译开始>
// 阻塞接收。
# <翻译结束>


<原文开始>
// Push to buffer queue.
<原文结束>

# <翻译开始>
// 将元素推送到缓冲队列中。
# <翻译结束>

