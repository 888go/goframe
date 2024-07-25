// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package ghttp_test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

func ExampleServer_Run() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello world")
	})
	s.SetPort(8999)
	s.Run()
}

// 自定义保存文件的名称。 md5:ece0b5139469c1f3
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
