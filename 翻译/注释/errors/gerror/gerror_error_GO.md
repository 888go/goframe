
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
// Error is custom error for additional features.
<原文结束>

# <翻译开始>
// Error 是自定义错误，用于附加功能。 md5:6111a19ebbc88a60
# <翻译结束>


<原文开始>
// Stack array, which records the stack information when this error is created or wrapped.
<原文结束>

# <翻译开始>
// Stack数组，用于记录在创建或包装此错误时的堆栈信息。 md5:a86aaae3cd5b8beb
# <翻译结束>


<原文开始>
// Custom Error text when Error is created, might be empty when its code is not nil.
<原文结束>

# <翻译开始>
// 当创建错误时自定义错误文本，当错误代码不为nil时可能会为空。 md5:42ad3f2e459c445e
# <翻译结束>


<原文开始>
// Error code if necessary.
<原文结束>

# <翻译开始>
// 如果需要，错误代码。 md5:92ec689bb1dbbb33
# <翻译结束>


<原文开始>
// Filtering key for current error module paths.
<原文结束>

# <翻译开始>
	// 当前错误模块路径的过滤键。 md5:c4b4987b15b2caf0
# <翻译结束>


<原文开始>
// goRootForFilter is used for stack filtering in development environment purpose.
<原文结束>

# <翻译开始>
	// goRootForFilter 用于开发环境中栈过滤的目的。 md5:df10489e92979e5e
# <翻译结束>


<原文开始>
// Error implements the interface of Error, it returns all the error as string.
<原文结束>

# <翻译开始>
// Error 实现了 Error 接口，它返回所有的错误信息作为字符串。 md5:916d521fe191e82f
# <翻译结束>


<原文开始>
// Cause returns the root cause error.
<原文结束>

# <翻译开始>
// Cause返回根本原因错误。 md5:c43631d8af1a0815
# <翻译结束>


<原文开始>
// Other Error that implements ApiCause interface.
<原文结束>

# <翻译开始>
				// 实现ApiCause接口的其他错误。 md5:50f12dd0449fe932
# <翻译结束>


<原文开始>
			// return loop
			//
			// To be compatible with Case of https://github.com/pkg/errors.
<原文结束>

# <翻译开始>
			// 返回循环
			//
			// 以兼容 https:			//github.com/pkg/errors 中的 Case 情况。
			// md5:a923900fc4a93e9d
# <翻译结束>


<原文开始>
// Current creates and returns the current level error.
// It returns nil if current level error is nil.
<原文结束>

# <翻译开始>
// Current 创建并返回当前级别的错误。如果当前级别错误为 nil，则返回 nil。
// md5:d8b26e22ec63a837
# <翻译结束>


<原文开始>
// Unwrap is alias of function `Next`.
// It is just for implements for stdlib errors.Unwrap from Go version 1.17.
<原文结束>

# <翻译开始>
// Unwrap 是函数 `Next` 的别名。
// 它只是为了实现自 Go 1.17 版本的stdlib库中的 `errors.Unwrap`。
// md5:4ab7dcc4181801cd
# <翻译结束>


<原文开始>
// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
<原文结束>

# <翻译开始>
// Equal 判断当前错误 `err` 是否等于目标错误 `target`。请注意，在默认的错误比较中，如果两个错误的 `code` 和 `text` 都相同，那么它们将被视为相等。
// md5:6256ec44e7b04b0e
# <翻译结束>


<原文开始>
	// Code should be the same.
	// Note that if both errors have `nil` code, they are also considered equal.
<原文结束>

# <翻译开始>
	// 代码应该保持不变。
	// 注意，如果两个错误的代码都为`nil`，那么它们也会被视为相等。
	// md5:9cd5037f48adc142
# <翻译结束>


<原文开始>
// Text should be the same.
<原文结束>

# <翻译开始>
	// 文本内容应保持一致。 md5:950f9f350f074b9c
# <翻译结束>


<原文开始>
// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.
<原文结束>

# <翻译开始>
// Is 判断当前错误 `err` 是否在其嵌套错误中包含目标错误 `target`。这是为了实现从 Go 1.17 版本开始的标准库中的 errors.Is 接口。
// md5:dfc92c8d3ba58133
# <翻译结束>

