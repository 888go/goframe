// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 时间类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/internal/json"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gutil "github.com/888go/goframe/util/gutil"
)

func Test_New(t *testing.T) {
	// time.Time
	gtest.C(t, func(t *gtest.T) {
		timeNow := time.Now()
		timeTemp := gtime.X创建(timeNow)
		t.Assert(timeTemp.Time.UnixNano(), timeNow.UnixNano())

		timeTemp1 := gtime.X创建()
		t.Assert(timeTemp1.Time, time.Time{})
	})
	// string
	gtest.C(t, func(t *gtest.T) {
		timeNow := gtime.X创建并按当前时间()
		timeTemp := gtime.X创建(timeNow.String())
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), timeNow.Time.Format("2006-01-02 15:04:05"))
	})
	gtest.C(t, func(t *gtest.T) {
		timeNow := gtime.X创建并按当前时间()
		timeTemp := gtime.X创建(timeNow.X取文本时间戳微秒())
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), timeNow.Time.Format("2006-01-02 15:04:05"))
	})
	// int64
	gtest.C(t, func(t *gtest.T) {
		timeNow := gtime.X创建并按当前时间()
		timeTemp := gtime.X创建(timeNow.X取时间戳微秒())
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), timeNow.Time.Format("2006-01-02 15:04:05"))
	})
	// short datetime.
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建("2021-2-9 08:01:21")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")

		timeTemp = gtime.X创建("2021-02-09 08:01:21", []byte("Y-m-d H:i:s"))
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")

		timeTemp = gtime.X创建([]byte("2021-02-09 08:01:21"))
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")

		timeTemp = gtime.X创建([]byte("2021-02-09 08:01:21"), "Y-m-d H:i:s")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")

		timeTemp = gtime.X创建([]byte("2021-02-09 08:01:21"), []byte("Y-m-d H:i:s"))
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2021-02-09 08:01:21")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2021-02-09 08:01:21")
	})
	//
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gtime.X创建(gtime.Time{}), nil)
		t.Assert(gtime.X创建(&gtime.Time{}), nil)
	})

	// unconventional
	gtest.C(t, func(t *gtest.T) {

		var testUnconventionalDates = []string{
			"2006-01.02",
			"2006.01-02",
		}

		for _, item := range testUnconventionalDates {
			timeTemp := gtime.X创建(item)
			t.Assert(timeTemp.X取时间戳毫秒(), 0)
			t.Assert(timeTemp.X取文本时间戳毫秒(), "")
			t.Assert(timeTemp.String(), "")
		}
	})
}

func Test_Nil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var t1 *gtime.Time
		t.Assert(t1.String(), "")
	})
	gtest.C(t, func(t *gtest.T) {
		var t1 gtime.Time
		t.Assert(t1.String(), "")
	})
}

func Test_NewFromStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2006-01-02 15:04:05")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := gtime.X创建并从文本("2006.0102")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})
}

func Test_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t1 := gtime.X创建并从文本("2006-01-02 15:04:05")
		t.Assert(t1.String(), "2006-01-02 15:04:05")
		t.Assert(fmt.Sprintf("%s", t1), "2006-01-02 15:04:05")

		t2 := *t1
		t.Assert(t2.String(), "2006-01-02 15:04:05")
		t.Assert(fmt.Sprintf("{%s}", t2.String()), "{2006-01-02 15:04:05}")
	})
}

func Test_NewFromStrFormat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按给定格式文本("2006-01-02 15:04:05", "Y-m-d H:i:s")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := gtime.X创建并按给定格式文本("2006-01-02 15:04:05", "aabbcc")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})

	gtest.C(t, func(t *gtest.T) {
		t1 := gtime.X创建并按给定格式文本("2019/2/1", "Y/n/j")
		t.Assert(t1.X取格式文本("Y-m-d"), "2019-02-01")

		t2 := gtime.X创建并按给定格式文本("2019/10/12", "Y/n/j")
		t.Assert(t2.X取格式文本("Y-m-d"), "2019-10-12")
	})
}

func Test_NewFromStrLayout(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按Layout格式文本("2006-01-02 15:04:05", "2006-01-02 15:04:05")
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := gtime.X创建并按Layout格式文本("2006-01-02 15:04:05", "aabbcc")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})
}

func Test_NewFromTimeStamp(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从时间戳(1554459846000)
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2019-04-05 18:24:06")
		timeTemp1 := gtime.X创建并从时间戳(0)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "0001-01-01 00:00:00")
		timeTemp2 := gtime.X创建并从时间戳(155445984)
		t.Assert(timeTemp2.X取格式文本("Y-m-d H:i:s"), "1974-12-05 11:26:24")
	})
}

func Test_Time_Second(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		t.Assert(timeTemp.X取秒(), timeTemp.Time.Second())
	})
}

func Test_Time_IsZero(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ti *gtime.Time = nil
		t.Assert(ti.IsZero(), true)
	})
}

func Test_Time_AddStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gt := gtime.X创建("2018-08-08 08:08:08")
		gt1, err := gt.X增加文本时长("10T")
		t.Assert(gt1, nil)
		t.AssertNE(err, nil)
	})
}

func Test_Time_Equal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var t1 *gtime.Time = nil
		var t2 = gtime.X创建()
		t.Assert(t1.X是否相等(t2), false)
		t.Assert(t1.X是否相等(t1), true)
		t.Assert(t2.X是否相等(t1), false)
	})
}

func Test_Time_After(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var t1 *gtime.Time = nil
		var t2 = gtime.X创建()
		t.Assert(t1.X是否之后(t2), false)
		t.Assert(t2.X是否之后(t1), true)
	})
}

