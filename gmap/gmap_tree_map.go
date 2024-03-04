// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package gmap

import (
	"github.com/gogf/gf/v2/container/gtree"
)

// TreeMap基于红黑树实现，是RedBlackTree的别名。
type TreeMap = gtree.RedBlackTree

// NewTreeMap 创建一个使用自定义比较器的树形映射。
// 参数`safe`用于指定是否在并发安全的情况下使用树，其默认值为false。
func NewTreeMap(comparator func(v1, v2 interface{}) int, safe ...bool) *TreeMap {
	return gtree.NewRedBlackTree(comparator, safe...)
}

// NewTreeMapFrom 通过自定义比较器和`data`映射实例化一个新的树形映射。
// 注意，参数`data`映射将被直接设置为底层数据映射（非深度复制），
// 因此在外部修改该映射时可能存在并发安全问题。
// 参数`safe`用于指定是否使用线程安全的树形结构，默认情况下为false。
func NewTreeMapFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *TreeMap {
	return gtree.NewRedBlackTreeFrom(comparator, data, safe...)
}
