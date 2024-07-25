
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// NewCode creates and returns an error that has error code and given text.
<原文结束>

# <翻译开始>
// NewCode 创建并返回一个具有错误代码和给定文本的错误。 md5:5f88f8ae1151acac
# <翻译结束>


<原文开始>
// NewCodef returns an error that has error code and formats as the given format and args.
<原文结束>

# <翻译开始>
// NewCodef 返回一个具有错误代码，并按照给定格式和参数格式化的错误。 md5:bb6b90ee5a4ce175
# <翻译结束>


<原文开始>
// NewCodeSkip creates and returns an error which has error code and is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// NewCodeSkip 创建并返回一个带有错误码的错误，该错误根据给定的文本格式化。
// 参数 `skip` 指定了跳过的堆栈调用者数量。 md5:5c3aabed2ce89e0c
# <翻译结束>


<原文开始>
// NewCodeSkipf returns an error that has error code and formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// NewCodeSkipf 返回一个具有指定错误代码和格式化参数的错误。
// 参数 `skip` 指定了要跳过的调用堆栈数量。 md5:ccd3b74e8b4f8acc
# <翻译结束>


<原文开始>
// WrapCode wraps error with code and text.
// It returns nil if given err is nil.
<原文结束>

# <翻译开始>
// WrapCode 将错误与代码和文本一起包装。
// 如果给定的 err 为 nil，它将返回 nil。 md5:5e09a5ffb6fa4e21
# <翻译结束>


<原文开始>
// WrapCodef wraps error with code and format specifier.
// It returns nil if given `err` is nil.
<原文结束>

# <翻译开始>
// WrapCodef 将错误与代码和格式化占位符一起包装。
// 如果给定的 `err` 为 nil，它将返回 nil。 md5:ef3a7436eb342ff6
# <翻译结束>


<原文开始>
// WrapCodeSkip wraps error with code and text.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// WrapCodeSkip 用于给错误附加代码和文本信息。
// 如果给定的err为nil，该函数将返回nil。
// 参数 `skip` 指定了要跳过的堆栈调用者数量。 md5:5ee348edd866b587
# <翻译结束>


<原文开始>
// WrapCodeSkipf wraps error with code and text that is formatted with given format and args.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// WrapCodeSkipf 使用给定的格式和参数将错误包装成带有代码和文本的错误。
// 如果给定的err为nil，它将返回nil。
// 参数`skip`指定了要跳过的调用者堆栈的数量。 md5:00fbaefc556da645
# <翻译结束>


<原文开始>
// Code returns the error code of `current error`.
// It returns `CodeNil` if it has no error code neither it does not implement interface Code.
<原文结束>

# <翻译开始>
// Code 函数返回当前错误的错误代码。
// 如果它没有错误代码并且也没有实现 Code 接口，它将返回 CodeNil。 md5:33b7429f6f7b3dbe
# <翻译结束>


<原文开始>
// HasCode checks and reports whether `err` has `code` in its chaining errors.
<原文结束>

# <翻译开始>
// HasCode 检查并报告 `err` 的链式错误中是否包含 `code`。 md5:5d1b8286d1872717
# <翻译结束>

