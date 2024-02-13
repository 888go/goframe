// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 时间类_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gutil"
)

func Test_New(t *testing.T) {
	// time.Time
	单元测试类.C(t, func(t *单元测试类.T) {
		timeNow := time.Now()
		timeTemp := 时间类.X创建(timeNow)
		t.Assert(timeTemp.Time.UnixNano(), timeNow.UnixNano())

		timeTemp1 := 时间类.X创建()
		t.Assert(timeTemp1.Time, time.Time{})
	})
	// string
	单元测试类.C(t, func(t *单元测试类.T) {
		timeNow := 时间类.X创建并按当前时间()
		timeTemp := 时间类.X创建(timeNow.String())
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), timeNow.Time.Format("2006-01-02 15:04:05"))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeNow := 时间类.X创建并按当前时间()
		timeTemp := 时间类.X创建(timeNow.X取文本时间戳微秒())
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), timeNow.Time.Format("2006-01-02 15:04:05"))
	})
	// int64
	单元测试类.C(t, func(t *单元测试类.T) {
		timeNow := 时间类.X创建并按当前时间()
		timeTemp := 时间类.X创建(timeNow.X取时间戳微秒())
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), timeNow.Time.Format("2006-01-02 15:04:05"))
	})
	// short datetime.
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建("2021-2-9 08:01:21")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")

		timeTemp = 时间类.X创建("2021-02-09 08:01:21", []byte("Y-m-d H:i:s"))
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")

		timeTemp = 时间类.X创建([]byte("2021-02-09 08:01:21"))
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")

		timeTemp = 时间类.X创建([]byte("2021-02-09 08:01:21"), "Y-m-d H:i:s")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")

		timeTemp = 时间类.X创建([]byte("2021-02-09 08:01:21"), []byte("Y-m-d H:i:s"))
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")
	})
	//
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(时间类.X创建(时间类.Time{}), nil)
		t.Assert(时间类.X创建(&时间类.Time{}), nil)
	})

	// unconventional
	单元测试类.C(t, func(t *单元测试类.T) {

		var testUnconventionalDates = []string{
			"2006-01.02",
			"2006.01-02",
		}

		for _, item := range testUnconventionalDates {
			timeTemp := 时间类.X创建(item)
			t.Assert(timeTemp.X取时间戳毫秒(), 0)
			t.Assert(timeTemp.X取文本时间戳毫秒(), "")
			t.Assert(timeTemp.String(), "")
		}
	})
}

func Test_Nil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var t1 *时间类.Time
		t.Assert(t1.String(), "")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var t1 时间类.Time
		t.Assert(t1.String(), "")
	})
}

func Test_NewFromStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2006-01-02 15:04:05")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := 时间类.X创建并从文本("2006.0102")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})
}

func Test_String(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t1 := 时间类.X创建并从文本("2006-01-02 15:04:05")
		t.Assert(t1.String(), "2006-01-02 15:04:05")
		t.Assert(fmt.Sprintf("%s", t1), "2006-01-02 15:04:05")

		t2 := *t1
		t.Assert(t2.String(), "2006-01-02 15:04:05")
		t.Assert(fmt.Sprintf("{%s}", t2.String()), "{2006-01-02 15:04:05}")
	})
}

func Test_NewFromStrFormat(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按给定格式文本("2006-01-02 15:04:05", "Y-m-d H:i:s")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := 时间类.X创建并按给定格式文本("2006-01-02 15:04:05", "aabbcc")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t1 := 时间类.X创建并按给定格式文本("2019/2/1", "Y/n/j")
		t.Assert(t1.X取格式文本("Y-m-d"), "2019-02-01")

		t2 := 时间类.X创建并按给定格式文本("2019/10/12", "Y/n/j")
		t.Assert(t2.X取格式文本("Y-m-d"), "2019-10-12")
	})
}

func Test_NewFromStrLayout(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按Layout格式文本("2006-01-02 15:04:05", "2006-01-02 15:04:05")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := 时间类.X创建并按Layout格式文本("2006-01-02 15:04:05", "aabbcc")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})
}

func Test_NewFromTimeStamp(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从时间戳(1554459846000)
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2019-04-05 18:24:06")
		timeTemp1 := 时间类.X创建并从时间戳(0)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "0001-01-01 00:00:00")
		timeTemp2 := 时间类.X创建并从时间戳(155445984)
		t.Assert(timeTemp2.X取格式文本("Y-m-d H:i:s"), "1974-12-05 11:26:24")
	})
}

func Test_Time_Second(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		t.Assert(timeTemp.X取秒(), timeTemp.Time.Second())
	})
}

func Test_Time_IsZero(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var ti *时间类.Time = nil
		t.Assert(ti.IsZero(), true)
	})
}

func Test_Time_AddStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		gt := 时间类.X创建("2018-08-08 08:08:08")
		gt1, err := gt.X增加文本时长("10T")
		t.Assert(gt1, nil)
		t.AssertNE(err, nil)
	})
}

func Test_Time_Equal(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var t1 *时间类.Time = nil
		var t2 = 时间类.X创建()
		t.Assert(t1.X是否相等(t2), false)
		t.Assert(t1.X是否相等(t1), true)
		t.Assert(t2.X是否相等(t1), false)
	})
}

func Test_Time_After(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var t1 *时间类.Time = nil
		var t2 = 时间类.X创建()
		t.Assert(t1.X是否之后(t2), false)
		t.Assert(t2.X是否之后(t1), true)
	})
}

