// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin
import (
	"errors"
	)
// RuleBankCard 实现了 `bank-card` 规则：
// 银行卡号。
//
// 格式：银行卡号
type RuleBankCard struct{}

func init() {
	Register(RuleBankCard{})
}

func (r RuleBankCard) Name() string {
	return "bank-card"
}

func (r RuleBankCard) Message() string {
	return "The {field} value `{value}` is not a valid bank card number"
}

func (r RuleBankCard) Run(in RunInput) error {
	if r.checkLuHn(in.Value.String()) {
		return nil
	}
	return errors.New(in.Message)
}

// checkLuHn 使用LUHN算法校验`value`。
// 通常用于银行卡号的验证。
func (r RuleBankCard) checkLuHn(value string) bool {
	var (
		sum     = 0
		nDigits = len(value)
		parity  = nDigits % 2
	)
	for i := 0; i < nDigits; i++ {
		var digit = int(value[i] - 48)
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
