// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类

// X取文本值或取默认值 检查并返回值，根据参数`param`是否可用。
// 如果`param`可用，它返回`param[0]`，否则返回`def`。
// md5:dca14c4157e86ece
func X取文本值或取默认值(默认值 string, 待判断变量 ...string) string {
	value := 默认值
	if len(待判断变量) > 0 && 待判断变量[0] != "" {
		value = 待判断变量[0]
	}
	return value
}

// X取值或取默认值 检查并根据参数 `param` 是否存在返回值。
// 如果 `param` 可用，它将返回 `param[0]`，否则返回 `def`。
// md5:19ff7265ff047831
func X取值或取默认值(默认值 interface{}, 待判断变量 ...interface{}) interface{} {
	value := 默认值
	if len(待判断变量) > 0 && 待判断变量[0] != "" {
		value = 待判断变量[0]
	}
	return value
}
