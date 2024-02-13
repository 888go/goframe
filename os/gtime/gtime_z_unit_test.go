// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 时间类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func Test_TimestampStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertGT(len(时间类.X取文本时间戳毫秒()), 0)
		t.AssertGT(len(时间类.X取文本时间戳微秒()), 0)
		t.AssertGT(len(时间类.X取文本时间戳纳秒()), 0)
	})
}

func Test_Nanosecond(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		nanos := 时间类.X取时间戳纳秒()
		timeTemp := time.Unix(0, nanos)
		t.Assert(nanos, timeTemp.UnixNano())
	})
}

func Test_Microsecond(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		micros := 时间类.X取时间戳微秒()
		timeTemp := time.Unix(0, micros*1e3)
		t.Assert(micros, timeTemp.UnixNano()/1e3)
	})
}

func Test_Millisecond(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		millis := 时间类.X取时间戳毫秒()
		timeTemp := time.Unix(0, millis*1e6)
		t.Assert(millis, timeTemp.UnixNano()/1e6)
	})
}

func Test_Second(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := 时间类.X取时间戳秒()
		timeTemp := time.Unix(s, 0)
		t.Assert(s, timeTemp.Unix())
	})
}

func Test_Date(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(时间类.Date(), time.Now().Format("2006-01-02"))
	})
}

func Test_Datetime(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		datetime := 时间类.X取当前日期时间()
		timeTemp, err := 时间类.X转换文本(datetime, "Y-m-d H:i:s")
		if err != nil {
			t.Error("test fail")
		}
		t.Assert(datetime, timeTemp.Time.Format("2006-01-02 15:04:05"))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp, err := 时间类.X转换文本("")
		t.Assert(err, nil)
		t.AssertLT(timeTemp.Unix(), 0)
		timeTemp, err = 时间类.X转换文本("2006-01")
		t.AssertNE(err, nil)
		t.Assert(timeTemp, nil)
	})
}

func Test_ISO8601(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		iso8601 := 时间类.X取当前日期时间ISO8601()
		t.Assert(iso8601, 时间类.X创建并按当前时间().X取格式文本("c"))
	})
}

func Test_RFC822(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		rfc822 := 时间类.X取当前日期时间RFC822()
		t.Assert(rfc822, 时间类.X创建并按当前时间().X取格式文本("r"))
	})
}

func Test_StrToTime(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		// 正确的日期时间字符串。
		var testDateTimes = []string{
			"2006-01-02 15:04:05",
			"2006/01/02 15:04:05",
			"2006.01.02 15:04:05.000",
			"2006.01.02 - 15:04:05",
			"2006.01.02 15:04:05 +0800 CST",
			"2006-01-02T20:05:06+05:01:01",
			"2006-01-02T14:03:04Z01:01:01",
			"2006-01-02T15:04:05Z",
			"02-jan-2006 15:04:05",
			"02/jan/2006 15:04:05",
			"02.jan.2006 15:04:05",
			"02.jan.2006:15:04:05",
		}

		for _, item := range testDateTimes {
			timeTemp, err := 时间类.X转换文本(item)
			t.AssertNil(err)
			t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")
		}

		// 正确的日期字符串
		var testDates = []string{
			"2006.01.02",
			"2006.01.02 00:00",
			"2006.01.02 00:00:00.000",
		}

		for _, item := range testDates {
			timeTemp, err := 时间类.X转换文本(item)
			t.AssertNil(err)
			t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 00:00:00")
		}

		// 正确的时间字符串。
		var testTimes = g.MapStrStr{
			"16:12:01":     "15:04:05",
			"16:12:01.789": "15:04:05.000",
		}

		for k, v := range testTimes {
			time1, err := 时间类.X转换文本(k)
			t.AssertNil(err)
			time2, err := time.ParseInLocation(v, k, time.Local)
			t.AssertNil(err)
			t.Assert(time1.Time, time2)
		}

		// formatToStdLayout
		var testDateFormats = []string{
			"Y-m-d H:i:s",
			"\\T\\i\\m\\e Y-m-d H:i:s",
			"Y-m-d H:i:s\\",
			"Y-m-j G:i:s.u",
			"Y-m-j G:i:su",
		}

		var testDateFormatsResult = []string{
			"2007-01-02 15:04:05",
			"Time 2007-01-02 15:04:05",
			"2007-01-02 15:04:05",
			"2007-01-02 15:04:05.000",
			"2007-01-02 15:04:05.000",
		}

		for index, item := range testDateFormats {
			timeTemp, err := 时间类.X转换文本(testDateFormatsResult[index], item)
			if err != nil {
				t.Error("test fail")
			}
			t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05.000"), "2007-01-02 15:04:05.000")
		}

		// 异常日期列表
		var testDatesFail = []string{
			"2006.01",
			"06..02",
		}

		for _, item := range testDatesFail {
			_, err := 时间类.X转换文本(item)
			if err == nil {
				t.Error("test fail")
			}
		}

		// test err
		_, err := 时间类.X转换文本("2006-01-02 15:04:05", "aabbccdd")
		if err == nil {
			t.Error("test fail")
		}
	})
}

