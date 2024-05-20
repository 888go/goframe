
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
// Package gproc implements management and communication for processes.
<原文结束>

# <翻译开始>
// gproc 包实现了对进程的管理和通信功能。. md5:2bdecc6699345c91
# <翻译结束>


<原文开始>
// processPid is the pid of current process.
<原文结束>

# <翻译开始>
// processPid 是当前进程的进程ID。. md5:72add22026a94fdb
# <翻译结束>


<原文开始>
// processStartTime is the start time of current process.
<原文结束>

# <翻译开始>
// processStartTime 是当前进程的启动时间。. md5:447a3fe1c369aced
# <翻译结束>


<原文开始>
// Pid returns the pid of current process.
<原文结束>

# <翻译开始>
// Pid返回当前进程的pid。. md5:547eaf09253b67f9
# <翻译结束>


<原文开始>
// PPid returns the custom parent pid if exists, or else it returns the system parent pid.
<原文结束>

# <翻译开始>
// PPid 返回自定义的父进程ID，如果存在的话，否则返回系统的父进程ID。. md5:177a13dad5ed9a39
# <翻译结束>


<原文开始>
// PPidOS returns the system parent pid of current process.
// Note that the difference between PPidOS and PPid function is that the PPidOS returns
// the system ppid, but the PPid functions may return the custom pid by gproc if the custom
// ppid exists.
<原文结束>

# <翻译开始>
// PPidOS 返回当前进程的系统父进程ID。
// 请注意，PPidOS与PPid函数的区别在于，PPidOS返回系统的父进程ID，而如果存在自定义父进程ID，PPid函数可能会返回由gproc设置的自定义进程ID。
// md5:f6f56ec93bfd6b19
# <翻译结束>


<原文开始>
// IsChild checks and returns whether current process is a child process.
// A child process is forked by another gproc process.
<原文结束>

# <翻译开始>
// IsChild 检查并返回当前进程是否是子进程。
// 子进程是由另一个gproc进程 fork() 创建的。
// md5:9ec53f2cdad75233
# <翻译结束>


<原文开始>
// SetPPid sets custom parent pid for current process.
<原文结束>

# <翻译开始>
// SetPPid 设置当前进程的自定义父进程ID。. md5:6da79f2272f63e59
# <翻译结束>


<原文开始>
// StartTime returns the start time of current process.
<原文结束>

# <翻译开始>
// StartTime 返回当前进程的启动时间。. md5:322d4b9a3dae1290
# <翻译结束>


<原文开始>
// Uptime returns the duration which current process has been running
<原文结束>

# <翻译开始>
// Uptime 返回当前进程已经运行的持续时间. md5:105744cf83fdec5c
# <翻译结束>


<原文开始>
// SearchBinary searches the binary `file` in current working folder and PATH environment.
<原文结束>

# <翻译开始>
// SearchBinary 在当前工作目录和PATH环境变量中搜索名为`file`的二进制文件。. md5:56a48fa45711f1c2
# <翻译结束>


<原文开始>
// Check if it is absolute path of exists at current working directory.
<原文结束>

# <翻译开始>
// 检查它是否是当前工作目录下存在的绝对路径。. md5:5c4a5911487345cd
# <翻译结束>


<原文开始>
// SearchBinaryPath searches the binary `file` in PATH environment.
<原文结束>

# <翻译开始>
// SearchBinaryPath 在PATH环境变量中搜索二进制文件`file`。. md5:2762ea99f9622d59
# <翻译结束>

