// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 时间类_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	gtime "github.com/888go/goframe/os/gtime"
)

func ExampleNew_Basic() {
	curTime := "2018-08-08 08:08:08"
	timer, _ := time.Parse("2006-01-02 15:04:05", curTime)
	t1 := gtime.X创建(&timer)
	t2 := gtime.X创建(curTime)
	t3 := gtime.X创建(curTime, "Y-m-d H:i:s")
	t4 := gtime.X创建(curTime)
	t5 := gtime.X创建(1533686888)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)
	fmt.Println(t5)

	// Output:
	// 2018-08-08 08:08:08
	// 2018-08-08 08:08:08
	// 2018-08-08 08:08:08
	// 2018-08-08 08:08:08
	// 2018-08-08 08:08:08
}

func ExampleNew_WithFormat() {
	fmt.Println(gtime.X创建("20220629133225", "YmdHis").X取格式文本("Y-m-d H:i:s"))

	// Output:
	// 2022-06-29 13:32:25
}

// 现在创建并返回一个表示当前时间的对象。 md5:1cfc3114797b1f98
func ExampleNow() {
	t := gtime.X创建并按当前时间()
	fmt.Println(t)

	// May Output:
	// 2021-11-06 13:41:08
}

// NewFromTime 根据给定的time.Time对象创建并返回一个Time对象。 md5:e1cf178ea024f53b
func ExampleNewFromTime() {
	timer, _ := time.Parse("2006-01-02 15:04:05", "2018-08-08 08:08:08")
	nTime := gtime.X创建并按Time(timer)

	fmt.Println(nTime)

	// Output:
	// 2018-08-08 08:08:08
}

// NewFromStr 根据给定的字符串创建并返回一个 Time 对象。
// 注意，如果发生错误，它将返回 nil。
// md5:4687b38a27582a12
func ExampleNewFromStr() {
	t := gtime.X创建并从文本("2018-08-08 08:08:08")

	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:08
}

// NewFromStrFormat 通过给定的字符串和自定义格式（如：Y-m-d H:i:s）创建并返回一个Time对象。
// 注意，如果发生错误，它将返回nil。
// md5:ed9966a0a8156f1d
func ExampleNewFromStrFormat() {
	t := gtime.X创建并按给定格式文本("2018-08-08 08:08:08", "Y-m-d H:i:s")
	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:08
}

// NewFromStrLayout 根据给定的字符串和标准库格式（如：2006-01-02 15:04:05）创建并返回一个Time对象。
// 注意，如果出现错误，它将返回nil。
// md5:027f4d0876baa1a8
func ExampleNewFromStrLayout() {
	t := gtime.X创建并按Layout格式文本("2018-08-08 08:08:08", "2006-01-02 15:04:05")
	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:08
}

// NewFromTimeStamp 根据给定的时间戳创建并返回一个 Time 对象，
// 该时间戳可以是秒到纳秒的精度。
// 例如：1600443866 和 1600443866199266000 都被视为有效的时间戳数值。
// md5:6a84edd691c97a4f
func ExampleNewFromTimeStamp() {
	t1 := gtime.X创建并从时间戳(1533686888)
	t2 := gtime.X创建并从时间戳(1533686888000)

	fmt.Println(t1.String() == t2.String())
	fmt.Println(t1)

	// Output:
	// true
	// 2018-08-08 08:08:08
}

// Timestamp 返回时间戳，以秒为单位。 md5:52f3b8b0088c2fab
func ExampleTime_Timestamp() {
	t := gtime.X取时间戳秒()

	fmt.Println(t)

	// May output:
	// 1533686888
}

// Timestamp 返回以毫秒为单位的时间戳。 md5:b4836efd766d4f28
func ExampleTime_TimestampMilli() {
	t := gtime.X取时间戳毫秒()

	fmt.Println(t)

	// May output:
	// 1533686888000
}

// Timestamp 返回时间戳，以微秒为单位。 md5:92d47303429ab4d0
func ExampleTime_TimestampMicro() {
	t := gtime.X取时间戳微秒()

	fmt.Println(t)

	// May output:
	// 1533686888000000
}

// Timestamp 返回纳秒级的时间戳。 md5:5f8d54218fb362c4
func ExampleTime_TimestampNano() {
	t := gtime.X取时间戳纳秒()

	fmt.Println(t)

	// May output:
	// 1533686888000000
}

