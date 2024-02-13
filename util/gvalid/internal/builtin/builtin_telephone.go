// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/text/gregex"
)

// RuleTelephone 实现了 `telephone` 规则：
// "XXXX-XXXXXXX"
// "XXXX-XXXXXXXX"
// "XXX-XXXXXXX"
// "XXX-XXXXXXXX"
// "XXXXXXX"
// "XXXXXXXX"
// 格式：电话号码
// 这段Go语言代码的注释翻译成中文注释如下：
// ```go
// RuleTelephone 实现了电话号码规则：
// 格式可能为：
// "XXXX-XXXXXXX"（四位区号-七位本地号码）
// "XXXX-XXXXXXXX"（四位区号-八位本地号码）
// "XXX-XXXXXXX"（三位区号-七位本地号码）
// "XXX-XXXXXXXX"（三位区号-八位本地号码）
// "XXXXXXX"（七位纯数字号码）
// "XXXXXXXX"（八位纯数字号码）
// 格式：电话号码
type RuleTelephone struct{}

func init() {
	Register(RuleTelephone{})
}

func (r RuleTelephone) Name() string {
	return "telephone"
}

func (r RuleTelephone) Message() string {
	return "The {field} value `{value}` is not a valid telephone number"
}

func (r RuleTelephone) Run(in RunInput) error {
	ok := 正则类.X是否匹配文本(
		`^((\d{3,4})|\d{3,4}-)?\d{7,8}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
