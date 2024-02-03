// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtest

import (
	"testing"
)

// T 是测试用例管理对象。
type T struct {
	*testing.T
}

// Assert 检查 `value` 和 `expect` 是否相等。
func (t *T) Assert(value, expect interface{}) {
	Assert(value, expect)
}

// AssertEQ 检查 `value` 和 `expect` 是否相等，包括它们的 TYPE（类型）。
func (t *T) AssertEQ(value, expect interface{}) {
	AssertEQ(value, expect)
}

// AssertNE 检查 `value` 和 `expect` 是否不相等。
func (t *T) AssertNE(value, expect interface{}) {
	AssertNE(value, expect)
}

// AssertNQ 检查 `value` 和 `expect` 是否不相等，包括它们的类型。
func (t *T) AssertNQ(value, expect interface{}) {
	AssertNQ(value, expect)
}

// AssertGT 检查 `value` 是否大于 `expect`。
// 注意，只有字符串、整数和浮点类型可以通过 AssertGT 进行比较，
// 其他类型是无效的。
func (t *T) AssertGT(value, expect interface{}) {
	AssertGT(value, expect)
}

// AssertGE 检查 `value` 是否大于或等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以使用 AssertGTE 进行比较，其他类型无效。
func (t *T) AssertGE(value, expect interface{}) {
	AssertGE(value, expect)
}

// AssertLT 检查 `value` 是否小于等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以使用 AssertLT 进行比较，
// 其他类型无效。
func (t *T) AssertLT(value, expect interface{}) {
	AssertLT(value, expect)
}

// AssertLE 检查 `value` 是否小于等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以通过 AssertLTE 进行比较，
// 其他类型是无效的。
func (t *T) AssertLE(value, expect interface{}) {
	AssertLE(value, expect)
}

// AssertIN 检查 `value` 是否在 `expect` 内。
// 其中 `expect` 应为一个切片，
// 但 `value` 可以是切片或基本类型变量。
func (t *T) AssertIN(value, expect interface{}) {
	AssertIN(value, expect)
}

// AssertNI 检查 `value` 是否不在 `expect` 中。
// `expect` 应该是一个切片，
// 但 `value` 可以是切片或基本类型变量。
func (t *T) AssertNI(value, expect interface{}) {
	AssertNI(value, expect)
}

// AssertNil 断言 `value` 为 nil。
func (t *T) AssertNil(value interface{}) {
	AssertNil(value)
}

// Error 使用给定的`message`引发panic异常。
func (t *T) Error(message ...interface{}) {
	Error(message...)
}

// Fatal将`message`打印到标准错误输出（stderr）并退出进程。
func (t *T) Fatal(message ...interface{}) {
	Fatal(message...)
}
