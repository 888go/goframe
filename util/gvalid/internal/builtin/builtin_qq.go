// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
	)
// RuleQQ 实现了 `qq` 规则：
// 腾讯QQ号码。
//
// 格式：qq
type RuleQQ struct{}

func init() {
	Register(RuleQQ{})
}

func (r RuleQQ) Name() string {
	return "qq"
}

func (r RuleQQ) Message() string {
	return "The {field} value `{value}` is not a valid QQ number"
}

func (r RuleQQ) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^[1-9][0-9]{4,}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
