
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
// Print prints `v` with newline using fmt.Sprintln.
// The parameter `v` can be multiple variables.
<原文结束>

# <翻译开始>
// Print 使用 fmt.Sprintln 函数打印变量 `v` 及其换行。参数 `v` 可以是多个变量。
# <翻译结束>


<原文开始>
// Printf prints `v` with format `format` using fmt.Sprintf.
// The parameter `v` can be multiple variables.
<原文结束>

# <翻译开始>
// Printf 通过使用 fmt.Sprintf 格式化方法打印变量 `v`。
// 参数 `v` 可以是多个变量。
# <翻译结束>


<原文开始>
// Fatal prints the logging content with [FATA] header and newline, then exit the current process.
<原文结束>

# <翻译开始>
// Fatal 打印带有[FATA]头部和换行符的日志内容，然后退出当前进程。
# <翻译结束>


<原文开始>
// Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.
<原文结束>

# <翻译开始>
// Fatalf 打印日志内容，其头部为[FATA]，采用自定义格式，并在末尾添加换行符，然后退出当前进程。
# <翻译结束>


<原文开始>
// Panic prints the logging content with [PANI] header and newline, then panics.
<原文结束>

# <翻译开始>
// Panic 会打印带有[PANI]头部和换行符的日志内容，然后触发 panic。
# <翻译结束>


<原文开始>
// Panicf prints the logging content with [PANI] header, custom format and newline, then panics.
<原文结束>

# <翻译开始>
// Panicf 函数会打印带有[PANI]头部、自定义格式以及换行符的日志内容，然后触发 panic。
# <翻译结束>


<原文开始>
// Info prints the logging content with [INFO] header and newline.
<原文结束>

# <翻译开始>
// Info打印日志内容，前面带有[INFO]头部和换行符。
# <翻译结束>


<原文开始>
// Infof prints the logging content with [INFO] header, custom format and newline.
<原文结束>

# <翻译开始>
// Infof 格式化并打印日志内容，带有 [INFO] 标头、自定义格式及换行。
// ```go
// Infof 根据提供的格式和参数，以 [INFO] 头部形式输出格式化信息，并在末尾添加换行。
// 示例用法：
// log.Infof("用户 %s 登录成功", username)
# <翻译结束>


<原文开始>
// Debug prints the logging content with [DEBU] header and newline.
<原文结束>

# <翻译开始>
// Debug 以 [DEBU] 标头和换行符打印日志内容。
# <翻译结束>


<原文开始>
// Debugf prints the logging content with [DEBU] header, custom format and newline.
<原文结束>

# <翻译开始>
// Debugf按照[DEBU]头部，自定义格式及换行打印日志内容。
# <翻译结束>


<原文开始>
// Notice prints the logging content with [NOTI] header and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Notice 以 [NOTI] 标头和换行符打印日志内容。
// 如果启用了堆栈功能，还会打印调用堆栈信息。
# <翻译结束>


<原文开始>
// Noticef prints the logging content with [NOTI] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Noticef函数以[NOTI]头、自定义格式及换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用堆栈信息。
# <翻译结束>


<原文开始>
// Warning prints the logging content with [WARN] header and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Warning 以 [WARN] 标头和换行符打印日志内容。
// 如果启用了堆栈功能，它还会打印调用者堆栈信息。
# <翻译结束>


<原文开始>
// Warningf prints the logging content with [WARN] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Warningf函数会打印带有[WARN]头的记录内容，使用自定义格式并添加换行符。
// 如果启用了堆栈功能，它还会打印调用堆栈信息。
# <翻译结束>


<原文开始>
// Error prints the logging content with [ERRO] header and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Error 以 [ERRO] 标头和换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
# <翻译结束>


<原文开始>
// Errorf prints the logging content with [ERRO] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Errorf 函数以 [ERRO] 标头、自定义格式及换行符打印日志内容。
// 若启用了堆栈追踪功能，它还会打印调用者堆栈信息。
# <翻译结束>


<原文开始>
// Critical prints the logging content with [CRIT] header and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Critical 以 [CRIT] 标头和换行符打印日志内容。
// 如果启用了堆栈追踪功能，还会打印调用堆栈信息。
# <翻译结束>


<原文开始>
// Criticalf prints the logging content with [CRIT] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Criticalf以[CRIT]头部、自定义格式和换行符打印日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
# <翻译结束>


<原文开始>
// checkLevel checks whether the given `level` could be output.
<原文结束>

# <翻译开始>
// checkLevel 检查给定的 `level` 是否可以输出。
# <翻译结束>

