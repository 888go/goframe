// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gfile

import (
	"bytes"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Search 通过名称 `name` 在以下路径中搜索文件，按照优先级进行：优先搜索路径（prioritySearchPaths）、当前工作目录（Pwd()）、自身目录（SelfDir()）和主包路径（MainPkgPath()）。如果找到，则返回文件的绝对路径；如果没有找到，则返回空字符串。
// md5:4251b6145a87bd38
func Search(name string, prioritySearchPaths ...string) (realPath string, err error) {
	// 检查它是否是绝对路径。 md5:fcbf7e8f8e3d02b7
	realPath = RealPath(name)
	if realPath != "" {
		return
	}
	// Search paths array.
	array := garray.NewStrArray()
	array.Append(prioritySearchPaths...)
	array.Append(Pwd(), SelfDir())
	if path := MainPkgPath(); path != "" {
		array.Append(path)
	}
	// Remove repeated items.
	array.Unique()
	// Do the searching.
	array.RLockFunc(func(array []string) {
		path := ""
		for _, v := range array {
			path = RealPath(v + Separator + name)
			if path != "" {
				realPath = path
				break
			}
		}
	})
	// 如果搜索失败，它将返回格式化的错误信息。 md5:fb9f11e1e7a4e8fc
	if realPath == "" {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteString(fmt.Sprintf(`cannot find "%s" in following paths:`, name))
		array.RLockFunc(func(array []string) {
			for k, v := range array {
				buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
			}
		})
		err = gerror.New(buffer.String())
	}
	return
}
