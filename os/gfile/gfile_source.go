// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile
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
func MainPkgPath() string {
	// 这仅适用于源代码开发环境。
	if goRootForFilter == "" {
		return ""
	}
	path := mainPkgPath.Val()
	if path != "" {
		return path
	}
	var lastFile string
	for i := 1; i < 10000; i++ {
		if pc, file, _, ok := runtime.Caller(i); ok {
			if goRootForFilter != "" && len(file) >= len(goRootForFilter) && file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			if Ext(file) != ".go" {
				continue
			}
			lastFile = file
// 检查是否在包初始化函数中被调用，
// 在这种情况下，此处无法获取主包路径，
// 因此仅返回一个值以便进行后续检查。
			if fn := runtime.FuncForPC(pc); fn != nil {
				array := gstr.Split(fn.Name(), ".")
				if array[0] != "main" {
					continue
				}
			}
			if gregex.IsMatchString(`package\s+main\s+`, GetContents(file)) {
				mainPkgPath.Set(Dir(file))
				return Dir(file)
			}
		} else {
			break
		}
	}
// 如果仍然无法找到包main的路径，
// 它会递归地搜索最后一个go文件所在的目录及其父目录。
// 这通常对于业务项目进行单元测试的情况是必要的。
	if lastFile != "" {
		for path = Dir(lastFile); len(path) > 1 && Exists(path) && path[len(path)-1] != os.PathSeparator; {
			files, _ := ScanDir(path, "*.go")
			for _, v := range files {
				if gregex.IsMatchString(`package\s+main\s+`, GetContents(v)) {
					mainPkgPath.Set(path)
					return path
				}
			}
			path = Dir(path)
		}
	}
	return ""
}
