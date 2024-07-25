// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gutil

import (
	"github.com/gogf/gf/v2/internal/deepcopy"
)

// Copy 返回 v 的深拷贝。
//
// Copy 无法复制结构体中的未导出字段（小写字段名）。未导出字段不能被 Go 运行时反射，因此无法进行任何数据拷贝。 md5:3460c0a1329d334a
func Copy(src interface{}) (dst interface{}) {
	return deepcopy.Copy(src)
}
