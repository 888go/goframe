// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
)

// ```go
// RuleBankCard 实现了 `bank-card` 规则：
// 银行卡号。
//
// 格式：bank-card
// ```
//
// 这段Go代码的注释是描述一个名为RuleBankCard的实现，它与验证银行卡号相关的规则有关。规则的格式是"bank-card"。
// md5:c9583cdff892228a
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

// checkLuHn 使用LUHN算法检查`value`。
// 通常用于银行卡号的验证。
// md5:fac6db232cdfe191
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
