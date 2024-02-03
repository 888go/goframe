// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gutil

// GetOrDefaultStr 检查并根据参数 `param` 是否可用返回值。
// 如果 `param[0]` 可用，返回 `param[0]`，否则返回 `def`。
func GetOrDefaultStr(def string, param ...string) string {
	value := def
	if len(param) > 0 && param[0] != "" {
		value = param[0]
	}
	return value
}

// GetOrDefaultAny 检查并根据参数 `param` 是否可用返回值。
// 如果 `param[0]` 可用，它将返回 `param[0]`；否则返回 `def`。
func GetOrDefaultAny(def interface{}, param ...interface{}) interface{} {
	value := def
	if len(param) > 0 && param[0] != "" {
		value = param[0]
	}
	return value
}
