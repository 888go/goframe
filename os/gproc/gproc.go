// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gproc实现了对进程的管理和通信功能。
package 进程类

import (
	"os"
	"runtime"
	"time"
	
	"github.com/888go/goframe/os/genv"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

const (
	envKeyPPid            = "GPROC_PPID"
	tracingInstrumentName = "github.com/gogf/gf/v2/os/gproc.Process"
)

var (
	processPid       = os.Getpid() // processPid 是当前进程的进程ID。
	processStartTime = time.Now()  // processStartTime 是当前进程的启动时间。
)

// Pid 返回当前进程的进程ID。
func Pid() int {
	return processPid
}

// PPid 返回自定义父进程ID（如果存在），否则返回系统父进程ID。
func PPid() int {
	if !IsChild() {
		return Pid()
	}
	ppidValue := os.Getenv(envKeyPPid)
	if ppidValue != "" && ppidValue != "0" {
		return 转换类.X取整数(ppidValue)
	}
	return PPidOS()
}

// PPidOS 返回当前进程的系统父进程ID。
// 注意，PPidOS 和 PPid 函数之间的区别在于：PPidOS 返回的是系统的父进程ID，
// 但若存在自定义父进程ID，PPid 函数可能会返回由 gproc 提供的自定义进程ID。
func PPidOS() int {
	return os.Getppid()
}

// IsChild 检查并返回当前进程是否为子进程。
// 子进程是由另一个 gproc 进程 fork（派生）出来的。
func IsChild() bool {
	ppidValue := os.Getenv(envKeyPPid)
	return ppidValue != "" && ppidValue != "0"
}

// SetPPid 设置当前进程的自定义父进程ID。
func SetPPid(ppid int) error {
	if ppid > 0 {
		return os.Setenv(envKeyPPid, 转换类.String(ppid))
	} else {
		return os.Unsetenv(envKeyPPid)
	}
}

// StartTime 返回当前进程的启动时间。
func StartTime() time.Time {
	return processStartTime
}

// Uptime 返回当前进程已经运行的时间间隔
func Uptime() time.Duration {
	return time.Since(processStartTime)
}

// SearchBinary 在当前工作目录和PATH环境变量中搜索二进制文件 `file`。
func SearchBinary(file string) string {
	// 检查给定路径是否为绝对路径，或者在当前工作目录中是否存在。
	if 文件类.X是否存在(file) {
		return file
	}
	return SearchBinaryPath(file)
}

// SearchBinaryPath在PATH环境变量中搜索二进制文件`file`。
func SearchBinaryPath(file string) string {
	array := ([]string)(nil)
	switch runtime.GOOS {
	case "windows":
		envPath := 环境变量类.X取值("PATH", 环境变量类.X取值("Path")).String()
		if 文本类.X是否包含(envPath, ";") {
			array = 文本类.X分割并忽略空值(envPath, ";")
		} else if 文本类.X是否包含(envPath, ":") {
			array = 文本类.X分割并忽略空值(envPath, ":")
		}
		if 文件类.X路径取扩展名(file) != ".exe" {
			file += ".exe"
		}

	default:
		array = 文本类.X分割并忽略空值(环境变量类.X取值("PATH").String(), ":")
	}
	if len(array) > 0 {
		path := ""
		for _, v := range array {
			path = v + 文件类.Separator + file
			if 文件类.X是否存在(path) && 文件类.X是否为文件(path) {
				return path
			}
		}
	}
	return ""
}
