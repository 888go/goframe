
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// C creates a unit testing case.
// The parameter `t` is the pointer to testing.T of stdlib (*testing.T).
// The parameter `f` is the closure function for unit testing case.
<原文结束>

# <翻译开始>
// C 创建一个单元测试用例。
// 参数 `t` 是标准库 (*testing.T) 的 testing.T 指针。
// 参数 `f` 是用于单元测试的闭包函数。
// md5:0a3ae380343ea962
# <翻译结束>


<原文开始>
// Assert checks `value` and `expect` EQUAL.
<原文结束>

# <翻译开始>
// Assert 检查 `value` 和 `expect` 是否相等。 md5:eaeea7c4fe0d764e
# <翻译结束>


<原文开始>
// AssertEQ checks `value` and `expect` EQUAL, including their TYPES.
<原文结束>

# <翻译开始>
// AssertEQ 检查 `value` 和 `expect` 是否相等，包括它们的类型。 md5:31097fa6b823a25a
# <翻译结束>


<原文开始>
// AssertNE checks `value` and `expect` NOT EQUAL.
<原文结束>

# <翻译开始>
// AssertNE 检查 `value` 和 `expect` 是否不相等。 md5:418e91b330bc944f
# <翻译结束>


<原文开始>
// AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.
<原文结束>

# <翻译开始>
// AssertNQ 检查 `value` 和 `expect` 是否不相等，包括它们的类型。 md5:bb13af00897290db
# <翻译结束>


<原文开始>
// AssertGT checks `value` is GREATER THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertGT,
// others are invalid.
<原文结束>

# <翻译开始>
// AssertGT 检查 `value` 是否大于 `expect`。
// 注意，只有字符串、整数和浮点数类型能使用 AssertGT 进行比较，
// 其他类型是无效的。
// md5:647270894818c6c7
# <翻译结束>


<原文开始>
// AssertGE checks `value` is GREATER OR EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertGTE,
// others are invalid.
<原文结束>

# <翻译开始>
// AssertGE 检查 `value` 是否大于或等于 `expect`。
// 请注意，只有字符串、整数和浮点数类型可以使用 AssertGE 进行比较，其他类型是无效的。
// md5:3227e007891ed72e
# <翻译结束>


<原文开始>
// AssertLT checks `value` is LESS EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertLT,
// others are invalid.
<原文结束>

# <翻译开始>
// AssertLT 检查 `value` 是否小于等于 `expect`。
// 注意，只有字符串、整数和浮点类型可以通过 AssertLT 进行比较，其他类型无效。
// md5:784a9db44c03122b
# <翻译结束>


<原文开始>
// AssertLE checks `value` is LESS OR EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertLTE,
// others are invalid.
<原文结束>

# <翻译开始>
// AssertLE 检查 `value` 是否小于或等于 `expect`。
// 请注意，只有字符串、整数和浮点类型可以通过 AssertLTE 进行比较，其他类型的值是无效的。
// md5:bca4df91bef4e152
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
// `expect` 应该是一个切片，
// 但是 `value` 可以是切片或基本类型变量。
// TODO: 添加对 map 的支持。
// TODO: gconv.Strings(0) 不应该转换为 `[0]`。
// md5:d8391e0c6cba6480
# <翻译结束>


<原文开始>
// AssertNI checks `value` is NOT IN `expect`.
// The `expect` should be a slice,
// but the `value` can be a slice or a basic type variable.
// TODO map support.
<原文结束>

# <翻译开始>
// AssertNI 检查 `value` 不在 `expect` 中。
// `expect` 应该是一个切片，
// 但是 `value` 可以是切片或基本类型变量。
// TODO 增加对 map 的支持。
// md5:483febd56930eb64
# <翻译结束>


<原文开始>
// Error panics with given `message`.
<原文结束>

# <翻译开始>
// 使用给定的`message`引发错误恐慌。 md5:6ddb84d91c681d1f
# <翻译结束>


<原文开始>
// Fatal prints `message` to stderr and exit the process.
<原文结束>

# <翻译开始>
// Fatal 将 `message` 打印到 stderr 并退出进程。 md5:15e177961f66ebe7
# <翻译结束>


<原文开始>
// compareMap compares two maps, returns nil if they are equal, or else returns error.
<原文结束>

# <翻译开始>
// compareMap 比较两个地图，如果它们相等则返回nil，否则返回错误。 md5:fd402375a76c3a4a
# <翻译结束>


<原文开始>
				// Turn two interface maps to the same type for comparison.
				// Direct use of rvValue.MapIndex(key).Interface() will panic
				// when the key types are inconsistent.
<原文结束>

# <翻译开始>
// 将两个接口映射转换为相同类型以便进行比较。
// 直接使用rvValue.MapIndex(key).Interface() 当键的类型不一致时，会导致恐慌。
// md5:ae85735772c34002
# <翻译结束>


<原文开始>
// AssertNil asserts `value` is nil.
<原文结束>

# <翻译开始>
// AssertNil 断言 `value` 为 nil。 md5:94a00206ff503e10
# <翻译结束>


<原文开始>
// DataPath retrieves and returns the testdata path of current package,
// which is used for unit testing cases only.
// The optional parameter `names` specifies the sub-folders/sub-files,
// which will be joined with current system separator and returned with the path.
<原文结束>

# <翻译开始>
// DataPath获取并返回当前包的测试数据路径，仅用于单元测试。
// 可选参数`names`指定了子文件夹/子文件，将与当前系统的分隔符连接，并与路径一起返回。
// md5:55efb430c9f8a73f
# <翻译结束>


<原文开始>
// DataContent retrieves and returns the file content for specified testdata path of current package
<原文结束>

# <翻译开始>
// DataContent 从当前包的特定测试数据路径中检索并返回文件内容. md5:26224495ddbd389e
# <翻译结束>

