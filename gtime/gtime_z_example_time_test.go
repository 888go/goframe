// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtime_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
	
	"github.com/888go/goframe/gtime"
)

func ExampleNew_Basic() {
	curTime := "2018-08-08 08:08:08"
	timer, _ := time.Parse("2006-01-02 15:04:05", curTime)
	t1 := gtime.New(&timer)
	t2 := gtime.New(curTime)
	t3 := gtime.New(curTime, "Y-m-d H:i:s")
	t4 := gtime.New(curTime)
	t5 := gtime.New(1533686888)

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
	fmt.Println(gtime.New("20220629133225", "YmdHis").Format("Y-m-d H:i:s"))

	// Output:
	// 2022-06-29 13:32:25
}

// Now 创建并返回一个表示当前时间的时间对象。
func ExampleNow() {
	t := gtime.Now()
	fmt.Println(t)

	// May Output:
	// 2021-11-06 13:41:08
}

// NewFromTime 根据给定的 time.Time 对象创建并返回一个 Time 对象。
func ExampleNewFromTime() {
	timer, _ := time.Parse("2006-01-02 15:04:05", "2018-08-08 08:08:08")
	nTime := gtime.NewFromTime(timer)

	fmt.Println(nTime)

	// Output:
	// 2018-08-08 08:08:08
}

// NewFromStr 通过给定的字符串创建并返回一个 Time 对象。
// 注意，如果发生错误，它将返回 nil。
func ExampleNewFromStr() {
	t := gtime.NewFromStr("2018-08-08 08:08:08")

	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:08
}

// NewFromStrFormat 根据给定的字符串和自定义格式（如：Y-m-d H:i:s）创建并返回一个Time对象。
// 需要注意，如果发生错误，它将返回nil。
func ExampleNewFromStrFormat() {
	t := gtime.NewFromStrFormat("2018-08-08 08:08:08", "Y-m-d H:i:s")
	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:08
}

// NewFromStrLayout 根据给定的字符串和标准库布局（如：2006-01-02 15:04:05）创建并返回一个Time对象。
// 需要注意的是，如果发生错误，它将返回nil。
func ExampleNewFromStrLayout() {
	t := gtime.NewFromStrLayout("2018-08-08 08:08:08", "2006-01-02 15:04:05")
	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:08
}

// NewFromTimeStamp 根据给定的时间戳创建并返回一个 Time 对象，
// 时间戳可以是秒级到纳秒级精度。
// 例如：1600443866 和 1600443866199266000 都被认为是有效的时间戳数值。
func ExampleNewFromTimeStamp() {
	t1 := gtime.NewFromTimeStamp(1533686888)
	t2 := gtime.NewFromTimeStamp(1533686888000)

	fmt.Println(t1.String() == t2.String())
	fmt.Println(t1)

	// Output:
	// true
	// 2018-08-08 08:08:08
}

// Timestamp 返回以秒为单位的时间戳。
func ExampleTime_Timestamp() {
	t := gtime.Timestamp()

	fmt.Println(t)

	// May output:
	// 1533686888
}

// Timestamp 返回以毫秒为单位的时间戳。
func ExampleTime_TimestampMilli() {
	t := gtime.TimestampMilli()

	fmt.Println(t)

	// May output:
	// 1533686888000
}

// Timestamp 返回以微秒为单位的的时间戳。
func ExampleTime_TimestampMicro() {
	t := gtime.TimestampMicro()

	fmt.Println(t)

	// May output:
	// 1533686888000000
}

// Timestamp 返回以纳秒为单位的时间戳。
func ExampleTime_TimestampNano() {
	t := gtime.TimestampNano()

	fmt.Println(t)

	// May output:
	// 1533686888000000
}

// TimestampStr 是一个便捷方法，用于获取并返回以字符串形式表示的秒级时间戳。
func ExampleTime_TimestampStr() {
	t := gtime.TimestampStr()

	fmt.Println(reflect.TypeOf(t))

	// Output:
	// string
}

// Month 返回由 t 指定的年份中的月份。
func ExampleTime_Month() {
	gt := gtime.New("2018-08-08 08:08:08")
	t1 := gt.Month()

	fmt.Println(t1)

	// Output:
	// 8
}

// Second 返回给定时间 t 的分钟内第二个偏移量，
// 范围在 [0, 59] 内。
func ExampleTime_Second() {
	gt := gtime.New("2018-08-08 08:08:08")
	t1 := gt.Second()

	fmt.Println(t1)

	// Output:
	// 8
}

// String 将当前时间对象转换为字符串并返回。
func ExampleTime_String() {
	gt := gtime.New("2018-08-08 08:08:08")
	t1 := gt.String()

	fmt.Println(t1)
	fmt.Println(reflect.TypeOf(t1))

	// Output:
	// 2018-08-08 08:08:08
	// string
}

// IsZero 判断 t 是否代表零时间点，即公元1年1月1日 00:00:00 UTC。
func ExampleTime_IsZero() {
	gt := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt.IsZero())

	// Output:
	// false
}

// Add 将持续时间添加到当前时间。
func ExampleTime_Add() {
	gt := gtime.New("2018-08-08 08:08:08")
	gt1 := gt.Add(time.Duration(10) * time.Second)

	fmt.Println(gt1)

	// Output:
	// 2018-08-08 08:08:18
}

// AddStr parses the given duration as string and adds it to current time.
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ExampleTime_AddStr() {
	gt := gtime.New("2018-08-08 08:08:08")
	gt1, _ := gt.AddStr("10s")

	fmt.Println(gt1)

	// Output:
	// 2018-08-08 08:08:18
}

