// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gsel

// defaultBuilder 是全局使用的默认构建器。 md5:d52d961d6f061a1b
var defaultBuilder = NewBuilderRoundRobin()

// SetBuilder 设置全局使用的默认构建器。 md5:c9730a1fbeb24fc6
func SetBuilder(builder Builder) {
	defaultBuilder = builder
}

// GetBuilder 返回全局使用的默认构建器。 md5:34c08237f2d0e6f4
func GetBuilder() Builder {
	return defaultBuilder
}
