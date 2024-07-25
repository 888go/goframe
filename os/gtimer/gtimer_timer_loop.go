// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gtimer

import "time"

// 使用独立的goroutine启动ticker。 md5:ce88d07fc9fbc1b3
func (t *Timer) loop() {
	go func() {
		var (
			currentTimerTicks   int64
			timerIntervalTicker = time.NewTicker(t.options.Interval)
		)
		defer timerIntervalTicker.Stop()
		for {
			select {
			case <-timerIntervalTicker.C:
				// 检查计时器状态。 md5:3099f00b46caf3a0
				switch t.status.Val() {
				case StatusRunning:
					// Timer proceeding.
					if currentTimerTicks = t.ticks.Add(1); currentTimerTicks >= t.queue.NextPriority() {
						t.proceed(currentTimerTicks)
					}

				case StatusStopped:
					// Do nothing.

				case StatusClosed:
					// Timer exits.
					return
				}
			}
		}
	}()
}

// proceed 函数负责处理定时任务的检查和运行逻辑。 md5:ddd19ba4b91ab76a
func (t *Timer) proceed(currentTimerTicks int64) {
	var (
		value interface{}
	)
	for {
		value = t.queue.Pop()
		if value == nil {
			break
		}
		entry := value.(*Entry)
		// 它检查是否满足ticks的需求。 md5:a861e677e940bac4
		if jobNextTicks := entry.nextTicks.Val(); currentTimerTicks < jobNextTicks {
			// 如果当前的ticks不满足作业的运行ticks要求，它会将作业推回。 md5:cdd563e1c262dcfa
			t.queue.Push(entry, entry.nextTicks.Val())
			break
		}
		// 它检查任务运行要求，然后进行异步运行。 md5:a7927f4887033102
		entry.doCheckAndRunByTicks(currentTimerTicks)
		// 状态检查：推回还是忽略它。 md5:16de30b2fa2279cf
		if entry.Status() != StatusClosed {
			// 它将任务重新推入队列以供下次运行。 md5:c330359de56328e4
			t.queue.Push(entry, entry.nextTicks.Val())
		}
	}
}
