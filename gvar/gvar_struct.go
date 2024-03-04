// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvar

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// Struct 将 `v` 的值映射到 `pointer`。
// 参数 `pointer` 应该是指向结构体实例的指针。
// 参数 `mapping` 用于指定键到属性的映射规则。
func (v *Var) Struct(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Struct(v.Val(), pointer, mapping...)
}

// Structs 将 `v` 转换并以给定的结构体切片形式返回。
func (v *Var) Structs(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Structs(v.Val(), pointer, mapping...)
}
