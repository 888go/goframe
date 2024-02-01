// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins
import (
	"github.com/888go/goframe/os/gcfg"
	)
// Config 返回一个使用默认设置的 View 实例。
// 参数 `name` 是该实例的名称。
func Config(name ...string) *gcfg.Config {
	return gcfg.Instance(name...)
}
