// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 定时类

import (
	"context"
	"testing"
	"time"
)

var (
	ctx   = context.TODO()
	timer = X创建()
)

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timer.X加入循环任务(ctx, time.Hour, func(ctx context.Context) {

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
		timer.X开始工作()
		timer.X暂停工作()
	}
}
