// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdebug

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

const (
	maxCallerDepth = 1000
	stackFilterKey = "/debug/gdebug/gdebug"
)

var (
	goRootForFilter  = runtime.GOROOT() // goRootForFilter 用于堆栈过滤的目的。
	binaryVersion    = ""               // 当前运行二进制文件的版本（uint64进制表示）。
	binaryVersionMd5 = ""               // 当前运行二进制文件的版本（MD5）。
	selfPath         = ""               // 当前运行的二进制文件绝对路径。
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
	// 初始化内部包变量：selfPath。
	selfPath, _ = exec.LookPath(os.Args[0])
	if selfPath != "" {
		selfPath, _ = filepath.Abs(selfPath)
	}
	if selfPath == "" {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
}

// Caller 返回调用者函数的名称以及其所在的绝对文件路径及行号。
func Caller(skip ...int) (function string, path string, line int) {
	return CallerWithFilter(nil, skip...)
}

// CallerWithFilter 返回调用者函数名称以及绝对文件路径及其行号。
//
// 参数`filters`用于过滤调用者路径。
func CallerWithFilter(filters []string, skip ...int) (function string, path string, line int) {
	var (
		number = 0
		ok     = true
	)
	if len(skip) > 0 {
		number = skip[0]
	}
	pc, file, line, start := callerFromIndex(filters)
	if start != -1 {
		for i := start + number; i < maxCallerDepth; i++ {
			if i != start {
				pc, file, line, ok = runtime.Caller(i)
			}
			if ok {
				if filterFileByFilters(file, filters) {
					continue
				}
				function = ""
				if fn := runtime.FuncForPC(pc); fn == nil {
					function = "unknown"
				} else {
					function = fn.Name()
				}
				return function, file, line
			} else {
				break
			}
		}
	}
	return "", "", -1
}

// callerFromIndex 返回调用位置及其相关信息，但排除了 debug 包内的信息。
//
// 非常注意：返回的索引值应为 `index - 1`，作为调用者开始的位置点。
func callerFromIndex(filters []string) (pc uintptr, file string, line int, index int) {
	var ok bool
	for index = 0; index < maxCallerDepth; index++ {
		if pc, file, line, ok = runtime.Caller(index); ok {
			if filterFileByFilters(file, filters) {
				continue
			}
			if index > 0 {
				index--
			}
			return
		}
	}
	return 0, "", -1, -1
}

func filterFileByFilters(file string, filters []string) (filtered bool) {
	// 过滤空文件。
	if file == "" {
		return true
	}
	// 过滤gdebug包的调用。
	if strings.Contains(file, stackFilterKey) {
		return true
	}
	for _, filter := range filters {
		if filter != "" && strings.Contains(file, filter) {
			return true
		}
	}
	// GOROOT filter.
	if goRootForFilter != "" && len(file) >= len(goRootForFilter) && file[0:len(goRootForFilter)] == goRootForFilter {
		// 这是 gf(golang frame) 项目在 GitHub 上的一个 issue 链接，具体为第 2047 号问题。
		fileSeparator := file[len(goRootForFilter)]
		if fileSeparator == filepath.Separator || fileSeparator == '\\' || fileSeparator == '/' {
			return true
		}
	}
	return false
}

// CallerPackage 返回调用者的包名。
func CallerPackage() string {
	function, _, _ := Caller()
	indexSplit := strings.LastIndexByte(function, '/')
	if indexSplit == -1 {
		return function[:strings.IndexByte(function, '.')]
	} else {
		leftPart := function[:indexSplit+1]
		rightPart := function[indexSplit+1:]
		indexDot := strings.IndexByte(function, '.')
		rightPart = rightPart[:indexDot-1]
		return leftPart + rightPart
	}
}

// CallerFunction 返回调用者函数的名称。
func CallerFunction() string {
	function, _, _ := Caller()
	function = function[strings.LastIndexByte(function, '/')+1:]
	function = function[strings.IndexByte(function, '.')+1:]
	return function
}

// CallerFilePath 返回调用者所在的文件路径。
func CallerFilePath() string {
	_, path, _ := Caller()
	return path
}

// CallerDirectory 返回调用者所在的目录。
func CallerDirectory() string {
	_, path, _ := Caller()
	return filepath.Dir(path)
}

// CallerFileLine 返回调用者所在的文件路径及行号。
func CallerFileLine() string {
	_, path, line := Caller()
	return fmt.Sprintf(`%s:%d`, path, line)
}

// CallerFileLineShort 返回调用者所在的文件名及行号。
func CallerFileLineShort() string {
	_, path, line := Caller()
	return fmt.Sprintf(`%s:%d`, filepath.Base(path), line)
}

// FuncPath 返回给定函数 `f` 的完整函数路径。
func FuncPath(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// FuncName 返回给定函数 `f` 的函数名称。
func FuncName(f interface{}) string {
	path := FuncPath(f)
	if path == "" {
		return ""
	}
	index := strings.LastIndexByte(path, '/')
	if index < 0 {
		index = strings.LastIndexByte(path, '\\')
	}
	return path[index+1:]
}
