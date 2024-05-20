
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
// Print prints `v` with newline using fmt.Sprintln.
// The parameter `v` can be multiple variables.
<原文结束>

# <翻译开始>
// Print 使用 fmt.Sprintln 函数打印 `v` 并添加换行符。
// 参数 `v` 可以是多个变量。
// md5:6c0b3b96234f77ce
# <翻译结束>


<原文开始>
// Printf prints `v` with format `format` using fmt.Sprintf.
// The parameter `v` can be multiple variables.
<原文结束>

# <翻译开始>
// Printf 使用 fmt.Sprintf 根据格式 `format` 打印 `v`。
// 参数 `v` 可以是多个变量。
// md5:e3b6ab3d8750ad4c
# <翻译结束>


<原文开始>
// Fatal prints the logging content with [FATA] header and newline, then exit the current process.
<原文结束>

# <翻译开始>
// Fatal 以[FATA]标题和换行符打印日志内容，然后退出当前进程。. md5:4061b6551d7137a1
# <翻译结束>


<原文开始>
// Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.
<原文结束>

# <翻译开始>
// Fatalf 打印带有 [FATA] 头部、自定义格式和换行符的日志内容，然后退出当前进程。. md5:cbaf3fb7e2b92df9
# <翻译结束>


<原文开始>
// Panic prints the logging content with [PANI] header and newline, then panics.
<原文结束>

# <翻译开始>
// Panic 在输出带有 [PANI] 头部和换行符的日志内容后，引发 panic。. md5:bcde7bf5ff45073a
# <翻译结束>


<原文开始>
// Panicf prints the logging content with [PANI] header, custom format and newline, then panics.
<原文结束>

# <翻译开始>
// Panicf 函数打印带有 [PANI] 标头、自定义格式和换行的日志内容，然后触发恐慌（panic）。. md5:bf1df3ad28ded71f
# <翻译结束>


<原文开始>
// Info prints the logging content with [INFO] header and newline.
<原文结束>

# <翻译开始>
// Info 打印带有 "[INFO]" 标头和换行符的日志内容。. md5:1ed8e917ecca5ef4
# <翻译结束>


<原文开始>
// Infof prints the logging content with [INFO] header, custom format and newline.
<原文结束>

# <翻译开始>
// Infof 打印带有 [INFO] 标头、自定义格式和换行符的日志内容。. md5:fda1e57b2e2089d7
# <翻译结束>


<原文开始>
// Debug prints the logging content with [DEBU] header and newline.
<原文结束>

# <翻译开始>
// Debug 打印带有 [DEBU] 标头和换行符的日志内容。. md5:7a64f5ebf48d4f92
# <翻译结束>


<原文开始>
// Debugf prints the logging content with [DEBU] header, custom format and newline.
<原文结束>

# <翻译开始>
// Debugf 函数打印带有 [DEBU] 标头、自定义格式化内容和换行符的日志信息。. md5:0c0164b88b59a40c
# <翻译结束>


<原文开始>
// Notice prints the logging content with [NOTI] header and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// ```go
// Notice 打印日志内容，前缀为 [NOTI] 并在末尾添加换行符。
// 如果启用了堆栈跟踪功能，它还会打印调用者堆栈信息。
// ```
// md5:c36d548c618d1251
# <翻译结束>


<原文开始>
// Noticef prints the logging content with [NOTI] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Noticef 打印带有 [NOTI] 标头的日志内容，自定义格式和换行符。如果启用了堆栈功能，它还会打印调用者堆栈信息。
// md5:f9d4f5af91856cd9
# <翻译结束>


<原文开始>
// Warning prints the logging content with [WARN] header and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Warning 打印带有 [WARN] 头部和换行符的日志内容。如果启用了堆栈功能，它还会打印调用者堆栈信息。
// md5:8142c86f6be53ee0
# <翻译结束>


<原文开始>
// Warningf prints the logging content with [WARN] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Warningf 函数打印带有 [WARN] 标头的记录内容、自定义格式化字符串以及换行。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
// md5:025f0baa4a1f8600
# <翻译结束>


<原文开始>
// Error prints the logging content with [ERRO] header and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Error 打印带有 [ERRO] 标头和换行符的日志内容。
// 如果启用了堆栈跟踪功能，它还会打印调用者堆栈信息。
// md5:f2aa6f6c0e4d2061
# <翻译结束>


<原文开始>
// Errorf prints the logging content with [ERRO] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Errorf 函数会打印带有 [ERRO] 标头的日志内容，使用自定义格式并添加换行符。如果启用了堆栈功能，它还会打印调用者堆栈信息。
// md5:4a90789d1de07943
# <翻译结束>


<原文开始>
// Critical prints the logging content with [CRIT] header and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Critical 打印带有 [CRIT] 头部和换行符的日志内容。如果启用了堆栈功能，它还会打印调用者堆栈信息。
// md5:f9fb0900ff8f602f
# <翻译结束>


<原文开始>
// Criticalf prints the logging content with [CRIT] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
<原文结束>

# <翻译开始>
// Criticalf 函数打印带有 [CRIT] 标头、自定义格式和换行的日志内容。
// 如果启用了堆栈追踪功能，它还会打印调用者堆栈信息。
// md5:fa381bbe7b0465d0
# <翻译结束>


<原文开始>
// checkLevel checks whether the given `level` could be output.
<原文结束>

# <翻译开始>
// checkLevel 检查给定的 `level` 是否可以输出。. md5:59e82f73882a5ac4
# <翻译结束>

