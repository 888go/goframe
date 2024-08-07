// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类_test

import (
	"testing"
	"time"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_Time(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.X取时长(""), time.Duration(int64(0)))
		t.AssertEQ(gconv.X取gtime时间类(""), gtime.X创建())
		t.AssertEQ(gconv.X取gtime时间类(nil), nil)
	})

	gtest.C(t, func(t *gtest.T) {
		s := "2011-10-10 01:02:03.456"
		t.AssertEQ(gconv.X取gtime时间类(s), gtime.X创建并从文本(s))
		t.AssertEQ(gconv.X取时间(nil), time.Time{})
		t.AssertEQ(gconv.X取时间(s), gtime.X创建并从文本(s).Time)
		t.AssertEQ(gconv.X取时长(100), 100*time.Nanosecond)
	})
	gtest.C(t, func(t *gtest.T) {
		s := "01:02:03.456"
		t.AssertEQ(gconv.X取gtime时间类(s).Hour(), 1)
		t.AssertEQ(gconv.X取gtime时间类(s).Minute(), 2)
		t.AssertEQ(gconv.X取gtime时间类(s).X取秒(), 3)
		t.AssertEQ(gconv.X取gtime时间类(s), gtime.X创建并从文本(s))
		t.AssertEQ(gconv.X取时间(s), gtime.X创建并从文本(s).Time)
	})
	gtest.C(t, func(t *gtest.T) {
		s := "0000-01-01 01:02:03"
		t.AssertEQ(gconv.X取gtime时间类(s).Year(), 0)
		t.AssertEQ(gconv.X取gtime时间类(s).X取月份(), 1)
		t.AssertEQ(gconv.X取gtime时间类(s).Day(), 1)
		t.AssertEQ(gconv.X取gtime时间类(s).Hour(), 1)
		t.AssertEQ(gconv.X取gtime时间类(s).Minute(), 2)
		t.AssertEQ(gconv.X取gtime时间类(s).X取秒(), 3)
		t.AssertEQ(gconv.X取gtime时间类(s), gtime.X创建并从文本(s))
		t.AssertEQ(gconv.X取时间(s), gtime.X创建并从文本(s).Time)
	})
	gtest.C(t, func(t *gtest.T) {
		t1 := gtime.X创建并从文本("2021-05-21 05:04:51.206547+00")
		t2 := gconv.X取gtime时间类(gvar.X创建(t1))
		t3 := gvar.X创建(t1).X取gtime时间类()
		t.AssertEQ(t1, t2)
		t.AssertEQ(t1.X取本地时区(), t2.X取本地时区())
		t.AssertEQ(t1, t3)
		t.AssertEQ(t1.X取本地时区(), t3.X取本地时区())
	})
}

func Test_Time_Slice_Attribute(t *testing.T) {
	type SelectReq struct {
		Arr []*gtime.Time
		One *gtime.Time
	}
	gtest.C(t, func(t *gtest.T) {
		var s *SelectReq
		err := gconv.Struct(g.Map{
			"arr": g.Slice别名{"2021-01-12 12:34:56", "2021-01-12 12:34:57"},
			"one": "2021-01-12 12:34:58",
		}, &s)
		t.AssertNil(err)
		t.Assert(s.One, "2021-01-12 12:34:58")
		t.Assert(s.Arr[0], "2021-01-12 12:34:56")
		t.Assert(s.Arr[1], "2021-01-12 12:34:57")
	})
}
