// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类

import (
	"bytes"
	"fmt"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Search searches file by name `name` in following paths with priority:
// prioritySearchPaths, Pwd()、SelfDir()、MainPkgPath().
// It returns the absolute file path of `name` if found, or en empty string if not found.
func X查找(文件名 string, 优先级查找路径 ...string) (路径 string, 错误 error) {
	// 检查是否为绝对路径。
	路径 = X取绝对路径且效验(文件名)
	if 路径 != "" {
		return
	}
	// 搜索路径数组。
	array := garray.NewStrArray()
	array.Append(优先级查找路径...)
	array.Append(X取当前工作目录(), X取当前进程目录())
	if path := X取main路径(); path != "" {
		array.Append(path)
	}
	// 移除重复的项。
	array.Unique()
	// Do the searching.
	array.RLockFunc(func(array []string) {
		path := ""
		for _, v := range array {
			path = X取绝对路径且效验(v + Separator + 文件名)
			if path != "" {
				路径 = path
				break
			}
		}
	})
	// 如果搜索失败，它将返回格式化的错误信息。
	if 路径 == "" {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteString(fmt.Sprintf(`cannot find "%s" in following paths:`, 文件名))
		array.RLockFunc(func(array []string) {
			for k, v := range array {
				buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
			}
		})
		错误 = gerror.New(buffer.String())
	}
	return
}
