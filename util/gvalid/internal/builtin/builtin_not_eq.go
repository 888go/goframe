// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

// RuleNotEq 实现了 `not-eq` 规则：
// 当前值应不同于字段的值。
//
// 格式：not-eq:field
type RuleNotEq struct{}

func init() {
	Register(RuleNotEq{})
}

func (r RuleNotEq) Name() string {
	return "not-eq"
}

func (r RuleNotEq) Message() string {
	return "The {field} value `{value}` must not be equal to field {field1} value `{value1}`"
}

func (r RuleNotEq) Run(in RunInput) error {
	return RuleDifferent{}.Run(in)
}
