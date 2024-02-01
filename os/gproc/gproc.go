// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gproc实现了对进程的管理和通信功能。
package gproc
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
		return gconv.Int(ppidValue)
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
		return os.Setenv(envKeyPPid, gconv.String(ppid))
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
	if gfile.Exists(file) {
		return file
	}
	return SearchBinaryPath(file)
}

// SearchBinaryPath在PATH环境变量中搜索二进制文件`file`。
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
