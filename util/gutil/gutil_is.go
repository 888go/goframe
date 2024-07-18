// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gutil

import (
	"reflect"

	"github.com/gogf/gf/v2/internal/empty"
)

// IsEmpty 检查给定的 `value` 是否为空。
// 如果 `value` 为以下情况，返回 false：整数0，bool 值 false，长度为0的切片/映射，或 nil；
// 否则，返回 true。
// md5:8e21f87627e70ce3
// ff:
// value:
func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}

// IsTypeOf 检查"value"和"valueInExpectType"的类型是否相等，并返回结果。 md5:e1d0bdccffd973a1
// ff:
// value:
// valueInExpectType:
func IsTypeOf(value, valueInExpectType interface{}) bool {
	return reflect.TypeOf(value) == reflect.TypeOf(valueInExpectType)
}
