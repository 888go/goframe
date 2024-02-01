// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv
import (
	"unsafe"
	)
// UnsafeStrToBytes 将字符串转换为 []byte，不进行内存拷贝。
// 注意，如果你确定在将来绝对不会使用 `s` 变量，
// 你可以使用这个不安全函数来实现高性能的类型转换。
func UnsafeStrToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// UnsafeBytesToStr 将 []byte 类型转换为 string 类型，且不进行内存拷贝操作。
// 注意：只有当你完全确定在未来绝不会再使用变量 `b` 时，
// 才可以使用这个不安全的函数以实现高性能的类型转换。
func UnsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
