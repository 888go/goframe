
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
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Error is custom error for additional features.
<原文结束>

# <翻译开始>
// Error 是自定义错误类型，用于提供额外功能。
# <翻译结束>


<原文开始>
// Stack array, which records the stack information when this error is created or wrapped.
<原文结束>

# <翻译开始>
// Stack 数组，用于记录当此错误创建或被包裹时的堆栈信息。
# <翻译结束>


<原文开始>
// Custom Error text when Error is created, might be empty when its code is not nil.
<原文结束>

# <翻译开始>
// 当 Error 被创建时的自定义错误文本，如果其代码不为空，则可能为空。
# <翻译结束>







<原文开始>
// Filtering key for current error module paths.
<原文结束>

# <翻译开始>
// 用于当前错误模块路径的过滤键。
# <翻译结束>


<原文开始>
// goRootForFilter is used for stack filtering in development environment purpose.
<原文结束>

# <翻译开始>
// goRootForFilter 用于在开发环境目的中进行堆栈过滤。
# <翻译结束>


<原文开始>
// Error implements the interface of Error, it returns all the error as string.
<原文结束>

# <翻译开始>
// Error 实现了 Error 接口，它将所有错误以字符串形式返回。
# <翻译结束>


<原文开始>
// Cause returns the root cause error.
<原文结束>

# <翻译开始>
// Cause 返回根本原因错误。
# <翻译结束>







<原文开始>
// Other Error that implements ApiCause interface.
<原文结束>

# <翻译开始>
// 其他实现了ApiCause接口的错误类型。
# <翻译结束>


<原文开始>
			// return loop
			//
			// To be compatible with Case of https://github.com/pkg/errors.
<原文结束>

# <翻译开始>
// 返回循环
//
// 为了与 https://github.com/pkg/errors 包中的 Case 兼容。
# <翻译结束>


<原文开始>
// Current creates and returns the current level error.
// It returns nil if current level error is nil.
<原文结束>

# <翻译开始>
// Current函数创建并返回当前层级的错误信息。
// 如果当前层级的错误信息为nil，则该函数返回nil。
# <翻译结束>


<原文开始>
// Unwrap is alias of function `Next`.
// It is just for implements for stdlib errors.Unwrap from Go version 1.17.
<原文结束>

# <翻译开始>
// Unwrap 是函数 `Next` 的别名。
// 它仅为实现从 Go 1.17 版本开始的stdlib errors.Unwrap 接口而存在。
# <翻译结束>


<原文开始>
// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
<原文结束>

# <翻译开始>
// Equal 判断当前错误 `err` 是否与目标错误 `target` 相等。
// 请注意，在默认的 `Error` 比较方式下，
// 如果两个错误的 `code` 和 `text` 都相同，则认为它们是相同的错误。
# <翻译结束>


<原文开始>
	// Code should be the same.
	// Note that if both errors have `nil` code, they are also considered equal.
<原文结束>

# <翻译开始>
// 代码应当保持一致。
// 注意，如果两个错误的code都是`nil`，则认为它们也是相等的。
# <翻译结束>







<原文开始>
// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.
<原文结束>

# <翻译开始>
// Is 报告当前错误 `err` 在其链式错误中是否包含错误 `target`。
// 这只是为了实现从 Go 1.17 版本开始的stdlib errors.Is功能。
# <翻译结束>


<原文开始>
// Error code if necessary.
<原文结束>

# <翻译开始>
// 如果必要，此处为错误代码。
# <翻译结束>


<原文开始>
// Internal Error struct.
<原文结束>

# <翻译开始>
// 内部错误结构体。
# <翻译结束>


<原文开始>
// Text should be the same.
<原文结束>

# <翻译开始>
// 文本内容应该保持一致。
# <翻译结束>

