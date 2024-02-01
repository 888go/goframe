
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
// utilPProf is the PProf interface implementer.
<原文结束>

# <翻译开始>
// utilPProf 是 PProf 接口的实现者。
# <翻译结束>


<原文开始>
// StartPProfServer starts and runs a new server for pprof.
<原文结束>

# <翻译开始>
// StartPProfServer 启动并运行一个新的 pprof 服务端。
# <翻译结束>


<原文开始>
// EnablePProf enables PProf feature for server.
<原文结束>

# <翻译开始>
// EnablePProf 启用服务器的 PProf 功能。
# <翻译结束>


<原文开始>
// EnablePProf enables PProf feature for server of specified domain.
<原文结束>

# <翻译开始>
// EnablePProf 启用指定域名服务器的 PProf 功能。
# <翻译结束>


<原文开始>
// Index shows the PProf index page.
<原文结束>

# <翻译开始>
// Index 显示 PProf 索引页面。
# <翻译结束>


<原文开始>
// Cmdline responds with the running program's
// command line, with arguments separated by NUL bytes.
// The package initialization registers it as /debug/pprof/cmdline.
<原文结束>

# <翻译开始>
// Cmdline 函数响应运行程序的命令行参数，其中各个参数由 NUL 字节分隔。
// 包初始化时会将其注册为 /debug/pprof/cmdline 路径。
# <翻译结束>


<原文开始>
// Profile responds with the pprof-formatted cpu profile.
// Profiling lasts for duration specified in seconds GET parameter, or for 30 seconds if not specified.
// The package initialization registers it as /debug/pprof/profile.
<原文结束>

# <翻译开始>
// Profile 函数响应 pprof 格式的 CPU 分析报告。
// 分析的持续时间由 GET 参数中指定的秒数决定，如果未指定，则默认为 30 秒。
// 包初始化时会将其注册为 /debug/pprof/profile 路径。
# <翻译结束>


<原文开始>
// Symbol looks up the program counters listed in the request,
// responding with a table mapping program counters to function names.
// The package initialization registers it as /debug/pprof/symbol.
<原文结束>

# <翻译开始>
// Symbol 函数通过查询请求中列出的程序计数器，
// 并以程序计数器到函数名称的映射表作为响应。
// 包初始化时将其注册为 /debug/pprof/symbol 路径。
# <翻译结束>


<原文开始>
// Trace responds with the execution trace in binary form.
// Tracing lasts for duration specified in seconds GET parameter, or for 1 second if not specified.
// The package initialization registers it as /debug/pprof/trace.
<原文结束>

# <翻译开始>
// Trace 以二进制形式响应执行跟踪。
// 跟踪持续时间由GET参数中指定的秒数决定，如果未指定，则持续1秒。
// 包初始化时将其注册为/debug/pprof/trace。
# <翻译结束>

