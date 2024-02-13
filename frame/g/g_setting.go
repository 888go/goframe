// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package g

import (
	"github.com/888go/goframe/internal/utils"
)

// SetDebug 手动启用/禁用 GoFrame 内部日志功能。
// 注意：此函数不是并发安全的，需要注意 DATA RACE（数据竞争）问题，
// 这意味着你应该在程序启动阶段调用此函数，而不要在运行时调用。
func X设置debug(开启 bool) {
	utils.SetDebugEnabled(开启)
}
