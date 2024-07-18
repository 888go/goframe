// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package gmap

import (
	"github.com/gogf/gf/v2/container/gtree"
)

// 基于红黑树的TreeMap，是RedBlackTree的别名。 md5:9f16a90eb8bdf4c1
type TreeMap = gtree.RedBlackTree

// NewTreeMap 使用自定义比较器创建一个树形映射。
// 参数 `safe` 用于指定是否在并发安全环境下使用树，其默认值为 false。
// md5:fde3476bb95496c2
// ff:创建红黑树Map
// comparator:回调函数
// v1:
// v2:
// safe:并发安全
func NewTreeMap(comparator func(v1, v2 interface{}) int, safe ...bool) *TreeMap {
	return gtree.NewRedBlackTree(comparator, safe...)
}

// NewTreeMapFrom使用自定义比较器和`data`映射创建一个树形映射。
// 注意，`data`映射将被设置为底层数据映射（不进行深拷贝），在外部更改映射时可能会存在并发安全问题。
// 参数`safe`用于指定是否使用并发安全的树，默认为false。
// md5:2421c85842b1f367
// ff:创建红黑树Map并从Map
// comparator:回调函数
// v1:
// v2:
// data:map值
// safe:并发安全
func NewTreeMapFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *TreeMap {
	return gtree.NewRedBlackTreeFrom(comparator, data, safe...)
}
