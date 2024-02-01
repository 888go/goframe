// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"strings"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
	)
// RuleDifferent 实现了 `different` 规则：
// 当前值应与指定字段的值不同。
//
// 格式：different:field
type RuleDifferent struct{}

func init() {
	Register(RuleDifferent{})
}

func (r RuleDifferent) Name() string {
	return "different"
}

func (r RuleDifferent) Message() string {
	return "The {field} value `{value}` must be different from field {field1} value `{value1}`"
}

func (r RuleDifferent) Run(in RunInput) error {
	var (
		ok    = true
		value = in.Value.String()
	)
	fieldName, fieldValue := gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
	if fieldValue != nil {
		if in.Option.CaseInsensitive {
			ok = !strings.EqualFold(value, gconv.String(fieldValue))
		} else {
			ok = strings.Compare(value, gconv.String(fieldValue)) != 0
		}
	}
	if ok {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
