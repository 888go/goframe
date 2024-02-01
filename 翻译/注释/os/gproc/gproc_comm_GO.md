
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
// MsgRequest is the request structure for process communication.
<原文结束>

# <翻译开始>
// MsgRequest 是用于进程间通信的请求结构体。
# <翻译结束>


<原文开始>
// MsgResponse is the response structure for process communication.
<原文结束>

# <翻译开始>
// MsgResponse 是用于进程间通信的响应结构体。
# <翻译结束>


<原文开始>
// Default folder name for storing pid to port mapping files.
<原文结束>

# <翻译开始>
// 默认的文件夹名称，用于存储pid到端口映射的文件。
# <翻译结束>


<原文开始>
// Starting port number for receiver listening.
<原文结束>

# <翻译开始>
// 接收器监听的起始端口号。
# <翻译结束>


<原文开始>
// Max size for each message queue of the group.
<原文结束>

# <翻译开始>
// 每个组的消息队列的最大尺寸。
# <翻译结束>


<原文开始>
	// commReceiveQueues is the group name to queue map for storing received data.
	// The value of the map is type of *gqueue.Queue.
<原文结束>

# <翻译开始>
// commReceiveQueues 是用于存储接收到数据的组名到队列映射。
// 该映射的值类型为 *gqueue.Queue。
# <翻译结束>


<原文开始>
// commPidFolderPath specifies the folder path storing pid to port mapping files.
<原文结束>

# <翻译开始>
// commPidFolderPath 指定存储进程ID到端口映射文件的文件夹路径。
# <翻译结束>


<原文开始>
// commPidFolderPathOnce is used for lazy calculation for `commPidFolderPath` is necessary.
<原文结束>

# <翻译开始>
// commPidFolderPathOnce 用于延迟计算，只有在必要时才计算 `commPidFolderPath`。
# <翻译结束>


<原文开始>
// getConnByPid creates and returns a TCP connection for specified pid.
<原文结束>

# <翻译开始>
// getConnByPid 为指定的pid创建并返回一个TCP连接。
# <翻译结束>


<原文开始>
// getPortByPid returns the listening port for specified pid.
// It returns 0 if no port found for the specified pid.
<原文结束>

# <翻译开始>
// getPortByPid 根据指定的进程id返回其监听的端口号。
// 如果指定的pid没有找到对应的端口，则返回0。
# <翻译结束>


<原文开始>
// getCommFilePath returns the pid to port mapping file path for given pid.
<原文结束>

# <翻译开始>
// getCommFilePath 根据给定的pid返回其对应的端口映射文件路径。
# <翻译结束>


<原文开始>
// getCommPidFolderPath retrieves and returns the available directory for storing pid mapping files.
<原文结束>

# <翻译开始>
// getCommPidFolderPath 获取并返回可用于存储pid映射文件的可用目录。
# <翻译结束>


<原文开始>
// Message group name.
<原文结束>

# <翻译开始>
// 消息组名称。
# <翻译结束>


<原文开始>
// 1: OK; Other: Error.
<原文结束>

# <翻译开始>
// 1: 表示成功；其它值：表示错误
# <翻译结束>


<原文开始>
// Default group name.
<原文结束>

# <翻译开始>
// 默认分组名称。
# <翻译结束>

