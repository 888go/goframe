// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类_test

import (
	"fmt"
	"time"

	gfile "github.com/888go/goframe/os/gfile"
)

func ExampleGetContentsWithCache() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.X取临时目录("gfile_example_cache")
		tempFile = gfile.X路径生成(tempDir, fileName)
	)

	// write contents
	gfile.X写入文本(tempFile, "goframe example content")

	// 它使用1分钟的缓存过期时间读取文件内容，
	// 这意味着在接下来的一分钟内，如果没有进行任何IO操作，它将从缓存中读取。
	// md5:2d9221dfe7c2f44a
	fmt.Println(gfile.X缓存读文本(tempFile, time.Minute))

		// 写入新内容将清除其缓存. md5:cdefd2fa84d5ae75
	gfile.X写入文本(tempFile, "new goframe example content")

		// 文件内容更改后，清除缓存会有一些延迟。 md5:7f776df808d0e69c
	time.Sleep(time.Second * 1)

	// read contents
	fmt.Println(gfile.X缓存读文本(tempFile))

	// May Output:
	// goframe example content
	// new goframe example content
}
