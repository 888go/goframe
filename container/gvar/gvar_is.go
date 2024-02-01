// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvar
import (
	"github.com/888go/goframe/internal/utils"
	)
// IsNil 检查 `v` 是否为 nil。
func (v *Var) IsNil() bool {
	return utils.IsNil(v.Val())
}

// IsEmpty 检查 `v` 是否为空。
func (v *Var) IsEmpty() bool {
	return utils.IsEmpty(v.Val())
}

// IsInt 检查 `v` 是否为 int 类型。
func (v *Var) IsInt() bool {
	return utils.IsInt(v.Val())
}

// IsUint 检查 `v` 是否为 uint 类型。
func (v *Var) IsUint() bool {
	return utils.IsUint(v.Val())
}

// IsFloat 检查 `v` 是否为浮点类型。
func (v *Var) IsFloat() bool {
	return utils.IsFloat(v.Val())
}

// IsSlice 检查 `v` 是否为切片类型。
func (v *Var) IsSlice() bool {
	return utils.IsSlice(v.Val())
}

// IsMap 检查 `v` 是否为 map 类型。
func (v *Var) IsMap() bool {
	return utils.IsMap(v.Val())
}

// IsStruct 检查 `v` 是否为结构体类型。
func (v *Var) IsStruct() bool {
	return utils.IsStruct(v.Val())
}
