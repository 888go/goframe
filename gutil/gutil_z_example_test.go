// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/888go/goframe/gutil"
)

func ExampleSliceInsertBefore() {
	s1 := g.Slice{
		0, 1, 2, 3, 4,
	}
	s2 := 工具类.SliceInsertBefore(s1, 1, 8, 9)
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
	s2 := 工具类.SliceInsertAfter(s1, 1, 8, 9)
	fmt.Println(s1)
	fmt.Println(s2)

	// Output:
	// [0 1 2 3 4]
	// [0 1 8 9 2 3 4]
}
