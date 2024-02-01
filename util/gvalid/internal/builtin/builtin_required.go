// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"reflect"
	
	"github.com/888go/goframe/util/gconv"
	)
// RuleRequired 实现了 `required` 规则。
// 格式：required
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
