// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 定时cron类_test

import (
	"context"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gcron "github.com/888go/goframe/os/gcron"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestCron_Entry_Operations(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			cron  = gcron.New()
			array = garray.New(true)
		)
		cron.DelayAddTimes(ctx, 500*time.Millisecond, "* * * * * *", 2, func(ctx context.Context) {
			array.Append(1)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.Len(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(3000 * time.Millisecond)
		t.Assert(array.Len(), 2)
		t.Assert(cron.Size(), 0)
	})

	gtest.C(t, func(t *gtest.T) {
		var (
			cron  = gcron.New()
			array = garray.New(true)
		)
		entry, err1 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append(1)
		})
		t.Assert(err1, nil)
		t.Assert(array.Len(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.Len(), 1)
		t.Assert(cron.Size(), 1)
		entry.Stop()
		time.Sleep(5000 * time.Millisecond)
		t.Assert(array.Len(), 1)
		t.Assert(cron.Size(), 1)
		entry.Start()
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
		t.Assert(cron.Size(), 1)
		entry.Close()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(cron.Size(), 0)
	})
}
