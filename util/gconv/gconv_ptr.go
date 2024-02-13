// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

// PtrAny 创建并返回一个指向此值的 interface{} 类型指针变量。
func X取any指针(值 interface{}) *interface{} {
	return &值
}

// PtrString 创建并返回一个指向此值的字符串指针变量。
func X取文本指针(值 interface{}) *string {
	v := String(值)
	return &v
}

// PtrBool 创建并返回一个指向此值的布尔指针变量。
func X取布尔指针(值 interface{}) *bool {
	v := X取布尔(值)
	return &v
}

// PtrInt 创建并返回一个指向该值的整数指针变量。
func X取整数指针(值 interface{}) *int {
	v := X取整数(值)
	return &v
}

// PtrInt8 创建并返回一个指向此值的 int8 类型指针变量。
func X取整数8位指针(值 interface{}) *int8 {
	v := X取整数8位(值)
	return &v
}

// PtrInt16 创建并返回一个指向此值的int16指针变量。
func X取整数16位指针(值 interface{}) *int16 {
	v := X取整数16位(值)
	return &v
}

// PtrInt32 创建并返回一个指向此值的int32指针变量。
func X取整数32位指针(值 interface{}) *int32 {
	v := X取整数32位(值)
	return &v
}

// PtrInt64 创建并返回一个指向此值的 int64 类型指针变量。
func X取整数64位指针(值 interface{}) *int64 {
	v := X取整数64位(值)
	return &v
}

// PtrUint 创建并返回一个指向此值的uint指针变量。
func X取正整数指针(值 interface{}) *uint {
	v := X取正整数(值)
	return &v
}

// PtrUint8 创建并返回一个指向此值的uint8类型的指针变量。
func X取正整数8位指针(值 interface{}) *uint8 {
	v := X取正整数8位(值)
	return &v
}

// PtrUint16 创建并返回一个指向此值的uint16类型的指针变量。
func X取正整数16位指针(值 interface{}) *uint16 {
	v := X取正整数16位(值)
	return &v
}

// PtrUint32 创建并返回一个指向此值的 uint32 类型指针变量。
func X取正整数32位指针(值 interface{}) *uint32 {
	v := X取正整数32位(值)
	return &v
}

// PtrUint64 创建并返回一个指向此值的uint64类型的指针变量。
func X取正整数64位指针(值 interface{}) *uint64 {
	v := X取正整数64位(值)
	return &v
}

// PtrFloat32 创建并返回一个指向此值的float32类型的指针变量。
func X取小数32位指针(值 interface{}) *float32 {
	v := X取小数32位(值)
	return &v
}

// PtrFloat64 创建并返回一个指向此值的 float64 类型指针变量。
func X取小数64位指针(值 interface{}) *float64 {
	v := X取小数64位(值)
	return &v
}
