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

	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// RuleSame 实现了 `same` 规则：
// 值应该与字段的值相同。
//
// 格式：same:field
// md5:433074ee67f413a5
type RuleSame struct{}

func init() {
	Register(RuleSame{})
}

func (r RuleSame) Name() string {
	return "same"
}

func (r RuleSame) Message() string {
	return "The {field} value `{value}` must be the same as field {field1} value `{value1}`"
}

func (r RuleSame) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	fieldName, fieldValue := gutil.MapPossibleItemByKey(in.Data.X取Map(), in.RulePattern)
	if fieldValue != nil {
		if in.Option.CaseInsensitive {
			ok = strings.EqualFold(value, gconv.String(fieldValue))
		} else {
			ok = strings.Compare(value, gconv.String(fieldValue)) == 0
		}
	}
	if !ok {
		return errors.New(gstr.Map替换(in.Message, map[string]string{
			"{field1}": fieldName,
			"{value1}": gconv.String(fieldValue),
		}))
	}
	return nil
}
