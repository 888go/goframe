// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 时间类_test

import (
	"fmt"

	gtime "github.com/888go/goframe/os/gtime"
)

// New 创建并返回一个带有给定参数的时间对象。
// 可选参数可以是：time.Time（时间类型）、string（字符串）或 integer（整数）。
// md5:6bd8b77e3852fdf3
func ExampleSetTimeZone() {
	gtime.X设置时区("Asia/Shanghai")
	fmt.Println(gtime.X取当前日期时间())
	// May Output:
	// 2018-08-08 08:08:08
}

func ExampleTimestamp() {
	fmt.Println(gtime.X取时间戳秒())

	// May Output:
	// 1636359252
}

func ExampleTimestampMilli() {
	fmt.Println(gtime.X取时间戳毫秒())

	// May Output:
	// 1636359252000
}

func ExampleTimestampMicro() {
	fmt.Println(gtime.X取时间戳微秒())

	// May Output:
	// 1636359252000000
}

func ExampleTimestampNano() {
	fmt.Println(gtime.X取时间戳纳秒())

	// May Output:
	// 1636359252000000000
}

func ExampleTimestampStr() {
	fmt.Println(gtime.X取文本时间戳秒())

	// May Output:
	// 1636359252
}

func ExampleDate() {
	fmt.Println(gtime.Date())

	// May Output:
	// 2006-01-02
}

func ExampleDatetime() {
	fmt.Println(gtime.X取当前日期时间())

	// May Output:
	// 2006-01-02 15:04:05
}

func ExampleISO8601() {
	fmt.Println(gtime.X取当前日期时间ISO8601())

	// May Output:
	// 2006-01-02T15:04:05-07:00
}

func ExampleRFC822() {
	fmt.Println(gtime.X取当前日期时间RFC822())

	// May Output:
	// Mon, 02 Jan 06 15:04 MST
}

func ExampleStrToTime() {
	res, _ := gtime.X转换文本("2006-01-02T15:04:05-07:00", "Y-m-d H:i:s")
	fmt.Println(res)

	// May Output:
	// 2006-01-02 15:04:05
}

func ExampleConvertZone() {
	res, _ := gtime.X转换时区("2006-01-02 15:04:05", "Asia/Tokyo", "Asia/Shanghai")
	fmt.Println(res)

	// Output:
	// 2006-01-02 16:04:05
}

func ExampleStrToTimeFormat() {
	res, _ := gtime.StrToTimeFormat别名("2006-01-02 15:04:05", "Y-m-d H:i:s")
	fmt.Println(res)

	// Output:
	// 2006-01-02 15:04:05
}

func ExampleStrToTimeLayout() {
	res, _ := gtime.X转换文本Layout("2018-08-08", "2006-01-02")
	fmt.Println(res)

	// Output:
	// 2018-08-08 00:00:00
}

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h", "1d" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h", "d".
//
// Very note that it supports unit "d" more than function time.ParseDuration.
func ExampleParseDuration() {
	res, _ := gtime.X文本取时长("+10h")
	fmt.Println(res)

	// Output:
	// 10h0m0s
}

func ExampleTime_Format() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取格式文本("Y-m-d"))
	fmt.Println(gt1.X取格式文本("l"))
	fmt.Println(gt1.X取格式文本("F j, Y, g:i a"))
	fmt.Println(gt1.X取格式文本("j, n, Y"))
	fmt.Println(gt1.X取格式文本("h-i-s, j-m-y, it is w Day z"))
	fmt.Println(gt1.X取格式文本("D M j G:i:s T Y"))

	// Output:
	// 2018-08-08
	// Wednesday
	// August 8, 2018, 8:08 am
	// 8, 8, 2018
	// 08-08-08, 8-08-18, 0831 0808 3 Wedam18 219
	// Wed Aug 8 8:08:08 CST 2018
}

func ExampleTime_FormatNew() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X按格式取副本("Y-m-d"))
	fmt.Println(gt1.X按格式取副本("Y-m-d H:i"))

	// Output:
	// 2018-08-08 00:00:00
	// 2018-08-08 08:08:00
}

func ExampleTime_FormatTo() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X格式设置("Y-m-d"))

	// Output:
	// 2018-08-08 00:00:00
}

func ExampleTime_Layout() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取Layout格式文本("2006-01-02"))

	// Output:
	// 2018-08-08
}

func ExampleTime_LayoutNew() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本并按Layout格式("2006-01-02"))

	// Output:
	// 2018-08-08 00:00:00
}

func ExampleTime_LayoutTo() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X设置Layout格式("2006-01-02"))

	// Output:
	// 2018-08-08 00:00:00
}

func ExampleTime_IsLeapYear() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X是否为闰年())

	// Output:
	// false
}

func ExampleTime_DayOfYear() {
	gt1 := gtime.X创建("2018-01-08 08:08:08")

	fmt.Println(gt1.X取全年第几天())

	// Output:
	// 7
}

// DaysInMonth 返回当前月份的天数。 md5:0cd1f14a8bb1f8fc
func ExampleTime_DaysInMonth() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取当前月份总天数())

	// Output:
	// 31
}

// WeeksOfYear 返回当前年份中的当前周数。 md5:a51898ffdc6f00df
func ExampleTime_WeeksOfYear() {
	gt1 := gtime.X创建("2018-01-08 08:08:08")

	fmt.Println(gt1.X取全年第几星期())

	// Output:
	// 2
}

func ExampleTime_ToZone() {
	gt1 := gtime.X创建并按当前时间()
	gt2, _ := gt1.X转换时区("Asia/Shanghai")
	gt3, _ := gt1.X转换时区("Asia/Tokyo")

	fmt.Println(gt2)
	fmt.Println(gt3)

	// May Output:
	// 2021-11-11 17:10:10
	// 2021-11-11 18:10:10
}
