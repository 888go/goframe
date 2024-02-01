// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"github.com/gorilla/websocket"
	)
// WebSocket 对底层的websocket连接进行封装
// 并提供了便捷的功能函数。
type WebSocket struct {
	*websocket.Conn
}

const (
// WsMsgText TextMessage 表示一个文本数据消息。
// 文本消息负载被解释为UTF-8编码的文本数据。
	WsMsgText = websocket.TextMessage

	// WsMsgBinary BinaryMessage 表示一个二进制数据消息。
	WsMsgBinary = websocket.BinaryMessage

// WsMsgClose 关闭消息表示一个关闭控制消息。
// 可选的消息负载包含一个数字代码和文本内容。
// 使用 FormatCloseMessage 函数来格式化一个关闭消息的负载。
	WsMsgClose = websocket.CloseMessage

// WsMsgPing PingMessage 表示一个ping控制消息。
// 可选的消息负载是UTF-8编码的文本。
	WsMsgPing = websocket.PingMessage

// WsMsgPong 表示一个 pong 控制消息。
// 可选的消息负载是 UTF-8 编码的文本。
	WsMsgPong = websocket.PongMessage
)
