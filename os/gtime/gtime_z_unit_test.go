// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 时间类_test

import (
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_TimestampStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertGT(len(gtime.X取文本时间戳毫秒()), 0)
		t.AssertGT(len(gtime.X取文本时间戳微秒()), 0)
		t.AssertGT(len(gtime.X取文本时间戳纳秒()), 0)
	})
}

func Test_Nanosecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		nanos := gtime.X取时间戳纳秒()
		timeTemp := time.Unix(0, nanos)
		t.Assert(nanos, timeTemp.UnixNano())
	})
}

func Test_Microsecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		micros := gtime.X取时间戳微秒()
		timeTemp := time.Unix(0, micros*1e3)
		t.Assert(micros, timeTemp.UnixNano()/1e3)
	})
}

func Test_Millisecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		millis := gtime.X取时间戳毫秒()
		timeTemp := time.Unix(0, millis*1e6)
		t.Assert(millis, timeTemp.UnixNano()/1e6)
	})
}

func Test_Second(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gtime.X取时间戳秒()
		timeTemp := time.Unix(s, 0)
		t.Assert(s, timeTemp.Unix())
	})
}

func Test_Date(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gtime.Date(), time.Now().Format("2006-01-02"))
	})
}

func Test_Datetime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		datetime := gtime.X取当前日期时间()
		timeTemp, err := gtime.X转换文本(datetime, "Y-m-d H:i:s")
		if err != nil {
			t.Error("test fail")
		}
		t.Assert(datetime, timeTemp.Time.Format("2006-01-02 15:04:05"))
	})
	gtest.C(t, func(t *gtest.T) {
		timeTemp, err := gtime.X转换文本("")
		t.AssertNil(err)
		t.AssertLT(timeTemp.Unix(), 0)
		timeTemp, err = gtime.X转换文本("2006-01")
		t.AssertNE(err, nil)
		t.Assert(timeTemp, nil)
	})
}

func Test_ISO8601(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		iso8601 := gtime.X取当前日期时间ISO8601()
		t.Assert(iso8601, gtime.X创建并按当前时间().X取格式文本("c"))
	})
}

func Test_RFC822(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		rfc822 := gtime.X取当前日期时间RFC822()
		t.Assert(rfc822, gtime.X创建并按当前时间().X取格式文本("r"))
	})
}

func Test_StrToTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
				// 正确的日期时间字符串。 md5:30202903e049484f
		var testDateTimes = []string{
			"2006-01-02 15:04:05",
			"2006/01/02 15:04:05",
			"2006.01.02 15:04:05.000",
			"2006.01.02 - 15:04:05",
			"2006.01.02 15:04:05 +0800 CST",
			"2006-01-02T12:05:05+05:01",
			"2006-01-02T02:03:05-05:01",
			"2006-01-02T15:04:05",
			"02-jan-2006 15:04:05",
			"02/jan/2006 15:04:05",
			"02.jan.2006 15:04:05",
			"02.jan.2006:15:04:05",
		}
				// 保存先前的时区. md5:e98b311301c3985e
		local := *time.Local
		defer func() {
			*time.Local = local
		}()
		time.Local = time.FixedZone("Asia/Shanghai", 8*3600)

		for i, item := range testDateTimes {
			timeTemp, err := gtime.X转换文本(item)
			t.AssertNil(err)
			t.Log(i)
			t.Assert(timeTemp.Time.Local().Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")
		}

		// Correct date string,.
		var testDates = []string{
			"2006.01.02",
			"2006.01.02 00:00",
			"2006.01.02 00:00:00.000",
		}

		for _, item := range testDates {
			timeTemp, err := gtime.X转换文本(item)
			t.AssertNil(err)
			t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 00:00:00")
		}

		// Correct time string.
		var testTimes = g.MapStrStr{
			"16:12:01":     "15:04:05",
			"16:12:01.789": "15:04:05.000",
		}

		for k, v := range testTimes {
			time1, err := gtime.X转换文本(k)
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
			timeTemp, err := gtime.X转换文本(testDateFormatsResult[index], item)
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
			_, err := gtime.X转换文本(item)
			if err == nil {
				t.Error("test fail")
			}
		}

		// test err
		_, err := gtime.X转换文本("2006-01-02 15:04:05", "aabbccdd")
		if err == nil {
			t.Error("test fail")
		}
	})
}

func Test_ConvertZone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 现行时间
		nowUTC := time.Now().UTC()
		testZone := "America/Los_Angeles"

		// 转换为洛杉矶时间
		t1, err := gtime.X转换时区(nowUTC.Format("2006-01-02 15:04:05"), testZone, "")
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
	gtest.C(t, func(t *gtest.T) {
		// 现行时间
		nowUTC := time.Now().UTC()
		// t.Log(nowUTC.Unix())
		testZone := "errZone"

		// 错误时间输入
		_, err := gtime.X转换时区(nowUTC.Format("06..02 15:04:05"), testZone, "")
		if err == nil {
			t.Error("test fail")
		}
		// 错误时区输入
		_, err = gtime.X转换时区(nowUTC.Format("2006-01-02 15:04:05"), testZone, "")
		if err == nil {
			t.Error("test fail")
		}
		// 错误时区输入
		_, err = gtime.X转换时区(nowUTC.Format("2006-01-02 15:04:05"), testZone, testZone)
		if err == nil {
			t.Error("test fail")
		}
	})
}

func Test_ParseDuration(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		d, err := gtime.X文本取时长("1d")
		t.AssertNil(err)
		t.Assert(d.String(), "24h0m0s")
	})
	gtest.C(t, func(t *gtest.T) {
		d, err := gtime.X文本取时长("1d2h3m")
		t.AssertNil(err)
		t.Assert(d.String(), "26h3m0s")
	})
	gtest.C(t, func(t *gtest.T) {
		d, err := gtime.X文本取时长("-1d2h3m")
		t.AssertNil(err)
		t.Assert(d.String(), "-26h3m0s")
	})
	gtest.C(t, func(t *gtest.T) {
		d, err := gtime.X文本取时长("3m")
		t.AssertNil(err)
		t.Assert(d.String(), "3m0s")
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		d, err := gtime.X文本取时长("-1dd2h3m")
		t.AssertNE(err, nil)
		t.Assert(d.String(), "0s")
	})
}

func Test_ParseTimeFromContent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timeTemp := gtime.X解析文本("我是中文2006-01-02 15:04:05我也是中文", "Y-m-d H:i:s")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		timeTemp1 := gtime.X解析文本("我是中文2006-01-02 15:04:05我也是中文")
		t.Assert(timeTemp1.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		timeTemp2 := gtime.X解析文本("我是中文02.jan.2006 15:04:05我也是中文")
		t.Assert(timeTemp2.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		// test err
		timeTempErr := gtime.X解析文本("我是中文", "Y-m-d H:i:s")
		if timeTempErr != nil {
			t.Error("test fail")
		}
	})

	gtest.C(t, func(t *gtest.T) {
		timeStr := "2021-1-27 9:10:24"
		t.Assert(gtime.X解析文本(timeStr, "Y-n-d g:i:s").String(), "2021-01-27 09:10:24")
	})
}

func Test_FuncCost(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gtime.X取函数执行时长(func() {

		})
	})
}
