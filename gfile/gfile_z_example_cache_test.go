// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"fmt"
	"time"
	
	"github.com/888go/goframe/gfile"
)

func ExampleGetContentsWithCache() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = 文件类.X取临时目录("gfile_example_cache")
		tempFile = 文件类.X路径生成(tempDir, fileName)
	)

	// write contents
	文件类.X写入文本(tempFile, "goframe example content")

// 它以一分钟的缓存时长读取文件内容，这意味着在一分钟内如果没有任何IO操作，它将从缓存中读取而不是实际读取文件。
	fmt.Println(文件类.X缓存读文本(tempFile, time.Minute))

	// 写入新内容将会清除其缓存
	文件类.X写入文本(tempFile, "new goframe example content")

	// 在文件内容更改后，缓存清除会有一些延迟。
	time.Sleep(time.Second * 1)

	// read contents
	fmt.Println(文件类.X缓存读文本(tempFile))

	// May Output:
	// goframe example content
	// new goframe example content
}
