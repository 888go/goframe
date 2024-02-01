// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
	)
// RuleDomain 实现了 `domain` 规则：
// 域名。
//
// 格式：域名
type RuleDomain struct{}

func init() {
	Register(RuleDomain{})
}

func (r RuleDomain) Name() string {
	return "domain"
}

func (r RuleDomain) Message() string {
	return "The {field} value `{value}` is not a valid domain format"
}

func (r RuleDomain) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^([0-9a-zA-Z][0-9a-zA-Z\-]{0,62}\.)+([a-zA-Z]{0,62})$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
