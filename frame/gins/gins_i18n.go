// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins
import (
	"github.com/888go/goframe/i18n/gi18n"
	)
// I18n 返回一个 gi18n.Manager 的实例。
// 参数 `name` 是该实例的名称。
func I18n(name ...string) *gi18n.Manager {
	return gi18n.Instance(name...)
}
