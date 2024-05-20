// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import "github.com/gorilla/websocket"

// WebSocket 包装了底层的websocket连接，并提供了方便的函数。
// 
// 警告：将来将被移除，请改用第三方websocket库。
// md5:85d39ed9e94efa8d
type WebSocket struct {
	*websocket.Conn
}

const (
// WsMsgText TextMessage 表示一个文本数据消息。
// 文本消息的有效负载被解释为UTF-8编码的文本数据。
// md5:2212894321ec5f57
	WsMsgText = websocket.TextMessage

	// WsMsgBinary 表示一个二进制数据消息。. md5:a65808b202eac553
	WsMsgBinary = websocket.BinaryMessage

// WsMsgClose CloseMessage 表示一个关闭控制消息。可选的消息负载包含一个数字代码和文本。使用 FormatCloseMessage 函数格式化关闭消息负载。
// md5:a469c715d4927f73
	WsMsgClose = websocket.CloseMessage

// WsMsgPing PingMessage 表示一个ping控制消息。可选的消息负载是UTF-8编码的文本。
// md5:be2bfff84d685414
	WsMsgPing = websocket.PingMessage

// WsMsgPong PongMessage 表示一个 pong 控制消息。
// 可选的消息负载是 UTF-8 编码的文本。
// md5:7fd652a30abef63d
	WsMsgPong = websocket.PongMessage
)
