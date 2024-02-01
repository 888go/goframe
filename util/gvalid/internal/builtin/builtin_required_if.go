// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"strings"
	
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
	)
// RuleRequiredIf 实现了 `required-if` 规则：
// 当且仅当所有给定的字段及其对应的值相等时，该字段才是必填项。
//
// 格式： required-if:field,value,...
// 示例： required-if: id,1,age,18
// 这段代码注释描述了一个名为`RuleRequiredIf`的Go语言实现，它遵循一个自定义验证规则——`required-if`。这个规则表示某个字段只有在其他指定字段具有特定值时才需要（即为必填）。具体来说，规则的格式是在字符串中以逗号分隔指定字段名和它们对应的值，例如："required-if: id,1,age,18"意味着如果id不等于1或age不等于18，则当前字段是必需填写的。
type RuleRequiredIf struct{}

func init() {
	Register(RuleRequiredIf{})
}

func (r RuleRequiredIf) Name() string {
	return "required-if"
}

func (r RuleRequiredIf) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredIf) Run(in RunInput) error {
	var (
		required   = false
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	// 它支持多个字段和值对。
	if len(array)%2 == 0 {
		for i := 0; i < len(array); {
			tk := array[i]
			tv := array[i+1]
			_, foundValue = gutil.MapPossibleItemByKey(in.Data.Map(), tk)
			if in.Option.CaseInsensitive {
				required = strings.EqualFold(tv, gconv.String(foundValue))
			} else {
				required = strings.Compare(tv, gconv.String(foundValue)) == 0
			}
			if required {
				break
			}
			i += 2
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
