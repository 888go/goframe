
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
// utilPProf is the PProf interface implementer.
<原文结束>

# <翻译开始>
// utilPProf是实现PProf接口的结构。. md5:61c1485646c2e81a
# <翻译结束>


<原文开始>
// StartPProfServer starts and runs a new server for pprof.
<原文结束>

# <翻译开始>
// StartPProfServer 启动并运行一个新的pprof服务器。. md5:4c0c47dfda03a84b
# <翻译结束>


<原文开始>
// EnablePProf enables PProf feature for server.
<原文结束>

# <翻译开始>
// EnablePProf 启用服务器的PProf功能。. md5:5603a60f147574d1
# <翻译结束>


<原文开始>
// EnablePProf enables PProf feature for server of specified domain.
<原文结束>

# <翻译开始>
// EnablePProf 为指定域名的服务器启用 PProf 功能。. md5:46c19e5f1d55beb1
# <翻译结束>


<原文开始>
// Index shows the PProf index page.
<原文结束>

# <翻译开始>
// Index 显示 PProf 的索引页面。. md5:606e9224f8418b6e
# <翻译结束>


<原文开始>
// Cmdline responds with the running program's
// command line, with arguments separated by NUL bytes.
// The package initialization registers it as /debug/pprof/cmdline.
<原文结束>

# <翻译开始>
// Cmdline 响应正在运行程序的命令行，参数之间用 NULL 字节分隔。包初始化时将其注册为 /debug/pprof/cmdline。
// md5:35f5d246119cca43
# <翻译结束>


<原文开始>
// Profile responds with the pprof-formatted cpu profile.
// Profiling lasts for duration specified in seconds GET parameter, or for 30 seconds if not specified.
// The package initialization registers it as /debug/pprof/profile.
<原文结束>

# <翻译开始>
// Profile 使用pprof格式返回CPU profiling信息。
// 如果GET参数指定了持续时间，那么 profiling 将持续该秒数；如果没有指定，则默认为30秒。
// 在包初始化时，它会注册为 "/debug/pprof/profile"。
// md5:11bd281949c0ba3c
# <翻译结束>


<原文开始>
// Symbol looks up the program counters listed in the request,
// responding with a table mapping program counters to function names.
// The package initialization registers it as /debug/pprof/symbol.
<原文结束>

# <翻译开始>
// Symbol 查找请求中列出的程序计数器，
// 并以映射表的形式响应，该映射表将程序计数器与函数名称关联起来。
// 包初始化时将其注册为 /debug/pprof/symbol 路由。
// md5:2944ed5cfe9e0c52
# <翻译结束>


<原文开始>
// Trace responds with the execution trace in binary form.
// Tracing lasts for duration specified in seconds GET parameter, or for 1 second if not specified.
// The package initialization registers it as /debug/pprof/trace.
<原文结束>

# <翻译开始>
// Trace 返回执行跟踪的二进制形式。
// 跟踪将持续指定的GET参数中的秒数，如果没有指定，则为1秒。
// 包初始化时将其注册为/debug/pprof/trace。
// md5:02830b4c9b48681f
# <翻译结束>

