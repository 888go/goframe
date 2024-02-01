// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
	)
// RuleBefore 实现了 `before` 规则：
// 该日期时间值应该在字段 `field` 的值之后。
//
// 格式：before:field
type RuleBefore struct{}

func init() {
	Register(RuleBefore{})
}

func (r RuleBefore) Name() string {
	return "before"
}

func (r RuleBefore) Message() string {
	return "The {field} value `{value}` must be before field {field1} value `{value1}`"
}

func (r RuleBefore) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.Before(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
