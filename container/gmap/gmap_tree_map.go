// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package map类

import (
	gtree "github.com/888go/goframe/container/gtree"
)

// 基于红黑树的TreeMap，是RedBlackTree的别名。 md5:9f16a90eb8bdf4c1
type TreeMap = gtree.RedBlackTree

// X创建红黑树Map 使用自定义比较器创建一个树形映射。
// 参数 `safe` 用于指定是否在并发安全环境下使用树，其默认值为 false。
// md5:fde3476bb95496c2
func X创建红黑树Map(回调函数 func(v1, v2 interface{}) int, 并发安全 ...bool) *TreeMap {
	return gtree.NewRedBlackTree(回调函数, 并发安全...)
}

// X创建红黑树Map并从Map使用自定义比较器和`data`映射创建一个树形映射。
// 注意，`data`映射将被设置为底层数据映射（不进行深拷贝），在外部更改映射时可能会存在并发安全问题。
// 参数`safe`用于指定是否使用并发安全的树，默认为false。
// md5:2421c85842b1f367
func X创建红黑树Map并从Map(回调函数 func(v1, v2 interface{}) int, map值 map[interface{}]interface{}, 并发安全 ...bool) *TreeMap {
	return gtree.NewRedBlackTreeFrom(回调函数, map值, 并发安全...)
}
