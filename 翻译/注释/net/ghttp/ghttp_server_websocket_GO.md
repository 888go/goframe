
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
// WebSocket wraps the underlying websocket connection
// and provides convenient functions.
<原文结束>

# <翻译开始>
// WebSocket 对底层的websocket连接进行封装
// 并提供了便捷的功能函数。
# <翻译结束>


<原文开始>
	// WsMsgText TextMessage denotes a text data message.
	// The text message payload is interpreted as UTF-8 encoded text data.
<原文结束>

# <翻译开始>
// WsMsgText TextMessage 表示一个文本数据消息。
// 文本消息负载被解释为UTF-8编码的文本数据。
# <翻译结束>


<原文开始>
// WsMsgBinary BinaryMessage denotes a binary data message.
<原文结束>

# <翻译开始>
// WsMsgBinary BinaryMessage 表示一个二进制数据消息。
# <翻译结束>


<原文开始>
	// WsMsgClose CloseMessage denotes a close control message.
	// The optional message payload contains a numeric code and text.
	// Use the FormatCloseMessage function to format a close message payload.
<原文结束>

# <翻译开始>
// WsMsgClose 关闭消息表示一个关闭控制消息。
// 可选的消息负载包含一个数字代码和文本内容。
// 使用 FormatCloseMessage 函数来格式化一个关闭消息的负载。
# <翻译结束>


<原文开始>
	// WsMsgPing PingMessage denotes a ping control message.
	// The optional message payload is UTF-8 encoded text.
<原文结束>

# <翻译开始>
// WsMsgPing PingMessage 表示一个ping控制消息。
// 可选的消息负载是UTF-8编码的文本。
# <翻译结束>


<原文开始>
	// WsMsgPong PongMessage denotes a pong control message.
	// The optional message payload is UTF-8 encoded text.
<原文结束>

# <翻译开始>
// WsMsgPong 表示一个 pong 控制消息。
// 可选的消息负载是 UTF-8 编码的文本。
# <翻译结束>

