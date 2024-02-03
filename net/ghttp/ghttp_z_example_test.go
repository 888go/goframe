// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test

import (
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gfile"
)

func ExampleServer_Run() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello world")
	})
	s.SetPort(8999)
	s.Run()
}

// 自定义保存文件名。
func ExampleUploadFile_Save() {
	s := g.Server()
	s.BindHandler("/upload", func(r *ghttp.Request) {
		file := r.GetUploadFile("TestFile")
		if file == nil {
			r.Response.Write("empty file")
			return
		}
		file.Filename = "MyCustomFileName.txt"
		fileName, err := file.Save(gfile.Temp())
		if err != nil {
			r.Response.Write(err)
			return
		}
		r.Response.Write(fileName)
	})
	s.SetPort(8999)
	s.Run()
}
