// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时类

import (
	"time"
)

// loop 使用独立的 goroutine 启动ticker。
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
				// 检查定时器状态。
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

// proceed 函数执行定时任务的检查和运行逻辑。
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
		// 它检查是否满足滴答（ticks）的要求。
		if jobNextTicks := entry.nextTicks.Val(); currentTimerTicks < jobNextTicks {
			// 如果当前的ticks数未达到其运行所需的ticks要求，则将该任务重新推回。
			t.queue.Push(entry, entry.nextTicks.Val())
			break
		}
		// 它检查作业运行需求，然后进行异步运行。
		entry.doCheckAndRunByTicks(currentTimerTicks)
		// 状态检查：将其推回或忽略。
		if entry.X取任务状态() != StatusClosed {
			// 它将任务推回到队列中以便下次运行。
			t.queue.Push(entry, entry.nextTicks.Val())
		}
	}
}