// TimestampStr 是一个方便的方法，它获取并返回时间戳（以秒为单位）的字符串形式。
// md5:f638769b91eb1dd5
func ExampleTime_TimestampStr() {
	t := gtime.X取文本时间戳秒()

	fmt.Println(reflect.TypeOf(t))

	// Output:
	// string
}

// Month 返回指定时间t的月份。 md5:84f113a801a5eb29
func ExampleTime_Month() {
	gt := gtime.X创建("2018-08-08 08:08:08")
	t1 := gt.X取月份()

	fmt.Println(t1)

	// Output:
	// 8
}

// Second返回由t指定的分钟内的第二个偏移量，范围在[0, 59]之间。
// md5:5666ae5cbf21989d
func ExampleTime_Second() {
	gt := gtime.X创建("2018-08-08 08:08:08")
	t1 := gt.X取秒()

	fmt.Println(t1)

	// Output:
	// 8
}

// String 返回当前时间对象作为字符串。 md5:4f5a1f3896ca049d
func ExampleTime_String() {
	gt := gtime.X创建("2018-08-08 08:08:08")
	t1 := gt.String()

	fmt.Println(t1)
	fmt.Println(reflect.TypeOf(t1))

	// Output:
	// 2018-08-08 08:08:08
	// string
}

// IsZero报告是否`t`表示零时间点，即UTC时间的1970年1月1日00:00:00。
// md5:4e2b46d4fa63a878
func ExampleTime_IsZero() {
	gt := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt.IsZero())

	// Output:
	// false
}

// Add 将持续时间添加到当前时间。 md5:8a845aeaaa064af4
func ExampleTime_Add() {
	gt := gtime.X创建("2018-08-08 08:08:08")
	gt1 := gt.X增加时长(time.Duration(10) * time.Second)

	fmt.Println(gt1)

	// Output:
	// 2018-08-08 08:08:18
}

// AddStr parses the given duration as string and adds it to current time.
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ExampleTime_AddStr() {
	gt := gtime.X创建("2018-08-08 08:08:08")
	gt1, _ := gt.X增加文本时长("10s")

	fmt.Println(gt1)

	// Output:
	// 2018-08-08 08:08:18
}

// AddDate 向时间添加年、月和日。 md5:643cfbc24c5bd938
func ExampleTime_AddDate() {
	var (
		year  = 1
		month = 2
		day   = 3
	)
	gt := gtime.X创建("2018-08-08 08:08:08")
	gt = gt.X增加时间(year, month, day)

	fmt.Println(gt)

	// Output:
	// 2019-10-11 08:08:08
}

// Round 返回将 t 四舍五入到 d 的倍数的结果（从零时间开始）。对于半等值，四舍五入行为向上取整。
// 如果 d 小于等于 0，Round 会返回 t 并移除任何单调时钟读数，但保持不变。
//
// Round 以绝对的自零时间以来的时间段进行操作；它不处理时间的呈现形式。因此，Round(Hour) 可能返回一个非零分钟的时间，具体取决于时间的 Location。
// md5:b2557220790fc058
func ExampleTime_Round() {
	gt := gtime.X创建("2018-08-08 08:08:08")
	t := gt.X向上舍入(time.Duration(10) * time.Second)

	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:10
}

// Truncate 返回将时间t向下舍入到d的倍数的结果（从零时间开始）。
// 如果d<=0，Truncate会返回t，但去除任何单调时钟读数，否则保持不变。
//
// Truncate是基于时间从零时间点起的绝对持续时间来进行操作的；
// 它并不作用于时间的展示形式。因此，Truncate(Hour)可能返回一个分钟数非零的时间，
// 这取决于该时间的位置信息（Location）。
// md5:f72e0e00b245e691
func ExampleTime_Truncate() {
	gt := gtime.X创建("2018-08-08 08:08:08")
	t := gt.X向下舍入(time.Duration(10) * time.Second)

	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:00
}

// Equal 函数报告 t 和 u 是否表示相同的时刻。
// 即使两个时间在不同的时区，它们也可以相等。
// 例如，CEST 的 6:00 +0200 和 UTC 的 4:00 是相等的。
// 查看 Time 类型的文档，了解使用 == 操作符比较时间值时可能遇到的问题；
// 大多数代码应使用 Equal 而非 ==。
// md5:a28e147d11d5fe0f
func ExampleTime_Equal() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")
	gt2 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X是否相等(gt2))

	// Output:
	// true
}

