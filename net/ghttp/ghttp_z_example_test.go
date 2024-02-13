// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gfile"
)

func ExampleServer_Run() {
	s := g.Http类()
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区("hello world")
	})
	s.X设置监听端口(8999)
	s.X启动服务()
}

// 自定义保存文件名。
func ExampleUploadFile_Save() {
	s := g.Http类()
	s.X绑定("/upload", func(r *http类.Request) {
		file := r.X取上传文件对象("TestFile")
		if file == nil {
			r.Response.X写响应缓冲区("empty file")
			return
		}
		file.Filename = "MyCustomFileName.txt"
		fileName, err := file.X保存(文件类.X取临时目录())
		if err != nil {
			r.Response.X写响应缓冲区(err)
			return
		}
		r.Response.X写响应缓冲区(fileName)
	})
	s.X设置监听端口(8999)
	s.X启动服务()
}
