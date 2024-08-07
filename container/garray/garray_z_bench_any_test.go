// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 切片类_test

import (
	"testing"

	garray "github.com/888go/goframe/container/garray"
)

type anySortedArrayItem struct {
	priority int64
	value    interface{}
}

var (
	anyArray       = garray.NewArray别名()
	anySortedArray = garray.X创建排序(func(a, b interface{}) int {
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
