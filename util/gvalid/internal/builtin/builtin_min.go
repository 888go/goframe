// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	"strconv"
	
	"github.com/888go/goframe/text/gstr"
	)
// RuleMin 实现了 `min` 规则：
// 大于等于 :min。它同时支持整数和浮点数。
//
// 格式：min:min
type RuleMin struct{}

func init() {
	Register(RuleMin{})
}

func (r RuleMin) Name() string {
	return "min"
}

func (r RuleMin) Message() string {
	return "The {field} value `{value}` must be equal or greater than {min}"
}

func (r RuleMin) Run(in RunInput) error {
	var (
		min, err1    = strconv.ParseFloat(in.RulePattern, 10)
		valueN, err2 = strconv.ParseFloat(in.Value.String(), 10)
	)
	if valueN < min || err1 != nil || err2 != nil {
		return errors.New(gstr.Replace(in.Message, "{min}", strconv.FormatFloat(min, 'f', -1, 64)))
	}
	return nil
}
