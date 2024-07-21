	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package gfile_test

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gfile"
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
