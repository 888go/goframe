// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类

import (
	"os"
	"runtime"
	"strings"
	
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
)

var (
	// goRootForFilter 用于堆栈过滤的目的。
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}

// MainPkgPath 返回包含入口函数 main 的 main 包的绝对文件路径。
//
// 该功能仅在开发环境中可用。
//
// 注意1：仅对源代码开发环境有效，即仅对生成此可执行文件的系统有效。
//
// 注意2：当首次调用该方法时，如果处于异步 goroutine 中，则该方法可能无法获取到 main 包的路径。
func X取main路径() string {
	// 这仅适用于源代码开发环境。
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
// 检查是否在包初始化函数中被调用，
// 在这种情况下，此处无法获取主包路径，
// 因此仅返回一个值以便进行后续检查。
			if fn := runtime.FuncForPC(pc); fn != nil {
				array := 文本类.X分割(fn.Name(), ".")
				if array[0] != "main" {
					continue
				}
			}
			if 正则类.X是否匹配文本(`package\s+main\s+`, X读文本(file)) {
				mainPkgPath.X设置值(X路径取父目录(file))
				return X路径取父目录(file)
			}
		} else {
			break
		}
	}
// 如果仍然无法找到包main的路径，
// 它会递归地搜索最后一个go文件所在的目录及其父目录。
// 这通常对于业务项目进行单元测试的情况是必要的。
	if lastFile != "" {
		for path = X路径取父目录(lastFile); len(path) > 1 && X是否存在(path) && path[len(path)-1] != os.PathSeparator; {
			files, _ := X枚举并含子目录名(path, "*.go")
			for _, v := range files {
				if 正则类.X是否匹配文本(`package\s+main\s+`, X读文本(v)) {
					mainPkgPath.X设置值(path)
					return path
				}
			}
			path = X路径取父目录(path)
		}
	}
	return ""
}
