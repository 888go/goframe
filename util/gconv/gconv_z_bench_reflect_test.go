// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

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
