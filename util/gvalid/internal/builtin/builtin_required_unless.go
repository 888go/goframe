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
// RuleRequiredUnless 实现了 `required-unless` 规则：
// 当且仅当所有给定的字段及其对应值不相等时，该字段才是必填的。
//
// 格式： required-unless:field,value,...
// 示例： required-unless:id,1,age,18
// 这段代码注释描述了一个名为`RuleRequiredUnless`的Go语言规则实现，该规则用于表单验证或者其他数据校验场景。具体而言，某个字段只有在其它指定字段与其对应的值不相等的情况下，才被视为必填项。注释中给出了该规则的格式示例，表明需要按照 "字段名,字段值,..." 的形式来配置所需的条件。例如："required-unless:id,1,age,18" 表示如果 id 不为 1 且 age 不为 18，则当前字段是必需的。
type RuleRequiredUnless struct{}

func init() {
	Register(RuleRequiredUnless{})
}

func (r RuleRequiredUnless) Name() string {
	return "required-unless"
}

func (r RuleRequiredUnless) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredUnless) Run(in RunInput) error {
	var (
		required   = true
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
				required = !strings.EqualFold(tv, gconv.String(foundValue))
			} else {
				required = strings.Compare(tv, gconv.String(foundValue)) != 0
			}
			if !required {
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