// Before 返回时间点 t 是否在 u 之前。 md5:36690a50c1e8d9d4
func ExampleTime_Before() {
	gt1 := gtime.X创建("2018-08-07")
	gt2 := gtime.X创建("2018-08-08")

	fmt.Println(gt1.X是否之前(gt2))

	// Output:
	// true
}

// After 判断时间点t是否在u之后。 md5:750eca8bb04e1a25
func ExampleTime_After() {
	gt1 := gtime.X创建("2018-08-07")
	gt2 := gtime.X创建("2018-08-08")

	fmt.Println(gt1.X是否之后(gt2))

	// Output:
	// false
}

// Sub 返回持续时间 t-u。如果结果超过了能存储在 Duration 类型中的最大（或最小）
// 值，那么将返回最大（或最小）的持续时间。
// 要计算 t-d（其中 d 为一个持续时间），请使用 t.Add(-d)。
// md5:c975e5087c03d3b9
func ExampleTime_Sub() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")
	gt2 := gtime.X创建("2018-08-08 08:08:10")

	fmt.Println(gt2.X取纳秒时长(gt1))

	// Output:
	// 2s
}

// StartOfMinute 克隆并返回一个新的时间，其中秒数被设置为0。 md5:dc10ea1284a17280
func ExampleTime_StartOfMinute() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本忽略秒())

	// Output:
	// 2018-08-08 08:08:00
}

func ExampleTime_StartOfHour() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本忽略分钟秒())

	// Output:
	// 2018-08-08 08:00:00
}

func ExampleTime_StartOfDay() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本忽略小时分钟秒())

	// Output:
	// 2018-08-08 00:00:00
}

func ExampleTime_StartOfWeek() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本周第一天())

	// Output:
	// 2018-08-05 00:00:00
}

func ExampleTime_StartOfQuarter() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本季度第一天())

	// Output:
	// 2018-07-01 00:00:00
}

func ExampleTime_StartOfHalf() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本半年第一天())

	// Output:
	// 2018-07-01 00:00:00
}

func ExampleTime_StartOfYear() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本年第一天())

	// Output:
	// 2018-01-01 00:00:00
}

func ExampleTime_EndOfMinute() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本59秒())

	// Output:
	// 2018-08-08 08:08:59
}

func ExampleTime_EndOfHour() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本59分59秒())

	// Output:
	// 2018-08-08 08:59:59
}

func ExampleTime_EndOfDay() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本23点59分59秒())

	// Output:
	// 2018-08-08 23:59:59
}

func ExampleTime_EndOfWeek() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本周末23点59分59秒())

	// Output:
	// 2018-08-11 23:59:59
}

func ExampleTime_EndOfMonth() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本月末23点59分59秒())

	// Output:
	// 2018-08-31 23:59:59
}

func ExampleTime_EndOfQuarter() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本季末23点59分59秒())

	// Output:
	// 2018-09-30 23:59:59
}

func ExampleTime_EndOfHalf() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本半年末23点59分59秒())

	// Output:
	// 2018-12-31 23:59:59
}

func ExampleTime_EndOfYear() {
	gt1 := gtime.X创建("2018-08-08 08:08:08")

	fmt.Println(gt1.X取副本年末23点59分59秒())

	// Output:
	// 2018-12-31 23:59:59
}

func ExampleTime_MarshalJSON() {
	type Person struct {
		Name     string      `json:"name"`
		Birthday *gtime.Time `json:"birthday"`
	}
	p := new(Person)
	p.Name = "goframe"
	p.Birthday = gtime.X创建("2018-08-08 08:08:08")
	j, _ := json.Marshal(p)
	fmt.Println(string(j))

	// Output:
	// {"name":"goframe","birthday":"2018-08-08 08:08:08"}
}

func ExampleTime_UnmarshalJSON() {
	type Person struct {
		Name     string      `json:"name"`
		Birthday *gtime.Time `json:"birthday"`
	}
	p := new(Person)
	src := `{"name":"goframe","birthday":"2018-08-08 08:08:08"}`
	json.Unmarshal([]byte(src), p)

	fmt.Println(p)

	// 输出
	// &{goframe 2018-08-08 08:08:08} 
	// 
	// 这段Go代码的注释表示这是一个输出（Output），内容是关于一个结构体（&{...}）的引用，该结构体名为goframe，包含了日期和时间信息（2018-08-08 08:08:08）。
	// md5:a93ddd4a9e34a1af
}
