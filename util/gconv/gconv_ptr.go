// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

// PtrAny creates and returns an interface{} pointer variable to this value.
// ff:取any指针
// any:值
func PtrAny(any interface{}) *interface{} {
	return &any
}

// PtrString creates and returns a string pointer variable to this value.
// ff:取文本指针
// any:值
func PtrString(any interface{}) *string {
	v := String(any)
	return &v
}

// PtrBool creates and returns a bool pointer variable to this value.
// ff:取布尔指针
// any:值
func PtrBool(any interface{}) *bool {
	v := Bool(any)
	return &v
}

// PtrInt creates and returns an int pointer variable to this value.
// ff:取整数指针
// any:值
func PtrInt(any interface{}) *int {
	v := Int(any)
	return &v
}

// PtrInt8 creates and returns an int8 pointer variable to this value.
// ff:取整数8位指针
// any:值
func PtrInt8(any interface{}) *int8 {
	v := Int8(any)
	return &v
}

// PtrInt16 creates and returns an int16 pointer variable to this value.
// ff:取整数16位指针
// any:值
func PtrInt16(any interface{}) *int16 {
	v := Int16(any)
	return &v
}

// PtrInt32 creates and returns an int32 pointer variable to this value.
// ff:取整数32位指针
// any:值
func PtrInt32(any interface{}) *int32 {
	v := Int32(any)
	return &v
}

// PtrInt64 creates and returns an int64 pointer variable to this value.
// ff:取整数64位指针
// any:值
func PtrInt64(any interface{}) *int64 {
	v := Int64(any)
	return &v
}

// PtrUint creates and returns an uint pointer variable to this value.
// ff:取正整数指针
// any:值
func PtrUint(any interface{}) *uint {
	v := Uint(any)
	return &v
}

// PtrUint8 creates and returns an uint8 pointer variable to this value.
// ff:取正整数8位指针
// any:值
func PtrUint8(any interface{}) *uint8 {
	v := Uint8(any)
	return &v
}

// PtrUint16 creates and returns an uint16 pointer variable to this value.
// ff:取正整数16位指针
// any:值
func PtrUint16(any interface{}) *uint16 {
	v := Uint16(any)
	return &v
}

// PtrUint32 creates and returns an uint32 pointer variable to this value.
// ff:取正整数32位指针
// any:值
func PtrUint32(any interface{}) *uint32 {
	v := Uint32(any)
	return &v
}

// PtrUint64 creates and returns an uint64 pointer variable to this value.
// ff:取正整数64位指针
// any:值
func PtrUint64(any interface{}) *uint64 {
	v := Uint64(any)
	return &v
}

// PtrFloat32 creates and returns a float32 pointer variable to this value.
// ff:取小数32位指针
// any:值
func PtrFloat32(any interface{}) *float32 {
	v := Float32(any)
	return &v
}

// PtrFloat64 creates and returns a float64 pointer variable to this value.
// ff:取小数64位指针
// any:值
func PtrFloat64(any interface{}) *float64 {
	v := Float64(any)
	return &v
}
