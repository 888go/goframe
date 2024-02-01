
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
// Package gsel provides selector definition and implements.
<原文结束>

# <翻译开始>
// Package gsel 提供选择器定义及其实现。
# <翻译结束>


<原文开始>
// Builder creates and returns selector in runtime.
<原文结束>

# <翻译开始>
// Builder在运行时创建并返回选择器。
# <翻译结束>


<原文开始>
// Selector for service balancer.
<原文结束>

# <翻译开始>
// 服务负载均衡器选择器。
# <翻译结束>


<原文开始>
// Pick selects and returns service.
<原文结束>

# <翻译开始>
// Pick 选择并返回服务。
# <翻译结束>


<原文开始>
// Update updates services into Selector.
<原文结束>

# <翻译开始>
// Update 将服务更新到Selector中。
# <翻译结束>


<原文开始>
// Nodes contains multiple Node.
<原文结束>

# <翻译开始>
// Nodes 包含多个 Node。
# <翻译结束>


<原文开始>
// DoneFunc is callback function when RPC invoke done.
<原文结束>

# <翻译开始>
// DoneFunc 是当 RPC 调用完成时的回调函数。
# <翻译结束>


<原文开始>
// DoneInfo contains additional information for done.
<原文结束>

# <翻译开始>
// DoneInfo 包含 done 的附加信息。
# <翻译结束>


<原文开始>
// Err is the rpc error the RPC finished with. It could be nil.
<原文结束>

# <翻译开始>
// Err 是 RPC 结束时的错误，可能为 nil。
# <翻译结束>


<原文开始>
// Trailer contains the metadata from the RPC's trailer, if present.
<原文结束>

# <翻译开始>
// Trailer 包含了RPC的尾部元数据（如果存在的话）。
# <翻译结束>


<原文开始>
// BytesSent indicates if any bytes have been sent to the server.
<原文结束>

# <翻译开始>
// BytesSent 表示是否已向服务器发送了任何字节。
# <翻译结束>


<原文开始>
// BytesReceived indicates if any byte has been received from the server.
<原文结束>

# <翻译开始>
// BytesReceived 表示是否已从服务器接收到任何字节。
# <翻译结束>


<原文开始>
	// ServerLoad is the load received from server. It's usually sent as part of
	// trailing metadata.
	//
	// The only supported type now is *orca_v1.LoadReport.
<原文结束>

# <翻译开始>
// ServerLoad表示从服务器接收到的负载信息。通常，它作为
// 附属元数据的一部分进行发送。
//
// 当前唯一支持的数据类型是*orca_v1.LoadReport。
# <翻译结束>


<原文开始>
// DoneInfoMD is a mapping from metadata keys to value array.
// Users should use the following two convenience functions New and Pairs to generate MD.
<原文结束>

# <翻译开始>
// DoneInfoMD 是一个从元数据键到值数组的映射。
// 用户应使用以下两个便捷函数 New 和 Pairs 来生成 MD。
# <翻译结束>


<原文开始>
// Len returns the number of items in md.
<原文结束>

# <翻译开始>
// Len 返回 md 中的项目数量。
# <翻译结束>


<原文开始>
	// Get obtains the values for a given key.
	//
	// k is converted to lowercase before searching in md.
<原文结束>

# <翻译开始>
// Get 方法用于获取给定键对应的值。
//
// 在md中搜索之前，会将k转换为小写。
# <翻译结束>


<原文开始>
	// Set sets the value of a given key with a slice of values.
	//
	// k is converted to lowercase before storing in md.
<原文结束>

# <翻译开始>
// Set 用于设置给定键的值为一个切片类型的值。
//
// 在存储到 md 之前，k 会被转换为小写。
# <翻译结束>


<原文开始>
	// Append adds the values to key k, not overwriting what was already stored at
	// that key.
	//
	// k is converted to lowercase before storing in md.
<原文结束>

# <翻译开始>
// Append 将值添加到键 k，但不会覆盖该键已存储的内容。
//
// 在存储到 md 之前，会将 k 转换为小写。
# <翻译结束>


<原文开始>
	// Delete removes the values for a given key k which is converted to lowercase
	// before removing it from md.
<原文结束>

# <翻译开始>
// Delete 删除给定键 k 的值，在从 md 中移除前，会将键 k 转换为小写。
# <翻译结束>


<原文开始>
// String formats and returns Nodes as string.
<原文结束>

# <翻译开始>
// String 方法格式化并以字符串形式返回 Nodes。
# <翻译结束>


<原文开始>
// Node is node interface.
<原文结束>

# <翻译开始>
// Node 是节点接口。
# <翻译结束>

