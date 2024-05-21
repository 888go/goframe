// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gutil

// GetOrDefaultStr checks and returns value according whether parameter `param` available.
// It returns `param[0]` if it is available, or else it returns `def`.

// ff:取文本值或取默认值
// param:待判断变量
// def:默认值
func GetOrDefaultStr(def string, param ...string) string {
	value := def
	if len(param) > 0 && param[0] != "" {
		value = param[0]
	}
	return value
}

// GetOrDefaultAny checks and returns value according whether parameter `param` available.
// It returns `param[0]` if it is available, or else it returns `def`.

// ff:取值或取默认值
// param:待判断变量
// def:默认值
func GetOrDefaultAny(def interface{}, param ...interface{}) interface{} {
	value := def
	if len(param) > 0 && param[0] != "" {
		value = param[0]
	}
	return value
}