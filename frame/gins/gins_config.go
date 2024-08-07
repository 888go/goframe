// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins

import (
	gcfg "github.com/888go/goframe/os/gcfg"
)

// Config返回一个具有默认设置的View实例。
// 参数`name`是实例的名称。
// md5:9ab5ade589362ad3
func Config(name ...string) *gcfg.Config {
	return gcfg.X取单例对象(name...)
}
