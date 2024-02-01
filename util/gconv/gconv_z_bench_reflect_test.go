// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gconv_test
import (
	"reflect"
	"testing"
	)
type testStruct struct {
	Id   int
	Name string
}

var ptr = []*testStruct{
	{
		Id:   1,
		Name: "test1",
	},
	{
		Id:   2,
		Name: "test2",
	},
}

func init() {
	for i := 1; i <= 1000; i++ {
		ptr = append(ptr, &testStruct{
			Id:   1,
			Name: "test1",
		})
	}
}

func Benchmark_Reflect_ValueOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(ptr)
	}
}

func Benchmark_Reflect_ValueOf_Kind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(ptr).Kind()
	}
}

func Benchmark_Reflect_ValueOf_Interface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(ptr).Interface()
	}
}

func Benchmark_Reflect_ValueOf_Len(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(ptr).Len()
	}
}
