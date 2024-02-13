// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 缓存类

import (
	"github.com/888go/goframe/os/gtime"
)

// IsExpired 检查 `item` 是否已过期。
func (item *adapterMemoryItem) IsExpired() bool {
// 注意这里应当使用大于或等于的判断
// 假设缓存时间仅为1毫秒

	return item.e < 时间类.X取时间戳毫秒()
}
