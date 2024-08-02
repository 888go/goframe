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

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// RuleRequiredIfAll 实现了 `required-if-all` 规则：
// 如果所有给定的字段及其值都相等，则为必填。
//
// 格式：required-if-all:field,value,...
// 示例：required-if-all:id,1,age,18
// md5:55b9b0db38b70d0e
type RuleRequiredIfAll struct{}

func init() {
	Register(RuleRequiredIfAll{})
}

func (r RuleRequiredIfAll) Name() string {
	return "required-if-all"
}

func (r RuleRequiredIfAll) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredIfAll) Run(in RunInput) error {
	var (
		required   = true
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
	for i := 0; i < len(array); {
		var (
			tk = array[i]
			tv = array[i+1]
			eq bool
		)
		_, foundValue = gutil.MapPossibleItemByKey(dataMap, tk)
		if in.Option.CaseInsensitive {
			eq = strings.EqualFold(tv, gconv.String(foundValue))
		} else {
			eq = strings.Compare(tv, gconv.String(foundValue)) == 0
		}
		if !eq {
			required = false
			break
		}
		i += 2
	}
	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
