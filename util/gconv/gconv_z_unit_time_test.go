// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Time(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X取时长(""), time.Duration(int64(0)))
		t.AssertEQ(转换类.X取gtime时间类(""), 时间类.X创建())
		t.AssertEQ(转换类.X取gtime时间类(nil), nil)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		s := "2011-10-10 01:02:03.456"
		t.AssertEQ(转换类.X取gtime时间类(s), 时间类.X创建并从文本(s))
		t.AssertEQ(转换类.X取时间(nil), time.Time{})
		t.AssertEQ(转换类.X取时间(s), 时间类.X创建并从文本(s).Time)
		t.AssertEQ(转换类.X取时长(100), 100*time.Nanosecond)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		s := "01:02:03.456"
		t.AssertEQ(转换类.X取gtime时间类(s).Hour(), 1)
		t.AssertEQ(转换类.X取gtime时间类(s).Minute(), 2)
		t.AssertEQ(转换类.X取gtime时间类(s).X取秒(), 3)
		t.AssertEQ(转换类.X取gtime时间类(s), 时间类.X创建并从文本(s))
		t.AssertEQ(转换类.X取时间(s), 时间类.X创建并从文本(s).Time)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		s := "0000-01-01 01:02:03"
		t.AssertEQ(转换类.X取gtime时间类(s).Year(), 0)
		t.AssertEQ(转换类.X取gtime时间类(s).X取月份(), 1)
		t.AssertEQ(转换类.X取gtime时间类(s).Day(), 1)
		t.AssertEQ(转换类.X取gtime时间类(s).Hour(), 1)
		t.AssertEQ(转换类.X取gtime时间类(s).Minute(), 2)
		t.AssertEQ(转换类.X取gtime时间类(s).X取秒(), 3)
		t.AssertEQ(转换类.X取gtime时间类(s), 时间类.X创建并从文本(s))
		t.AssertEQ(转换类.X取时间(s), 时间类.X创建并从文本(s).Time)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t1 := 时间类.X创建并从文本("2021-05-21 05:04:51.206547+00")
		t2 := 转换类.X取gtime时间类(泛型类.X创建(t1))
		t3 := 泛型类.X创建(t1).X取gtime时间类()
		t.AssertEQ(t1, t2)
		t.AssertEQ(t1.X取本地时区(), t2.X取本地时区())
		t.AssertEQ(t1, t3)
		t.AssertEQ(t1.X取本地时区(), t3.X取本地时区())
	})
}

func Test_Time_Slice_Attribute(t *testing.T) {
	type SelectReq struct {
		Arr []*时间类.Time
		One *时间类.Time
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var s *SelectReq
		err := 转换类.Struct(g.Map{
			"arr": g.Slice别名{"2021-01-12 12:34:56", "2021-01-12 12:34:57"},
			"one": "2021-01-12 12:34:58",
		}, &s)
		t.AssertNil(err)
		t.Assert(s.One, "2021-01-12 12:34:58")
		t.Assert(s.Arr[0], "2021-01-12 12:34:56")
		t.Assert(s.Arr[1], "2021-01-12 12:34:57")
	})
}

func Test_Issue2901(t *testing.T) {
	type GameApp2 struct {
		ForceUpdateTime *time.Time
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		src := map[string]interface{}{
			"FORCE_UPDATE_TIME": time.Now(),
		}
		m := GameApp2{}
		err := 转换类.Scan(src, &m)
		t.AssertNil(err)
	})
}
