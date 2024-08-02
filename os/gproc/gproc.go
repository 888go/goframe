// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gproc 包实现了对进程的管理和通信功能。 md5:2bdecc6699345c91
package 进程类

import (
	"os"
	"runtime"
	"time"

	genv "github.com/888go/goframe/os/genv"
	gfile "github.com/888go/goframe/os/gfile"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

const (
	envKeyPPid            = "GPROC_PPID"
	tracingInstrumentName = "github.com/gogf/gf/v2/os/gproc.Process"
)

var (
	processPid       = os.Getpid() // processPid 是当前进程的进程ID。 md5:72add22026a94fdb
	processStartTime = time.Now()  // processStartTime 是当前进程的启动时间。 md5:447a3fe1c369aced
)

// Pid返回当前进程的pid。 md5:547eaf09253b67f9
func Pid() int {
	return processPid
}

// PPid 返回自定义的父进程ID，如果存在的话，否则返回系统的父进程ID。 md5:177a13dad5ed9a39
func PPid() int {
	if !IsChild() {
		return Pid()
	}
	ppidValue := os.Getenv(envKeyPPid)
	if ppidValue != "" && ppidValue != "0" {
		return gconv.Int(ppidValue)
	}
	return PPidOS()
}

// PPidOS 返回当前进程的系统父进程ID。
// 请注意，PPidOS与PPid函数的区别在于，PPidOS返回系统的父进程ID，而如果存在自定义父进程ID，PPid函数可能会返回由gproc设置的自定义进程ID。
// md5:f6f56ec93bfd6b19
func PPidOS() int {
	return os.Getppid()
}

// IsChild 检查并返回当前进程是否是子进程。
// 子进程是由另一个gproc进程 fork() 创建的。
// md5:9ec53f2cdad75233
func IsChild() bool {
	ppidValue := os.Getenv(envKeyPPid)
	return ppidValue != "" && ppidValue != "0"
}

// SetPPid 设置当前进程的自定义父进程ID。 md5:6da79f2272f63e59
func SetPPid(ppid int) error {
	if ppid > 0 {
		return os.Setenv(envKeyPPid, gconv.String(ppid))
	} else {
		return os.Unsetenv(envKeyPPid)
	}
}

// StartTime 返回当前进程的启动时间。 md5:322d4b9a3dae1290
func StartTime() time.Time {
	return processStartTime
}

// Uptime 返回当前进程已经运行的持续时间. md5:105744cf83fdec5c
func Uptime() time.Duration {
	return time.Since(processStartTime)
}

// SearchBinary 在当前工作目录和PATH环境变量中搜索名为`file`的二进制文件。 md5:56a48fa45711f1c2
func SearchBinary(file string) string {
		// 检查它是否是当前工作目录下存在的绝对路径。 md5:5c4a5911487345cd
	if gfile.Exists(file) {
		return file
	}
	return SearchBinaryPath(file)
}

// SearchBinaryPath 在PATH环境变量中搜索二进制文件`file`。 md5:2762ea99f9622d59
func SearchBinaryPath(file string) string {
	array := ([]string)(nil)
	switch runtime.GOOS {
	case "windows":
		envPath := genv.Get("PATH", genv.Get("Path")).String()
		if gstr.Contains(envPath, ";") {
			array = gstr.SplitAndTrim(envPath, ";")
		} else if gstr.Contains(envPath, ":") {
			array = gstr.SplitAndTrim(envPath, ":")
		}
		if gfile.Ext(file) != ".exe" {
			file += ".exe"
		}

	default:
		array = gstr.SplitAndTrim(genv.Get("PATH").String(), ":")
	}
	if len(array) > 0 {
		path := ""
		for _, v := range array {
			path = v + gfile.Separator + file
			if gfile.Exists(path) && gfile.IsFile(path) {
				return path
			}
		}
	}
	return ""
}
