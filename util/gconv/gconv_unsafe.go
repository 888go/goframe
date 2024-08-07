// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"unsafe"
)

// X文本到字节集_非安全 将字符串转换为 []byte，其间不进行内存复制。
// 注意，如果您完全确定将来绝不会使用 `s` 变量，
// 您可以使用这个不安全的函数来实现高性能的类型转换。
// md5:2ea7b3663055237b
func X文本到字节集_非安全(文本 string) []byte {
	return *(*[]byte)(unsafe.Pointer(&文本))
}

// X字节集到文本_非安全 将 []byte 转换为 string，而不进行内存复制。
// 请注意，如果你确定将来绝不会使用 `b` 变量，
// 可以使用这个不安全的函数来实现高性能的类型转换。
// md5:1d73c9ff996784ae
func X字节集到文本_非安全(字节集 []byte) string {
	return *(*string)(unsafe.Pointer(&字节集))
}