// AddDate 向时间添加年、月和日。
func ExampleTime_AddDate() {
	var (
		year  = 1
		month = 2
		day   = 3
	)
	gt := gtime.New("2018-08-08 08:08:08")
	gt = gt.AddDate(year, month, day)

	fmt.Println(gt)

	// Output:
	// 2019-10-11 08:08:08
}

// Round 函数将 t 舍入到最接近 d 的倍数（以零时间点为基准）。
// 对于刚好位于中间值的舍入行为是向上舍入。
// 如果 d 小于等于 0，Round 函数将返回剥离了单调时钟读数但其他部分保持不变的 t。
//
// Round 函数针对的是以零时间为基准的绝对持续时间上的时间；
// 它并不作用在时间的表现形式上。因此，即使调用 Round(Hour)，
// 返回的时间也可能存在非零分钟值，这取决于时间所处的 Location（时区）。
func ExampleTime_Round() {
	gt := gtime.New("2018-08-08 08:08:08")
	t := gt.Round(time.Duration(10) * time.Second)

	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:10
}

// Truncate 方法将 t 向下舍入至 d 的倍数（以零时间点为基准）。
// 若 d 小于等于 0，Truncate 方法会返回剥离了单调时钟读数但其他部分保持不变的 t。
//
// Truncate 对时间进行操作时将其视为从零时间点开始的绝对持续时间；
// 它并不直接作用于时间的展示形式。因此，调用 Truncate(Hour) 可能会返回一个分钟不为零的时间，
// 具体取决于该时间的位置（Location）。
func ExampleTime_Truncate() {
	gt := gtime.New("2018-08-08 08:08:08")
	t := gt.Truncate(time.Duration(10) * time.Second)

	fmt.Println(t)

	// Output:
	// 2018-08-08 08:08:00
}

// Equal 判断 t 和 u 是否表示相同的时刻。
// 即使两个时间位于不同的时区，它们也可能相等。
// 例如，6:00 +0200 CEST（中欧夏令时）和 4:00 UTC 是相等的。
// 查看 Time 类型的文档，了解使用 == 操作符比较 Time 值时的陷阱；
// 大多数代码应使用 Equal 方法代替。
func ExampleTime_Equal() {
	gt1 := gtime.New("2018-08-08 08:08:08")
	gt2 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.Equal(gt2))

	// Output:
	// true
}

// Before 判断时间点 t 是否在时间点 u 之前。
func ExampleTime_Before() {
	gt1 := gtime.New("2018-08-07")
	gt2 := gtime.New("2018-08-08")

	fmt.Println(gt1.Before(gt2))

	// Output:
	// true
}

// After 判断时间点 t 是否在时间点 u 之后。
func ExampleTime_After() {
	gt1 := gtime.New("2018-08-07")
	gt2 := gtime.New("2018-08-08")

	fmt.Println(gt1.After(gt2))

	// Output:
	// false
}

// Sub 计算并返回时间段 t-u。如果结果超出了 Duration 类型能够存储的最大（或最小）值，
// 则会返回最大（或最小）的有效持续时间。
// 若要计算 t 与一个持续时间 d 的差值（t-d），请使用 t.Add(-d)。
func ExampleTime_Sub() {
	gt1 := gtime.New("2018-08-08 08:08:08")
	gt2 := gtime.New("2018-08-08 08:08:10")

	fmt.Println(gt2.Sub(gt1))

	// Output:
	// 2s
}

// StartOfMinute 复制并返回一个新的时间，其秒数设置为0。
func ExampleTime_StartOfMinute() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.StartOfMinute())

	// Output:
	// 2018-08-08 08:08:00
}

func ExampleTime_StartOfHour() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.StartOfHour())

	// Output:
	// 2018-08-08 08:00:00
}

func ExampleTime_StartOfDay() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.StartOfDay())

	// Output:
	// 2018-08-08 00:00:00
}

func ExampleTime_StartOfWeek() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.StartOfWeek())

	// Output:
	// 2018-08-05 00:00:00
}

func ExampleTime_StartOfQuarter() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.StartOfQuarter())

	// Output:
	// 2018-07-01 00:00:00
}

func ExampleTime_StartOfHalf() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.StartOfHalf())

	// Output:
	// 2018-07-01 00:00:00
}

func ExampleTime_StartOfYear() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.StartOfYear())

	// Output:
	// 2018-01-01 00:00:00
}

func ExampleTime_EndOfMinute() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.EndOfMinute())

	// Output:
	// 2018-08-08 08:08:59
}

func ExampleTime_EndOfHour() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.EndOfHour())

	// Output:
	// 2018-08-08 08:59:59
}

func ExampleTime_EndOfDay() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.EndOfDay())

	// Output:
	// 2018-08-08 23:59:59
}

func ExampleTime_EndOfWeek() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.EndOfWeek())

	// Output:
	// 2018-08-11 23:59:59
}

func ExampleTime_EndOfMonth() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.EndOfMonth())

	// Output:
	// 2018-08-31 23:59:59
}

func ExampleTime_EndOfQuarter() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.EndOfQuarter())

	// Output:
	// 2018-09-30 23:59:59
}

func ExampleTime_EndOfHalf() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.EndOfHalf())

	// Output:
	// 2018-12-31 23:59:59
}

func ExampleTime_EndOfYear() {
	gt1 := gtime.New("2018-08-08 08:08:08")

	fmt.Println(gt1.EndOfYear())

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
	p.Birthday = gtime.New("2018-08-08 08:08:08")
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
// （这段注释表明该段代码的执行结果会输出一个包含"goframe"和"2018-08-08 08:08:08"信息的数据结构，并以地址引用形式显示。）
}
