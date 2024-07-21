// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

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
	goRootForFilter  = runtime.GOROOT() // goRootForFilter 用于栈过滤目的。 md5:538cfd57e5493ca3
	binaryVersion    = ""               // 当前正在运行的二进制文件的版本（十六进制表示的uint64）。 md5:c6f3c11c0f44f372
	binaryVersionMd5 = ""               // 当前运行二进制文件的版本（MD5）。 md5:0e6faf25b85553ad
	selfPath         = ""               // 当前运行二进制文件的绝对路径。 md5:e56aa0343f75f24c
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
	// 初始化内部包变量：selfPath。 md5:8d1168ac7361bb54
	selfPath, _ = exec.LookPath(os.Args[0])
	if selfPath != "" {
		selfPath, _ = filepath.Abs(selfPath)
	}
	if selfPath == "" {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
}

// Caller 返回调用者函数的名称，以及包含该函数的绝对文件路径和行号。
// md5:ede3e19ac5afa26d
func Caller(skip ...int) (function string, path string, line int) {
	return CallerWithFilter(nil, skip...)
}

// CallerWithFilter 返回调用者函数名、绝对文件路径及其行号。
// 
// 参数 `filters` 用于过滤调用者的路径。
// md5:77e7b623dc180797
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

// callerFromIndex 返回调用位置及其相关信息，但不包括 debug 包的内容。
//
// 非常注意，返回的索引值应该是 `index - 1`，因为调用者的起点是从 1 开始的。
// md5:7a22ee9c6da468f5
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
	// Filter empty file.
	if file == "" {
		return true
	}
	// 过滤掉gdebug包的调用。 md5:c68547ef782edd13
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
		// https://github.com/gogf/gf/issues/2047
// 
// 这段注释是链接到GitHub上一个名为gf的项目中的问题编号2047。在Go语言中，这种注释通常用于引用外部资源、问题或讨论，以便其他开发者可以查看更多的上下文信息。 md5:3146de65d5a8eeb4
		fileSeparator := file[len(goRootForFilter)]
		if fileSeparator == filepath.Separator || fileSeparator == '\\' || fileSeparator == '/' {
			return true
		}
	}
	return false
}

// CallerPackage 返回调用者的包名。 md5:5ce61ae99065c96c
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

// CallerFunction 返回调用者函数的名称。 md5:af9ad1617f64a1c2
func CallerFunction() string {
	function, _, _ := Caller()
	function = function[strings.LastIndexByte(function, '/')+1:]
	function = function[strings.IndexByte(function, '.')+1:]
	return function
}

// CallerFilePath 返回调用者的文件路径。 md5:cf0e426c8a45ff1b
func CallerFilePath() string {
	_, path, _ := Caller()
	return path
}

// CallerDirectory 返回调用者的目录。 md5:13a2e4e3afd2554a
func CallerDirectory() string {
	_, path, _ := Caller()
	return filepath.Dir(path)
}

// CallerFileLine 返回调用者的文件路径和行号。 md5:94cba50c9cbd0bd5
func CallerFileLine() string {
	_, path, line := Caller()
	return fmt.Sprintf(`%s:%d`, path, line)
}

// CallerFileLineShort 返回调用者所在的文件名和行号。 md5:ca795c06dfcf9d18
func CallerFileLineShort() string {
	_, path, line := Caller()
	return fmt.Sprintf(`%s:%d`, filepath.Base(path), line)
}

// FuncPath 返回给定函数`f`的完整函数路径。 md5:fcff03839e6125e6
func FuncPath(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// FuncName 返回给定函数`f`的函数名称。 md5:6ccbd3c81a265a6e
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