func Test_Time_Sub(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var t1 *gtime.Time = nil
		var t2 = gtime.X创建()
		t.Assert(t1.X取纳秒时长(t2), time.Duration(0))
		t.Assert(t2.X取纳秒时长(t1), time.Duration(0))
	})
}

func Test_Time_Nanosecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		t.Assert(timeTemp.X取纳秒(), timeTemp.Time.Nanosecond())
	})
}

func Test_Time_Microsecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		t.Assert(timeTemp.X取微秒(), timeTemp.Time.Nanosecond()/1e3)
	})
}

func Test_Time_Millisecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		t.Assert(timeTemp.X取毫秒(), timeTemp.Time.Nanosecond()/1e6)
	})
}

func Test_Time_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		t.Assert(timeTemp.String(), timeTemp.Time.Format("2006-01-02 15:04:05"))
	})
}

func Test_Time_ISO8601(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		now := gtime.X创建并按当前时间()
		t.Assert(now.X取文本时间ISO8601(), now.X取格式文本("c"))
	})
}

func Test_Time_RFC822(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		now := gtime.X创建并按当前时间()
		t.Assert(now.X取文本时间RFC822(), now.X取格式文本("r"))
	})
}

func Test_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		timeTemp1 := timeTemp.X取副本()
		t.Assert(timeTemp.Time.Unix(), timeTemp1.Time.Unix())
	})
}

func Test_ToTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		t.Assert(timeTemp.Time.UnixNano(), timeTemp1.UnixNano())
	})
}

func Test_Add(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2006-01-02 15:04:05")
		timeTemp = timeTemp.X增加时长(time.Second)
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2006-01-02 15:04:06")
	})
}

func Test_ToZone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
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
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2006-01-02 15:04:05")
		timeTemp = timeTemp.X增加时间(1, 2, 3)
		t.Assert(timeTemp.X取格式文本("Y-m-d H:i:s"), "2007-03-05 15:04:05")
	})
}

func Test_UTC(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		timeTemp.X取UTC时区()
		t.Assert(timeTemp.UnixNano(), timeTemp1.UTC().UnixNano())
	})
}

func Test_Local(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		timeTemp.X取本地时区()
		t.Assert(timeTemp.UnixNano(), timeTemp1.Local().UnixNano())
	})
}

func Test_Round(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		timeTemp = timeTemp.X向上舍入(time.Hour)
		t.Assert(timeTemp.UnixNano(), timeTemp1.Round(time.Hour).UnixNano())
	})
}

func Test_Truncate(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并按当前时间()
		timeTemp1 := timeTemp.Time
		timeTemp = timeTemp.X向下舍入(time.Hour)
		t.Assert(timeTemp.UnixNano(), timeTemp1.Truncate(time.Hour).UnixNano())
	})
}

func Test_StartOfMinute(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本忽略秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-12 18:24:00")
	})
}

func Test_EndOfMinute(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 18:24:59.000")
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 18:24:59.999")
	})
}

func Test_StartOfHour(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本忽略分钟秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-12 18:00:00")
	})
}

func Test_EndOfHour(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 18:59:59.000")
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 18:59:59.999")
	})
}

func Test_StartOfDay(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本忽略小时分钟秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-12 00:00:00")
	})
}

func Test_EndOfDay(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 23:59:59.000")
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 23:59:59.999")
	})
}

func Test_StartOfWeek(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本周第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-06 00:00:00")
	})
}

func Test_EndOfWeek(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本周末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 23:59:59.000")
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本周末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-12 23:59:59.999")
	})
}

func Test_StartOfMonth(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本月第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-12-01 00:00:00")
	})
}

func Test_EndOfMonth(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本月末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.000")
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-12 18:24:06")
		timeTemp1 := timeTemp.X取副本月末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.999")
	})
}

func Test_StartOfQuarter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本季度第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-10-01 00:00:00")
	})
}

func Test_EndOfQuarter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本季末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.000")
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本季末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.999")
	})
}

func Test_StartOfHalf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本半年第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-07-01 00:00:00")
	})
}

func Test_EndOfHalf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本半年末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.000")
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本半年末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.999")
	})
}

func Test_StartOfYear(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本年第一天()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s"), "2020-01-01 00:00:00")
	})
}

func Test_EndOfYear(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本年末23点59分59秒()
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.000")
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X创建并从文本("2020-12-06 18:24:06")
		timeTemp1 := timeTemp.X取副本年末23点59分59秒(true)
		t.Assert(timeTemp1.X取格式文本("Y-m-d H:i:s.u"), "2020-12-31 23:59:59.999")
	})
}

func Test_OnlyTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		obj := gtime.X创建并从文本("18:24:06")
		t.Assert(obj.String(), "18:24:06")
	})
}

func Test_DeepCopy(t *testing.T) {
	type User struct {
		Id          int
		CreatedTime *gtime.Time
	}
	gtest.C(t, func(t *gtest.T) {
		u1 := &User{
			Id:          1,
			CreatedTime: gtime.X创建("2022-03-08T03:01:14+08:00"),
		}
		u2 := gutil.X深拷贝(u1).(*User)
		t.Assert(u1, u2)
	})
	// nil attribute.
	gtest.C(t, func(t *gtest.T) {
		u1 := &User{}
		u2 := gutil.X深拷贝(u1).(*User)
		t.Assert(u1, u2)
	})
	gtest.C(t, func(t *gtest.T) {
		var t1 *gtime.Time = nil
		t.Assert(t1.DeepCopy(), nil)
	})
}

func Test_UnmarshalJSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var t1 gtime.Time
		t.AssertNE(json.Unmarshal([]byte("{}"), &t1), nil)
	})
}
