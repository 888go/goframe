// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 转换类

import (
	"reflect"
	"testing"
)

type structType struct {
	Name  string
	Score int
}

var (
	structMap = map[string]interface{}{
		"name":  "gf",
		"score": 100,
	}
	structObj = structType{
		Name:  "john",
		Score: 60,
	}
	structPointer = &structType{
		Name:  "john",
		Score: 60,
	}
	structPointerNil *structType
	// struct slice
	structSliceNil []structType
	structSlice    = []structType{
		{Name: "john", Score: 60},
		{Name: "smith", Score: 100},
	}
	// 结构体指针切片
	structPointerSliceNil []*structType
	structPointerSlice    = []*structType{
		{Name: "john", Score: 60},
		{Name: "smith", Score: 100},
	}
)

func Benchmark_Struct_Basic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Struct(structMap, structPointer)
	}
}

// 将指针类型从*struct转换为**struct
func Benchmark_Reflect_PPStruct_PStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := reflect.ValueOf(&structPointerNil)
		v2 := reflect.ValueOf(structPointer)
// 如果v1的Kind（类型）是reflect.Ptr（指针类型），
// 那么进一步检查：
// 获取v1指向的元素值elem，如果elem的Type（类型）与v2的Type相同，
// 则将v2的值赋给elem。
// 这段代码实现了当v1是一个指向与v2相同类型的指针时，将v2的值赋给v1所指向的元素。
		v1.Elem().Set(v2)
	}
}

func Benchmark_Struct_PPStruct_PStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Struct(structPointer, &structPointerNil)
	}
}

// struct -> *struct
func Benchmark_Reflect_PStruct_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := reflect.ValueOf(structPointer)
		v2 := reflect.ValueOf(structObj)
// 如果v1的Kind（类型）是reflect.Ptr（指针类型），
// 那么进一步检查：
// 获取v1指向的元素值elem，如果elem的Type（类型）与v2的Type相同，
// 则将v2的值赋给elem。
// 这段代码实现了当v1是一个指向与v2相同类型的指针时，将v2的值赋给v1所指向的元素。
		v1.Elem().Set(v2)
	}
}

func Benchmark_Struct_PStruct_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Struct(structObj, structPointer)
	}
}

// 将切片结构体转换为指向切片结构体的指针
func Benchmark_Reflect_PStructs_Structs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := reflect.ValueOf(&structSliceNil)
		v2 := reflect.ValueOf(structSlice)
// 如果v1的Kind（类型）是reflect.Ptr（指针类型），
// 那么进一步检查：
// 获取v1指向的元素值elem，如果elem的Type（类型）与v2的Type相同，
// 则将v2的值赋给elem。
// 这段代码实现了当v1是一个指向与v2相同类型的指针时，将v2的值赋给v1所指向的元素。
		v1.Elem().Set(v2)
	}
}

func Benchmark_Structs_PStructs_Structs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Structs(structSlice, &structSliceNil)
	}
}

// 将指向结构体数组的指针转换为指向结构体指针数组的指针
func Benchmark_Reflect_PPStructs_PStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := reflect.ValueOf(&structPointerSliceNil)
		v2 := reflect.ValueOf(structPointerSlice)
// 如果v1的Kind（类型）是reflect.Ptr（指针类型），
// 那么进一步检查：
// 获取v1指向的元素值elem，如果elem的Type（类型）与v2的Type相同，
// 则将v2的值赋给elem。
// 这段代码实现了当v1是一个指向与v2相同类型的指针时，将v2的值赋给v1所指向的元素。
		v1.Elem().Set(v2)
	}
}

func Benchmark_Structs_PPStructs_PStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Structs(structPointerSlice, &structPointerSliceNil)
	}
}
