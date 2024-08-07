// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	"os"
	"runtime"
	"strings"

	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
)

var (
		// goRootForFilter 用于栈过滤目的。 md5:538cfd57e5493ca3
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}

// X取main路径 返回包含入口函数main的package main的绝对文件路径。
//
// 它仅在开发环境中可用。
//
// 注意1：仅对源代码开发环境有效，
// 即仅对生成此可执行文件的系统有效。
//
// 注意2：首次调用此方法时，如果处于异步goroutine中，
// 方法可能无法获取到main包的路径。
// md5:7fb1d2fdcb626f85
func X取main路径() string {
		// 仅供源代码开发环境使用。 md5:56e807aeb00eee19
	if goRootForFilter == "" {
		return ""
	}
	path := mainPkgPath.X取值()
	if path != "" {
		return path
	}
	var lastFile string
	for i := 1; i < 10000; i++ {
		if pc, file, _, ok := runtime.Caller(i); ok {
			if goRootForFilter != "" && len(file) >= len(goRootForFilter) && file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			if X路径取扩展名(file) != ".go" {
				continue
			}
			lastFile = file
			// 检查它是否在包初始化函数中被调用，
			// 在这种情况下，无法获取主包路径，
			// 所以直接返回，以便进行下一步检查。
			// md5:e583ee52c2832f4d
			if fn := runtime.FuncForPC(pc); fn != nil {
				array := gstr.X分割(fn.Name(), ".")
				if array[0] != "main" {
					continue
				}
			}
			if gregex.X是否匹配文本(`package\s+main\s+`, X读文本(file)) {
				mainPkgPath.X设置值(X路径取父目录(file))
				return X路径取父目录(file)
			}
		} else {
			break
		}
	}
	// 如果仍然无法找到main包的路径，它会递归地搜索最后一个go文件的目录及其父目录。这对于商业项目中的整数测试用例通常是必要的。
	// md5:5bee1ce703ae05d8
	if lastFile != "" {
		for path = X路径取父目录(lastFile); len(path) > 1 && X是否存在(path) && path[len(path)-1] != os.PathSeparator; {
			files, _ := X枚举并含子目录名(path, "*.go")
			for _, v := range files {
				if gregex.X是否匹配文本(`package\s+main\s+`, X读文本(v)) {
					mainPkgPath.X设置值(path)
					return path
				}
			}
			path = X路径取父目录(path)
		}
	}
	return ""
}
