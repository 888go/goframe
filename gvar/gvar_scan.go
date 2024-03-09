// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// Scan 会自动检测 `pointer` 的类型，并将 `params` 转换为 `pointer`。它支持对 `pointer`
// 进行转换的类型包括：`*map`、`*[]map`、`*[]*map`、`*struct`、`**struct`、`*[]struct` 和 `*[]*struct`。
//
// 详情请参阅 gconv.Scan。
func (v *Var) X取结构体指针(结构体指针 interface{}, 名称映射 ...map[string]string) error {
	return gconv.Scan(v.X取值(), 结构体指针, 名称映射...)
}
