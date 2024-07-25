// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

// RuleNotEq 实现了 `not-eq` 规则：
// 值应该与字段的值不同。
//
// 格式：not-eq:字段 md5:0a6bcd6ab9b5e298
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