func Test_ConvertZone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		// 现行时间
		nowUTC := time.Now().UTC()
		testZone := "America/Los_Angeles"

		// 转换为洛杉矶时间
		t1, err := 时间类.X转换时区(nowUTC.Format("2006-01-02 15:04:05"), testZone, "")
		if err != nil {
			t.Error("test fail")
		}

		// 使用洛杉矶时区解析上面转换后的时间
		laStr := t1.Time.Format("2006-01-02 15:04:05")
		loc, err := time.LoadLocation(testZone)
		t2, err := time.ParseInLocation("2006-01-02 15:04:05", laStr, loc)

		// 判断是否与现行时间匹配
		t.Assert(t2.UTC().Unix(), nowUTC.Unix())

	})

	// test err
	单元测试类.C(t, func(t *单元测试类.T) {
		// 现行时间
		nowUTC := time.Now().UTC()
		// t.Log(nowUTC.Unix()) // 输出当前UTC时间的Unix时间戳（单位为秒）
		testZone := "errZone"

		// 错误时间输入
		_, err := 时间类.X转换时区(nowUTC.Format("06..02 15:04:05"), testZone, "")
		if err == nil {
			t.Error("test fail")
		}
		// 错误时区输入
		_, err = 时间类.X转换时区(nowUTC.Format("2006-01-02 15:04:05"), testZone, "")
		if err == nil {
			t.Error("test fail")
		}
		// 错误时区输入
		_, err = 时间类.X转换时区(nowUTC.Format("2006-01-02 15:04:05"), testZone, testZone)
		if err == nil {
			t.Error("test fail")
		}
	})
}

func Test_ParseDuration(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		d, err := 时间类.X文本取时长("1d")
		t.AssertNil(err)
		t.Assert(d.String(), "24h0m0s")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		d, err := 时间类.X文本取时长("1d2h3m")
		t.AssertNil(err)
		t.Assert(d.String(), "26h3m0s")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		d, err := 时间类.X文本取时长("-1d2h3m")
		t.AssertNil(err)
		t.Assert(d.String(), "-26h3m0s")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		d, err := 时间类.X文本取时长("3m")
		t.AssertNil(err)
		t.Assert(d.String(), "3m0s")
	})
	// error
	单元测试类.C(t, func(t *单元测试类.T) {
		d, err := 时间类.X文本取时长("-1dd2h3m")
		t.AssertNE(err, nil)
		t.Assert(d.String(), "0s")
	})
}

func Test_ParseTimeFromContent(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		timeTemp := 时间类.X解析文本("我是中文2006-01-02 15:04:05我也是中文", "Y-m-d H:i:s")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		timeTemp1 := 时间类.X解析文本("我是中文2006-01-02 15:04:05我也是中文")
		t.Assert(timeTemp1.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		timeTemp2 := 时间类.X解析文本("我是中文02.jan.2006 15:04:05我也是中文")
		t.Assert(timeTemp2.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		// test err
		timeTempErr := 时间类.X解析文本("我是中文", "Y-m-d H:i:s")
		if timeTempErr != nil {
			t.Error("test fail")
		}
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		timeStr := "2021-1-27 9:10:24"
		t.Assert(时间类.X解析文本(timeStr, "Y-n-d g:i:s").String(), "2021-01-27 09:10:24")
	})
}

func Test_FuncCost(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		时间类.X取函数执行时长(func() {

		})
	})
}
