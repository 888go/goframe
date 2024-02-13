// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类

import (
	"github.com/888go/goframe/util/gutil"
)

// ListItemValues 通过 key `key` 获取并返回所有项（item）结构体或映射中的元素。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，
// 否则将返回一个空切片。
func (v *Var) X取结构数组或Map数组值(名称 interface{}) (值s []interface{}) {
	return 工具类.ListItemValues(v.X取值(), 名称)
}

// ListItemValuesUnique 通过键 `key` 获取并返回所有结构体或映射中的唯一元素。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，
// 否则将返回一个空切片。
func (v *Var) X取结构数组或Map数组值并去重(名称 string) []interface{} {
	return 工具类.ListItemValuesUnique(v.X取值(), 名称)
}
