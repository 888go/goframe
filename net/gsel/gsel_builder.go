// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsel

// defaultBuilder 是用于全局通用目的的默认构建器。
var defaultBuilder = NewBuilderRoundRobin()

// SetBuilder 设置全局默认的构建器，用于全局通用目的。
func SetBuilder(builder Builder) {
	defaultBuilder = builder
}

// GetBuilder 返回用于全局通用目的的默认构建器。
func GetBuilder() Builder {
	return defaultBuilder
}
