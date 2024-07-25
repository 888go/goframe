// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gvar

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// Struct 将值`v`映射到`pointer`。
// 参数`pointer`应为指向结构体实例的指针。
// 参数`mapping`用于指定键到属性的映射规则。 md5:a5bf066b3ef1c653
func (v *Var) Struct(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Struct(v.Val(), pointer, mapping...)
}

// Structs 将 `v` 转换并返回为指定结构体切片。 md5:396a4079aac15c40
func (v *Var) Structs(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Structs(v.Val(), pointer, mapping...)
}
