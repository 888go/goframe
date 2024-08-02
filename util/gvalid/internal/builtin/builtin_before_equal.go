// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// RuleBeforeEqual 实现了“before-equal”规则：
// 日期时间值应该大于或等于字段 `field` 的值。
//
// 格式：before-equal:field
// md5:f50af52a18b33834
type RuleBeforeEqual struct{}

func init() {
	Register(RuleBeforeEqual{})
}

func (r RuleBeforeEqual) Name() string {
	return "before-equal"
}

func (r RuleBeforeEqual) Message() string {
	return "The {field} value `{value}` must be before or equal to field {pattern}"
}

func (r RuleBeforeEqual) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.Before(fieldDatetime) || valueDatetime.Equal(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
