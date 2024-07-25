
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
// Header size for simple package protocol.
<原文结束>

# <翻译开始>
// 简单包协议的头部大小。 md5:6e3294120717ff1f
# <翻译结束>


<原文开始>
// Max header size for simple package protocol.
<原文结束>

# <翻译开始>
// 简单包协议的最大头部大小。 md5:b1e4f447dc182fa5
# <翻译结束>


<原文开始>
// PkgOption is package option for simple protocol.
<原文结束>

# <翻译开始>
// PkgOption是简单协议的包选项。 md5:f49ee7556a39be3e
# <翻译结束>


<原文开始>
	// HeaderSize is used to mark the data length for next data receiving.
	// It's 2 bytes in default, 4 bytes max, which stands for the max data length
	// from 65535 to 4294967295 bytes.
<原文结束>

# <翻译开始>
	// HeaderSize 用于标记接下来接收数据的长度。
	// 它默认为2字节，最大为4字节，代表数据的最大长度可以从65535字节到4294967295字节。 md5:cc02a98c94fddd76
# <翻译结束>


<原文开始>
	// MaxDataSize is the data field size in bytes for data length validation.
	// If it's not manually set, it'll automatically be set correspondingly with the HeaderSize.
<原文结束>

# <翻译开始>
	// MaxDataSize 是数据字段的字节大小，用于数据长度验证。
	// 如果未手动设置，它将根据HeaderSize自动相应设置。 md5:a47982162ce5ef01
# <翻译结束>


<原文开始>
// Retry policy when operation fails.
<原文结束>

# <翻译开始>
	// 操作失败时的重试策略。 md5:cd672b1b96adbbdd
# <翻译结束>


<原文开始>
// SendPkg send data using simple package protocol.
//
// Simple package protocol: DataLength(24bit)|DataField(variant)。
//
// Note that,
// 1. The DataLength is the length of DataField, which does not contain the header size.
// 2. The integer bytes of the package are encoded using BigEndian order.
<原文结束>

# <翻译开始>
// SendPkg 使用简单包协议发送数据。
//
// 简单包协议：DataLength(24位)|DataField(variant)。
//
// 注意，
// 1. DataLength 是 DataField 的长度，不包含头大小。
// 2. 包的整数字节使用大端序编码。 md5:daa39f4e32227d79
# <翻译结束>


<原文开始>
// SendPkgWithTimeout writes data to connection with timeout using simple package protocol.
<原文结束>

# <翻译开始>
// 使用简单包协议带超时时间地向连接发送数据。 md5:3f89f6011aed63bc
# <翻译结束>


<原文开始>
// SendRecvPkg writes data to connection and blocks reading response using simple package protocol.
<原文结束>

# <翻译开始>
// SendRecvPkg 使用简单的包协议将数据写入连接，并阻塞读取响应。 md5:c157760431f11896
# <翻译结束>


<原文开始>
// SendRecvPkgWithTimeout writes data to connection and reads response with timeout using simple package protocol.
<原文结束>

# <翻译开始>
// SendRecvPkgWithTimeout 使用简单包协议向连接写入数据，并在超时后读取响应。 md5:6da9109d534f7729
# <翻译结束>


<原文开始>
// RecvPkg receives data from connection using simple package protocol.
<原文结束>

# <翻译开始>
// RecvPkg 使用简单包协议从连接接收数据。 md5:cf1329c5df27539a
# <翻译结束>


<原文开始>
// It fills with zero if the header size is lesser than 4 bytes (uint32).
<原文结束>

# <翻译开始>
		// 如果头部大小小于4字节（uint32），则用零填充。 md5:5e9e147401796703
# <翻译结束>


<原文开始>
	// It here validates the size of the package.
	// It clears the buffer and returns error immediately if it validates failed.
<原文结束>

# <翻译开始>
	// 此处校验包的大小。
	// 如果校验失败，会立即清空缓冲区并返回错误。 md5:0871405b30986628
# <翻译结束>


<原文开始>
// RecvPkgWithTimeout reads data from connection with timeout using simple package protocol.
<原文结束>

# <翻译开始>
// RecvPkgWithTimeout 使用简单包协议，从连接中读取数据，同时设置超时。 md5:5e1d4882f4476862
# <翻译结束>


<原文开始>
// getPkgOption wraps and returns the PkgOption.
// If no option given, it returns a new option with default value.
<原文结束>

# <翻译开始>
// getPkgOption 包装并返回 PkgOption。
// 如果没有提供选项，则返回一个具有默认值的新选项。 md5:752809cff379479d
# <翻译结束>


<原文开始>
// math.MaxInt32 not math.MaxUint32
<原文结束>

# <翻译开始>
			// math.MaxInt32 而不是 math.MaxUint32. md5:11ed9a0830ca2d39
# <翻译结束>

