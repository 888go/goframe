// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcron_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	ctx = context.TODO()
)

func TestCron_Add_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		_, err1 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			g.Log().Print(ctx, "cron1")
			array.Append(1)
		})
		_, err2 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			g.Log().Print(ctx, "cron2")
			array.Append(1)
		}, "test")
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(cron.Size(), 2)
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.Len(), 4)
		cron.Close()
		time.Sleep(1300 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.Len(), fixedLength)
	})
}

func TestCron_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		cron.Add(ctx, "* * * * * *", func(ctx context.Context) {}, "add")
		// fmt.Println("开始", time.Now()). md5:518c980e118bdbbf
		cron.DelayAdd(ctx, time.Second, "* * * * * *", func(ctx context.Context) {}, "delay_add")
		t.Assert(cron.Size(), 1)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(cron.Size(), 2)

		cron.Remove("delay_add")
		t.Assert(cron.Size(), 1)

		entry1 := cron.Search("add")
		entry2 := cron.Search("test-none")
		t.AssertNE(entry1, nil)
		t.Assert(entry2, nil)
	})

	// test @ error
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		defer cron.Close()
		_, err := cron.Add(ctx, "@aaa", func(ctx context.Context) {}, "add")
		t.AssertNE(err, nil)
	})

	// test @every error
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		defer cron.Close()
		_, err := cron.Add(ctx, "@every xxx", func(ctx context.Context) {}, "add")
		t.AssertNE(err, nil)
	})
}

func TestCron_Remove(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append(1)
		}, "add")
		t.Assert(array.Len(), 0)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)

		cron.Remove("add")
		t.Assert(array.Len(), 1)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestCron_Add_FixedPattern(t *testing.T) {
	for i := 0; i < 5; i++ {
		doTestCronAddFixedPattern(t)
	}
}

func doTestCronAddFixedPattern(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			now    = time.Now()
			cron   = gcron.New()
			array  = garray.New(true)
			expect = now.Add(time.Second * 2)
		)
		defer cron.Close()

		var pattern = fmt.Sprintf(
			`%d %d %d %d %d %s`,
			expect.Second(), expect.Minute(), expect.Hour(), expect.Day(), expect.Month(), expect.Weekday().String(),
		)
		cron.SetLogger(g.Log())
		g.Log().Debugf(ctx, `pattern: %s`, pattern)
		_, err := cron.Add(ctx, pattern, func(ctx context.Context) {
			array.Append(1)
		})
		t.AssertNil(err)
		time.Sleep(3000 * time.Millisecond)
		g.Log().Debug(ctx, `current time`)
		t.Assert(array.Len(), 1)
	})
}

func TestCron_AddSingleton(t *testing.T) {
							// 未使用，可以删除. md5:7acdb570a3213eac
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		cron.Add(ctx, "* * * * * *", func(ctx context.Context) {}, "add")
		cron.DelayAdd(ctx, time.Second, "* * * * * *", func(ctx context.Context) {}, "delay_add")
		t.Assert(cron.Size(), 1)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(cron.Size(), 2)

		cron.Remove("delay_add")
		t.Assert(cron.Size(), 1)

		entry1 := cron.Search("add")
		entry2 := cron.Search("test-none")
		t.AssertNE(entry1, nil)
		t.Assert(entry2, nil)
	})
	// keep this
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		cron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append(1)
			time.Sleep(50 * time.Second)
		})
		t.Assert(cron.Size(), 1)
		time.Sleep(3500 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})

}

func TestCron_AddOnce1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		cron.AddOnce(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append(1)
		})
		cron.AddOnce(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append(1)
		})
		t.Assert(cron.Size(), 2)
		time.Sleep(2500 * time.Millisecond)
		t.Assert(array.Len(), 2)
		t.Assert(cron.Size(), 0)
	})
}

func TestCron_AddOnce2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		cron.AddOnce(ctx, "@every 2s", func(ctx context.Context) {
			array.Append(1)
		})
		t.Assert(cron.Size(), 1)
		time.Sleep(3000 * time.Millisecond)
		t.Assert(array.Len(), 1)
		t.Assert(cron.Size(), 0)
	})
}

func TestCron_AddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		_, _ = cron.AddTimes(ctx, "* * * * * *", 2, func(ctx context.Context) {
			array.Append(1)
		})
		time.Sleep(3500 * time.Millisecond)
		t.Assert(array.Len(), 2)
		t.Assert(cron.Size(), 0)
	})
}

func TestCron_DelayAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		cron.DelayAdd(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
			array.Append(1)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.Len(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
		t.Assert(cron.Size(), 1)
	})
}

func TestCron_DelayAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		cron.DelayAddSingleton(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(2200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		t.Assert(cron.Size(), 1)
	})
}

func TestCron_DelayAddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
		cron.DelayAddOnce(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
			array.Append(1)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.Len(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(2200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		t.Assert(cron.Size(), 0)
	})
}

func TestCron_DelayAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cron := gcron.New()
		array := garray.New(true)
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
}
