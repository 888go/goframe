// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strconv"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// RuleGTE 实现了 `gte` 规则：
// 大于或等于 `field`。
// 它支持整数和浮点数。
//
// 格式：gte:field md5:c3fa2ab775ef036b
type RuleGTE struct{}

func init() {
	Register(RuleGTE{})
}

func (r RuleGTE) Name() string {
	return "gte"
}

func (r RuleGTE) Message() string {
	return "The {field} value `{value}` must be greater than or equal to field {field1} value `{value1}`"
}

func (r RuleGTE) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		fieldValueN, err1     = strconv.ParseFloat(gconv.String(fieldValue), 10)
		valueN, err2          = strconv.ParseFloat(in.Value.String(), 10)
	)

	if valueN < fieldValueN || err1 != nil || err2 != nil {
		return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
			"{field1}": fieldName,
			"{value1}": gconv.String(fieldValue),
		}))
	}
	return nil
}
