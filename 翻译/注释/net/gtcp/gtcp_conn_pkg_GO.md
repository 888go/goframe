
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
// Header size for simple package protocol.
<原文结束>

# <翻译开始>
// 简单包协议的头部大小
# <翻译结束>


<原文开始>
// Max header size for simple package protocol.
<原文结束>

# <翻译开始>
// 简单包协议的最大头部大小
# <翻译结束>


<原文开始>
// PkgOption is package option for simple protocol.
<原文结束>

# <翻译开始>
// PkgOption 是用于简单协议的包选项。
# <翻译结束>


<原文开始>
	// HeaderSize is used to mark the data length for next data receiving.
	// It's 2 bytes in default, 4 bytes max, which stands for the max data length
	// from 65535 to 4294967295 bytes.
<原文结束>

# <翻译开始>
// HeaderSize 用于标记下一段数据的长度，用于接收数据时进行判断。
// 默认情况下占2字节，最大为4字节，可表示的数据长度范围从65535字节到4294967295字节。
# <翻译结束>


<原文开始>
	// MaxDataSize is the data field size in bytes for data length validation.
	// If it's not manually set, it'll automatically be set correspondingly with the HeaderSize.
<原文结束>

# <翻译开始>
// MaxDataSize 是用于数据长度验证的数据字段大小，单位为字节。
// 若未手动设置，它将根据 HeaderSize 自动进行相应的设置。
# <翻译结束>


<原文开始>
// Retry policy when operation fails.
<原文结束>

# <翻译开始>
// 当操作失败时的重试策略。
# <翻译结束>


<原文开始>
// SendPkgWithTimeout writes data to connection with timeout using simple package protocol.
<原文结束>

# <翻译开始>
// SendPkgWithTimeout 使用简单的包协议并设置超时，向连接写入数据。
# <翻译结束>


<原文开始>
// SendRecvPkg writes data to connection and blocks reading response using simple package protocol.
<原文结束>

# <翻译开始>
// SendRecvPkg 使用简单的包协议将数据写入连接，并阻塞等待读取响应。
# <翻译结束>


<原文开始>
// SendRecvPkgWithTimeout writes data to connection and reads response with timeout using simple package protocol.
<原文结束>

# <翻译开始>
// SendRecvPkgWithTimeout 使用简单包协议，以超时机制向连接写入数据并读取响应。
# <翻译结束>


<原文开始>
// RecvPkg receives data from connection using simple package protocol.
<原文结束>

# <翻译开始>
// RecvPkg 通过简单的包协议从连接中接收数据。
# <翻译结束>


<原文开始>
// It fills with zero if the header size is lesser than 4 bytes (uint32).
<原文结束>

# <翻译开始>
// 如果头部大小小于4字节（uint32），则用零填充。
# <翻译结束>


<原文开始>
	// It here validates the size of the package.
	// It clears the buffer and returns error immediately if it validates failed.
<原文结束>

# <翻译开始>
// 在这里，它验证包的大小。
// 如果验证失败，它会立即清除缓冲区并返回错误。
# <翻译结束>


<原文开始>
// RecvPkgWithTimeout reads data from connection with timeout using simple package protocol.
<原文结束>

# <翻译开始>
// RecvPkgWithTimeout 使用简单包协议，以超时方式从连接中读取数据。
# <翻译结束>


<原文开始>
// getPkgOption wraps and returns the PkgOption.
// If no option given, it returns a new option with default value.
<原文结束>

# <翻译开始>
// getPkgOption 包装并返回 PkgOption。
// 如果未提供选项，则返回一个具有默认值的新选项。
# <翻译结束>


<原文开始>
// math.MaxInt32 not math.MaxUint32
<原文结束>

# <翻译开始>
// math.MaxInt32 不是 math.MaxUint32
// （这段注释是在强调或纠正某个变量或函数的使用，指出这里使用的是32位整数的最大值常量 `math.MaxInt32`，而不是32位无符号整数的最大值常量 `math.MaxUint32`。）
# <翻译结束>

