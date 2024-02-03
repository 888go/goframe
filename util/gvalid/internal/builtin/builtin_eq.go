// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

// RuleEq 实现了 `eq` 规则：
// 值应与字段的值相同。
//
// 该规则与 `same` 规则执行相同的操作。
//
// 格式：eq:field
type RuleEq struct{}

func init() {
	Register(RuleEq{})
}

func (r RuleEq) Name() string {
	return "eq"
}

func (r RuleEq) Message() string {
	return "The {field} value `{value}` must be equal to field {field1} value `{value1}`"
}

func (r RuleEq) Run(in RunInput) error {
	return RuleSame{}.Run(in)
}
