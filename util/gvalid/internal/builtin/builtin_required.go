// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"reflect"

	gconv "github.com/888go/goframe/util/gconv"
)

// RuleRequired 实现了 `required` 规则。
// 格式：required
// md5:3ca664798f2da0f8
type RuleRequired struct{}

func init() {
	Register(RuleRequired{})
}

func (r RuleRequired) Name() string {
	return "required"
}

func (r RuleRequired) Message() string {
	return "The {field} field is required"
}

func (r RuleRequired) Run(in RunInput) error {
	if isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}

// isRequiredEmpty 检查并返回给定的值是否为空字符串。
// 请注意，如果给定的值为零整数，它将被视为非空。
// md5:746fb5e5fd71d752
func isRequiredEmpty(value interface{}) bool {
	reflectValue := reflect.ValueOf(value)
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	switch reflectValue.Kind() {
	case reflect.String, reflect.Map, reflect.Array, reflect.Slice:
		return reflectValue.Len() == 0
	}
	return gconv.String(value) == ""
}
