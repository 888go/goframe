// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 缓存类

import (
	gtime "github.com/888go/goframe/os/gtime"
)

// IsExpired 检查 `item` 是否已过期。 md5:9e46a52b25ea5be9
func (item *adapterMemoryItem) IsExpired() bool {
	// 注意这里应该使用大于或等于的判断
	// 想象一下缓存时间只有1毫秒。
	// md5:9d0401593ecbe5eb

	return item.e < gtime.TimestampMilli()
}
