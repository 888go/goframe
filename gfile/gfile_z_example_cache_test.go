// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test

import (
	"fmt"
	"time"
	
	"github.com/888go/goframe/gfile"
)

func ExampleGetContentsWithCache() {
	// init
	var (
		fileName = "gflie_example.txt"
		tempDir  = gfile.Temp("gfile_example_cache")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goframe example content")

// 它以一分钟的缓存时长读取文件内容，这意味着在一分钟内如果没有任何IO操作，它将从缓存中读取而不是实际读取文件。
	fmt.Println(gfile.GetContentsWithCache(tempFile, time.Minute))

	// 写入新内容将会清除其缓存
	gfile.PutContents(tempFile, "new goframe example content")

	// 在文件内容更改后，缓存清除会有一些延迟。
	time.Sleep(time.Second * 1)

	// read contents
	fmt.Println(gfile.GetContentsWithCache(tempFile))

	// May Output:
	// goframe example content
	// new goframe example content
}
