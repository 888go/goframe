// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	"strings"
	
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/util/gutil"
)

// RuleRequiredWith 实现了 `required-with` 规则：
// 当给定的任意字段非空时，该字段为必填。
//
// 格式： required-with:field1,field2,...
// 示例： required-with:id,name
type RuleRequiredWith struct{}

func init() {
	Register(RuleRequiredWith{})
}

func (r RuleRequiredWith) Name() string {
	return "required-with"
}

func (r RuleRequiredWith) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredWith) Run(in RunInput) error {
	var (
		required   = false
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	for i := 0; i < len(array); i++ {
		_, foundValue = 工具类.MapPossibleItemByKey(in.Data.X取Map(), array[i])
		if !empty.IsEmpty(foundValue) {
			required = true
			break
		}
	}

	if required && isRequiredEmpty(in.Value.X取值()) {
		return errors.New(in.Message)
	}
	return nil
}
