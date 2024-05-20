// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gconv

import "github.com/gogf/gf/v2/os/gtime"

// iVal 用于类型断言以实现String()方法。. md5:46c43d7149579432
type iVal interface {
	Val() interface{}
}

// iString 用于类型断言API，用于String()。. md5:8ec0af717c4f530e
type iString interface {
	String() string
}

// iBool 用于布尔类型断言的 API。. md5:59ff9e6dd6e8d3da
type iBool interface {
	Bool() bool
}

// iInt64 用于对 Int64() 函数的类型断言。. md5:db7ff0850aa64638
type iInt64 interface {
	Int64() int64
}

// iUint64 用于为 Uint64() 方法的类型断言提供支持。. md5:49a588d92caa8794
type iUint64 interface {
	Uint64() uint64
}

// iFloat32 用于在Float32()方法中进行类型断言的API。. md5:2187b83d5c94d667
type iFloat32 interface {
	Float32() float32
}

// iFloat64 用于Float64()类型的断言API。. md5:b38f89afc3709759
type iFloat64 interface {
	Float64() float64
}

// iError用于类型断言错误信息。. md5:ca9885066be22039
type iError interface {
	Error() string
}

// iBytes 用于类型断言 API，以支持 Bytes() 方法。. md5:f39d15d800efa326
type iBytes interface {
	Bytes() []byte
}

// iInterface 用于Interface()方法的类型断言接口。. md5:9daf47766ff28118
type iInterface interface {
	Interface() interface{}
}

// iInterfaces 用于类型断言接口，用于 Interfaces() 方法。. md5:711dc755f9cd4979
type iInterfaces interface {
	Interfaces() []interface{}
}

// iFloats 用于Floats()方法的类型断言。. md5:72d86b425f0484a9
type iFloats interface {
	Floats() []float64
}

// iInts 用于 Ints() 类型断言API。. md5:f310759e5276f31e
type iInts interface {
	Ints() []int
}

// iStrings 用于为 Strings() 方法提供类型断言的接口。. md5:fb5546612acb4787
type iStrings interface {
	Strings() []string
}

// iUints 用于Uints()的类型断言API。. md5:df1a889976394f51
type iUints interface {
	Uints() []uint
}

// iMapStrAny 是一个接口，支持将结构体参数转换为映射。. md5:cfd4642c77fca6ec
type iMapStrAny interface {
	MapStrAny() map[string]interface{}
}

// iUnmarshalValue 是用于自定义类型定制值赋值的接口。
// 注意，只有指针可以实现 iUnmarshalValue 接口。
// md5:be9b0c2575849208
type iUnmarshalValue interface {
	UnmarshalValue(interface{}) error
}

// iUnmarshalText 是自定义类型用于自定义值赋值的接口。
// 注意，只有指针类型可以实现 iUnmarshalText 接口。
// md5:cdd798fd0d1402d5
type iUnmarshalText interface {
	UnmarshalText(text []byte) error
}

// iUnmarshalText 是用于自定义类型以定制值赋予的接口。
// 注意，只有指针类型可以实现 iUnmarshalJSON 接口。
// md5:ea7d987eea1cf703
type iUnmarshalJSON interface {
	UnmarshalJSON(b []byte) error
}

// iSet 是自定义值赋值的接口。. md5:a36eda9131af6c27
type iSet interface {
	Set(value interface{}) (old interface{})
}

// iGTime是gtime.Time转换的接口。. md5:33093b8b6fff69af
type iGTime interface {
	GTime(format ...string) *gtime.Time
}
