// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package glog

import "github.com/gogf/gf/v2/container/gmap"

const (
	// DefaultName是实例使用的默认组名。. md5:ca8d8295a8c531f4
	DefaultName = "default"
)

var (
	// Instances map.
	instances = gmap.NewStrAnyMap(true)
)

// Instance 返回一个具有默认设置的 Logger 实例。
// 参数 `name` 是该实例的名称。
// md5:ff5711e68bbfdd35
func Instance(name ...string) *Logger {
	key := DefaultName
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	return instances.GetOrSetFuncLock(key, func() interface{} {
		return New()
	}).(*Logger)
}
