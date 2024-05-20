// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gclient

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketClient 封装了底层的websocket客户端连接，
// 并提供了便捷的功能方法。
// md5:82e7259d43b0aaee
type WebSocketClient struct {
	*websocket.Dialer
}

// NewWebSocket 创建并返回一个新的WebSocket客户端对象。. md5:03f6812fa459ed81
func NewWebSocket() *WebSocketClient {
	return &WebSocketClient{
		&websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 45 * time.Second,
		},
	}
}
