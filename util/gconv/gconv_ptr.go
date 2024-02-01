// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv

// PtrAny 创建并返回一个指向此值的 interface{} 类型指针变量。
func PtrAny(any interface{}) *interface{} {
	return &any
}

// PtrString 创建并返回一个指向此值的字符串指针变量。
func PtrString(any interface{}) *string {
	v := String(any)
	return &v
}

// PtrBool 创建并返回一个指向此值的布尔指针变量。
func PtrBool(any interface{}) *bool {
	v := Bool(any)
	return &v
}

// PtrInt 创建并返回一个指向该值的整数指针变量。
func PtrInt(any interface{}) *int {
	v := Int(any)
	return &v
}

// PtrInt8 创建并返回一个指向此值的 int8 类型指针变量。
func PtrInt8(any interface{}) *int8 {
	v := Int8(any)
	return &v
}

// PtrInt16 创建并返回一个指向此值的int16指针变量。
func PtrInt16(any interface{}) *int16 {
	v := Int16(any)
	return &v
}

// PtrInt32 创建并返回一个指向此值的int32指针变量。
func PtrInt32(any interface{}) *int32 {
	v := Int32(any)
	return &v
}

// PtrInt64 创建并返回一个指向此值的 int64 类型指针变量。
func PtrInt64(any interface{}) *int64 {
	v := Int64(any)
	return &v
}

// PtrUint 创建并返回一个指向此值的uint指针变量。
func PtrUint(any interface{}) *uint {
	v := Uint(any)
	return &v
}

// PtrUint8 创建并返回一个指向此值的uint8类型的指针变量。
func PtrUint8(any interface{}) *uint8 {
	v := Uint8(any)
	return &v
}

// PtrUint16 创建并返回一个指向此值的uint16类型的指针变量。
func PtrUint16(any interface{}) *uint16 {
	v := Uint16(any)
	return &v
}

// PtrUint32 创建并返回一个指向此值的 uint32 类型指针变量。
func PtrUint32(any interface{}) *uint32 {
	v := Uint32(any)
	return &v
}

// PtrUint64 创建并返回一个指向此值的uint64类型的指针变量。
func PtrUint64(any interface{}) *uint64 {
	v := Uint64(any)
	return &v
}

// PtrFloat32 创建并返回一个指向此值的float32类型的指针变量。
func PtrFloat32(any interface{}) *float32 {
	v := Float32(any)
	return &v
}

// PtrFloat64 创建并返回一个指向此值的 float64 类型指针变量。
func PtrFloat64(any interface{}) *float64 {
	v := Float64(any)
	return &v
}
