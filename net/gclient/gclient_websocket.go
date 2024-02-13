// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类

import (
	"net/http"
	"time"
	
	"github.com/gorilla/websocket"
)

// WebSocketClient 对底层的websocket客户端连接进行封装，
// 并提供了便捷的功能方法。
type WebSocketClient struct {
	*websocket.Dialer
}

// NewWebSocket 创建并返回一个新的 WebSocketClient 对象。
func X创建WebSocket() *WebSocketClient {
	return &WebSocketClient{
		&websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 45 * time.Second,
		},
	}
}
