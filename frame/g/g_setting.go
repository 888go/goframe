// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package g

import (
	"github.com/888go/goframe/internal/utils"
)

// X设置debug 手动启用/禁用 GoFrame 内部日志。
// 注意，此函数不保证并发安全，可能会导致数据竞争（DATA RACE），因此应在启动时调用此函数，而不是在运行时。
// md5:c69ba0c4065e5d84
func X设置debug(开启 bool) {
	utils.SetDebugEnabled(开启)
}
