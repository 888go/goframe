// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// RuleRequiredUnless 实现了 `required-unless` 规则：
// 当给定的字段及其值都不相等时，该字段是必需的。
//
// 格式：required-unless:field,value,...
// 示例：required-unless:id,1,age,18
// md5:3492b6f7c0cf1435
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
		dataMap    = in.Data.Map()
	)

		// 它支持多个字段和值对。 md5:42a9e200e1db00d5
	if len(array)%2 == 0 {
		for i := 0; i < len(array); {
			tk := array[i]
			tv := array[i+1]
			_, foundValue = gutil.MapPossibleItemByKey(dataMap, tk)
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
