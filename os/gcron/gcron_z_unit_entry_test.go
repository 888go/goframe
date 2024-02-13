// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时cron类_test

import (
	"context"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/gcron"
	"github.com/888go/goframe/test/gtest"
)

func TestCron_Entry_Operations(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			cron  = 定时cron类.New()
			array = 数组类.X创建(true)
		)
		cron.DelayAddTimes(ctx, 500*time.Millisecond, "* * * * * *", 2, func(ctx context.Context) {
			array.Append别名(1)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(3000 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		t.Assert(cron.Size(), 0)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			cron  = 定时cron类.New()
			array = 数组类.X创建(true)
		)
		entry, err1 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append别名(1)
		})
		t.Assert(err1, nil)
		t.Assert(array.X取长度(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		t.Assert(cron.Size(), 1)
		entry.Stop()
		time.Sleep(5000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		t.Assert(cron.Size(), 1)
		entry.Start()
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		t.Assert(cron.Size(), 1)
		entry.Close()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(cron.Size(), 0)
	})
}
