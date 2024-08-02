// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

// PtrAny 创建并返回一个指向此值的interface{}指针变量。 md5:12130f8892007df6
func PtrAny(any interface{}) *interface{} {
	return &any
}

// PtrString 创建并返回一个指向此值的字符串指针变量。 md5:8a45efa4c90feefd
func PtrString(any interface{}) *string {
	v := String(any)
	return &v
}

// PtrBool 创建并返回一个指向此值的布尔型指针变量。 md5:662b2b040324119f
func PtrBool(any interface{}) *bool {
	v := Bool(any)
	return &v
}

// PtrInt 创建并返回一个指向该值的int指针变量。 md5:aac4fd8dc5360539
func PtrInt(any interface{}) *int {
	v := Int(any)
	return &v
}

// PtrInt8 创建并返回一个指向此值的 int8 指针变量。 md5:135a8560671f447d
func PtrInt8(any interface{}) *int8 {
	v := Int8(any)
	return &v
}

// PtrInt16 创建并返回一个指向此值的 int16 指针变量。 md5:f1ea6c718962fd2e
func PtrInt16(any interface{}) *int16 {
	v := Int16(any)
	return &v
}

// PtrInt32 创建并返回一个指向此值的int32指针变量。 md5:7f7072cdeb72f52a
func PtrInt32(any interface{}) *int32 {
	v := Int32(any)
	return &v
}

// PtrInt64 创建并返回一个指向该值的int64指针变量。 md5:60417c5b51562e51
func PtrInt64(any interface{}) *int64 {
	v := Int64(any)
	return &v
}

// PtrUint 创建并返回一个指向此值的无符号整数指针变量。 md5:04c5b5df0b2baa8e
func PtrUint(any interface{}) *uint {
	v := Uint(any)
	return &v
}

// PtrUint8 创建并返回一个指向此值的 uint8 指针变量。 md5:c4901dd67ca1d339
func PtrUint8(any interface{}) *uint8 {
	v := Uint8(any)
	return &v
}

// PtrUint16 创建并返回一个指向此值的uint16指针变量。 md5:7fcfbf0260f97aa5
func PtrUint16(any interface{}) *uint16 {
	v := Uint16(any)
	return &v
}

// PtrUint32 创建并返回一个指向该值的uint32指针变量。 md5:95c5b9723ded3fd2
func PtrUint32(any interface{}) *uint32 {
	v := Uint32(any)
	return &v
}

// PtrUint64 创建并返回一个指向此值的 uint64 类型指针变量。 md5:85c9fc668348f455
func PtrUint64(any interface{}) *uint64 {
	v := Uint64(any)
	return &v
}

// PtrFloat32 创建并返回一个指向此值的 float32 指针变量。 md5:3ccf785d35432892
func PtrFloat32(any interface{}) *float32 {
	v := Float32(any)
	return &v
}

// PtrFloat64 创建并返回一个指向此值的float64指针变量。 md5:4ce193832d7c216e
func PtrFloat64(any interface{}) *float64 {
	v := Float64(any)
	return &v
}
