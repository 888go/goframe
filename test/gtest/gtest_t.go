// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtest//bm:单元测试类

import (
	"testing"
)

// T 是测试用例管理对象。 md5:48ef1adf72668d4b
type T struct {
	*testing.T
}

// Assert 检查 `value` 和 `expect` 是否相等。 md5:eaeea7c4fe0d764e
// ff:
// t:
// value:
// expect:
func (t *T) Assert(value, expect interface{}) {
	Assert(value, expect)
}

// AssertEQ 检查 `value` 和 `expect` 是否相等，包括它们的类型。 md5:31097fa6b823a25a
// ff:
// t:
// value:
// expect:
func (t *T) AssertEQ(value, expect interface{}) {
	AssertEQ(value, expect)
}

// AssertNE 检查 `value` 和 `expect` 是否不相等。 md5:418e91b330bc944f
// ff:
// t:
// value:
// expect:
func (t *T) AssertNE(value, expect interface{}) {
	AssertNE(value, expect)
}

// AssertNQ 检查 `value` 和 `expect` 是否不相等，包括它们的类型。 md5:bb13af00897290db
// ff:
// t:
// value:
// expect:
func (t *T) AssertNQ(value, expect interface{}) {
	AssertNQ(value, expect)
}

// AssertGT 检查 `value` 是否大于 `expect`。
// 注意，只有字符串、整数和浮点数类型能使用 AssertGT 进行比较，
// 其他类型是无效的。
// md5:647270894818c6c7
// ff:
// t:
// value:
// expect:
func (t *T) AssertGT(value, expect interface{}) {
	AssertGT(value, expect)
}

// AssertGE 检查 `value` 是否大于或等于 `expect`。
// 请注意，只有字符串、整数和浮点数类型可以使用 AssertGE 进行比较，其他类型是无效的。
// md5:3227e007891ed72e
// ff:
// t:
// value:
// expect:
func (t *T) AssertGE(value, expect interface{}) {
	AssertGE(value, expect)
}

// AssertLT 检查 `value` 是否小于等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以通过 AssertLT 进行比较，其他类型无效。
// md5:784a9db44c03122b
// ff:
// t:
// value:
// expect:
func (t *T) AssertLT(value, expect interface{}) {
	AssertLT(value, expect)
}

// AssertLE 检查 `value` 是否小于或等于 `expect`。
// 请注意，只有字符串、整数和浮点类型可以通过 AssertLTE 进行比较，其他类型的值是无效的。
// md5:bca4df91bef4e152
// ff:
// t:
// value:
// expect:
func (t *T) AssertLE(value, expect interface{}) {
	AssertLE(value, expect)
}

// AssertIN 检查 `value` 是否在 `expect` 中。
// `expect` 应该是一个切片，
// 但 `value` 可以是一个切片或一个基本类型变量。
// md5:596913e44fc64a93
// ff:
// t:
// value:
// expect:
func (t *T) AssertIN(value, expect interface{}) {
	AssertIN(value, expect)
}

// AssertNI 检查 `value` 不在 `expect` 列表中。
// `expect` 应该是一个切片，
// 而 `value` 可以是一个切片或基本类型变量。
// md5:8657bc61646e08fc
// ff:
// t:
// value:
// expect:
func (t *T) AssertNI(value, expect interface{}) {
	AssertNI(value, expect)
}

// AssertNil 断言 `value` 为 nil。 md5:94a00206ff503e10
// ff:
// t:
// value:
func (t *T) AssertNil(value interface{}) {
	AssertNil(value)
}

// 使用给定的`message`引发错误恐慌。 md5:6ddb84d91c681d1f
// ff:
// t:
// message:
func (t *T) Error(message ...interface{}) {
	Error(message...)
}

// Fatal 将 `message` 打印到 stderr 并退出进程。 md5:15e177961f66ebe7
// ff:
// t:
// message:
func (t *T) Fatal(message ...interface{}) {
	Fatal(message...)
}
