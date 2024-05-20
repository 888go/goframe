
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
// New creates and returns an error which is formatted from given text.
<原文结束>

# <翻译开始>
// New 创建并返回一个根据给定文本格式化的错误。. md5:de9ec7c958a945bb
# <翻译结束>


<原文开始>
// Newf returns an error that formats as the given format and args.
<原文结束>

# <翻译开始>
// Newf 返回一个根据给定格式和参数格式化的错误。. md5:bd62f35687f8bc83
# <翻译结束>


<原文开始>
// NewSkip creates and returns an error which is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// NewSkip 创建并返回一个根据给定文本格式化的错误。参数 `skip` 指定了要跳过的调用者堆栈数量。
// md5:22bec296ea4c17b0
# <翻译结束>


<原文开始>
// NewSkipf returns an error that formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// NewSkipf 返回一个按照给定格式和参数格式化的错误。
// 参数 `skip` 指定了跳过调用栈的层数。
// md5:82d8fef84b9d2ba0
# <翻译结束>


<原文开始>
// Wrap wraps error with text. It returns nil if given err is nil.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
<原文结束>

# <翻译开始>
// Wrap 使用文本包装错误。如果给定的 err 为 nil，则返回 nil。
// 注意，它不会丢失被包裹错误的错误码，因为它从被包裹的错误中继承了错误码。
// md5:e04f9222b50c8938
# <翻译结束>


<原文开始>
// Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the format specifier.
// It returns nil if given `err` is nil.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
<原文结束>

# <翻译开始>
// Wrapf 会在调用 Wrapf 的位置为错误 err 添加一个堆栈跟踪信息，并使用格式化指定器。
// 如果给定的 `err` 为 nil，它将返回 nil。
// 请注意，它不会丢失被包装错误的错误代码，因为它从错误中继承了错误代码。
// md5:cbfccfaa6fa0bee1
# <翻译结束>


<原文开始>
// WrapSkip wraps error with text. It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
<原文结束>

# <翻译开始>
// WrapSkip 使用文本包装错误。如果给定的 err 为 nil，它将返回 nil。
// 参数 `skip` 指定了跳过调用堆栈的层数。
// 注意，它不会丢失被包装错误的错误代码，因为它是从其继承错误代码的。
// md5:5f87402ce06c586b
# <翻译结束>


<原文开始>
// WrapSkipf wraps error with text that is formatted with given format and args. It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
<原文结束>

# <翻译开始>
// WrapSkipf 将错误用给定的格式和参数进行文本包装。如果给定的 err 为 nil，它将返回 nil。
// 参数 `skip` 指定了要跳过的调用栈层数。
// 注意，它不会丢失被包装错误的错误代码，因为它是从原始错误中继承错误代码的。
// md5:82d4f5ae39c67b27
# <翻译结束>

