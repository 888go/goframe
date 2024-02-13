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

// RuleGT 实现了 `gt` 规则：
// 大于 `field`。
// 它同时支持整数和浮点数类型。
//
// 格式：gt:field
type RuleGT struct{}

func init() {
	Register(RuleGT{})
}

func (r RuleGT) Name() string {
	return "gt"
}

func (r RuleGT) Message() string {
	return "The {field} value `{value}` must be greater than field {field1} value `{value1}`"
}

func (r RuleGT) Run(in RunInput) error {
	var (
		fieldName, fieldValue = 工具类.MapPossibleItemByKey(in.Data.X取Map(), in.RulePattern)
		fieldValueN, err1     = strconv.ParseFloat(转换类.String(fieldValue), 10)
		valueN, err2          = strconv.ParseFloat(in.Value.String(), 10)
	)

	if valueN <= fieldValueN || err1 != nil || err2 != nil {
		return errors.New(文本类.Map替换(in.Message, map[string]string{
			"{field1}": fieldName,
			"{value1}": 转换类.String(fieldValue),
		}))
	}
	return nil
}
