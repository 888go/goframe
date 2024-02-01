// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test
import (
	"fmt"
	
	"github.com/888go/goframe/os/gfile"
	)

func ExampleSortFiles() {
	files := []string{
		"/aaa/bbb/ccc.txt",
		"/aaa/bbb/",
		"/aaa/",
		"/aaa",
		"/aaa/ccc/ddd.txt",
		"/bbb",
		"/0123",
		"/ddd",
		"/ccc",
	}
	sortOut := gfile.SortFiles(files)
	fmt.Println(sortOut)

	// Output:
	// [/0123 /aaa /aaa/ /aaa/bbb/ /aaa/bbb/ccc.txt /aaa/ccc/ddd.txt /bbb /ccc /ddd]
}
