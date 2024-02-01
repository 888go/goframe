
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
// NewCode creates and returns an error that has error code and given text.
<原文结束>

# <翻译开始>
// NewCode创建并返回一个错误，该错误包含错误代码和给定的文本。
# <翻译结束>


<原文开始>
// NewCodef returns an error that has error code and formats as the given format and args.
<原文结束>

# <翻译开始>
// NewCodef返回一个错误，该错误包含错误代码，并使用给定的格式和参数进行格式化。
# <翻译结束>


<原文开始>
// NewCodeSkip creates and returns an error which has error code and is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// NewCodeSkip 创建并返回一个错误，该错误包含错误代码，并根据给定文本格式化而成。
// 参数`skip`用于指定跳过堆栈调用者的数量。
# <翻译结束>


<原文开始>
// NewCodeSkipf returns an error that has error code and formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// NewCodeSkipf 返回一个具有错误代码且格式化为给定格式和参数的错误。
// 参数 `skip` 指定了要跳过的堆栈调用者数量。
# <翻译结束>


<原文开始>
// WrapCode wraps error with code and text.
// It returns nil if given err is nil.
<原文结束>

# <翻译开始>
// WrapCode 将错误用代码和文本包装起来。
// 如果给出的 err 为 nil，则返回 nil。
# <翻译结束>


<原文开始>
// WrapCodef wraps error with code and format specifier.
// It returns nil if given `err` is nil.
<原文结束>

# <翻译开始>
// WrapCodef 将错误用代码和格式符包装起来。
// 如果给出的 `err` 为 nil，则返回 nil。
# <翻译结束>


<原文开始>
// WrapCodeSkip wraps error with code and text.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// WrapCodeSkip 通过代码和文本包装错误。
// 如果给出的 err 为 nil，则返回 nil。
// 参数 `skip` 指定了需要跳过的堆栈调用者数量。
# <翻译结束>


<原文开始>
// WrapCodeSkipf wraps error with code and text that is formatted with given format and args.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
<原文结束>

# <翻译开始>
// WrapCodeSkipf 对给定的错误 err 进行包装，添加代码和格式化后的文本信息。如果给定的 err 为 nil，则返回 nil。
// 参数 `skip` 指定了需要跳过的堆栈调用者数量。
# <翻译结束>


<原文开始>
// Code returns the error code of current error.
// It returns `CodeNil` if it has no error code neither it does not implement interface Code.
<原文结束>

# <翻译开始>
// Code 返回当前错误的错误代码。
// 如果该错误没有错误代码，或者未实现 Code 接口，它将返回 `CodeNil`。
# <翻译结束>


<原文开始>
// HasCode checks and reports whether `err` has `code` in its chaining errors.
<原文结束>

# <翻译开始>
// HasCode 检查并报告 `err` 在其链式错误中是否包含 `code`。
# <翻译结束>

