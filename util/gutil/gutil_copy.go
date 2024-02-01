// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gutil
import (
	"github.com/888go/goframe/internal/deepcopy"
	)
// Copy 返回v的深度拷贝。
//
// Copy无法复制结构体中的未导出字段（小写字段名）。
// 未导出字段不能被Go运行时反射，因此无法执行任何数据拷贝操作。
func Copy(src interface{}) (dst interface{}) {
	return deepcopy.Copy(src)
}
