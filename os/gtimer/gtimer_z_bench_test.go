// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtimer

import (
	"context"
	"testing"
	"time"
)

var (
	ctx   = context.TODO()
	timer = New()
)

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timer.Add(ctx, time.Hour, func(ctx context.Context) {

		})
	}
}

func Benchmark_PriorityQueue_Pop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timer.queue.Pop()
	}
}

func Benchmark_StartStop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timer.Start()
		timer.Stop()
	}
}
