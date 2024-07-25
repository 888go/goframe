// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

// RuleEq 实现了 `eq` 规则：
// 值应与字段的值相同。
//
// 此规则的行为与 `same` 规则相同。
//
// 格式：eq:field md5:82514ba4addceb19
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
