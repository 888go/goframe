// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 转换类

import (
	"reflect"
	"testing"
)

type structType struct {
	Name  string
	Score int
	Age   int
	ID    int
}

type structType8 struct {
	Name        string  `json:"name"   `
	CategoryId  string  `json:"category-Id" `
	Price       float64 `json:"price"    `
	Code        string  `json:"code"       `
	Image       string  `json:"image"   `
	Description string  `json:"description" `
	Status      int     `json:"status"   `
	IdType      int     `json:"id-type"`
	Score       int
	Age         int
	ID          int
}

var (
	structMap = map[string]interface{}{
		"name":  "gf",
		"score": 100,
		"Age":   98,
		"ID":    199,
	}

	structMapFields8 = map[string]interface{}{
		"name":  "gf",
		"score": 100,
		"Age":   98,
		"ID":    199,

		"category-Id": "1",
		"price":       198.09,
		"code":        "1",
		"image":       "https://goframe.org",
		"description": "This is the data for testing eight fields",
		"status":      1,
		"id-type":     2,
	}

	structObj = structType{
		Name:  "john",
		Score: 60,
		Age:   98,
		ID:    199,
	}
	structPointer = &structType{
		Name:  "john",
		Score: 60,
	}
	structPointer8   = &structType8{}
	structPointerNil *structType

	// struct slice
	structSliceNil []structType
	structSlice    = []structType{
		{Name: "john", Score: 60},
		{Name: "smith", Score: 100},
	}
	// struct pointer slice
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

func Benchmark_doStruct_Fields8_Basic_MapToStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doStruct(structMapFields8, structPointer8, map[string]string{}, "")
	}
}

// *struct -> **struct
func Benchmark_Reflect_PPStruct_PStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := reflect.ValueOf(&structPointerNil)
		v2 := reflect.ValueOf(structPointer)
		//如果v1的Kind()为指针类型 {
		// 如果elem是v1的元素（即指针所指向的对象），并且elem的Type()与v2相同 {
		// 将v2的值赋给elem
		//}
		// md5:f0c12588cbe6880e
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
		//如果v1的Kind()为指针类型 {
		// 如果elem是v1的元素（即指针所指向的对象），并且elem的Type()与v2相同 {
		// 将v2的值赋给elem
		//}
		// md5:f0c12588cbe6880e
		v1.Elem().Set(v2)
	}
}

func Benchmark_Struct_PStruct_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Struct(structObj, structPointer)
	}
}

// []struct -> *[]struct
func Benchmark_Reflect_PStructs_Structs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := reflect.ValueOf(&structSliceNil)
		v2 := reflect.ValueOf(structSlice)
		//如果v1的Kind()为指针类型 {
		// 如果elem是v1的元素（即指针所指向的对象），并且elem的Type()与v2相同 {
		// 将v2的值赋给elem
		//}
		// md5:f0c12588cbe6880e
		v1.Elem().Set(v2)
	}
}

func Benchmark_Structs_PStructs_Structs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Structs(structSlice, &structSliceNil)
	}
}

// 将切片的指针类型转换为指向切片元素的指针的指针类型. md5:3e13fcf1edb49ff6
func Benchmark_Reflect_PPStructs_PStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := reflect.ValueOf(&structPointerSliceNil)
		v2 := reflect.ValueOf(structPointerSlice)
		//如果v1的Kind()为指针类型 {
		// 如果elem是v1的元素（即指针所指向的对象），并且elem的Type()与v2相同 {
		// 将v2的值赋给elem
		//}
		// md5:f0c12588cbe6880e
		v1.Elem().Set(v2)
	}
}

func Benchmark_Structs_PPStructs_PStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Structs(structPointerSlice, &structPointerSliceNil)
	}
}
