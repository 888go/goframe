// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

// X取any指针 创建并返回一个指向此值的interface{}指针变量。 md5:12130f8892007df6
func X取any指针(值 interface{}) *interface{} {
	return &值
}

// X取文本指针 创建并返回一个指向此值的字符串指针变量。 md5:8a45efa4c90feefd
func X取文本指针(值 interface{}) *string {
	v := String(值)
	return &v
}

// X取布尔指针 创建并返回一个指向此值的布尔型指针变量。 md5:662b2b040324119f
func X取布尔指针(值 interface{}) *bool {
	v := X取布尔(值)
	return &v
}

// X取整数指针 创建并返回一个指向该值的int指针变量。 md5:aac4fd8dc5360539
func X取整数指针(值 interface{}) *int {
	v := X取整数(值)
	return &v
}

// X取整数8位指针 创建并返回一个指向此值的 int8 指针变量。 md5:135a8560671f447d
func X取整数8位指针(值 interface{}) *int8 {
	v := X取整数8位(值)
	return &v
}

// X取整数16位指针 创建并返回一个指向此值的 int16 指针变量。 md5:f1ea6c718962fd2e
func X取整数16位指针(值 interface{}) *int16 {
	v := X取整数16位(值)
	return &v
}

// X取整数32位指针 创建并返回一个指向此值的int32指针变量。 md5:7f7072cdeb72f52a
func X取整数32位指针(值 interface{}) *int32 {
	v := X取整数32位(值)
	return &v
}

// X取整数64位指针 创建并返回一个指向该值的int64指针变量。 md5:60417c5b51562e51
func X取整数64位指针(值 interface{}) *int64 {
	v := X取整数64位(值)
	return &v
}

// X取正整数指针 创建并返回一个指向此值的无符号整数指针变量。 md5:04c5b5df0b2baa8e
func X取正整数指针(值 interface{}) *uint {
	v := X取正整数(值)
	return &v
}

// X取正整数8位指针 创建并返回一个指向此值的 uint8 指针变量。 md5:c4901dd67ca1d339
func X取正整数8位指针(值 interface{}) *uint8 {
	v := X取正整数8位(值)
	return &v
}

// X取正整数16位指针 创建并返回一个指向此值的uint16指针变量。 md5:7fcfbf0260f97aa5
func X取正整数16位指针(值 interface{}) *uint16 {
	v := X取正整数16位(值)
	return &v
}

// X取正整数32位指针 创建并返回一个指向该值的uint32指针变量。 md5:95c5b9723ded3fd2
func X取正整数32位指针(值 interface{}) *uint32 {
	v := X取正整数32位(值)
	return &v
}

// X取正整数64位指针 创建并返回一个指向此值的 uint64 类型指针变量。 md5:85c9fc668348f455
func X取正整数64位指针(值 interface{}) *uint64 {
	v := X取正整数64位(值)
	return &v
}

// X取小数32位指针 创建并返回一个指向此值的 float32 指针变量。 md5:3ccf785d35432892
func X取小数32位指针(值 interface{}) *float32 {
	v := X取小数32位(值)
	return &v
}

// X取小数64位指针 创建并返回一个指向此值的float64指针变量。 md5:4ce193832d7c216e
func X取小数64位指针(值 interface{}) *float64 {
	v := X取小数64位(值)
	return &v
}
