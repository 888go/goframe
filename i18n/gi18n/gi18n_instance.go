// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gi18n
import (
	"github.com/888go/goframe/container/gmap"
	)
const (
	// DefaultName 是用于实例使用的默认组名。
	DefaultName = "default"
)

var (
// instances 是用于管理的实例映射，
// 通过名称来管理多个 i18n 实例。
	instances = gmap.NewStrAnyMap(true)
)

// Instance 返回一个 Resource 类型的实例。
// 参数 `name` 是该实例的名称。
func Instance(name ...string) *Manager {
	key := DefaultName
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	return instances.GetOrSetFuncLock(key, func() interface{} {
		return New()
	}).(*Manager)
}
