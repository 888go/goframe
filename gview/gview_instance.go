// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 模板类

import (
	"github.com/gogf/gf/v2/container/gmap"
)

const (
	// DefaultName 是用于实例使用的默认组名称。
	DefaultName = "default"
)

var (
	// Instances map.
	instances = gmap.NewStrAnyMap(true)
)

// Instance 返回一个使用默认设置的 View 实例。
// 参数 `name` 是该实例的名称。
func Instance(name ...string) *View {
	key := DefaultName
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	return instances.GetOrSetFuncLock(key, func() interface{} {
		return New()
	}).(*View)
}