func Test_Time_Sub(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var t1 *时间类.Time = nil
		var t2 = 时间类.X创建()
		t.Assert(t1.X取纳秒时长(t2), time.Duration(0))
		t.Assert(t2.X取纳秒时长(t1), time.Duration(0))
	})
}

func Test_Time_Nanosecond(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		t.Assert(timeTemp.X取纳秒(), timeTemp.Time.Nanosecond())
	})
}

func Test_Time_Microsecond(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		t.Assert(timeTemp.X取微秒(), timeTemp.Time.Nanosecond()/1e3)
	})
}

func Test_Time_Millisecond(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		t.Assert(timeTemp.X取毫秒(), timeTemp.Time.Nanosecond()/1e6)
	})
}

func Test_Time_String(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		t.Assert(timeTemp.String(), timeTemp.Time.Format("2006-01-02 15:04:05"))
	})
}

func Test_Time_ISO8601(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		now := 时间类.X创建并按当前时间()
		t.Assert(now.X取文本时间ISO8601(), now.X取格式文本("c"))
	})
}

func Test_Time_RFC822(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		now := 时间类.X创建并按当前时间()
		t.Assert(now.X取文本时间RFC822(), now.X取格式文本("r"))
	})
}

func Test_Clone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		timeTemp1 := timeTemp.X取副本()
		t.Assert(timeTemp.Time.Unix(), timeTemp1.Time.Unix())
	})
}

func Test_ToTime(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		t.Assert(timeTemp.Time.UnixNano(), timeTemp1.UnixNano())
	})
}

func Test_Add(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2006-01-02 15:04:05")
		timeTemp = timeTemp.X增加时长(time.Second)
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2006-01-02 15:04:06")
	})
}

func Test_ToZone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		timeTemp, _ = timeTemp.X转换时区("America/Los_Angeles")
		t.Assert(timeTemp.Time.Location().String(), "America/Los_Angeles")

		loc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			t.Error("test fail")
		}
		timeTemp = timeTemp.X转换时区Location(loc)
		t.Assert(timeTemp.Time.Location().String(), "Asia/Shanghai")

		timeTemp1, _ := timeTemp.X转换时区("errZone")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})
}

func Test_AddDate(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2006-01-02 15:04:05")
		timeTemp = timeTemp.X增加时间(1, 2, 3)
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2007-03-05 15:04:05")
	})
}

func Test_UTC(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		timeTemp.X取UTC时区()
		t.Assert(timeTemp.UnixNano(), timeTemp1.UTC().UnixNano())
	})
}

func Test_Local(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		timeTemp.X取本地时区()
		t.Assert(timeTemp.UnixNano(), timeTemp1.Local().UnixNano())
	})
}

func Test_Round(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		timeTemp = timeTemp.X向上舍入(time.Hour)
		t.Assert(timeTemp.UnixNano(), timeTemp1.Round(time.Hour).UnixNano())
	})
}

func Test_Truncate(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		timeTemp = timeTemp.X向下舍入(time.Hour)
		t.Assert(timeTemp.UnixNano(), timeTemp1.Truncate(time.Hour).UnixNano())
	})
}

func Test_StartOfMinute(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本忽略秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-12 18:24:00")
	})
}

func Test_EndOfMinute(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 18:24:59.000")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 18:24:59.999")
	})
}

func Test_StartOfHour(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本忽略分钟秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-12 18:00:00")
	})
}

func Test_EndOfHour(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 18:59:59.000")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 18:59:59.999")
	})
}

func Test_StartOfDay(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本忽略小时分钟秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-12 00:00:00")
	})
}

func Test_EndOfDay(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 23:59:59.000")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 23:59:59.999")
	})
}

func Test_StartOfWeek(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本周第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-06 00:00:00")
	})
}

func Test_EndOfWeek(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本周末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 23:59:59.000")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本周末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 23:59:59.999")
	})
}

func Test_StartOfMonth(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本月第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-01 00:00:00")
	})
}

func Test_EndOfMonth(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本月末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.000")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本月末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.999")
	})
}

func Test_StartOfQuarter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本季度第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-10-01 00:00:00")
	})
}

func Test_EndOfQuarter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本季末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.000")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本季末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.999")
	})
}

func Test_StartOfHalf(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本半年第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-07-01 00:00:00")
	})
}

func Test_EndOfHalf(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本半年末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.000")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本半年末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.999")
	})
}

func Test_StartOfYear(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本年第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-01-01 00:00:00")
	})
}

func Test_EndOfYear(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本年末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.000")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本年末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.999")
	})
}

func Test_OnlyTime(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		obj := 时间类.X创建并从文本("18:24:06")
		t.Assert(obj.String(), "18:24:06")
	})
}

func Test_DeepCopy(t *testing.T) {
	type User struct {
		Id          int
		CreatedTime *时间类.Time
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		u1 := &User{
			Id:          1,
			CreatedTime: 时间类.X创建("2022-03-08T03:01:14+08:00"),
		}
		u2 := 工具类.X深拷贝(u1).(*User)
		t.Assert(u1, u2)
	})
	// nil attribute.
	单元测试类.C(t, func(t *单元测试类.T) {
		u1 := &User{}
		u2 := 工具类.X深拷贝(u1).(*User)
		t.Assert(u1, u2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var t1 *时间类.Time = nil
		t.Assert(t1.DeepCopy(), nil)
	})
}

func Test_UnmarshalJSON(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var t1 时间类.Time
		t.AssertNE(json.Unmarshal([]byte("{}"), &t1), nil)
	})
}
