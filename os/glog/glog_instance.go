// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"github.com/888go/goframe/container/gmap"
)

const (
	// DefaultName 是用于实例使用的默认组名。
	DefaultName = "default"
)

var (
	// Instances map.
	instances = map类.X创建StrAny(true)
)

// Instance 返回一个具有默认设置的 Logger 实例。
// 参数 `name` 是该实例的名称。
func X取单例对象(名称 ...string) *Logger {
	key := DefaultName
	if len(名称) > 0 && 名称[0] != "" {
		key = 名称[0]
	}
	return instances.X取值或设置值_函数带锁(key, func() interface{} {
		return X创建()
	}).(*Logger)
}
