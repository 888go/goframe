// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_WebSocket(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X绑定("/ws", func(r *ghttp.Request) {
		ws, err := r.X升级为websocket请求()
		if err != nil {
			r.X退出当前()
		}
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}
			if err = ws.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf(
			"ws://127.0.0.1:%d/ws", s.X取已监听端口(),
		), nil)
		t.AssertNil(err)
		defer conn.Close()

		msg := []byte("hello")
		err = conn.WriteMessage(websocket.TextMessage, msg)
		t.AssertNil(err)

		mt, data, err := conn.ReadMessage()
		t.AssertNil(err)
		t.Assert(mt, websocket.TextMessage)
		t.Assert(data, msg)
	})
}
