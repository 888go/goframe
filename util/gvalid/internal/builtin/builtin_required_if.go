// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// RuleRequiredIf 实现了 `required-if` 规则：
// 如果给定的字段及其值中的任何一个相等，则为必填。
//
// 格式：required-if:field,value,...
// 示例：required-if:id,1,age,18 md5:fc6bf09f9de6e20b
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
		dataMap    = in.Data.Map()
	)
	if len(array)%2 != 0 {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid "%s" rule pattern: %s`,
			r.Name(),
			in.RulePattern,
		)
	}
	// 它支持多个字段和值对。 md5:42a9e200e1db00d5
	for i := 0; i < len(array); {
		var (
			tk = array[i]
			tv = array[i+1]
		)
		_, foundValue = gutil.MapPossibleItemByKey(dataMap, tk)
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
	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
