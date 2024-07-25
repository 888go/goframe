
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
// MsgRequest is the request structure for process communication.
<原文结束>

# <翻译开始>
// MsgRequest是进程间通信的请求结构体。 md5:aa294ed7aef773f3
# <翻译结束>


<原文开始>
// MsgResponse is the response structure for process communication.
<原文结束>

# <翻译开始>
// MsgResponse 是进程通信中的响应结构体。 md5:a2e9e35f8a32b58e
# <翻译结束>


<原文开始>
// Default folder name for storing pid to port mapping files.
<原文结束>

# <翻译开始>
// 默认的保存pid到端口映射文件的文件夹名称。 md5:64d7a3cc62fc8b3c
# <翻译结束>


<原文开始>
// Starting port number for receiver listening.
<原文结束>

# <翻译开始>
// 用于接收者监听的起始端口号。 md5:57cde4f483b095cf
# <翻译结束>


<原文开始>
// Max size for each message queue of the group.
<原文结束>

# <翻译开始>
// 集群中每个消息队列的最大大小。 md5:64e3f3ac37111858
# <翻译结束>


<原文开始>
	// commReceiveQueues is the group name to queue map for storing received data.
	// The value of the map is type of *gqueue.Queue.
<原文结束>

# <翻译开始>
	// commReceiveQueues 是一个用于存储接收到的数据的组名到队列的映射。
	// 该映射的值类型为*gqueue.Queue。
	// md5:adb11ba95544ea8c
# <翻译结束>


<原文开始>
// commPidFolderPath specifies the folder path storing pid to port mapping files.
<原文结束>

# <翻译开始>
	// commPidFolderPath 指定了存储 PID 到端口映射文件的文件夹路径。 md5:bc9b0e25bfe8ea53
# <翻译结束>


<原文开始>
// commPidFolderPathOnce is used for lazy calculation for `commPidFolderPath` is necessary.
<原文结束>

# <翻译开始>
	// commPidFolderPathOnce用于惰性计算`commPidFolderPath`，只有在必要时才进行。 md5:669e811a3607b61c
# <翻译结束>


<原文开始>
// getConnByPid creates and returns a TCP connection for specified pid.
<原文结束>

# <翻译开始>
// getConnByPid 为指定的 pid 创建并返回一个 TCP 连接。 md5:19b60bfdf9f18aa2
# <翻译结束>


<原文开始>
// getPortByPid returns the listening port for specified pid.
// It returns 0 if no port found for the specified pid.
<原文结束>

# <翻译开始>
// getPortByPid 根据指定的进程ID返回其监听的端口。
// 如果没有为指定的进程ID找到端口，则返回0。
// md5:1fc2deacfe985ab1
# <翻译结束>


<原文开始>
// getCommFilePath returns the pid to port mapping file path for given pid.
<原文结束>

# <翻译开始>
// getCommFilePath 返回给定pid的进程到端口映射文件路径。 md5:6b8e5776476edbb5
# <翻译结束>


<原文开始>
// getCommPidFolderPath retrieves and returns the available directory for storing pid mapping files.
<原文结束>

# <翻译开始>
// getCommPidFolderPath 获取并返回用于存储进程映射文件的可用目录。 md5:d871e38ee1ac7054
# <翻译结束>

