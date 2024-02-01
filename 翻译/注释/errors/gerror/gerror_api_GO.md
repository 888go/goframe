
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
// New creates and returns an error which is formatted from given text.
<原文结束>

# <翻译开始>
// New根据给定的文本创建并返回一个格式化后的错误。
# <翻译结束>


<原文开始>
// Newf returns an error that formats as the given format and args.
<原文结束>

# <翻译开始>
// Newf 返回一个错误，其格式化输出为给定的 format 和 args。
# <翻译结束>


<原文开始>
// NewSkip creates and returns an error which is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// NewSkip 创建并返回一个根据给定文本格式化的错误。
// 参数 `skip` 指定了跳过堆栈调用者的数量。
# <翻译结束>


<原文开始>
// NewSkipf returns an error that formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// NewSkipf 返回一个格式化为给定格式和参数的错误。
// 参数 `skip` 指定了跳过的调用栈层数量。
# <翻译结束>


<原文开始>
// Wrap wraps error with text. It returns nil if given err is nil.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
<原文结束>

# <翻译开始>
// Wrap 使用文本包装错误。如果给出的 err 为 nil，则返回 nil。
// 注意，它不会丢失被包装错误的错误码，因为它会继承该错误的错误码。
# <翻译结束>


<原文开始>
// Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the format specifier.
// It returns nil if given `err` is nil.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
<原文结束>

# <翻译开始>
// Wrapf 返回一个错误，该错误在Wrapf调用的位置添加了堆栈跟踪信息以及格式化说明符。
// 如果给出的`err`为nil，则返回nil。
// 注意，它不会丢失被包装错误的错误码，因为它会继承该错误的错误码。
# <翻译结束>


<原文开始>
// WrapSkip wraps error with text. It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
<原文结束>

# <翻译开始>
// WrapSkip 用于用文本包装错误。如果给定的 err 为 nil，则返回 nil。
// 参数 `skip` 指定了跳过的堆栈调用者数量。
// 注意，它不会丢失包装错误的错误码，因为它会继承该错误的错误码。
# <翻译结束>


<原文开始>
// WrapSkipf wraps error with text that is formatted with given format and args. It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
<原文结束>

# <翻译开始>
// WrapSkipf 通过给定的格式和参数对错误进行包装并添加文本描述。如果给定的err为nil，此函数将返回nil。
// 参数`skip`用于指定跳过堆栈调用者的数量。
// 注意，它不会丢失包装错误的错误码，因为它会继承该错误的错误码。
// 以下是翻译后更详尽的中文注释：
// ```go
// WrapSkipf 函数用于对传入的错误 err 进行包装，同时使用给定的 format 和 args 格式化输出附加的错误信息。
// 若传入的原始错误 err 为 nil，则 WrapSkipf 函数直接返回 nil。
// 参数 `skip` 指定了在获取堆栈信息时需要跳过的调用层级数量。
// 需要注意的是，WrapSkipf 在对错误进行包装的过程中，会保留原始错误的错误代码，因为它从原始错误中继承了该错误代码属性。
# <翻译结束>

