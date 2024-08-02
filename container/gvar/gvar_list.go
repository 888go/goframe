// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 泛型类

import (
	gutil "github.com/888go/goframe/util/gutil"
)

// ListItemValues 获取并返回所有项结构体/映射中键为`key`的元素。
// 注意，参数`list`应该是包含映射或结构体元素的切片类型，
// 否则将返回一个空切片。
// md5:22a160e15c6a6d6c
func (v *Var) ListItemValues(key interface{}) (values []interface{}) {
	return gutil.ListItemValues(v.Val(), key)
}

// ListItemValuesUnique 获取并返回具有键为`key`的所有结构体/映射的独特元素。
// 请注意，参数`list`应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
// md5:0f361d3ff901d0a1
func (v *Var) ListItemValuesUnique(key string) []interface{} {
	return gutil.ListItemValuesUnique(v.Val(), key)
}
