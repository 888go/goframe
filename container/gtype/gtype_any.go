// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtype

// Any是一个用于并发安全操作的any类型的结构体。 md5:40fc918c497f5cab
type Any = Interface

// NewAny 为任何类型创建并返回一个并发安全的对象，初始值为`value`。
// md5:5624706a34a7a1be
func NewAny(value ...any) *Any {
	t := &Any{}
	if len(value) > 0 && value[0] != nil {
		t.value.Store(value[0])
	}
	return t
}
