// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gins

import (
	"github.com/gogf/gf/v2/os/gres"
)

// Resource 返回一个 Resource 类型的实例。
// 参数 `name` 为该实例的名称。 md5:42e664c4b3a2bb54
func Resource(name ...string) *gres.Resource {
	return gres.Instance(name...)
}
