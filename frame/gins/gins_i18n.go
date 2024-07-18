// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins

import (
	"github.com/gogf/gf/v2/i18n/gi18n"
)

// I18n 返回一个 gi18n.Manager 的实例。
// 参数 `name` 是实例的名称。
// md5:cb8fb8e2c93c597b
// ff:
// name:
func I18n(name ...string) *gi18n.Manager {
	return gi18n.Instance(name...)
}
