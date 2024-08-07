// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 泛型类

import (
	"github.com/888go/goframe/internal/utils"
)

// X是否为Nil 检查 `v` 是否为 nil。 md5:af4db285987ff32d
func (v *Var) X是否为Nil() bool {
	return utils.IsNil(v.X取值())
}

// X是否为空 检查 `v` 是否为空。 md5:270630fa930d8a69
func (v *Var) X是否为空() bool {
	return utils.IsEmpty(v.X取值())
}

// X是否为整数 检查 `v` 是否为整数类型。 md5:2f04bd336f37dccf
func (v *Var) X是否为整数() bool {
	return utils.IsInt(v.X取值())
}

// X是否为正整数检查`v`是否为uint类型。 md5:b5f59074f32c46ac
func (v *Var) X是否为正整数() bool {
	return utils.IsUint(v.X取值())
}

// X是否为小数 检查 `v` 是否为浮点类型。 md5:b61eead751ffcf77
func (v *Var) X是否为小数() bool {
	return utils.IsFloat(v.X取值())
}

// X是否为切片 检查 `v` 是否为切片类型。 md5:a71074d53c0be209
func (v *Var) X是否为切片() bool {
	return utils.IsSlice(v.X取值())
}

// X是否为Map 检查 `v` 是否为map类型。 md5:95b395f907d9b23f
func (v *Var) X是否为Map() bool {
	return utils.IsMap(v.X取值())
}

// X是否为结构 检查 `v` 是否为结构体类型。 md5:9b667cd3b80530d9
func (v *Var) X是否为结构() bool {
	return utils.IsStruct(v.X取值())
}
