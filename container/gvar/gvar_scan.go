// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 泛型类

import (
	gconv "github.com/888go/goframe/util/gconv"
)

// X取结构体指针 自动检查`pointer`的类型，并将`params`转换为`pointer`。它支持`pointer`类型，包括`*map`、`*[]map`、`*[]*map`、`*struct`、`**struct`、`[]struct`和`[]*struct`。
// 
// 参见gconv.X取结构体指针。
// md5:7f6cfec69999319d
func (v *Var) X取结构体指针(结构体指针 interface{}, 名称映射 ...map[string]string) error {
	return gconv.Scan(v.X取值(), 结构体指针, 名称映射...)
}
