// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gutil_test

import (
	"fmt"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/util/gutil"
)

func ExampleSliceInsertBefore() {
	s1 := g.Slice{
		0, 1, 2, 3, 4,
	}
	s2 := gutil.SliceInsertBefore(s1, 1, 8, 9)
	fmt.Println(s1)
	fmt.Println(s2)

	// Output:
	// [0 1 2 3 4]
	// [0 8 9 1 2 3 4]
}

func ExampleSliceInsertAfter() {
	s1 := g.Slice{
		0, 1, 2, 3, 4,
	}
	s2 := gutil.SliceInsertAfter(s1, 1, 8, 9)
	fmt.Println(s1)
	fmt.Println(s2)

	// Output:
	// [0 1 2 3 4]
	// [0 1 8 9 2 3 4]
}
