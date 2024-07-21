// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcron

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/internal/intlog"
)

// getAndUpdateLastCheckTimestamp 检查、修复并返回在几秒钟内有延迟修复的最后时间戳。 md5:617d53ed6d0eee3f
func (s *cronSchedule) getAndUpdateLastCheckTimestamp(ctx context.Context, t time.Time) int64 {
	var (
		currentTimestamp   = t.Unix()
		lastCheckTimestamp = s.lastCheckTimestamp.Val()
	)
	switch {
// 通常情况下，定时器在同一秒内触发，但毫秒数不同。
// 例如：
// lastCheckTimestamp: 2024-03-26 19:47:34.000
// currentTimestamp:   2024-03-26 19:47:34.999
// md5:7ad3ec347d1a6583
	case
		lastCheckTimestamp == currentTimestamp:
		lastCheckTimestamp += 1

// 经常发生的情况，没有延迟。
// 示例：
// lastCheckTimestamp: 2024年03月26日 19时47分34秒.000
// currentTimestamp:   2024年03月26日 19时47分35秒.000
// md5:1ed300ef7b928611
	case
		lastCheckTimestamp == currentTimestamp-1:
		lastCheckTimestamp = currentTimestamp

	// 可容忍的延迟时间为3秒。
	// 例如：
	// lastCheckTimestamp: 2024-03-26 19:47:31.000、2024-03-26 19:47:32.000
	// currentTimestamp:   2024-03-26 19:47:34.000
	// md5:21934a048bbfddaf
	case
		lastCheckTimestamp == currentTimestamp-2,
		lastCheckTimestamp == currentTimestamp-3:
		lastCheckTimestamp += 1

	// 延迟太多，它忽略了修复，定时任务可能不会被触发。 md5:5d550b27269fafbf
	default:
		// 延迟时间过长，让我们将最后的timestamp更新为当前时间。 md5:7b051dda466c96cf
		intlog.Printf(
			ctx,
			`too much latency, last timestamp "%d", current "%d", latency "%d"`,
			lastCheckTimestamp, currentTimestamp, currentTimestamp-lastCheckTimestamp,
		)
		lastCheckTimestamp = currentTimestamp
	}
	s.lastCheckTimestamp.Set(lastCheckTimestamp)
	return lastCheckTimestamp
}
