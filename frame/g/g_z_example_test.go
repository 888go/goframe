// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package g_test

import (
	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
)

func ExampleServer() {
	// A hello world example.
	s := g.Http类()
	s.X绑定("/", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("hello world")
	})
	s.X设置监听端口(8999)
	s.X启动服务()
}
