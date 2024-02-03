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

// RuleRequiredWithout 实现了 `required-without` 规则：
// 如果给定的任意字段为空，则该字段为必填。
//
// 格式： required-without:field1,field2,...
// 示例：required-without:id,name
type RuleRequiredWithout struct{}

func init() {
	Register(RuleRequiredWithout{})
}

func (r RuleRequiredWithout) Name() string {
	return "required-without"
}

func (r RuleRequiredWithout) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredWithout) Run(in RunInput) error {
	var (
		required   = false
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	for i := 0; i < len(array); i++ {
		_, foundValue = gutil.MapPossibleItemByKey(in.Data.Map(), array[i])
		if empty.IsEmpty(foundValue) {
			required = true
			break
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
