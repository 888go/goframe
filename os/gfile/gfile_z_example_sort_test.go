// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package 文件类_test

import (
	"fmt"

	gfile "github.com/888go/goframe/os/gfile"
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
