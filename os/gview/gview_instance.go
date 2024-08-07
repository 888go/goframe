// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 模板类

import (
	gmap "github.com/888go/goframe/container/gmap"
)

const (
		// DefaultName是实例使用的默认组名。 md5:ca8d8295a8c531f4
	DefaultName = "default"
)

var (
	// Instances map.
	instances = gmap.X创建StrAny(true)
)

// Instance 返回一个具有默认设置的 View 实例。
// 参数 `name` 是该实例的名称。
// md5:7ac19173e044fa8f
func Instance(name ...string) *View {
	key := DefaultName
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	return instances.X取值或设置值_函数带锁(key, func() interface{} {
		return New()
	}).(*View)
}
