// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时cron类_test

import (
	"context"
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcron"
	"github.com/888go/goframe/test/gtest"
)

var (
	ctx = context.TODO()
)

func TestCron_Add_Close(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		_, err1 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			g.X日志类().X输出(ctx, "cron1")
			array.Append别名(1)
		})
		_, err2 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			g.X日志类().X输出(ctx, "cron2")
			array.Append别名(1)
		}, "test")
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(cron.Size(), 2)
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.X取长度(), 4)
		cron.Close()
		time.Sleep(1300 * time.Millisecond)
		fixedLength := array.X取长度()
		time.Sleep(1300 * time.Millisecond)
		t.Assert(array.X取长度(), fixedLength)
	})
}

func TestCron_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		cron.Add(ctx, "* * * * * *", func(ctx context.Context) {}, "add")
		// 打印输出 "start" 及当前时间（使用 time.Now() 函数获取）
// ```go
// fmt.Println("start", time.Now())
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
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		defer cron.Close()
		_, err := cron.Add(ctx, "@aaa", func(ctx context.Context) {}, "add")
		t.AssertNE(err, nil)
	})

	// test @every error
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		defer cron.Close()
		_, err := cron.Add(ctx, "@every xxx", func(ctx context.Context) {}, "add")
		t.AssertNE(err, nil)
	})
}

func TestCron_Remove(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append别名(1)
		}, "add")
		t.Assert(array.X取长度(), 0)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)

		cron.Remove("add")
		t.Assert(array.X取长度(), 1)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})
}

func TestCron_Add_FixedPattern(t *testing.T) {
	for i := 0; i < 5; i++ {
		doTestCronAddFixedPattern(t)
	}
}

func doTestCronAddFixedPattern(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			now    = time.Now()
			cron   = 定时cron类.New()
			array  = 数组类.X创建(true)
			expect = now.Add(time.Second * 2)
		)
		defer cron.Close()

		var pattern = fmt.Sprintf(
			`%d %d %d %d %d %s`,
			expect.Second(), expect.Minute(), expect.Hour(), expect.Day(), expect.Month(), expect.Weekday().String(),
		)
		cron.SetLogger(g.X日志类())
		g.X日志类().X输出并格式化DEBU(ctx, `pattern: %s`, pattern)
		_, err := cron.Add(ctx, pattern, func(ctx context.Context) {
			array.Append别名(1)
		})
		t.AssertNil(err)
		time.Sleep(3000 * time.Millisecond)
		g.X日志类().X输出DEBU(ctx, `current time`)
		t.Assert(array.X取长度(), 1)
	})
}

func TestCron_AddSingleton(t *testing.T) {
	// 未使用，可以移除
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
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
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		cron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(50 * time.Second)
		})
		t.Assert(cron.Size(), 1)
		time.Sleep(3500 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})

}

func TestCron_AddOnce1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		cron.AddOnce(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append别名(1)
		})
		cron.AddOnce(ctx, "* * * * * *", func(ctx context.Context) {
			array.Append别名(1)
		})
		t.Assert(cron.Size(), 2)
		time.Sleep(2500 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		t.Assert(cron.Size(), 0)
	})
}

func TestCron_AddOnce2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		cron.AddOnce(ctx, "@every 2s", func(ctx context.Context) {
			array.Append别名(1)
		})
		t.Assert(cron.Size(), 1)
		time.Sleep(3000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		t.Assert(cron.Size(), 0)
	})
}

func TestCron_AddTimes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		_, _ = cron.AddTimes(ctx, "* * * * * *", 2, func(ctx context.Context) {
			array.Append别名(1)
		})
		time.Sleep(3500 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		t.Assert(cron.Size(), 0)
	})
}

func TestCron_DelayAdd(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		cron.DelayAdd(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
			array.Append别名(1)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		t.Assert(cron.Size(), 1)
	})
}

func TestCron_DelayAddSingleton(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		cron.DelayAddSingleton(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
			array.Append别名(1)
			time.Sleep(10 * time.Second)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(2200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		t.Assert(cron.Size(), 1)
	})
}

func TestCron_DelayAddOnce(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
		cron.DelayAddOnce(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
			array.Append别名(1)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.X取长度(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(2200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		t.Assert(cron.Size(), 0)
	})
}

func TestCron_DelayAddTimes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cron := 定时cron类.New()
		array := 数组类.X创建(true)
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
}
