// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gtype 提供高性能且并发安全的基本变量类型。
package 安全变量类

// New 是 NewInterface 的别名。
// 请参阅 NewInterface。
func New(value ...interface{}) *Interface {
	return NewInterface(value...)
}
