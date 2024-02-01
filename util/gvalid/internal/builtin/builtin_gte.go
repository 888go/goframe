// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"strconv"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
	)
// RuleGTE 实现了 `gte` 规则：
// 大于或等于 `field`。
// 它同时支持整数和浮点数。
//
// 格式：gte:field
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
