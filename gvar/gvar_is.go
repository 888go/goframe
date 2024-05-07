// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类

import (
	"github.com/888go/goframe/gvar/internal/utils"
)

// IsNil 检查 `v` 是否为 nil。
func (v *Var) X是否为Nil() bool {
	return utils.IsNil(v.X取值())
}

// IsEmpty 检查 `v` 是否为空。
func (v *Var) X是否为空() bool {
	return utils.IsEmpty(v.X取值())
}

// IsInt 检查 `v` 是否为 int 类型。
func (v *Var) X是否为整数() bool {
	return utils.IsInt(v.X取值())
}

// IsUint 检查 `v` 是否为 uint 类型。
func (v *Var) X是否为正整数() bool {
	return utils.IsUint(v.X取值())
}

// IsFloat 检查 `v` 是否为浮点类型。
func (v *Var) X是否为小数() bool {
	return utils.IsFloat(v.X取值())
}

// IsSlice 检查 `v` 是否为切片类型。
func (v *Var) X是否为切片() bool {
	return utils.IsSlice(v.X取值())
}

// IsMap 检查 `v` 是否为 map 类型。
func (v *Var) X是否为Map() bool {
	return utils.IsMap(v.X取值())
}

// IsStruct 检查 `v` 是否为结构体类型。
func (v *Var) X是否为结构() bool {
	return utils.IsStruct(v.X取值())
}
