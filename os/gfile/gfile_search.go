// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	"bytes"
	"fmt"

	garray "github.com/888go/goframe/container/garray"
	gerror "github.com/888go/goframe/errors/gerror"
)

// X查找 通过名称 `name` 在以下路径中搜索文件，按照优先级进行：优先搜索路径（prioritySearchPaths）、当前工作目录（Pwd()）、自身目录（SelfDir()）和主包路径（MainPkgPath()）。如果找到，则返回文件的绝对路径；如果没有找到，则返回空字符串。
// md5:4251b6145a87bd38
func X查找(文件名 string, 优先级查找路径 ...string) (路径 string, 错误 error) {
		// 检查它是否是绝对路径。 md5:fcbf7e8f8e3d02b7
	路径 = X取绝对路径且效验(文件名)
	if 路径 != "" {
		return
	}
	// Search paths array.
	array := garray.X创建文本()
	array.Append别名(优先级查找路径...)
	array.Append别名(X取当前工作目录(), X取当前进程目录())
	if path := X取main路径(); path != "" {
		array.Append别名(path)
	}
	// Remove repeated items.
	array.X去重()
	// Do the searching.
	array.X遍历读锁定(func(array []string) {
		path := ""
		for _, v := range array {
			path = X取绝对路径且效验(v + Separator + 文件名)
			if path != "" {
				路径 = path
				break
			}
		}
	})
		// 如果搜索失败，它将返回格式化的错误信息。 md5:fb9f11e1e7a4e8fc
	if 路径 == "" {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteString(fmt.Sprintf(`cannot find "%s" in following paths:`, 文件名))
		array.X遍历读锁定(func(array []string) {
			for k, v := range array {
				buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
			}
		})
		错误 = gerror.X创建(buffer.String())
	}
	return
}
