// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类

// GetOrDefaultStr 检查并根据参数 `param` 是否可用返回值。
// 如果 `param[0]` 可用，返回 `param[0]`，否则返回 `def`。
func X取文本值或取默认值(默认值 string, 待判断变量 ...string) string {
	value := 默认值
	if len(待判断变量) > 0 && 待判断变量[0] != "" {
		value = 待判断变量[0]
	}
	return value
}

// GetOrDefaultAny 检查并根据参数 `param` 是否可用返回值。
// 如果 `param[0]` 可用，它将返回 `param[0]`；否则返回 `def`。
func X取值或取默认值(默认值 interface{}, 待判断变量 ...interface{}) interface{} {
	value := 默认值
	if len(待判断变量) > 0 && 待判断变量[0] != "" {
		value = 待判断变量[0]
	}
	return value
}
