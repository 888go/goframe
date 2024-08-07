// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gfile "github.com/888go/goframe/os/gfile"
)

func ExampleServer_Run() {
	s := g.Http类()
	s.X绑定("/", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区("hello world")
	})
	s.X设置监听端口(8999)
	s.X启动服务()
}

// 自定义保存文件的名称。 md5:ece0b5139469c1f3
func ExampleUploadFile_Save() {
	s := g.Http类()
	s.X绑定("/upload", func(r *ghttp.Request) {
		file := r.X取上传文件对象("TestFile")
		if file == nil {
			r.X响应.X写响应缓冲区("empty file")
			return
		}
		file.Filename = "MyCustomFileName.txt"
		fileName, err := file.X保存(gfile.X取临时目录())
		if err != nil {
			r.X响应.X写响应缓冲区(err)
			return
		}
		r.X响应.X写响应缓冲区(fileName)
	})
	s.X设置监听端口(8999)
	s.X启动服务()
}
