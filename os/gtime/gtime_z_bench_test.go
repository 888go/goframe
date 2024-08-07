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

	gtime "github.com/888go/goframe/os/gtime"
)

func Benchmark_Timestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X取时间戳秒()
	}
}

func Benchmark_TimestampMilli(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X取时间戳毫秒()
	}
}

func Benchmark_TimestampMicro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X取时间戳微秒()
	}
}

func Benchmark_TimestampNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X取时间戳纳秒()
	}
}

func Benchmark_StrToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X转换文本("2018-02-09T20:46:17.897Z")
	}
}

func Benchmark_StrToTime_Format(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X转换文本("2018-02-09 20:46:17.897", "Y-m-d H:i:su")
	}
}

func Benchmark_StrToTime_Layout(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X转换文本Layout("2018-02-09T20:46:17.897Z", time.RFC3339)
	}
}

func Benchmark_ParseTimeFromContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X解析文本("2018-02-09T20:46:17.897Z")
	}
}

func Benchmark_NewFromTimeStamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X创建并从时间戳(1542674930)
	}
}

func Benchmark_Date(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.Date()
	}
}

func Benchmark_Datetime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.X取当前日期时间()
	}
}
