// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 数组类_test

import (
	"testing"
	
	"github.com/888go/goframe/garray"
)

type anySortedArrayItem struct {
	priority int64
	value    interface{}
}

var (
	anyArray       = 数组类.NewArray别名()
	anySortedArray = 数组类.X创建排序(func(a, b interface{}) int {
		return int(a.(anySortedArrayItem).priority - b.(anySortedArrayItem).priority)
	})
)

func Benchmark_AnyArray_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anyArray.Append别名(i)
	}
}

func Benchmark_AnySortedArray_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anySortedArray.X入栈右(anySortedArrayItem{
			priority: int64(i),
			value:    i,
		})
	}
}
