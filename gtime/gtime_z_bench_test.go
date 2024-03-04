// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtime_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/gtime"
)

func Benchmark_Timestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.Timestamp()
	}
}

func Benchmark_TimestampMilli(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.TimestampMilli()
	}
}

func Benchmark_TimestampMicro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.TimestampMicro()
	}
}

func Benchmark_TimestampNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.TimestampNano()
	}
}

func Benchmark_StrToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.StrToTime("2018-02-09T20:46:17.897Z")
	}
}

func Benchmark_StrToTime_Format(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.StrToTime("2018-02-09 20:46:17.897", "Y-m-d H:i:su")
	}
}

func Benchmark_StrToTime_Layout(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.StrToTimeLayout("2018-02-09T20:46:17.897Z", time.RFC3339)
	}
}

func Benchmark_ParseTimeFromContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.ParseTimeFromContent("2018-02-09T20:46:17.897Z")
	}
}

func Benchmark_NewFromTimeStamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.NewFromTimeStamp(1542674930)
	}
}

func Benchmark_Date(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.Date()
	}
}

func Benchmark_Datetime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gtime.Datetime()
	}
}
