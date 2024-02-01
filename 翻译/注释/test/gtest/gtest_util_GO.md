
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// C creates a unit testing case.
// The parameter `t` is the pointer to testing.T of stdlib (*testing.T).
// The parameter `f` is the closure function for unit testing case.
<原文结束>

# <翻译开始>
// C 创建一个单元测试用例。
// 参数 `t` 是指向标准库 testing.T 的指针（*testing.T）。
// 参数 `f` 是用于单元测试用例的闭包函数。
# <翻译结束>


<原文开始>
// Assert checks `value` and `expect` EQUAL.
<原文结束>

# <翻译开始>
// Assert 检查 `value` 和 `expect` 是否相等。
# <翻译结束>


<原文开始>
// AssertEQ checks `value` and `expect` EQUAL, including their TYPES.
<原文结束>

# <翻译开始>
// AssertEQ 检查 `value` 和 `expect` 是否相等，包括它们的 TYPE（类型）。
# <翻译结束>


<原文开始>
// AssertNE checks `value` and `expect` NOT EQUAL.
<原文结束>

# <翻译开始>
// AssertNE 检查 `value` 和 `expect` 是否不相等。
# <翻译结束>


<原文开始>
// AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.
<原文结束>

# <翻译开始>
// AssertNQ 检查 `value` 和 `expect` 是否不相等，包括它们的类型。
# <翻译结束>


<原文开始>
// AssertGT checks `value` is GREATER THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertGT,
// others are invalid.
<原文结束>

# <翻译开始>
// AssertGT 检查 `value` 是否大于 `expect`。
// 注意，只有字符串、整数和浮点类型可以通过 AssertGT 进行比较，
// 其他类型是无效的。
# <翻译结束>


<原文开始>
// AssertGE checks `value` is GREATER OR EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertGTE,
// others are invalid.
<原文结束>

# <翻译开始>
// AssertGE 检查 `value` 是否大于或等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以使用 AssertGTE 进行比较，其他类型无效。
# <翻译结束>


<原文开始>
// AssertLT checks `value` is LESS EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertLT,
// others are invalid.
<原文结束>

# <翻译开始>
// AssertLT 检查 `value` 是否小于等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以使用 AssertLT 进行比较，
// 其他类型无效。
# <翻译结束>


<原文开始>
// AssertLE checks `value` is LESS OR EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertLTE,
// others are invalid.
<原文结束>

# <翻译开始>
// AssertLE 检查 `value` 是否小于等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以通过 AssertLTE 进行比较，
// 其他类型是无效的。
# <翻译结束>


<原文开始>
// AssertIN checks `value` is IN `expect`.
// The `expect` should be a slice,
// but the `value` can be a slice or a basic type variable.
// TODO map support.
// TODO: gconv.Strings(0) is not [0]
<原文结束>

# <翻译开始>
// AssertIN 检查 `value` 是否在 `expect` 中。
// `expect` 应该是一个切片类型，
// 但 `value` 可以是切片类型或基本类型变量。
// TODO: 添加对 map 类型的支持。
// TODO: gconv.Strings(0) 的结果不是 [0]
# <翻译结束>


<原文开始>
// AssertNI checks `value` is NOT IN `expect`.
// The `expect` should be a slice,
// but the `value` can be a slice or a basic type variable.
// TODO map support.
<原文结束>

# <翻译开始>
// AssertNI 检查 `value` 是否不在 `expect` 中。
// `expect` 应该是一个切片，
// 但 `value` 可以是切片或基本类型变量。
// TODO: 添加对 map 的支持。
# <翻译结束>


<原文开始>
// Error panics with given `message`.
<原文结束>

# <翻译开始>
// Error 使用给定的`message`引发panic异常。
# <翻译结束>


<原文开始>
// Fatal prints `message` to stderr and exit the process.
<原文结束>

# <翻译开始>
// Fatal将`message`打印到标准错误输出（stderr）并退出进程。
# <翻译结束>


<原文开始>
// compareMap compares two maps, returns nil if they are equal, or else returns error.
<原文结束>

# <翻译开始>
// compareMap 比较两个映射，如果它们相等则返回 nil，否则返回错误。
# <翻译结束>


<原文开始>
				// Turn two interface maps to the same type for comparison.
				// Direct use of rvValue.MapIndex(key).Interface() will panic
				// when the key types are inconsistent.
<原文结束>

# <翻译开始>
// 将两个接口映射转换为同一类型以便进行比较。
// 若直接使用rvValue.MapIndex(key).Interface()，当键类型不一致时会触发 panic。
# <翻译结束>


<原文开始>
// AssertNil asserts `value` is nil.
<原文结束>

# <翻译开始>
// AssertNil 断言 `value` 为 nil。
# <翻译结束>


<原文开始>
// DataPath retrieves and returns the testdata path of current package,
// which is used for unit testing cases only.
// The optional parameter `names` specifies the sub-folders/sub-files,
// which will be joined with current system separator and returned with the path.
<原文结束>

# <翻译开始>
// DataPath 获取并返回当前包的 testdata 路径，
// 该路径仅用于单元测试用例。
// 可选参数 `names` 指定子文件夹/子文件，
// 这些名称将与当前系统分隔符连接，并与路径一起返回。
# <翻译结束>


<原文开始>
// DataContent retrieves and returns the file content for specified testdata path of current package
<原文结束>

# <翻译开始>
// DataContent 函数用于获取并返回当前包中指定testdata路径下的文件内容
# <翻译结束>

