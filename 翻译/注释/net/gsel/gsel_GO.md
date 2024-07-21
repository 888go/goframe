
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
// Package gsel provides selector definition and implements.
<原文结束>

# <翻译开始>
// 包 gsel 提供选择器的定义和实现。 md5:009a7633736cc0b4
# <翻译结束>


<原文开始>
// Builder creates and returns selector in runtime.
<原文结束>

# <翻译开始>
// Builder 在运行时创建并返回选择器。 md5:ffd073ea24966ab1
# <翻译结束>


<原文开始>
// Selector for service balancer.
<原文结束>

# <翻译开始>
// 用于服务均衡器的选择器。 md5:6331f98beb651e36
# <翻译结束>


<原文开始>
// Pick selects and returns service.
<原文结束>

# <翻译开始>
// Pick 选择并返回服务。 md5:2719544a395c30f1
# <翻译结束>


<原文开始>
// Update updates services into Selector.
<原文结束>

# <翻译开始>
// Update 更新服务到Selector中。 md5:8ca4ca3b3f8963c1
# <翻译结束>


<原文开始>
// Node is node interface.
<原文结束>

# <翻译开始>
// Node 是节点接口。 md5:333117a374d3b386
# <翻译结束>


<原文开始>
// Nodes contains multiple Node.
<原文结束>

# <翻译开始>
// Nodes 包含多个 Node。 md5:96478357456f6f66
# <翻译结束>


<原文开始>
// DoneFunc is callback function when RPC invoke done.
<原文结束>

# <翻译开始>
// DoneFunc 是RPC调用完成时的回调函数。 md5:571d85b6ab17bac9
# <翻译结束>


<原文开始>
// DoneInfo contains additional information for done.
<原文结束>

# <翻译开始>
// DoneInfo包含done的额外信息。 md5:fe59ad436c6d2598
# <翻译结束>


<原文开始>
// Err is the rpc error the RPC finished with. It could be nil.
<原文结束>

# <翻译开始>
// Err 是RPC完成时伴随的错误。它可能为nil（即没有错误）。 md5:6f92d17caab8ccb5
# <翻译结束>


<原文开始>
// Trailer contains the metadata from the RPC's trailer, if present.
<原文结束>

# <翻译开始>
// Trailer 包含了如果存在的话，RPC调用的追踪元数据。 md5:291f630111ba0776
# <翻译结束>


<原文开始>
// BytesSent indicates if any bytes have been sent to the server.
<原文结束>

# <翻译开始>
// BytesSent 表示是否已将任何字节发送到服务器。 md5:d9e7a45a59ef93ae
# <翻译结束>


<原文开始>
// BytesReceived indicates if any byte has been received from the server.
<原文结束>

# <翻译开始>
// BytesReceived表示是否已从服务器接收任何字节。 md5:76160cbbde4b2264
# <翻译结束>


<原文开始>
	// ServerLoad is the load received from server. It's usually sent as part of
	// trailing metadata.
	//
	// The only supported type now is *orca_v1.LoadReport.
<原文结束>

# <翻译开始>
	// ServerLoad 是从服务器接收到的负载。它通常作为尾部元数据的一部分发送。
	//
	// 目前唯一支持的类型是 *orca_v1.LoadReport。
	// md5:631ea6798c4714cb
# <翻译结束>


<原文开始>
// DoneInfoMD is a mapping from metadata keys to value array.
// Users should use the following two convenience functions New and Pairs to generate MD.
<原文结束>

# <翻译开始>
// DoneInfoMD 是一个从元数据键映射到值数组的映射。
// 用户应使用以下两个便利函数 New 和 Pairs 来生成 MD。
// md5:46c19d6d882119ab
# <翻译结束>


<原文开始>
// Len returns the number of items in md.
<原文结束>

# <翻译开始>
// Len 返回md中的项目数量。 md5:dd1e6c01d830a7fd
# <翻译结束>


<原文开始>
	// Get obtains the values for a given key.
	//
	// k is converted to lowercase before searching in md.
<原文结束>

# <翻译开始>
	// Get 获取指定键的值。
	//
	// 在对 md 进行搜索之前，k 会被转换为小写。
	// md5:cd83da607a1a524c
# <翻译结束>


<原文开始>
	// Set sets the value of a given key with a slice of values.
	//
	// k is converted to lowercase before storing in md.
<原文结束>

# <翻译开始>
	// Set 使用一系列值设置给定键的值。
	//
	// 在存储到 md 之前，k 将被转换为小写。
	// md5:723075d936727645
# <翻译结束>


<原文开始>
	// Append adds the values to key k, not overwriting what was already stored at
	// that key.
	//
	// k is converted to lowercase before storing in md.
<原文结束>

# <翻译开始>
	// Append 将值添加到键 k，但不会覆盖已经存储在该键上的内容。
	//
	// 在存储到 md 中之前，k 会被转换为小写。
	// md5:8b748588b95754ce
# <翻译结束>


<原文开始>
	// Delete removes the values for a given key k which is converted to lowercase
	// before removing it from md.
<原文结束>

# <翻译开始>
	// Delete 删除给定键 k 的值，该键在从 md 中删除之前会被转换为小写。
	// md5:fa165ee7e187c245
# <翻译结束>


<原文开始>
// String formats and returns Nodes as string.
<原文结束>

# <翻译开始>
// String 方法将节点格式化并返回为字符串。 md5:8aee26fc061ca942
# <翻译结束>

