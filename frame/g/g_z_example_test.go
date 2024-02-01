// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package g_test
import (
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	)

func ExampleServer() {
	// 一个简单的“你好，世界”示例。
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello world")
	})
	s.SetPort(8999)
	s.Run()
}
