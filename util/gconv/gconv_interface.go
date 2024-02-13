// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"github.com/888go/goframe/os/gtime"
)

// iVal 用于在进行类型断言时作为 String() 方法的 API。
type iVal interface {
	X取值() interface{}
}

// iString 用于在进行类型断言时，配合 String() 方法使用。
type iString interface {
	String() string
}

// iBool 用于 Bool() 函数的类型断言API。
type iBool interface {
	X取布尔() bool
}

// iInt64 用于在类型断言 API 中对 Int64() 方法进行操作。
type iInt64 interface {
	X取整数64位() int64
}

// iUint64 用于对 Uint64() 方法进行类型断言。
type iUint64 interface {
	X取正整数64位() uint64
}

// iFloat32 用于进行类型断言，以调用 Float32() 函数。
type iFloat32 interface {
	X取小数32位() float32
}

// iFloat64 用于进行类型断言，以支持 Float64() 方法的调用。
type iFloat64 interface {
	X取小数64位() float64
}

// iError 用于对 Error() 方法进行类型断言。
type iError interface {
	Error() string
}

// iBytes 用于对 Bytes() 方法进行类型断言。
type iBytes interface {
	X取字节集() []byte
}

// iInterface 用于 Interface() 方法的类型断言接口。
type iInterface interface {
	Interface() interface{}
}

// iInterfaces 用于对 Interfaces() 方法进行类型断言。
type iInterfaces interface {
	X取any数组() []interface{}
}

// iFloats 用于进行类型断言，以配合 Floats() API 使用。
type iFloats interface {
	X取小数数组() []float64
}

// iInts 用于 Ints() 方法的类型断言API。
type iInts interface {
	X取整数数组() []int
}

// iStrings 用于对 Strings() 方法进行类型断言。
type iStrings interface {
	X取文本数组() []string
}

// iUints 用于 Uints() 函数的类型断言接口。
type iUints interface {
	X取正整数数组() []uint
}

// iMapStrAny 是支持将结构体参数转换为映射的接口。
type iMapStrAny interface {
	X取MapStrAny() map[string]interface{}
}

// iUnmarshalValue 是一个接口，用于为自定义类型定制值赋值功能。
// 注意，只有指针类型可以实现 iUnmarshalValue 接口。
type iUnmarshalValue interface {
	UnmarshalValue(interface{}) error
}

// iUnmarshalText 是为自定义类型定制值赋值的接口。
// 注意，只有指针类型可以实现 iUnmarshalText 接口。
// 这段Go语言代码注释翻译成中文如下：
// ```go
// iUnmarshalText 是一个用于自定义类型以个性化实现值赋值的接口。
// 需要注意的是，只有指针类型才能实现 iUnmarshalText 接口。
type iUnmarshalText interface {
	UnmarshalText(text []byte) error
}

// iUnmarshalText 是用于自定义类型定制值赋值的接口。
// 注意，只有指针类型才能实现 iUnmarshalJSON 接口。
type iUnmarshalJSON interface {
	UnmarshalJSON(b []byte) error
}

// iSet 是用于自定义值赋值的接口。
type iSet interface {
	X设置值(value interface{}) (old interface{})
}

// iGTime是用于gtime.Time转换的接口。
type iGTime interface {
	X取gtime时间类(format ...string) *时间类.Time
}
