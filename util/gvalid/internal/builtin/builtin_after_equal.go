// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// RuleAfterEqual 实现了 `after-equal` 规则：
// 日期时间值应该在字段 `field` 的值之后或等于该值。
//
// 格式：after-equal:field md5:b4f07a240c382495
type RuleAfterEqual struct{}

func init() {
	Register(RuleAfterEqual{})
}

func (r RuleAfterEqual) Name() string {
	return "after-equal"
}

func (r RuleAfterEqual) Message() string {
	return "The {field} value `{value}` must be after or equal to field {field1} value `{value1}`"
}

func (r RuleAfterEqual) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.After(fieldDatetime) || valueDatetime.Equal(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
