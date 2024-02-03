// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcron

import (
	"context"
	"time"
	
	"github.com/888go/goframe/internal/intlog"
)

// getAndUpdateLastTimestamp 检查并更新修复时间戳，返回在某些秒数内具有延迟修复的最后一个时间戳。
func (s *cronSchedule) getAndUpdateLastTimestamp(ctx context.Context, t time.Time) int64 {
	var (
		currentTimestamp = t.Unix()
		lastTimestamp    = s.lastTimestamp.Val()
	)
	switch {
	case
		lastTimestamp == currentTimestamp:
		lastTimestamp += 1

	case
		lastTimestamp == currentTimestamp-1:
		lastTimestamp = currentTimestamp

	case
		lastTimestamp == currentTimestamp-2,
		lastTimestamp == currentTimestamp-3:
		lastTimestamp += 1

	default:
		// 延迟过长，让我们将最后的时间戳更新为当前时间戳。
		intlog.Printf(
			ctx,
			`too much delay, last timestamp "%d", current "%d"`,
			lastTimestamp, currentTimestamp,
		)
		lastTimestamp = currentTimestamp
	}
	s.lastTimestamp.Set(lastTimestamp)
	return lastTimestamp
}
